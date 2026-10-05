import { useState, useEffect, useCallback } from 'react'
import ProjectItem from './ProjectItem'
import { projectService } from '../services/projectService'
import './Projects.css'

const FALLBACK_PROJECTS = [
  {
    id: 1,
    title: 'Safeguarding Reporting Platform',
    description:
      'A secure, confidential reporting platform designed for anonymous safeguarding disclosures, encrypted submission workflows, and case audit records.',
    category: 'Full-Stack Web Application',
    year: 2026,
    technologies: ['React', 'JavaScript', 'Go', 'SQLite', 'REST APIs', 'Security'],
    image: '/images/projects/safeguarding.svg',
    githubUrl: 'https://github.com/VALENTINE-it/safeguarding-platform',
    liveUrl: 'https://safeguarding.example.com',
  },
  {
    id: 2,
    title: 'Career Guidance Platform',
    description:
      'An educational guidance system that assesses academic performance and technical interests to recommend career pathways, skill milestones, and mentorship.',
    category: 'Education Technology',
    year: 2026,
    technologies: ['React', 'JavaScript', 'Go', 'REST APIs', 'SQLite'],
    image: '/images/projects/career-guidance.svg',
    githubUrl: 'https://github.com/VALENTINE-it/career-guidance',
    liveUrl: '',
  },
  {
    id: 3,
    title: 'Personal Developer Portfolio',
    description:
      'A digital editorial portfolio and professional showcase engineered with React and Go, featuring responsive fluid typography, clean REST architecture, and SQLite persistence.',
    category: 'Full-Stack Portfolio',
    year: 2026,
    technologies: ['React', 'Vite', 'JavaScript', 'Go', 'REST API', 'SQLite'],
    image: '/images/projects/portfolio.svg',
    githubUrl: 'https://github.com/VALENTINE-it/MyPortfolio',
    liveUrl: '',
  },
]

const ORGANISATIONS = [
  {
    id: 'hopereach',
    name: 'HopeReach',
    role: 'Full-Stack Developer & Technical Lead',
    status: 'Currently Working',
    period: '2026 – Present',
    logo: '/images/organisations/hopereach.jpeg',
    description:
      'Developing and maintaining vital digital solutions, including the anonymous safeguarding reporting platform, secure data architecture in Go, React UI development, and robust data privacy systems.',
    technologies: ['React', 'Go', 'SQLite', 'REST APIs', 'Data Privacy', 'Security'],
    isCurrent: true,
  },
  {
    id: 'community-connectivity',
    name: 'Community Internet & Tech Initiatives',
    role: 'Network & Systems Specialist',
    status: 'Collaborator',
    period: '2025 – 2026',
    icon: 'fas fa-network-wired',
    description:
      'Configured network routing schemas, bandwidth management policies, and local infrastructure to facilitate dependable internet connectivity and digital literacy.',
    technologies: ['Networking', 'Routing', 'Bandwidth Management', 'Troubleshooting'],
    isCurrent: false,
  },
  {
    id: 'web3-research',
    name: 'Web3 & Distributed Systems Lab',
    role: 'Blockchain Technology Developer',
    status: 'Milestone',
    period: '2025',
    icon: 'fas fa-cubes',
    description:
      'Engineered smart contract architectures, explored distributed ledger protocols, and built verifiable decentralized application components.',
    technologies: ['Smart Contracts', 'Web3', 'Distributed Ledgers', 'Cryptography'],
    isCurrent: false,
  },
]

const FILTER_CATEGORIES = ['All', 'Full-Stack', 'Backend', 'Education']

function ProjectTimeline() {
  const [projects, setProjects] = useState(FALLBACK_PROJECTS)
  const [activeFilter, setActiveFilter] = useState('All')

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

  const filteredProjects = projects.filter((project) => {
    if (activeFilter === 'All') return true
    const cat = (project.category || '').toLowerCase()
    return cat.includes(activeFilter.toLowerCase())
  })

  return (
    <>
      {/* 1. Accomplishments / Experience Timeline (#experience) */}
      <div className="experience" id="experience">
        <div className="container">
          <div className="section-header text-center">
            <p>What I Achieved?</p>
            <h2>Accomplishments</h2>
          </div>

          {/* Organisations Showcase */}
          <div className="orgs-showcase">
            <div className="orgs-header">
              <h3 className="orgs-title">
                <i className="fas fa-building"></i> Organisations & Experience
              </h3>
              <p className="orgs-subtitle">
                Organisations I have contributed to and projects I am currently actively engineering.
              </p>
            </div>

            <div className="orgs-grid">
              {ORGANISATIONS.map((org) => (
                <div
                  key={org.id}
                  className={`org-card ${org.isCurrent ? 'org-card-current' : ''}`}
                >
                  <div className="org-card-header">
                    <div className="org-logo-wrapper">
                      {org.logo ? (
                        <img
                          src={org.logo}
                          alt={`${org.name} Logo`}
                          className="org-logo-img"
                        />
                      ) : (
                        <div className="org-logo-icon">
                          <i className={org.icon}></i>
                        </div>
                      )}
                    </div>
                    <div className="org-header-meta">
                      <span
                        className={`org-status-badge ${
                          org.isCurrent ? 'status-current' : 'status-past'
                        }`}
                      >
                        {org.isCurrent && <span className="status-pulse-dot"></span>}
                        {org.status}
                      </span>
                      <span className="org-period">{org.period}</span>
                    </div>
                  </div>

                  <div className="org-card-body">
                    <h3 className="org-name">{org.name}</h3>
                    <h4 className="org-role">{org.role}</h4>
                    <p className="org-desc">{org.description}</p>
                    <div className="org-tech-stack">
                      {org.technologies.map((t) => (
                        <span key={t} className="org-tech-pill">
                          {t}
                        </span>
                      ))}
                    </div>
                  </div>
                </div>
              ))}
            </div>
          </div>

          <div
            className="section-header text-center"
            style={{ marginTop: '60px', marginBottom: '20px' }}
          >
            <p>Milestone Timeline</p>
            <h2>Project Milestones</h2>
          </div>

          <div className="timeline">
            {projects.map((project, index) => (
              <ProjectItem
                key={project.id || index}
                project={project}
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

          {/* Portfolio Filter Pills */}
          <ul id="portfolio-filter">
            {FILTER_CATEGORIES.map((filter) => (
              <li
                key={filter}
                className={activeFilter === filter ? 'filter-active' : ''}
                onClick={() => setActiveFilter(filter)}
              >
                {filter}
              </li>
            ))}
          </ul>

          {/* Portfolio Item Cards */}
          <div className="portfolio-grid">
            {filteredProjects.map((project) => (
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
                    <h3 title={project.title}>{project.title}</h3>
                    <a
                      className="btn"
                      href={project.liveUrl || project.githubUrl || '#'}
                      target="_blank"
                      rel="noopener noreferrer"
                      title="View Details"
                      aria-label={`View ${project.title}`}
                    >
                      +
                    </a>
                  </div>
                </div>
              </div>
            ))}
          </div>
        </div>
      </div>
    </>
  )
}

export default ProjectTimeline
