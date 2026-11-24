import { useEffect } from 'react';
import { GSC_VERIFICATION, SITE_INFO, SITE_URL } from '../config/seoConfig';
import { getCanonicalUrl } from '../utils/routing';
import { buildPageJsonLd, PORTFOLIO_PROJECTS } from '../utils/structuredData';

function upsertMeta(attrName, attrValue, content) {
  if (!content) return;
  let el = document.head.querySelector(`meta[${attrName}="${attrValue}"]`);
  if (!el) {
    el = document.createElement('meta');
    el.setAttribute(attrName, attrValue);
    document.head.appendChild(el);
  }
  el.setAttribute('content', content);
}

function upsertLink(rel, href) {
  if (!href) return;
  let el = document.head.querySelector(`link[rel="${rel}"]`);
  if (!el) {
    el = document.createElement('link');
    el.setAttribute('rel', rel);
    document.head.appendChild(el);
  }
  el.setAttribute('href', href);
}

function upsertJsonLd(payload) {
  const existing = document.head.querySelector('script[type="application/ld+json"][data-seo-managed="true"]');
  if (!payload) {
    if (existing) existing.remove();
    return;
  }

  const scriptEl = existing || document.createElement('script');
  scriptEl.type = 'application/ld+json';
  scriptEl.setAttribute('data-seo-managed', 'true');
  scriptEl.textContent = JSON.stringify(payload);
  if (!existing) {
    document.head.appendChild(scriptEl);
  }
}

/**
 * Synchronizes document head metadata, canonical URL, Open Graph, Twitter cards,
 * and Schema.org JSON-LD whenever the active route or projects change.
 */
export function useSEO(routeConfig, projects = PORTFOLIO_PROJECTS) {
  useEffect(() => {
    if (!routeConfig) return;

    const isNotFound = routeConfig.key === 'notFound';
    const canonicalUrl = isNotFound ? '' : getCanonicalUrl(routeConfig.path, SITE_URL);
    const ogImageUrl = `${SITE_URL}${SITE_INFO.defaultOgImage}`;

    document.title = routeConfig.title;

    upsertMeta('name', 'description', routeConfig.description);
    upsertMeta('name', 'robots', routeConfig.robots || 'index, follow');

    if (GSC_VERIFICATION) {
      upsertMeta('name', 'google-site-verification', GSC_VERIFICATION);
    }

    const existingCanonical = document.head.querySelector('link[rel="canonical"]');
    if (isNotFound) {
      if (existingCanonical) existingCanonical.remove();
    } else {
      upsertLink('canonical', canonicalUrl);
    }

    // Open Graph
    upsertMeta('property', 'og:type', routeConfig.ogType || 'website');
    upsertMeta('property', 'og:site_name', `${SITE_INFO.name} Portfolio`);
    upsertMeta('property', 'og:locale', SITE_INFO.locale);
    upsertMeta('property', 'og:title', routeConfig.title);
    upsertMeta('property', 'og:description', routeConfig.description);
    if (canonicalUrl) {
      upsertMeta('property', 'og:url', canonicalUrl);
    }
    upsertMeta('property', 'og:image', ogImageUrl);
    upsertMeta('property', 'og:image:width', '1200');
    upsertMeta('property', 'og:image:height', '630');
    upsertMeta('property', 'og:image:alt', SITE_INFO.defaultOgImageAlt);

    // Twitter / X Card
    upsertMeta('name', 'twitter:card', 'summary_large_image');
    upsertMeta('name', 'twitter:title', routeConfig.title);
    upsertMeta('name', 'twitter:description', routeConfig.description);
    upsertMeta('name', 'twitter:image', ogImageUrl);
    upsertMeta('name', 'twitter:image:alt', SITE_INFO.defaultOgImageAlt);

    // JSON-LD Structured Data
    const jsonLd = buildPageJsonLd(routeConfig, SITE_URL, projects);
    upsertJsonLd(jsonLd);
  }, [routeConfig, projects]);
}
