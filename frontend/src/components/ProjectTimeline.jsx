import { useState, useEffect, useCallback } from 'react'
import ProjectItem from './ProjectItem'
import { projectService } from '../services/projectService'
import './Projects.css'

// Organizations under Accomplishments (managed one by one)
const ORGANISATIONS = [
  {
    id: 'hopereach',
    name: 'HopeReach',
    role: 'Founder and CEO',
    period: '2026',
    logo: '/images/hopereach.jpeg',
    description:
      'HopeReach is a community-driven NGO based in Kisumu, Kenya, dedicated to empowering vulnerable communities through compassionate support and sustainable solutions—addressing food security (500,000+ meals served), transitional housing (250+ individuals housed), education & skills development, and mental health advocacy.',
    technologies: ['React', 'JavaScript', 'Go', 'REST APIs', 'UI/UX', 'Community Platforms'],
    liveUrl: 'https://hope-reach-project-ngo-gdq7.onrender.com/',
  },
  {
    id: 'afrinit',
    name: 'Afrinit',
    role: 'Co-Founder',
    period: '2026',
    logo: '/images/afrinit.jpeg',
    description:
      'Afrinit is a technology company delivering end-to-end software development and comprehensive tech-related solutions, specializing primarily in building robust custom software, responsive web applications, and digital systems.',
    technologies: ['Software Development', 'Full-Stack Engineering', 'Web Applications', 'Tech Solutions'],
    liveUrl: '#',
  },
  {
    id: 'zone01',
    name: 'Zone01 Kisumu',
    role: 'Software Developer',
    period: '2026',
    logo: '/images/zone01.png',
    description:
      'Zone01 Kisumu is an innovative tech talent accelerator in Kisumu, Kenya, powered by the 01Edu peer-to-peer learning model. It delivers intensive, tuition-free, project-based training in full-stack software engineering, systems programming, Go, JavaScript, and collaborative problem-solving to cultivate top-tier engineering talent.',
    technologies: ['Software Engineering', 'Full-Stack Development', 'Go', 'JavaScript', 'System Architecture', 'Peer-to-Peer Learning'],
    liveUrl: 'https://www.zone01kisumu.ke/',
  },
]

const FALLBACK_PROJECTS = [
  {
    id: 1,
    title: 'Safeguarding Reporting Platform',
    description:
      'A secure, confidential reporting platform designed for anonymous safeguarding disclosures, encrypted submission workflows, and case audit records.',
    category: 'Full-Stack Web Application',
    year: 2026,
    technologies: ['React', 'JavaScript', 'Go', 'SQLite', 'REST APIs', 'Security'],
    image: '/images/projects/safeguarding.png',
    githubUrl: 'https://github.com/VALENTINE-it/safeguarding-project',
    liveUrl: 'https://safeguarding-app-1.onrender.com/',
  },
  {
    id: 2,
    title: 'Nemo',
    description:
      'An award-winning IoT security system recognized at the Kijani Space Hackathon, engineered to secure, monitor, and protect fish cages in Lake Victoria through real-time telemetry and intrusion alerts.',
    category: 'IoT & Embedded Security',
    year: 2026,
    technologies: ['IoT', 'Embedded Systems', 'Sensors', 'Telemetry', 'Hardware Security', 'Go'],
    image: '/images/projects/fish-cage-iot.jpeg',
    githubUrl: 'https://github.com/nyabokegrace/nemo',
    liveUrl: '',
  },
  {
    id: 3,
    title: 'Anga Guard',
    description:
      'A decentralized dMRV oracle and SME ESG platform empowering Western Kenya smallholders with verifiable biochar carbon removal credits compliant with Kenya National Carbon Registry.',
    category: 'Web3 & Climate Tech Oracle',
    year: 2026,
    technologies: ['dMRV Oracle', 'Web3', 'React', 'Go', 'IoT Telemetry', 'Smart Contracts', 'ESG Analytics'],
    image: '/images/projects/angaguard.png',
    githubUrl: 'https://github.com/ClayMichael2004/angaguard',
    liveUrl: 'https://angaguard-d96o.onrender.com/',
  },
]

function ProjectTimeline() {
  const [projects, setProjects] = useState(FALLBACK_PROJECTS)

  const fetchProjects = useCallback(async () => {
    try {
      const data = await projectService.getProjects()
      if (Array.isArray(data) && data.length > 0) {
        setProjects(data)
      }
    } catch {
      // Fallback is already initialized
    }
  }, [])

  useEffect(() => {
    fetchProjects()
  }, [fetchProjects])

  return (
    <>
      {/* 1. Accomplishments Section (#experience) — Featuring Organisations */}
      <div className="experience" id="experience">
        <div className="container">
          <div className="section-header text-center">
            <p>Where I&apos;ve Worked &amp; Contributing</p>
            <h2>Accomplishments</h2>
          </div>

          <div className="timeline">
            {ORGANISATIONS.map((org, index) => (
              <ProjectItem
                key={org.id || index}
                item={org}
                isLeft={index % 2 === 0}
              />
            ))}
          </div>
        </div>
      </div>

      {/* 2. Portfolio Gallery Grid (#portfolio) */}
      <div className="portfolio" id="portfolio">
        <div className="container">
          <div className="section-header text-center">
            <p>My Portfolio</p>
            <h2>Featured Works</h2>
          </div>

          {/* Portfolio Item Cards */}
          <div className="portfolio-grid">
            {projects.map((project) => {
              const techs = Array.isArray(project.technologies)
                ? project.technologies
                : typeof project.technologies === 'string'
                ? JSON.parse(project.technologies || '[]')
                : []

              return (
                <div key={project.id} className="portfolio-item">
                  <div className="portfolio-wrap">
                    <div className="portfolio-img">
                      <img
                        src={project.image || '/images/projects/portfolio.svg'}
                        alt={project.title}
                        loading="lazy"
                      />
                    </div>
                    <div className="portfolio-text">
                      <div className="portfolio-info">
                        <h3 title={project.title}>{project.title}</h3>
                        {project.category && (
                          <span className="portfolio-card-category">{project.category}</span>
                        )}
                      </div>
                      <a
                        className="btn"
                        href={project.liveUrl || project.githubUrl || '#'}
                        target={project.liveUrl || project.githubUrl ? '_blank' : undefined}
                        rel="noopener noreferrer"
                        title="View Details"
                        aria-label={`View ${project.title}`}
                      >
                        +
                      </a>
                    </div>

                    {/* Full Details Reveal on Hover */}
                    <div className="portfolio-overlay">
                      <div className="portfolio-overlay-header">
                        {project.category && (
                          <span className="portfolio-overlay-cat">{project.category}</span>
                        )}
                        {project.year && (
                          <span className="portfolio-overlay-year">{project.year}</span>
                        )}
                      </div>
                      <h3 className="portfolio-overlay-title">{project.title}</h3>
                      <p className="portfolio-overlay-desc">{project.description}</p>

                      {techs.length > 0 && (
                        <div className="portfolio-overlay-techs">
                          {techs.map((tech) => (
                            <span key={tech} className="portfolio-tech-tag">
                              {tech}
                            </span>
                          ))}
                        </div>
                      )}

                      <div className="portfolio-overlay-actions">
                        {project.liveUrl && (
                          <a
                            href={project.liveUrl}
                            target="_blank"
                            rel="noopener noreferrer"
                            className="btn btn-sm"
                            aria-label={`Live demo for ${project.title}`}
                          >
                            Live Site <i className="fas fa-external-link-alt" style={{ marginLeft: 5 }}></i>
                          </a>
                        )}
                        {project.githubUrl && (
                          <a
                            href={project.githubUrl}
                            target="_blank"
                            rel="noopener noreferrer"
                            className="btn btn-secondary btn-sm"
                            aria-label={`GitHub repo for ${project.title}`}
                          >
                            GitHub <i className="fab fa-github" style={{ marginLeft: 5 }}></i>
                          </a>
                        )}
                      </div>
                    </div>
                  </div>
                </div>
              )
            })}
          </div>

          {/* View Profile on GitHub */}
          <div className="portfolio-more text-center">
            <a
              href="https://github.com/VALENTINE-it"
              target="_blank"
              rel="noopener noreferrer"
              className="btn btn-primary"
              aria-label="Visit my GitHub profile"
            >
              View GitHub Profile <i className="fab fa-github" style={{ marginLeft: 8 }}></i>
            </a>
          </div>
        </div>
      </div>
    </>
  )
}

export default ProjectTimeline
