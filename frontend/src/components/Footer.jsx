import './Footer.css'

export default function Footer() {
  const scrollToTop = (e) => {
    e.preventDefault()
    window.scrollTo({ top: 0, behavior: 'smooth' })
  }

  return (
    <footer className="footer">
      <div className="container">
        <div className="footer-top">
          {/* Developer Profile Info */}
          <div className="footer-brand">
            <h2 className="footer-name">VALENTINE OMONDI AWILI</h2>
            <p className="footer-title">Full-Stack Developer</p>
            <p className="footer-bio">
              Building practical, high-performance web applications with React, JavaScript, Go, and SQLite.
            </p>
          </div>

          {/* Social and Connect Links */}
          <div className="footer-nav">
            <span className="footer-heading">Connect</span>
            <a
              href="https://github.com/VALENTINE-it"
              target="_blank"
              rel="noopener noreferrer"
              className="footer-link"
              aria-label="GitHub profile"
            >
              GitHub ↗
            </a>
            <a
              href="https://linkedin.com/in/valentine-awili"
              target="_blank"
              rel="noopener noreferrer"
              className="footer-link"
              aria-label="LinkedIn profile"
            >
              LinkedIn ↗
            </a>
            <a
              href="mailto:valentineawili@gmail.com"
              className="footer-link"
              aria-label="Send email to Valentine"
            >
              Email ↗
            </a>
          </div>
        </div>

        {/* Bottom Metadata & Copyright */}
        <div className="footer-bottom">
          <p className="footer-copyright">
            © {new Date().getFullYear()} Valentine Omondi Awili. All rights reserved.
          </p>
          <a
            href="#home"
            className="footer-back-to-top"
            onClick={scrollToTop}
          >
            Back to Top ↑
          </a>
        </div>
      </div>
    </footer>
  )
}
