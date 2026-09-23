import './About.css'

const PROFICIENCY_BARS = [
  { label: 'Frontend Development (React, JavaScript, Vite, CSS)', percentage: 95 },
  { label: 'Backend Systems (Go, REST APIs, Architecture)', percentage: 90 },
  { label: 'Databases & Storage (SQLite, SQL, Normalization)', percentage: 85 },
  { label: 'Security & DevOps (Linux, Git, Zero-Trust Validation)', percentage: 88 },
]

export default function About() {
  const scrollToContact = (e) => {
    e.preventDefault()
    const target = document.getElementById('contact')
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
      window.history.pushState(null, '', '#contact')
    }
  }

  return (
    <div className="about" id="about">
      <div className="container-fluid">
        <div className="about-row">
          {/* Left Column: Image */}
          <div className="about-col-img">
            <div className="about-img">
              <img
                src="/images/about.svg"
                alt="Valentine Omondi Awili Workspace"
                loading="lazy"
              />
            </div>
          </div>

          {/* Right Column: Narrative Content & Skill Bars */}
          <div className="about-col-content">
            <div className="about-content">
              <div className="section-header text-left">
                <p>Learn About Me</p>
                <h2>Full-Stack Developer</h2>
              </div>

              <div className="about-text">
                <p>
                  I am <strong>Valentine Omondi Awili</strong>, a dedicated full-stack software engineer
                  passionate about engineering reliable, scalable, and responsive digital systems. My approach
                  combines clean architectural discipline with modern, intuitive user interfaces.
                </p>
                <p>
                  My engineering journey spans developing concurrent, high-throughput backend services
                  in <strong>Go</strong> to crafting expressive client interfaces with <strong>React</strong>.
                  With a solid grounding in computer science principles, algorithms, and secure systems design,
                  I focus on writing code that is performant, maintainable, and built to last.
                </p>
                <p>
                  Whether architecting RESTful microservices, designing persistent database schemas in SQLite,
                  or refining user interactions, I strive for purposeful craftsmanship over unnecessary bloat.
                </p>
              </div>

              {/* Skill Proficiency Progress Bars */}
              <div className="skills">
                {PROFICIENCY_BARS.map((item) => (
                  <div key={item.label} className="skill-wrapper">
                    <div className="skill-name">
                      <p>{item.label}</p>
                      <p>{item.percentage}%</p>
                    </div>
                    <div className="progress">
                      <div
                        className="progress-bar"
                        role="progressbar"
                        style={{ width: `${item.percentage}%` }}
                        aria-valuenow={item.percentage}
                        aria-valuemin="0"
                        aria-valuemax="100"
                      ></div>
                    </div>
                  </div>
                ))}
              </div>

              <a className="btn" href="#contact" onClick={scrollToContact}>
                Get In Touch
              </a>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}
