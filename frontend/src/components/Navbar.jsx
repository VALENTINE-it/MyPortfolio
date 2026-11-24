import { useState, useEffect } from 'react';
import './Navbar.css';

const NAV_ITEMS = [
  { path: '/', sectionId: 'home', label: 'Home' },
  { path: '/about', sectionId: 'about', label: 'About' },
  { path: '/#experience', sectionId: 'experience', label: 'Accomplishments' },
  { path: '/projects', sectionId: 'portfolio', label: 'Portfolio' },
  { path: '/skills', sectionId: 'skills', label: 'Skills' },
  { path: '/contact', sectionId: 'contact', label: 'Contact' },
];

export default function Navbar({ currentPath = '/', onNavigate }) {
  const [activeSection, setActiveSection] = useState('home');
  const [isSticky, setIsSticky] = useState(false);
  const [mobileMenuOpen, setMobileMenuOpen] = useState(false);

  // Track scroll position for sticky navbar
  useEffect(() => {
    const handleScroll = () => {
      setIsSticky(window.scrollY > 40);
    };

    window.addEventListener('scroll', handleScroll, { passive: true });
    handleScroll();
    return () => window.removeEventListener('scroll', handleScroll);
  }, []);

  // Observe active sections for nav highlighting on Home
  useEffect(() => {
    if (currentPath !== '/') return;

    const observerOptions = {
      root: null,
      rootMargin: '-20% 0px -50% 0px',
      threshold: 0,
    };

    const observerCallback = (entries) => {
      entries.forEach((entry) => {
        if (entry.isIntersecting) {
          setActiveSection(entry.target.id);
        }
      });
    };

    const observer = new IntersectionObserver(observerCallback, observerOptions);
    NAV_ITEMS.forEach((item) => {
      const el = document.getElementById(item.sectionId);
      if (el) observer.observe(el);
    });

    return () => observer.disconnect();
  }, [currentPath]);

  const handleLinkClick = (e, item) => {
    setMobileMenuOpen(false);
    if (typeof onNavigate === 'function') {
      e.preventDefault();
      onNavigate(item.path);
    }
  };

  const isNavItemActive = (item) => {
    if (currentPath === '/') {
      return activeSection === item.sectionId;
    }
    return currentPath === item.path;
  };

  return (
    <nav
      className={`navbar ${isSticky || currentPath !== '/' ? 'nav-sticky' : ''}`}
      aria-label="Primary Navigation"
    >
      <div className="container-fluid navbar-inner">
        <a
          href="/"
          className="navbar-brand"
          onClick={(e) => handleLinkClick(e, NAV_ITEMS[0])}
          aria-label="Valentine Omondi Awili Home"
        >
          Valentine
        </a>

        {/* Mobile Toggler */}
        <button
          type="button"
          className="navbar-toggler"
          onClick={() => setMobileMenuOpen((prev) => !prev)}
          aria-expanded={mobileMenuOpen}
          aria-label="Toggle navigation"
        >
          <i className={mobileMenuOpen ? 'fas fa-times' : 'fas fa-bars'} aria-hidden="true"></i>
        </button>

        {/* Nav Links */}
        <div className={`navbar-collapse ${mobileMenuOpen ? 'show' : ''}`}>
          <div className="navbar-nav">
            {NAV_ITEMS.map((item) => {
              const active = isNavItemActive(item);
              return (
                <a
                  key={item.path}
                  href={item.path}
                  className={`nav-link ${active ? 'active' : ''}`}
                  aria-current={currentPath === item.path ? 'page' : undefined}
                  onClick={(e) => handleLinkClick(e, item)}
                >
                  {item.label}
                </a>
              );
            })}
          </div>
        </div>
      </div>
    </nav>
  );
}
