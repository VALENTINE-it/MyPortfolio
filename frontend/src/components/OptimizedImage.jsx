const KNOWN_WEBP_ASSETS = new Set([
  '/images/hero-profile.webp',
  '/images/about-me.webp',
  '/images/afrinit.webp',
  '/images/hopereach.webp',
  '/images/zone01.webp',
  '/images/projects/angaguard.webp',
  '/images/projects/safeguarding.webp',
  '/images/projects/fish-cage-iot.webp',
]);

const WEBP_ENABLED_EXTENSIONS = /\.(jpe?g|png)$/i;

/**
 * Resolves a verified local .webp sibling path when available.
 */
function getWebpSource(src) {
  if (!src || !WEBP_ENABLED_EXTENSIONS.test(src)) {
    return null;
  }
  const candidate = src.replace(WEBP_ENABLED_EXTENSIONS, '.webp');
  return KNOWN_WEBP_ASSETS.has(candidate) ? candidate : null;
}

/**
 * Renders a <picture> with a WebP <source> (when available) and fallback <img>
 * carrying explicit dimensions and loading hints to prevent CLS and optimize LCP.
 */
export default function OptimizedImage({
  src,
  alt,
  width,
  height,
  loading = 'lazy',
  decoding = 'async',
  fetchPriority,
  className,
  style,
}) {
  const webpSrc = getWebpSource(src);

  const imgProps = {
    src,
    alt,
    width,
    height,
    loading,
    decoding,
    className,
    style,
  };

  if (fetchPriority) {
    imgProps.fetchpriority = fetchPriority;
  }

  if (!webpSrc) {
    return <img {...imgProps} />;
  }

  return (
    <picture style={{ display: 'contents' }}>
      <source srcSet={webpSrc} type="image/webp" />
      <img {...imgProps} />
    </picture>
  );
}
