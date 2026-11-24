import fs from 'node:fs/promises';
import path from 'node:path';
import { fileURLToPath, pathToFileURL } from 'node:url';

const __dirname = path.dirname(fileURLToPath(import.meta.url));
const frontendDir = path.resolve(__dirname, '..');
const distDir = path.join(frontendDir, 'dist');
const ssrDir = path.join(frontendDir, 'dist-ssr');

const rawSiteUrl = process.env.VITE_SITE_URL || 'https://valentineawili.dev';
const SITE_URL = rawSiteUrl.replace(/\/+$/, '');
const GSC_VERIFICATION = process.env.VITE_GSC_VERIFICATION || '';
const BUILD_DATE = new Date().toISOString().slice(0, 10);

const ROUTES_TO_PRERENDER = [
  { routePath: '/', outputFile: 'index.html', changefreq: 'monthly', priority: '1.0' },
  { routePath: '/about', outputFile: path.join('about', 'index.html'), changefreq: 'monthly', priority: '0.8' },
  { routePath: '/projects', outputFile: path.join('projects', 'index.html'), changefreq: 'weekly', priority: '0.9' },
  { routePath: '/skills', outputFile: path.join('skills', 'index.html'), changefreq: 'monthly', priority: '0.8' },
  { routePath: '/contact', outputFile: path.join('contact', 'index.html'), changefreq: 'yearly', priority: '0.7' },
  { routePath: '/404', outputFile: '404.html', indexable: false },
];

function escapeAttr(value) {
  return String(value)
    .replace(/&/g, '&amp;')
    .replace(/"/g, '&quot;')
    .replace(/</g, '&lt;')
    .replace(/>/g, '&gt;');
}

function escapeHtml(value) {
  return String(value)
    .replace(/&/g, '&amp;')
    .replace(/</g, '&lt;')
    .replace(/>/g, '&gt;');
}

function buildHeadTags({ routeConfig, canonicalUrl, ogImageUrl, jsonLd }) {
  const tags = [
    `<title>${escapeHtml(routeConfig.title)}</title>`,
    `<meta name="description" content="${escapeAttr(routeConfig.description)}" />`,
    `<meta name="robots" content="${escapeAttr(routeConfig.robots || 'index, follow')}" />`,
    `<meta name="author" content="Valentine Omondi Awili" />`,
    `<meta name="theme-color" content="#EF233C" />`,
  ];

  if (GSC_VERIFICATION) {
    tags.push(`<meta name="google-site-verification" content="${escapeAttr(GSC_VERIFICATION)}" />`);
  }

  if (canonicalUrl) {
    tags.push(`<link rel="canonical" href="${escapeAttr(canonicalUrl)}" />`);
  }

  tags.push(
    `<meta property="og:type" content="${escapeAttr(routeConfig.ogType || 'website')}" />`,
    `<meta property="og:site_name" content="Valentine Omondi Awili Portfolio" />`,
    `<meta property="og:locale" content="en_US" />`,
    `<meta property="og:title" content="${escapeAttr(routeConfig.title)}" />`,
    `<meta property="og:description" content="${escapeAttr(routeConfig.description)}" />`,
    `<meta property="og:image" content="${escapeAttr(ogImageUrl)}" />`,
    `<meta property="og:image:width" content="1200" />`,
    `<meta property="og:image:height" content="630" />`,
    `<meta property="og:image:alt" content="Valentine Omondi Awili — Full-Stack Developer Portfolio" />`,
    `<meta name="twitter:card" content="summary_large_image" />`,
    `<meta name="twitter:title" content="${escapeAttr(routeConfig.title)}" />`,
    `<meta name="twitter:description" content="${escapeAttr(routeConfig.description)}" />`,
    `<meta name="twitter:image" content="${escapeAttr(ogImageUrl)}" />`,
    `<meta name="twitter:image:alt" content="Valentine Omondi Awili — Full-Stack Developer Portfolio" />`
  );

  if (canonicalUrl) {
    tags.push(`<meta property="og:url" content="${escapeAttr(canonicalUrl)}" />`);
  }

  if (jsonLd) {
    tags.push(
      `<script type="application/ld+json" data-seo-managed="true">${JSON.stringify(jsonLd)}</script>`
    );
  }

  return tags.join('\n    ');
}

function injectIntoTemplate(template, rendered) {
  // Strip fallback SEO meta/title/canonical/og/twitter from template before injecting route-specific tags
  let cleaned = template
    .replace(/<title>[\s\S]*?<\/title>\s*/i, '')
    .replace(/<meta\s+name="(?:description|robots|author|theme-color|google-site-verification|twitter:[^"]+)"[^>]*>\s*/gi, '')
    .replace(/<meta\s+property="og:[^"]+"[^>]*>\s*/gi, '')
    .replace(/<link\s+rel="canonical"[^>]*>\s*/gi, '')
    .replace(/<script\s+type="application\/ld\+json"[^>]*>[\s\S]*?<\/script>\s*/gi, '');

  const seoHeadBlock = buildHeadTags(rendered);
  cleaned = cleaned.replace(
    /(<meta\s+name="viewport"[^>]*>)/i,
    `$1\n    ${seoHeadBlock}`
  );

  return cleaned.replace(
    '<div id="root"></div>',
    `<div id="root">${rendered.html}</div>`
  );
}

async function run() {
  const templatePath = path.join(distDir, 'index.html');
  const template = await fs.readFile(templatePath, 'utf-8');

  const ssrEntryUrl = pathToFileURL(path.join(ssrDir, 'entry-server.js')).href;
  const { renderRoute } = await import(ssrEntryUrl);

  const sitemapEntries = [];

  for (const route of ROUTES_TO_PRERENDER) {
    const rendered = renderRoute(route.routePath, SITE_URL);
    const finalHtml = injectIntoTemplate(template, rendered);
    const outPath = path.join(distDir, route.outputFile);

    await fs.mkdir(path.dirname(outPath), { recursive: true });
    await fs.writeFile(outPath, finalHtml, 'utf-8');
    console.log(`[prerender] Wrote ${route.outputFile} (${route.routePath})`);

    if (route.indexable !== false && rendered.canonicalUrl) {
      sitemapEntries.push({
        loc: rendered.canonicalUrl,
        lastmod: BUILD_DATE,
        changefreq: route.changefreq,
        priority: route.priority,
      });
    }
  }

  // Generate sitemap.xml in dist/
  const sitemapXml = [
    '<?xml version="1.0" encoding="UTF-8"?>',
    '<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">',
    ...sitemapEntries.map(
      (entry) => `  <url>
    <loc>${escapeHtml(entry.loc)}</loc>
    <lastmod>${entry.lastmod}</lastmod>
    <changefreq>${entry.changefreq}</changefreq>
    <priority>${entry.priority}</priority>
  </url>`
    ),
    '</urlset>',
    '',
  ].join('\n');

  await fs.writeFile(path.join(distDir, 'sitemap.xml'), sitemapXml, 'utf-8');
  console.log(`[prerender] Wrote sitemap.xml (${sitemapEntries.length} URLs)`);

  // Generate robots.txt in dist/
  const robotsTxt = [
    'User-agent: *',
    'Allow: /',
    'Disallow: /api/',
    'Disallow: /metrics',
    'Disallow: /health/',
    '',
    `Sitemap: ${SITE_URL}/sitemap.xml`,
    '',
  ].join('\n');

  await fs.writeFile(path.join(distDir, 'robots.txt'), robotsTxt, 'utf-8');
  console.log('[prerender] Wrote robots.txt');

  // Remove temporary SSR bundle
  await fs.rm(ssrDir, { recursive: true, force: true });
}

run().catch((err) => {
  console.error('[prerender] Failed:', err);
  process.exit(1);
});
