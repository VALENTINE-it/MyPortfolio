import './Hero.css'

export default function Hero() {
  const scrollTo = (e, sectionId) => {
    e.preventDefault()
    const target = document.getElementById(sectionId)
    if (target) {
      target.scrollIntoView({ behavior: 'smooth' })
      window.history.pushState(null, '', `#${sectionId}`)
    }
  }

  return (
    <section id="home" className="hero-section container">
      <div className="hero-grid">
        {/* Left Column: Hero Text Content */}
        <div className="hero-content">
          <span className="hero-pretitle animate-fade-up">I&apos;m</span>

          <h1 className="hero-name animate-fade-up animate-delay-1">
            <span>VALENTINE</span>
            <span>OMONDI AWILI</span>
          </h1>

          <div className="hero-title animate-fade-up animate-delay-2">
            Full-Stack Developer
          </div>

          <p className="hero-description animate-fade-up animate-delay-2">
            I build modern, responsive and practical web applications using
            React, JavaScript and Go. Driven by performance, clean architecture, and
            editorial user experiences.
          </p>

          {/* Action CTAs */}
          <div className="hero-actions animate-fade-up animate-delay-3">
            <a
              href="#projects"
              className="btn btn-primary"
              onClick={(e) => scrollTo(e, 'projects')}
            >
              View My Work
            </a>
            <a
              href="#contact"
              className="btn btn-outline"
              onClick={(e) => scrollTo(e, 'contact')}
            >
              Contact Me
            </a>
          </div>

          {/* Social Links */}
          <div className="hero-socials animate-fade-up animate-delay-3">
            <a
              href="https://github.com/VALENTINE-it"
              target="_blank"
              rel="noopener noreferrer"
              className="hero-social-link"
              aria-label="GitHub profile"
            >
              GitHub ↗
            </a>
            <a
              href="https://linkedin.com/in/valentine-awili"
              target="_blank"
              rel="noopener noreferrer"
              className="hero-social-link"
              aria-label="LinkedIn profile"
            >
              LinkedIn ↗
            </a>
            <a
              href="mailto:valentineawili@gmail.com"
              className="hero-social-link"
              aria-label="Email Valentine"
            >
              Email ↗
            </a>
          </div>
        </div>

        {/* Right Column: Editorial Profile Image */}
        <div className="hero-image-wrapper animate-image-reveal">
          <div className="hero-image-frame">
            <img
              src="/images/profile.svg"
              alt="Valentine Omondi Awili portrait"
              loading="eager"
              width="800"
              height="1000"
            />
          </div>
        </div>
      </div>

      {/* Editorial Section Transition Indicator */}
      <div>
        <a
          href="#about"
          className="hero-scroll-indicator"
          onClick={(e) => scrollTo(e, 'about')}
          aria-label="Scroll down to About section"
        >
          <span>Learn About Me</span>
          <span className="scroll-arrow" aria-hidden="true">↓</span>
        </a>
      </div>
    </section>
  )
}
