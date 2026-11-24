import { SITE_INFO, SITE_URL } from '../config/seoConfig';
import { getCanonicalUrl } from './routing';

/**
 * Factual portfolio projects used for Schema.org CreativeWork / ItemList output
 * and static fallback rendering.
 */
export const PORTFOLIO_PROJECTS = [
  {
    id: 1,
    title: 'Safeguarding Reporting Platform',
    description:
      'A secure, confidential reporting platform designed for anonymous safeguarding disclosures, encrypted submission workflows, and case audit records.',
    category: 'Full-Stack Web Application',
    year: 2026,
    technologies: ['React', 'JavaScript', 'Go', 'SQLite', 'REST APIs', 'Security'],
    image: '/images/projects/safeguarding.png',
    githubUrl: 'https://github.com/VALENTINE-it/safeguarding-project',
    liveUrl: 'https://safeguarding-app-1.onrender.com/',
  },
  {
    id: 2,
    title: 'Nemo',
    description:
      'An IoT security system recognized at the Kijani Space Hackathon, engineered to secure, monitor, and protect fish cages in Lake Victoria through real-time telemetry and intrusion alerts.',
    category: 'IoT & Embedded Security',
    year: 2026,
    technologies: ['IoT', 'Embedded Systems', 'Sensors', 'Telemetry', 'Hardware Security', 'Go'],
    image: '/images/projects/fish-cage-iot.jpeg',
    githubUrl: 'https://github.com/nyabokegrace/nemo',
    liveUrl: '',
  },
  {
    id: 3,
    title: 'Anga Guard',
    description:
      'A decentralized dMRV oracle and SME ESG platform empowering Western Kenya smallholders with verifiable biochar carbon removal credits compliant with Kenya National Carbon Registry.',
    category: 'Web3 & Climate Tech Oracle',
    year: 2026,
    technologies: ['dMRV Oracle', 'Web3', 'React', 'Go', 'IoT Telemetry', 'Smart Contracts', 'ESG Analytics'],
    image: '/images/projects/angaguard.png',
    githubUrl: 'https://github.com/ClayMichael2004/angaguard',
    liveUrl: 'https://angaguard-d96o.onrender.com/',
  },
  {
    id: 4,
    title: 'Career Guidance Platform',
    description:
      'An interactive educational advisory application helping secondary students in Western Kenya discover tertiary courses and career pathways matching their academic strengths.',
    category: 'Full-Stack EdTech System',
    year: 2026,
    technologies: ['React', 'JavaScript', 'Go', 'SQLite', 'REST APIs', 'Analytics'],
    image: '/images/projects/career-guidance.svg',
    githubUrl: 'https://github.com/VALENTINE-it/career-guidance',
    liveUrl: '',
  },
  {
    id: 5,
    title: 'HopeReach Community Platform',
    description:
      'A digital NGO operations and outreach platform delivering meal tracking, housing intake logs, and volunteer mobilization.',
    category: 'Full-Stack Web Application',
    year: 2026,
    technologies: ['React', 'Go', 'SQLite', 'REST APIs', 'UI/UX', 'Community Platforms'],
    image: '/images/hopereach.jpeg',
    githubUrl: 'https://github.com/VALENTINE-it/hope-reach',
    liveUrl: 'https://hope-reach-project-ngo-gdq7.onrender.com/',
  },
  {
    id: 6,
    title: 'Valentine Omondi Developer Portfolio',
    description:
      'A high-performance editorial developer portfolio showcasing architectural systems, real-time telemetry projects, and full-stack engineering achievements with Go and React.',
    category: 'Portfolio & System Showcase',
    year: 2026,
    technologies: ['React', 'Vite', 'JavaScript', 'Go', 'REST APIs', 'SQLite'],
    image: '/images/projects/portfolio.svg',
    githubUrl: 'https://github.com/VALENTINE-it/portfolio',
    liveUrl: '',
  },
];

export function buildPersonSchema(baseUrl = SITE_URL) {
  const canonicalHome = getCanonicalUrl('/', baseUrl);
  return {
    '@type': 'Person',
    '@id': `${canonicalHome}#person`,
    name: SITE_INFO.name,
    url: canonicalHome,
    image: `${baseUrl}/images/hero-profile.jpeg`,
    jobTitle: SITE_INFO.jobTitle,
    description:
      'Full-stack software developer and blockchain technology developer based in Kenya, building web applications with React, JavaScript, Go, REST APIs, and SQLite.',
    alumniOf: {
      '@type': 'CollegeOrUniversity',
      name: 'Maseno University',
    },
    homeLocation: {
      '@type': 'Place',
      address: {
        '@type': 'PostalAddress',
        addressLocality: 'Nairobi',
        addressCountry: 'KE',
      },
    },
    knowsAbout: [
      'React',
      'JavaScript',
      'Go',
      'REST APIs',
      'SQLite',
      'Full-Stack Web Development',
      'Blockchain Technology',
      'Smart Contracts',
      'Computer Networking',
      'Linux',
    ],
    sameAs: [SITE_INFO.githubUrl, SITE_INFO.linkedinUrl],
  };
}

export function buildWebSiteSchema(baseUrl = SITE_URL) {
  const canonicalHome = getCanonicalUrl('/', baseUrl);
  return {
    '@type': 'WebSite',
    '@id': `${canonicalHome}#website`,
    url: canonicalHome,
    name: `${SITE_INFO.name} — Developer Portfolio`,
    inLanguage: 'en',
    publisher: {
      '@id': `${canonicalHome}#person`,
    },
  };
}

export function buildBreadcrumbSchema(routeConfig, baseUrl = SITE_URL) {
  const items = [
    {
      '@type': 'ListItem',
      position: 1,
      name: 'Home',
      item: getCanonicalUrl('/', baseUrl),
    },
  ];

  if (routeConfig.path !== '/' && routeConfig.key !== 'notFound') {
    items.push({
      '@type': 'ListItem',
      position: 2,
      name: routeConfig.breadcrumbName,
      item: getCanonicalUrl(routeConfig.path, baseUrl),
    });
  }

  return {
    '@type': 'BreadcrumbList',
    itemListElement: items,
  };
}

export function buildProjectsItemListSchema(projects = PORTFOLIO_PROJECTS, baseUrl = SITE_URL) {
  const canonicalHome = getCanonicalUrl('/', baseUrl);
  return {
    '@type': 'ItemList',
    name: 'Featured Software Engineering Projects',
    itemListElement: projects.map((project, index) => {
      const work = {
        '@type': 'SoftwareSourceCode',
        name: project.title,
        description: project.description,
        author: { '@id': `${canonicalHome}#person` },
        programmingLanguage: Array.isArray(project.technologies) ? project.technologies : [],
        datePublished: String(project.year || 2026),
      };
      if (project.githubUrl) {
        work.codeRepository = project.githubUrl;
      }
      if (project.liveUrl) {
        work.url = project.liveUrl;
      }
      if (project.image) {
        work.image = `${baseUrl}${project.image}`;
      }
      return {
        '@type': 'ListItem',
        position: index + 1,
        item: work,
      };
    }),
  };
}

/**
 * Generates a complete Schema.org JSON-LD payload for a given route.
 */
export function buildPageJsonLd(routeConfig, baseUrl = SITE_URL, projects = PORTFOLIO_PROJECTS) {
  if (!routeConfig || routeConfig.key === 'notFound') {
    return null;
  }

  const canonicalUrl = getCanonicalUrl(routeConfig.path, baseUrl);
  const canonicalHome = getCanonicalUrl('/', baseUrl);

  const webPage = {
    '@type': routeConfig.key === 'about' ? 'ProfilePage' : routeConfig.key === 'projects' ? 'CollectionPage' : routeConfig.key === 'contact' ? 'ContactPage' : 'WebPage',
    '@id': `${canonicalUrl}#webpage`,
    url: canonicalUrl,
    name: routeConfig.title,
    description: routeConfig.description,
    inLanguage: 'en',
    isPartOf: { '@id': `${canonicalHome}#website` },
    about: { '@id': `${canonicalHome}#person` },
  };

  const graph = [
    buildPersonSchema(baseUrl),
    buildWebSiteSchema(baseUrl),
    webPage,
    buildBreadcrumbSchema(routeConfig, baseUrl),
  ];

  if (routeConfig.key === 'home' || routeConfig.key === 'projects') {
    graph.push(buildProjectsItemListSchema(projects, baseUrl));
  }

  return {
    '@context': 'https://schema.org',
    '@graph': graph,
  };
}
