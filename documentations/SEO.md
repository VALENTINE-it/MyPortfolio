# SEO.md — Search Engine Optimization Architecture & Operations Guide

## 1. SEO Architecture Overview

The portfolio remains a **React 18 + Vite** frontend paired with a **Go REST API + SQLite** backend (`backend/`), deployed via **AWS Amplify** (`amplify.yml`) or **Nginx** (`nginx/nginx.conf`).

To achieve full crawlability and unique per-page metadata without migrating to Next.js or adding runtime dependencies (`react-router-dom`, `react-helmet-async`):

1. **Zero-Dependency Client Routing (`src/utils/routing.js`)**: Uses the browser History API (`pushState` / `popstate`) across five canonical public routes (`/`, `/about`, `/projects`, `/skills`, `/contact`) plus a non-indexed 404 route (`/404`), while keeping `/` as the full continuous editorial portfolio experience.
2. **Dynamic Head Synchronization (`src/hooks/useSEO.js`)**: Updates `<title>`, `<meta name="description">`, `<meta name="robots">`, `<link rel="canonical">`, Open Graph tags, Twitter/X tags, and a single `data-seo-managed="true"` JSON-LD `<script>` in place—preventing duplicate tags during client-side navigation.
3. **Build-Time Static Prerendering (`src/entry-server.jsx` + `scripts/prerender-seo.mjs`)**: Uses React 18's built-in `react-dom/server` (`renderToString`) during `npm run build` to emit static, content-populated HTML files:
   - `dist/index.html` (`/`)
   - `dist/about/index.html` (`/about`)
   - `dist/projects/index.html` (`/projects`)
   - `dist/skills/index.html` (`/skills`)
   - `dist/contact/index.html` (`/contact`)
   - `dist/404.html` (`/404`, `noindex, nofollow`)
4. **Client Hydration (`src/main.jsx`)**: Uses `hydrateRoot` when pre-rendered markup is present in `#root` and `createRoot` during local `vite dev`.

---

## 2. Metadata & Keyword Strategy

All route metadata is centralized in `frontend/src/config/seoConfig.js`:

| Route | Search Intent | `<title>` Focus |
| :--- | :--- | :--- |
| `/` | Primary portfolio & personal brand | Valentine Omondi Awili — Full-Stack Developer in Kenya (React & Go) |
| `/about` | Biography, education (Maseno University), Kisumu & Nairobi | About Valentine Omondi Awili — Software Developer, Kisumu & Nairobi |
| `/projects` | Software engineering work, Web3, IoT, full-stack systems | Software Development Projects — Valentine Omondi Awili |
| `/skills` | Tech stack verification (React, Go, SQLite, Blockchain, Linux) | Technical Skills & Stack — React, Go, SQLite & Blockchain |
| `/contact` | Hiring, freelance inquiries, collaboration | Contact Valentine Omondi Awili — Web Development & Collaboration |
| `/404` | Non-existent URLs | Page Not Found (`noindex, nofollow`, no canonical) |

- Fallback metadata for `/` lives in `frontend/index.html` so even non-prerendered dev environments have valid head tags.
- All descriptions, titles, and keywords reflect verified technologies and experience from the portfolio (`React`, `JavaScript`, `Go`, `REST APIs`, `SQLite`, `Blockchain`, `Nairobi`, `Kisumu`, `Kenya`, `Maseno University`).

---

## 3. Sitemap Generation

- `frontend/scripts/prerender-seo.mjs` generates `dist/sitemap.xml` at build time using `VITE_SITE_URL` (default fallback: `https://valentineawili.dev`) and the build date (`YYYY-MM-DD`) as `<lastmod>`.
- `frontend/public/sitemap.xml` serves as a static fallback for local dev (`vite dev`).
- **Inclusion policy**: Only the 5 canonical, public, indexable URLs (`/`, `/about`, `/projects`, `/skills`, `/contact`) are listed. Private, API, health, metrics, and 404 routes (`/api/*`, `/api/admin/*`, `/metrics`, `/health/*`, `/404`) are strictly excluded.

---

## 4. Robots.txt Configuration

- Generated into `dist/robots.txt` during `npm run build` (with static fallback in `frontend/public/robots.txt`):
  - `User-agent: *`
  - `Allow: /`
  - `Disallow: /api/`
  - `Disallow: /metrics`
  - `Disallow: /health/`
  - `Sitemap: <VITE_SITE_URL>/sitemap.xml`
- **Security note**: `robots.txt` is crawler guidance only and is **never** relied upon for security. Administrative and sensitive backend endpoints remain enforced by Go middleware (`RequireAdmin`, `CSRFProtection`, `CORS`, `IPRateLimiter`) and Nginx access restrictions.

---

## 5. Structured Data (JSON-LD)

`frontend/src/utils/structuredData.js` generates Schema.org `@graph` JSON-LD using only verifiable portfolio facts:
- **`Person`**: Name, job title, Maseno University (`alumniOf`), Nairobi, Kenya (`homeLocation`), `knowsAbout`, and `sameAs` (GitHub & LinkedIn).
- **`WebSite`**: Portfolio site linked to `@id` `#person`.
- **`WebPage` / `ProfilePage` / `CollectionPage` / `ContactPage`**: Route-specific page entity.
- **`BreadcrumbList`**: `Home` → `<Page>` hierarchy on subpages.
- **`ItemList` of `SoftwareSourceCode`**: Rendered on `/` and `/projects` for the portfolio's projects.

---

## 6. Canonical URL Policy

1. Set `VITE_SITE_URL` in your production build environment (e.g., AWS Amplify Environment Variables: `VITE_SITE_URL=https://your-domain.com`, no trailing slash).
2. `normalizePathname()` strips `/index.html` and trailing slashes on non-root paths (`/about/` → `/about`).
3. `nginx/nginx.conf` issues HTTP `301` redirects from `/path/` to `/path` and serves `404.html` with HTTP `404` status for unknown URLs.
4. `404.html` explicitly emits `noindex, nofollow` and omits `<link rel="canonical">`.

---

## 7. Google Search Console Setup (Manual Steps)

Verification tokens must **not** be invented or hardcoded. To connect Google Search Console:

1. **Set Production Site URL**: In AWS Amplify Console → *App settings* → *Environment variables* (or your build environment), set:
   - `VITE_SITE_URL` = your exact HTTPS production URL (e.g. `https://valentineawili.dev` or `https://main.<app-id>.amplifyapp.com`).
2. **Verify Domain Ownership** (choose one):
   - *Option A (Recommended — DNS TXT)*: Add the property as a **Domain** property in Google Search Console and add the provided TXT record at your DNS registrar.
   - *Option B (HTML Meta Tag)*: Copy the verification token from Google Search Console (URL-prefix property) and set `VITE_GSC_VERIFICATION=<your-token>` in your build environment variables, then redeploy. `prerender-seo.mjs` and `useSEO.js` will automatically inject `<meta name="google-site-verification" content="..." />`.
3. **Configure AWS Amplify Rewrite for 404s**:
   - In Amplify Console → *Hosting* → *Rewrites and redirects*, ensure you have a `404 (Rewrite)` rule `/<*>` → `/404.html` (`404`) rather than a `200` rewrite to `/index.html`, so unknown URLs return a true HTTP `404`.
4. **Submit Sitemap**: In Search Console → *Sitemaps*, submit `sitemap.xml` (`https://your-domain.com/sitemap.xml`).
5. **Inspect Key URLs**: Use *URL Inspection* on `/`, `/about`, `/projects`, `/skills`, and `/contact` and request indexing.

---

## 8. Adding SEO Metadata to a New Page

1. Add a new entry to `ROUTE_CONFIGS` and `PUBLIC_ROUTES` in `frontend/src/config/seoConfig.js` with a unique `title`, `description`, `path`, and `breadcrumbName`.
2. Create the page component in `frontend/src/pages/` with `<PageHeader>` (providing exactly one `<h1>` and breadcrumbs).
3. Wire the route case in `frontend/src/App.jsx` and add navigation links in `Navbar.jsx` / `Footer.jsx`.
4. Add the route entry to `ROUTES_TO_PRERENDER` in `frontend/scripts/prerender-seo.mjs` and `PUBLIC_PAGES` in `frontend/scripts/validate-seo.mjs`.
5. Run `npm run build && npm run validate:seo`.

---

## 9. Validating SEO Changes

Run from `frontend/`:

```bash
VITE_SITE_URL=https://valentineawili.dev npm run build
npm run validate:seo
```

`validate-seo.mjs` verifies:
- Unique `<title>` and `<meta name="description">` per public page
- Single `<h1>` and sequential heading hierarchy (`h1` → `h2` → `h3` → `h4`)
- Single accurate `<link rel="canonical">` per public page (and none on `404.html`)
- Valid Open Graph, Twitter Card, and Schema.org JSON-LD blocks
- `<img>` `alt` attributes and explicit `width`/`height`
- Non-empty prerendered HTML inside `#root`
- Valid `dist/sitemap.xml` (no private/404 routes) and `dist/robots.txt`
