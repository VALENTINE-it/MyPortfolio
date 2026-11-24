import { ROUTE_CONFIGS, SITE_URL } from '../config/seoConfig';

/**
 * Normalizes a URL pathname by collapsing trailing slashes (except root '/')
 * and stripping '/index.html'.
 */
export function normalizePathname(pathname = '/') {
  if (!pathname) return '/';
  let clean = pathname.split('?')[0].split('#')[0];
  if (clean === '/index.html') return '/';
  if (clean.endsWith('/index.html')) {
    clean = clean.slice(0, -'/index.html'.length) || '/';
  }
  if (clean.length > 1 && clean.endsWith('/')) {
    clean = clean.slice(0, -1);
  }
  return clean || '/';
}

/**
 * Resolves a pathname into one of the known portfolio route configurations.
 */
export function getRouteConfig(pathname = '/') {
  const normalized = normalizePathname(pathname);
  const match = Object.values(ROUTE_CONFIGS).find(
    (route) => route.key !== 'notFound' && route.path === normalized
  );
  return match || ROUTE_CONFIGS.notFound;
}

/**
 * Builds an absolute canonical URL for a route path.
 */
export function getCanonicalUrl(path, baseUrl = SITE_URL) {
  const cleanBase = baseUrl.replace(/\/+$/, '');
  const normalized = normalizePathname(path);
  return normalized === '/' ? `${cleanBase}/` : `${cleanBase}${normalized}`;
}

/**
 * Performs client-side History API navigation without full page reloads.
 */
export function navigateTo(targetPath, onRouteChange) {
  if (typeof window === 'undefined') return;

  const [rawPath, hash] = targetPath.split('#');
  const nextPath = normalizePathname(rawPath || '/');
  const currentPath = normalizePathname(window.location.pathname);

  if (nextPath !== currentPath) {
    const nextUrl = hash ? `${nextPath}#${hash}` : nextPath;
    window.history.pushState(null, '', nextUrl);
    if (typeof onRouteChange === 'function') {
      onRouteChange(nextPath);
    }
  }

  if (hash) {
    window.requestAnimationFrame(() => {
      const target = document.getElementById(hash);
      if (target) {
        const offset = 65;
        const bodyRect = document.body.getBoundingClientRect().top;
        const elementRect = target.getBoundingClientRect().top;
        window.scrollTo({
          top: elementRect - bodyRect - offset,
          behavior: 'smooth',
        });
      }
    });
  } else {
    window.scrollTo({ top: 0, behavior: 'smooth' });
  }
}
