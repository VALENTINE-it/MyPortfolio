import { useState } from 'react'
import './About.css'

export default function About() {
  const [isExpanded, setIsExpanded] = useState(false)

  return (
    <div className="about" id="about">
      <div className="container-fluid">
        <div className="about-row">
          {/* Left Column: Image */}
          <div className="about-col-img">
            <div className="about-img">
              <img
                src="/images/about-me.jpeg"
                alt="Valentine Omondi Awili Workspace"
                loading="lazy"
              />
            </div>
          </div>

          {/* Right Column: Narrative Content */}
          <div className="about-col-content">
            <div className="about-content">
              <div className="section-header text-left">
                <p>Learn About Me</p>
                <h2>Full-Stack Developer</h2>
              </div>

              <div className="about-text">
                <p>
                  I am <strong>Valentine Omondi Awili</strong>, a dedicated full-stack software engineer
                  and <strong>blockchain technology developer</strong> passionate about engineering reliable, scalable, and responsive digital systems. My approach
                  combines clean architectural discipline with modern, intuitive user interfaces.
                </p>
                <p>
                  My engineering journey spans developing concurrent, high-throughput backend services
                  in <strong>Go</strong> to crafting expressive client interfaces with <strong>React</strong>, alongside decentralized Web3 solutions and smart contract development.
                </p>
                <p>
                  With a solid grounding in computer science principles and defense-in-depth architecture,
                  I am <strong>currently learning cybersecurity</strong> to expand my capabilities in vulnerability analysis, threat modeling, and building hardened, attack-resilient applications.
                </p>

                <div className="about-section-block">
                  <h3 className="about-subtitle">Building Solutions That Matter</h3>
                  <p>
                    My projects reflect my interest in using technology to address everyday challenges. From an anonymous safeguarding reporting platform to a career guidance application, I’m developing tools that help people access support and make informed decisions. These projects are strengthening my skills in application design, databases, authentication, and user privacy.
                  </p>
                </div>

                {isExpanded && (
                  <div className="about-expanded-content">
                    <div className="about-section-block">
                      <h3 className="about-subtitle">My Networking Journey</h3>
                      <p>
                        I’m also looking to venture into computer networking and internet connectivity. I want to understand how networks are designed, configured, secured, and maintained—from routers and IP addressing to bandwidth management and troubleshooting. My longer-term goal is to explore providing reliable internet access to homes and communities, bringing together my software development skills and practical networking knowledge.
                      </p>
                    </div>

                    <div className="about-section-block">
                      <h3 className="about-subtitle">Where I’m Heading</h3>
                      <p>
                        My goal is to grow into a well-rounded technology professional who can build applications and understand the infrastructure that keeps them running. Through my ICT studies, practical projects, and continued learning, I’m working toward a career that combines software development, networking, and cybersecurity.
                      </p>
                    </div>
                  </div>
                )}
              </div>

              <div className="about-actions">
                <button
                  type="button"
                  className="btn-read-more"
                  onClick={() => setIsExpanded(!isExpanded)}
                  aria-expanded={isExpanded}
                >
                  {isExpanded ? 'Read Less' : 'Read More'}
                  <i className={`fas fa-chevron-${isExpanded ? 'up' : 'down'}`}></i>
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}
