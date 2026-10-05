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
            {projects.map((project) => (
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
