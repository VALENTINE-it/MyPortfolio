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
  {
    id: 4,
    title: 'Career Guidance Platform',
    description:
      'An interactive educational advisory application helping secondary students in Western Kenya discover tertiary courses and career pathways matching their academic strengths.',
    category: 'Full-Stack EdTech System',
    year: 2026,
    technologies: ['React', 'JavaScript', 'Go', 'SQLite', 'REST APIs', 'Analytics'],
    image: '/images/projects/career-guidance.svg',
    githubUrl: 'https://github.com/VALENTINE-it/career-guidance',
    liveUrl: '',
  },
  {
    id: 5,
    title: 'HopeReach Community Platform',
    description:
      'A comprehensive digital NGO operations and outreach platform delivering real-time meal tracking (500k+ meals served), housing intake logs, and volunteer mobilization.',
    category: 'Full-Stack Web Application',
    year: 2026,
    technologies: ['React', 'Go', 'SQLite', 'REST APIs', 'UI/UX', 'Community Platforms'],
    image: '/images/hopereach.jpeg',
    githubUrl: 'https://github.com/VALENTINE-it/hope-reach',
    liveUrl: 'https://hope-reach-project-ngo-gdq7.onrender.com/',
  },
  {
    id: 6,
    title: 'Valentine Omondi Developer Portfolio',
    description:
      'A high-performance editorial developer portfolio showcasing architectural systems, real-time telemetry projects, and full-stack engineering achievements with Go and React.',
    category: 'Portfolio & System Showcase',
    year: 2026,
    technologies: ['React', 'Vite', 'JavaScript', 'Go', 'REST APIs', 'SQLite'],
    image: '/images/projects/portfolio.svg',
    githubUrl: 'https://github.com/VALENTINE-it/portfolio',
    liveUrl: '',
  },
]

const PROJECTS_PER_PAGE = 3

function ProjectTimeline() {
  const [projects, setProjects] = useState(FALLBACK_PROJECTS)
  const [currentPage, setCurrentPage] = useState(1)

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

  const totalPages = Math.max(1, Math.ceil(projects.length / PROJECTS_PER_PAGE))
  const startIndex = (currentPage - 1) * PROJECTS_PER_PAGE
  const visibleProjects = projects.slice(startIndex, startIndex + PROJECTS_PER_PAGE)

  const handlePageChange = (newPage) => {
    if (newPage < 1 || newPage > totalPages || newPage === currentPage) return
    setCurrentPage(newPage)
    const el = document.getElementById('portfolio')
    if (el) {
      const offset = 70
      const bodyRect = document.body.getBoundingClientRect().top
      const elementRect = el.getBoundingClientRect().top
      window.scrollTo({
        top: elementRect - bodyRect - offset,
        behavior: 'smooth',
      })
    }
  }

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
            {visibleProjects.map((project) => {
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

          {/* Projects Pagination Component */}
          {totalPages > 1 && (
            <div className="portfolio-pagination" role="navigation" aria-label="Projects pagination">
              <div className="pagination-info">
                <span>
                  Showing {startIndex + 1}–{Math.min(startIndex + PROJECTS_PER_PAGE, projects.length)} of {projects.length} Projects
                </span>
              </div>

              <div className="pagination-controls">
                <button
                  type="button"
                  className="pagination-btn pagination-prev"
                  onClick={() => handlePageChange(currentPage - 1)}
                  disabled={currentPage === 1}
                  aria-label="Previous page"
                >
                  <i className="fas fa-chevron-left"></i> Previous
                </button>

                <div className="pagination-pages">
                  {Array.from({ length: totalPages }, (_, i) => i + 1).map((pageNum) => (
                    <button
                      key={pageNum}
                      type="button"
                      className={`pagination-page-btn ${currentPage === pageNum ? 'active' : ''}`}
                      onClick={() => handlePageChange(pageNum)}
                      aria-label={`Page ${pageNum}`}
                      aria-current={currentPage === pageNum ? 'page' : undefined}
                    >
                      {pageNum}
                    </button>
                  ))}
                </div>

                <button
                  type="button"
                  className="pagination-btn pagination-next"
                  onClick={() => handlePageChange(currentPage + 1)}
                  disabled={currentPage === totalPages}
                  aria-label="Next page"
                >
                  Next <i className="fas fa-chevron-right"></i>
                </button>
              </div>
            </div>
          )}

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
