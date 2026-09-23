import './About.css'

const WHAT_I_DO_ITEMS = [
  {
    number: '01',
    title: 'Full-Stack Web Development',
    description: 'Engineering responsive, modern web applications from frontend interfaces in React to high-performance Go backends.',
  },
  {
    number: '02',
    title: 'REST API Design & Architecture',
    description: 'Designing clean, modular, and maintainable RESTful services following standard HTTP conventions and robust validation.',
  },
  {
    number: '03',
    title: 'Database Architecture & SQLite',
    description: 'Constructing normalized database schemas, parameterized SQL queries, indexing, and persistent storage pipelines.',
  },
  {
    number: '04',
    title: 'Security-First Engineering',
    description: 'Implementing defense-in-depth principles, strict input sanitization, rate limiting, and zero-trust backend boundaries.',
  },
  {
    number: '05',
    title: 'Performance & Editorial Interfaces',
    description: 'Crafting lightweight, accessible user experiences using fluid typography, semantic markup, and zero unnecessary bloat.',
  },
]

export default function About() {
  return (
    <section id="about" className="section container about-section">
      <div className="about-header">
        <span className="label-editorial">Biography &amp; Focus</span>
        <h2 className="heading-editorial">ABOUT ME</h2>
      </div>

      <div className="about-grid">
        {/* Left Column: Large Editorial Image */}
        <div className="about-image-column">
          <div className="about-image-frame">
            <img
              src="/images/about.svg"
              alt="Valentine Omondi Awili workspace and engineering concept"
              loading="lazy"
              width="900"
              height="1100"
            />
          </div>
        </div>

        {/* Right Column: Editorial Narrative & What I Do */}
        <div className="about-story-column">
          {/* Who I Am */}
          <div className="about-block">
            <h3 className="about-subheading">Who I Am</h3>
            <p className="about-paragraph">
              I am <strong>Valentine Omondi Awili</strong>, a dedicated full-stack software developer
              focused on building robust, dependable, and high-quality web applications. I approach
              software engineering with an emphasis on clarity, reliability, and thoughtful digital
              craftsmanship.
            </p>
          </div>

          {/* Education */}
          <div className="about-block">
            <h3 className="about-subheading">Education</h3>
            <p className="about-paragraph">
              With a solid educational foundation in computer science and technology disciplines, my
              studies provided deep grounding in computer systems, algorithms, networking architecture,
              and software engineering methodology.
            </p>
          </div>

          {/* Development Journey */}
          <div className="about-block">
            <h3 className="about-subheading">Development Journey</h3>
            <p className="about-paragraph">
              My engineering journey began with exploring how computing platforms communicate and scale.
              Over time, I expanded from foundational web development into backend systems programming,
              specializing in <strong>Go</strong> for concurrent backend services and <strong>React</strong> for
              expressive client-side interfaces.
            </p>
          </div>

          {/* Current Focus */}
          <div className="about-block">
            <h3 className="about-subheading">Current Professional Focus</h3>
            <p className="about-paragraph">
              Presently, I am focused on architecting resilient full-stack systems, clean RESTful APIs,
              lightweight relational persistence with <strong>SQLite</strong>, and editorial web experiences
              that prioritize speed, accessibility, and purposeful aesthetics over generic templates.
            </p>
          </div>

          {/* What I Do Section */}
          <div className="what-i-do-container">
            <h3 className="what-i-do-title">What I Do</h3>
            <div className="what-i-do-list">
              {WHAT_I_DO_ITEMS.map((item) => (
                <div key={item.number} className="what-i-do-item">
                  <span className="what-i-do-number">{item.number}</span>
                  <div className="what-i-do-content">
                    <h4>{item.title}</h4>
                    <p>{item.description}</p>
                  </div>
                </div>
              ))}
            </div>
          </div>
        </div>
      </div>
    </section>
  )
}
