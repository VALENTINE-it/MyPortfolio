import { useState, useEffect, useCallback } from 'react'
import './PagePagination.css'

const SECTIONS = [
  { id: 'home', number: '01', title: 'Home' },
  { id: 'about', number: '02', title: 'About' },
  { id: 'experience', number: '03', title: 'Accomplishments' },
  { id: 'portfolio', number: '04', title: 'Portfolio' },
  { id: 'skills', number: '05', title: 'Skills' },
  { id: 'contact', number: '06', title: 'Contact' },
]

export default function PagePagination() {
  const [activeSection, setActiveSection] = useState('home')
  const [isVisible, setIsVisible] = useState(false)

  // Track scroll position to reveal pagination once scrolled past top banner
  useEffect(() => {
    const handleScroll = () => {
      setIsVisible(window.scrollY > 80)
    }

    window.addEventListener('scroll', handleScroll, { passive: true })
    handleScroll()
    return () => window.removeEventListener('scroll', handleScroll)
  }, [])

  // Observe active section
  useEffect(() => {
    const observerOptions = {
      root: null,
      rootMargin: '-30% 0px -40% 0px',
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
    SECTIONS.forEach((sec) => {
      const el = document.getElementById(sec.id)
      if (el) observer.observe(el)
    })

    return () => observer.disconnect()
  }, [])

  const currentIndex = SECTIONS.findIndex((s) => s.id === activeSection)
  const currentSection = SECTIONS[currentIndex] || SECTIONS[0]

  const scrollToSection = useCallback((sectionId) => {
    const el = document.getElementById(sectionId)
    if (el) {
      const offset = 65
      const bodyRect = document.body.getBoundingClientRect().top
      const elementRect = el.getBoundingClientRect().top
      const offsetPosition = elementRect - bodyRect - offset

      window.scrollTo({
        top: offsetPosition,
        behavior: 'smooth',
      })
      setActiveSection(sectionId)
      window.history.pushState(null, '', `#${sectionId}`)
    }
  }, [])

  const handlePrevPage = () => {
    if (currentIndex > 0) {
      scrollToSection(SECTIONS[currentIndex - 1].id)
    }
  }

  const handleNextPage = () => {
    if (currentIndex < SECTIONS.length - 1) {
      scrollToSection(SECTIONS[currentIndex + 1].id)
    }
  }

  return (
    <nav
      className={`page-pagination ${isVisible ? 'visible' : ''}`}
      aria-label="Website Page Pagination"
    >
      {/* Previous Page Button */}
      <button
        type="button"
        className="page-pagination-arrow"
        onClick={handlePrevPage}
        disabled={currentIndex === 0}
        aria-label="Previous section page"
        title={currentIndex > 0 ? `Go to ${SECTIONS[currentIndex - 1].title}` : 'Start of page'}
      >
        <i className="fas fa-chevron-up"></i>
      </button>

      {/* Numbered Page Track */}
      <ul className="page-pagination-list">
        {SECTIONS.map((sec, idx) => {
          const isActive = activeSection === sec.id
          return (
            <li key={sec.id} className="page-pagination-item">
              <button
                type="button"
                className={`page-pagination-btn ${isActive ? 'active' : ''}`}
                onClick={() => scrollToSection(sec.id)}
                aria-label={`Page ${sec.number}: ${sec.title}`}
                aria-current={isActive ? 'page' : undefined}
              >
                <span className="page-dot"></span>
                <span className="page-number">{sec.number}</span>
                <span className="page-tooltip">{sec.number} — {sec.title}</span>
              </button>
            </li>
          )
        })}
      </ul>

      {/* Next Page Button */}
      <button
        type="button"
        className="page-pagination-arrow"
        onClick={handleNextPage}
        disabled={currentIndex === SECTIONS.length - 1}
        aria-label="Next section page"
        title={currentIndex < SECTIONS.length - 1 ? `Go to ${SECTIONS[currentIndex + 1].title}` : 'End of page'}
      >
        <i className="fas fa-chevron-down"></i>
      </button>

      {/* Active Page Counter Indicator */}
      <div className="page-pagination-counter">
        <span className="page-counter-current">{currentSection.number}</span>
        <span className="page-counter-divider">/</span>
        <span className="page-counter-total">{String(SECTIONS.length).padStart(2, '0')}</span>
      </div>
    </nav>
  )
}
