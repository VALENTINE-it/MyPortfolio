import { useState, useEffect } from 'react'
import './Footer.css'

export default function Footer() {
  const [showBackToTop, setShowBackToTop] = useState(false)

  useEffect(() => {
    const handleScroll = () => {
      setShowBackToTop(window.scrollY > 250)
    }

    window.addEventListener('scroll', handleScroll, { passive: true })
    return () => window.removeEventListener('scroll', handleScroll)
  }, [])

  const scrollToTop = (e) => {
    e.preventDefault()
    window.scrollTo({ top: 0, behavior: 'smooth' })
  }

  return (
    <>
      <footer className="footer">
        <div className="container-fluid">
          <div className="container">
            <div className="footer-info">
              <h2>Valentine Omondi Awili</h2>
              <h3>Full-Stack Software Developer · Kisumu, Kenya</h3>

              <div className="footer-menu">
                <p>valentineawili@gmail.com</p>
                <p>React &amp; Go Specialist</p>
                <p>REST API &amp; SQLite Architecture</p>
              </div>

              <div className="footer-social">
                <a
                  href="https://github.com/VALENTINE-it"
                  target="_blank"
                  rel="noopener noreferrer"
                  aria-label="GitHub"
                >
                  <i className="fab fa-github"></i>
                </a>
                <a
                  href="https://www.linkedin.com/in/valentine-omondi-434aa12ab/"
                  target="_blank"
                  rel="noopener noreferrer"
                  aria-label="LinkedIn"
                >
                  <i className="fab fa-linkedin-in"></i>
                </a>
                <a
                  href="mailto:valentineawili@gmail.com"
                  aria-label="Email"
                >
                  <i className="fas fa-envelope"></i>
                </a>
              </div>
            </div>
          </div>

          <div className="container copyright">
            <p>
              &copy; {new Date().getFullYear()} <a href="#home" onClick={scrollToTop}>Valentine Omondi Awili</a>, All Rights Reserved | Portfolio inspired by Syed Ameenulla Hussaini
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
          <i className="fa fa-chevron-up"></i>
        </button>
      )}
    </>
  )
}
