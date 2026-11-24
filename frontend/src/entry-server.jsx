import { renderToString } from 'react-dom/server';
import App from './App.jsx';
import { getRouteConfig, getCanonicalUrl } from './utils/routing';
import { buildPageJsonLd } from './utils/structuredData';
import { SITE_INFO, SITE_URL } from './config/seoConfig';

/**
 * Renders the application HTML and route-specific SEO head tags for build-time prerendering.
 */
export function renderRoute(urlPath, siteUrl = SITE_URL) {
  const routeConfig = getRouteConfig(urlPath);
  const html = renderToString(<App initialPath={urlPath} />);
  const canonicalUrl =
    routeConfig.key === 'notFound' ? '' : getCanonicalUrl(routeConfig.path, siteUrl);
  const jsonLd = buildPageJsonLd(routeConfig, siteUrl);
  const ogImageUrl = `${siteUrl}${SITE_INFO.defaultOgImage}`;

  return {
    html,
    routeConfig,
    canonicalUrl,
    ogImageUrl,
    jsonLd,
  };
}
