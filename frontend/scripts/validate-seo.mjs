import fs from 'node:fs/promises';
import path from 'node:path';
import { fileURLToPath } from 'node:url';

const __dirname = path.dirname(fileURLToPath(import.meta.url));
const distDir = path.resolve(__dirname, '..', 'dist');

const PUBLIC_PAGES = [
  { file: 'index.html', expectedPath: '/' },
  { file: path.join('about', 'index.html'), expectedPath: '/about' },
  { file: path.join('projects', 'index.html'), expectedPath: '/projects' },
  { file: path.join('skills', 'index.html'), expectedPath: '/skills' },
  { file: path.join('contact', 'index.html'), expectedPath: '/contact' },
];

const errors = [];
const seenTitles = new Set();
const seenDescriptions = new Set();
const seenCanonicals = new Set();

function assert(condition, message) {
  if (!condition) {
    errors.push(message);
  }
}

async function validatePage(page) {
  const fullPath = path.join(distDir, page.file);
  const html = await fs.readFile(fullPath, 'utf-8');

  const titleMatch = html.match(/<title>([\s\S]*?)<\/title>/i);
  const title = titleMatch ? titleMatch[1].trim() : '';
  assert(title.length > 10, `${page.file}: missing or too short <title>`);
  assert(!seenTitles.has(title), `${page.file}: duplicate <title> "${title}"`);
  seenTitles.add(title);

  const descMatch = html.match(/<meta\s+name="description"\s+content="([^"]*)"/i);
  const description = descMatch ? descMatch[1].trim() : '';
  assert(description.length >= 50, `${page.file}: missing or too short meta description`);
  assert(!seenDescriptions.has(description), `${page.file}: duplicate meta description`);
  seenDescriptions.add(description);

  const robotsMatch = html.match(/<meta\s+name="robots"\s+content="([^"]*)"/i);
  assert(robotsMatch && robotsMatch[1].includes('index') && !robotsMatch[1].includes('noindex'), `${page.file}: must allow indexing`);

  const canonicalMatches = [...html.matchAll(/<link\s+rel="canonical"\s+href="([^"]*)"/gi)];
  assert(canonicalMatches.length === 1, `${page.file}: expected exactly 1 canonical link, found ${canonicalMatches.length}`);
  if (canonicalMatches[0]) {
    const canonical = canonicalMatches[0][1];
    assert(!seenCanonicals.has(canonical), `${page.file}: duplicate canonical "${canonical}"`);
    seenCanonicals.add(canonical);
    const expectedSuffix = page.expectedPath === '/' ? '/' : page.expectedPath;
    assert(canonical.endsWith(expectedSuffix), `${page.file}: canonical "${canonical}" does not end with "${expectedSuffix}"`);
  }

  assert(/<meta\s+property="og:title"/i.test(html), `${page.file}: missing og:title`);
  assert(/<meta\s+property="og:description"/i.test(html), `${page.file}: missing og:description`);
  assert(/<meta\s+property="og:image"[^>]*og-preview\.png/i.test(html), `${page.file}: missing og:image`);
  assert(/<meta\s+name="twitter:card"/i.test(html), `${page.file}: missing twitter:card`);

  // JSON-LD check
  const jsonLdMatch = html.match(/<script\s+type="application\/ld\+json"[^>]*>([\s\S]*?)<\/script>/i);
  assert(Boolean(jsonLdMatch), `${page.file}: missing JSON-LD script`);
  if (jsonLdMatch) {
    try {
      const parsed = JSON.parse(jsonLdMatch[1]);
      assert(parsed['@context'] === 'https://schema.org', `${page.file}: invalid JSON-LD @context`);
      assert(Array.isArray(parsed['@graph']) && parsed['@graph'].length >= 3, `${page.file}: JSON-LD @graph incomplete`);
    } catch (e) {
      errors.push(`${page.file}: invalid JSON-LD JSON (${e.message})`);
    }
  }

  // Heading hierarchy & single H1 check
  const headings = [...html.matchAll(/<h([1-6])\b[^>]*>([\s\S]*?)<\/h\1>/gi)].map((m) => ({
    level: Number(m[1]),
    text: m[2].replace(/<[^>]+>/g, '').trim(),
  }));
  const h1Count = headings.filter((h) => h.level === 1).length;
  assert(h1Count === 1, `${page.file}: expected exactly 1 <h1>, found ${h1Count}`);
  for (let i = 1; i < headings.length; i++) {
    const jump = headings[i].level - headings[i - 1].level;
    assert(jump <= 1, `${page.file}: skipped heading level from h${headings[i - 1].level} ("${headings[i - 1].text}") to h${headings[i].level} ("${headings[i].text}")`);
  }

  // Image alt & dimension checks
  const imgs = [...html.matchAll(/<img\b([^>]*)>/gi)];
  for (const img of imgs) {
    const attrs = img[1];
    assert(/\balt="[^"]*"/i.test(attrs), `${page.file}: <img> missing alt attribute (${attrs})`);
    assert(/\bwidth="\d+"/i.test(attrs) && /\bheight="\d+"/i.test(attrs), `${page.file}: <img> missing width/height (${attrs})`);
  }

  // Verify prerendered content inside #root is non-empty
  assert(!/<div id="root"><\/div>/i.test(html), `${page.file}: #root is empty (prerender missing)`);
}

async function validate404() {
  const html = await fs.readFile(path.join(distDir, '404.html'), 'utf-8');
  assert(/<meta\s+name="robots"\s+content="noindex,\s*nofollow"/i.test(html), '404.html: missing noindex, nofollow');
  assert(!/<link\s+rel="canonical"/i.test(html), '404.html: must not declare a canonical link');
  const h1Count = [...html.matchAll(/<h1\b/gi)].length;
  assert(h1Count === 1, `404.html: expected 1 <h1>, found ${h1Count}`);
}

async function validateSitemapAndRobots() {
  const sitemap = await fs.readFile(path.join(distDir, 'sitemap.xml'), 'utf-8');
  assert(sitemap.startsWith('<?xml version="1.0" encoding="UTF-8"?>'), 'sitemap.xml: missing XML declaration');
  assert(sitemap.includes('<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">'), 'sitemap.xml: invalid urlset');

  const locs = [...sitemap.matchAll(/<loc>(.*?)<\/loc>/g)].map((m) => m[1]);
  assert(locs.length === PUBLIC_PAGES.length, `sitemap.xml: expected ${PUBLIC_PAGES.length} URLs, found ${locs.length}`);
  for (const loc of locs) {
    assert(!/\/api\/|\/admin|\/404|\/metrics|\/health/i.test(loc), `sitemap.xml: forbidden URL "${loc}"`);
  }

  const robots = await fs.readFile(path.join(distDir, 'robots.txt'), 'utf-8');
  assert(/User-agent:\s*\*/i.test(robots), 'robots.txt: missing User-agent');
  assert(/Allow:\s*\//i.test(robots), 'robots.txt: missing Allow: /');
  assert(/Sitemap:\s*https?:\/\/.+\/sitemap\.xml/i.test(robots), 'robots.txt: missing Sitemap directive');
}

async function run() {
  for (const page of PUBLIC_PAGES) {
    await validatePage(page);
  }
  await validate404();
  await validateSitemapAndRobots();

  if (errors.length > 0) {
    console.error('SEO validation failed:');
    for (const err of errors) {
      console.error(`  - ${err}`);
    }
    process.exit(1);
  }

  console.log(`SEO validation passed (${PUBLIC_PAGES.length} public pages + 404.html + sitemap.xml + robots.txt).`);
}

run().catch((err) => {
  console.error('SEO validation error:', err);
  process.exit(1);
});
