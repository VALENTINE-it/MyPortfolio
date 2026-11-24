import { useState, useEffect } from 'react';
import './Footer.css';

const FOOTER_ROUTES = [
  { path: '/', label: 'Home' },
  { path: '/about', label: 'About' },
  { path: '/projects', label: 'Projects' },
  { path: '/skills', label: 'Skills' },
  { path: '/contact', label: 'Contact' },
];

export default function Footer({ onNavigate }) {
  const [showBackToTop, setShowBackToTop] = useState(false);

  useEffect(() => {
    const handleScroll = () => {
      setShowBackToTop(window.scrollY > 250);
    };

    window.addEventListener('scroll', handleScroll, { passive: true });
    return () => window.removeEventListener('scroll', handleScroll);
  }, []);

  const scrollToTop = (e) => {
    e.preventDefault();
    window.scrollTo({ top: 0, behavior: 'smooth' });
  };

  const handleLinkClick = (e, path) => {
    if (typeof onNavigate === 'function') {
      e.preventDefault();
      onNavigate(path);
    }
  };

  return (
    <>
      <footer className="footer">
        <div className="container-fluid">
          <div className="container">
            <div className="footer-info">
              <p className="footer-name">Valentine Omondi Awili</p>
              <p className="footer-role">Full-Stack Software Developer · Nairobi, Kenya</p>

              <div className="footer-menu">
                <p>valentineawili@gmail.com</p>
                <p>React &amp; Go Specialist</p>
                <p>REST API &amp; SQLite Architecture</p>
              </div>

              <nav className="footer-nav" aria-label="Footer Navigation">
                {FOOTER_ROUTES.map((route) => (
                  <a
                    key={route.path}
                    href={route.path}
                    onClick={(e) => handleLinkClick(e, route.path)}
                  >
                    {route.label}
                  </a>
                ))}
              </nav>

              <div className="footer-social">
                <a
                  href="https://github.com/VALENTINE-it"
                  target="_blank"
                  rel="noopener noreferrer"
                  aria-label="GitHub"
                >
                  <i className="fab fa-github" aria-hidden="true"></i>
                </a>
                <a
                  href="https://www.linkedin.com/in/valentine-omondi-434aa12ab/"
                  target="_blank"
                  rel="noopener noreferrer"
                  aria-label="LinkedIn"
                >
                  <i className="fab fa-linkedin-in" aria-hidden="true"></i>
                </a>
                <a
                  href="https://wa.me/254758013777"
                  target="_blank"
                  rel="noopener noreferrer"
                  aria-label="WhatsApp"
                >
                  <i className="fab fa-whatsapp" aria-hidden="true"></i>
                </a>
                <a
                  href="mailto:valentineawili@gmail.com"
                  aria-label="Email"
                >
                  <i className="fas fa-envelope" aria-hidden="true"></i>
                </a>
              </div>
            </div>
          </div>

          <div className="container copyright">
            <p>
              &copy; 2026 <a href="/" onClick={(e) => handleLinkClick(e, '/')}>Valentine Omondi Awili</a>, All Rights Reserved | Portfolio inspired by Syed Ameenulla Hussaini
            </p>
          </div>
        </div>
      </footer>

      {/* Floating Back to Top Button */}
      {showBackToTop && (
        <button
          type="button"
          className="back-to-top"
          onClick={scrollToTop}
          aria-label="Back to top"
        >
          <i className="fa fa-chevron-up" aria-hidden="true"></i>
        </button>
      )}
    </>
  );
}
