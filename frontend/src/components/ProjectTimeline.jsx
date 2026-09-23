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
