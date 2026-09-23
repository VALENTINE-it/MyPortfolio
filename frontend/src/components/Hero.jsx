import { useState, useEffect } from 'react'
import './Hero.css'

const TYPED_STRINGS = [
  'Full-Stack Developer',
  'React & JavaScript Specialist',
  'Go Backend Engineer',
  'REST API & SQLite Architect',
  'Passionate Problem Solver',
]

export default function Hero() {
  const [currentText, setCurrentText] = useState('')
  const [stringIndex, setStringIndex] = useState(0)
  const [isDeleting, setIsDeleting] = useState(false)

  // Typing effect loop
  useEffect(() => {
    const fullText = TYPED_STRINGS[stringIndex]
    const speed = isDeleting ? 40 : 80

    const timer = setTimeout(() => {
      if (!isDeleting) {
        // Typing characters
        const nextText = fullText.slice(0, currentText.length + 1)
        setCurrentText(nextText)

        // Finished typing full word
        if (nextText === fullText) {
          setTimeout(() => setIsDeleting(true), 1500)
        }
      } else {
        // Deleting characters
        const nextText = fullText.slice(0, currentText.length - 1)
        setCurrentText(nextText)

        // Finished deleting
        if (nextText === '') {
          setIsDeleting(false)
          setStringIndex((prev) => (prev + 1) % TYPED_STRINGS.length)
        }
      }
    }, speed)

    return () => clearTimeout(timer)
  }, [currentText, isDeleting, stringIndex])

  const scrollTo = (e, sectionId) => {
    e.preventDefault()
    const target = document.getElementById(sectionId)
    if (target) {
      const offset = 65
      const bodyRect = document.body.getBoundingClientRect().top
      const elementRect = target.getBoundingClientRect().top
      const elementPosition = elementRect - bodyRect
      const offsetPosition = elementPosition - offset

      window.scrollTo({
        top: offsetPosition,
        behavior: 'smooth',
      })
      window.history.pushState(null, '', `#${sectionId}`)
    }
  }

  return (
    <div className="hero" id="home">
      <div className="container-fluid">
        <div className="hero-row">
          {/* Left Column: Hero Text */}
          <div className="hero-col-text">
            <div className="hero-content">
              <div className="hero-text">
                <p>I&apos;m</p>
                <h1 className="hero-name">Valentine Omondi Awili</h1>
                <h2 className="hero-typed">
                  <span>{currentText}</span>
                  <span className="typed-cursor">|</span>
                </h2>
              </div>
              <div className="hero-btn">
                <a
                  className="btn"
                  href="#portfolio"
                  onClick={(e) => scrollTo(e, 'portfolio')}
                >
                  Portfolio
                </a>
                <a
                  className="btn"
                  href="#contact"
                  onClick={(e) => scrollTo(e, 'contact')}
                >
                  Contact Me
                </a>
              </div>
              <div className="hero-social">
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

          {/* Right Column: Hero Image */}
          <div className="hero-col-image">
            <div className="hero-image">
              <img
                src="/images/profile.svg"
                alt="Valentine Omondi Awili"
                loading="eager"
              />
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}
