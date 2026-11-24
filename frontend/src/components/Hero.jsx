import { useState, useEffect } from 'react';
import OptimizedImage from './OptimizedImage';
import './Hero.css';

const TYPED_STRINGS = [
  'Full-Stack Developer',
  'Blockchain Technology Developer',
  'React & JavaScript Specialist',
  'Go Backend Engineer',
  'REST API & SQLite Architect',
  'Passionate Problem Solver',
];

export default function Hero() {
  const [currentText, setCurrentText] = useState(TYPED_STRINGS[0]);
  const [stringIndex, setStringIndex] = useState(0);
  const [isDeleting, setIsDeleting] = useState(false);
  const [isPaused, setIsPaused] = useState(true);

  // Typing effect loop initialized with full primary role for immediate LCP & SEO text
  useEffect(() => {
    const fullText = TYPED_STRINGS[stringIndex];

    if (isPaused) {
      const pauseTimer = setTimeout(() => {
        setIsPaused(false);
        setIsDeleting(true);
      }, 1800);
      return () => clearTimeout(pauseTimer);
    }

    const speed = isDeleting ? 40 : 80;

    const timer = setTimeout(() => {
      if (!isDeleting) {
        const nextText = fullText.slice(0, currentText.length + 1);
        setCurrentText(nextText);

        if (nextText === fullText) {
          setIsPaused(true);
        }
      } else {
        const nextText = fullText.slice(0, currentText.length - 1);
        setCurrentText(nextText);

        if (nextText === '') {
          setIsDeleting(false);
          setStringIndex((prev) => (prev + 1) % TYPED_STRINGS.length);
        }
      }
    }, speed);

    return () => clearTimeout(timer);
  }, [currentText, isDeleting, isPaused, stringIndex]);

  const scrollTo = (e, sectionId) => {
    e.preventDefault();
    const target = document.getElementById(sectionId);
    if (target) {
      const offset = 65;
      const bodyRect = document.body.getBoundingClientRect().top;
      const elementRect = target.getBoundingClientRect().top;
      const elementPosition = elementRect - bodyRect;
      const offsetPosition = elementPosition - offset;

      window.scrollTo({
        top: offsetPosition,
        behavior: 'smooth',
      });
      window.history.pushState(null, '', `#${sectionId}`);
    }
  };

  return (
    <section className="hero" id="home" aria-label="Introduction">
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
                  <span className="typed-cursor" aria-hidden="true">|</span>
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

          {/* Right Column: Hero Image */}
          <div className="hero-col-image">
            <div className="hero-image">
              <OptimizedImage
                src="/images/hero-profile.jpeg"
                alt="Valentine Omondi Awili, Full-Stack Developer"
                width={420}
                height={540}
                loading="eager"
                decoding="async"
                fetchPriority="high"
              />
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}
