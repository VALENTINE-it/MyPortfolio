import { useState, useEffect } from 'react'
import './Navbar.css'

const NAV_ITEMS = [
  { id: 'home', label: 'HOME' },
  { id: 'about', label: 'ABOUT' },
  { id: 'projects', label: 'PROJECTS' },
  { id: 'skills', label: 'SKILLS' },
  { id: 'contact', label: 'CONTACT' },
]

export default function Navbar() {
  const [activeSection, setActiveSection] = useState('home')
  const [isScrolled, setIsScrolled] = useState(false)
  const [mobileMenuOpen, setMobileMenuOpen] = useState(false)

  // Track scroll position for navbar styling & active section
  useEffect(() => {
    const handleScroll = () => {
      setIsScrolled(window.scrollY > 40)
    }

    window.addEventListener('scroll', handleScroll, { passive: true })
    return () => window.removeEventListener('scroll', handleScroll)
  }, [])

  // Observe active sections
  useEffect(() => {
    const observerOptions = {
      root: null,
      rootMargin: '-20% 0px -60% 0px',
      threshold: 0,
    }

    const observerCallback = (entries) => {
      entries.forEach((entry) => {
        if (entry.isIntersecting) {
          setActiveSection(entry.target.id)
        }
      })
    }

    const observer = new IntersectionObserver(observerCallback, observerOptions)
    NAV_ITEMS.forEach((item) => {
      const el = document.getElementById(item.id)
      if (el) observer.observe(el)
    })

    return () => observer.disconnect()
  }, [])

  // Close mobile menu on Escape key
  useEffect(() => {
    const handleKeyDown = (e) => {
      if (e.key === 'Escape' && mobileMenuOpen) {
        setMobileMenuOpen(false)
      }
    }
    window.addEventListener('keydown', handleKeyDown)
    return () => window.removeEventListener('keydown', handleKeyDown)
  }, [mobileMenuOpen])

  const scrollToSection = (e, sectionId) => {
    e.preventDefault()
    setMobileMenuOpen(false)

    const el = document.getElementById(sectionId)
    if (el) {
      el.scrollIntoView({ behavior: 'smooth' })
      setActiveSection(sectionId)
      window.history.pushState(null, '', `#${sectionId}`)
    }
  }

  return (
    <header className={`navbar ${isScrolled ? 'navbar-scrolled' : ''}`}>
      <div className="container navbar-container">
        {/* Brand Logo */}
        <a
          href="#home"
          className="navbar-logo"
          onClick={(e) => scrollToSection(e, 'home')}
          aria-label="Valentine Omondi Awili Portfolio Home"
        >
          VALENTINE
        </a>

        {/* Desktop Navigation */}
        <nav className="navbar-links" aria-label="Main Navigation">
          {NAV_ITEMS.map((item) => (
            <a
              key={item.id}
              href={`#${item.id}`}
              className={`nav-link ${activeSection === item.id ? 'active' : ''}`}
              onClick={(e) => scrollToSection(e, item.id)}
            >
              {item.label}
            </a>
          ))}
        </nav>

        {/* Mobile Toggle Button */}
        <button
          type="button"
          className="mobile-menu-btn"
          onClick={() => setMobileMenuOpen((prev) => !prev)}
          aria-expanded={mobileMenuOpen}
          aria-label="Toggle navigation menu"
        >
          {mobileMenuOpen ? 'CLOSE' : 'MENU'}
        </button>
      </div>

      {/* Mobile Menu Dropdown */}
      <nav
        className={`mobile-nav ${mobileMenuOpen ? 'open' : ''}`}
        aria-label="Mobile Navigation"
      >
        {NAV_ITEMS.map((item) => (
          <a
            key={item.id}
            href={`#${item.id}`}
            className={`nav-link ${activeSection === item.id ? 'active' : ''}`}
            onClick={(e) => scrollToSection(e, item.id)}
          >
            {item.label}
          </a>
        ))}
      </nav>
    </header>
  )
}
