/**
 * Centralized SEO & route configuration for Valentine Omondi Awili's portfolio.
 * Set VITE_SITE_URL in the environment to match the deployed canonical origin.
 */

const rawSiteUrl =
  (import.meta.env.VITE_SITE_URL || '').trim() || 'https://valentineawili.dev';

export const SITE_URL = rawSiteUrl.replace(/\/+$/, '');

export const GSC_VERIFICATION =
  (import.meta.env.VITE_GSC_VERIFICATION || '').trim();

export const SITE_INFO = {
  name: 'Valentine Omondi Awili',
  brand: 'Valentine',
  jobTitle: 'Full-Stack Developer',
  email: 'valentineawili@gmail.com',
  location: 'Nairobi & Kisumu, Kenya',
  githubUrl: 'https://github.com/VALENTINE-it',
  linkedinUrl: 'https://www.linkedin.com/in/valentine-omondi-434aa12ab/',
  defaultOgImage: '/images/og-preview.png',
  defaultOgImageAlt: 'Valentine Omondi Awili — Full-Stack Developer Portfolio',
  locale: 'en_US',
};

export const ROUTE_CONFIGS = {
  home: {
    key: 'home',
    path: '/',
    label: 'Home',
    title: 'Valentine Omondi Awili — Full-Stack Developer in Kenya (React & Go)',
    description:
      'Personal developer portfolio of Valentine Omondi Awili, a full-stack software developer in Kenya specializing in React, Go, REST APIs, SQLite, and Web3 systems.',
    robots: 'index, follow',
    ogType: 'website',
    breadcrumbName: 'Home',
  },
  about: {
    key: 'about',
    path: '/about',
    label: 'About',
    title: 'About Valentine Omondi Awili — Software Developer, Kisumu & Nairobi',
    description:
      'Learn about Valentine Omondi Awili, ICT student at Maseno University and full-stack software developer building React, Go, blockchain, and resilient web systems in Kenya.',
    robots: 'index, follow',
    ogType: 'profile',
    breadcrumbName: 'About',
  },
  projects: {
    key: 'projects',
    path: '/projects',
    label: 'Projects',
    title: 'Software Development Projects — Valentine Omondi Awili',
    description:
      'Explore full-stack web applications, IoT security telemetry, and Web3 projects built by Valentine Omondi Awili using React, Go, SQLite, and REST APIs.',
    robots: 'index, follow',
    ogType: 'website',
    breadcrumbName: 'Projects',
  },
  skills: {
    key: 'skills',
    path: '/skills',
    label: 'Skills',
    title: 'Technical Skills & Stack — React, Go, SQLite & Blockchain',
    description:
      'Core engineering skills of Valentine Omondi Awili across React frontend development, Go backend APIs, SQLite databases, networking, Linux, and Web3.',
    robots: 'index, follow',
    ogType: 'website',
    breadcrumbName: 'Skills',
  },
  contact: {
    key: 'contact',
    path: '/contact',
    label: 'Contact',
    title: 'Contact Valentine Omondi Awili — Web Development & Collaboration',
    description:
      'Get in touch with Valentine Omondi Awili for full-stack web application development, software engineering roles, and technical collaborations in Kenya and remotely.',
    robots: 'index, follow',
    ogType: 'website',
    breadcrumbName: 'Contact',
  },
  notFound: {
    key: 'notFound',
    path: '/404',
    label: 'Page Not Found',
    title: 'Page Not Found | Valentine Omondi Awili',
    description:
      'The requested page could not be found on Valentine Omondi Awili’s developer portfolio. Navigate back to Home, Projects, Skills, or Contact.',
    robots: 'noindex, nofollow',
    ogType: 'website',
    breadcrumbName: 'Not Found',
  },
};

export const PUBLIC_ROUTES = [
  ROUTE_CONFIGS.home,
  ROUTE_CONFIGS.about,
  ROUTE_CONFIGS.projects,
  ROUTE_CONFIGS.skills,
  ROUTE_CONFIGS.contact,
];
