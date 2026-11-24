import { useState, useEffect, useCallback } from 'react';
import ProjectItem from './ProjectItem';
import OptimizedImage from './OptimizedImage';
import { projectService } from '../services/projectService';
import { PORTFOLIO_PROJECTS } from '../utils/structuredData';
import './Projects.css';

// Organizations under Accomplishments
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
    liveUrl: '',
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
];

const PROJECT_IMAGE_DIMENSIONS = {
  '/images/projects/safeguarding.png': { width: 1200, height: 800 },
  '/images/projects/fish-cage-iot.jpeg': { width: 640, height: 640 },
  '/images/projects/angaguard.png': { width: 1076, height: 711 },
  '/images/projects/career-guidance.svg': { width: 800, height: 500 },
  '/images/hopereach.jpeg': { width: 800, height: 791 },
  '/images/projects/portfolio.svg': { width: 800, height: 500 },
};

const PROJECTS_PER_PAGE = 3;

function ProjectTimeline({ isStandalonePage = false, onNavigate }) {
  const [projects, setProjects] = useState(PORTFOLIO_PROJECTS);
  const [currentPage, setCurrentPage] = useState(1);

  const fetchProjects = useCallback(async () => {
    try {
      const data = await projectService.getProjects();
      if (Array.isArray(data) && data.length > 0) {
        setProjects(data);
      }
    } catch {
      // Fallback is already initialized
    }
  }, []);

  useEffect(() => {
    fetchProjects();
  }, [fetchProjects]);

  const showAll = isStandalonePage;
  const totalPages = showAll ? 1 : Math.max(1, Math.ceil(projects.length / PROJECTS_PER_PAGE));
  const startIndex = showAll ? 0 : (currentPage - 1) * PROJECTS_PER_PAGE;
  const visibleProjects = showAll
    ? projects
    : projects.slice(startIndex, startIndex + PROJECTS_PER_PAGE);

  const handlePageChange = (newPage) => {
    if (newPage < 1 || newPage > totalPages || newPage === currentPage) return;
    setCurrentPage(newPage);
    const el = document.getElementById('portfolio');
    if (el) {
      const offset = 70;
      const bodyRect = document.body.getBoundingClientRect().top;
      const elementRect = el.getBoundingClientRect().top;
      window.scrollTo({
        top: elementRect - bodyRect - offset,
        behavior: 'smooth',
      });
    }
  };

  const handleLinkClick = (e, path) => {
    if (typeof onNavigate === 'function') {
      e.preventDefault();
      onNavigate(path);
    }
  };

  return (
    <>
      {/* 1. Accomplishments Section (#experience) — Featuring Organisations */}
      <section className="experience" id="experience" aria-label="Accomplishments and Organizations">
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
      </section>

      {/* 2. Portfolio Gallery Grid (#portfolio) */}
      <section className="portfolio" id="portfolio" aria-label="Featured Software Projects">
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
                : [];

              const imgSrc = project.image || '/images/projects/portfolio.svg';
              const dims = PROJECT_IMAGE_DIMENSIONS[imgSrc] || { width: 800, height: 533 };
              const targetUrl = project.liveUrl || project.githubUrl || 'https://github.com/VALENTINE-it';

              return (
                <article key={project.id} className="portfolio-item">
                  <div className="portfolio-wrap">
                    <div className="portfolio-img">
                      <OptimizedImage
                        src={imgSrc}
                        alt={`${project.title} — ${project.category || 'Software project'}`}
                        width={dims.width}
                        height={dims.height}
                        loading="lazy"
                        decoding="async"
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
                        href={targetUrl}
                        target="_blank"
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
                      <h4 className="portfolio-overlay-title">{project.title}</h4>
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
                            Live Site <i className="fas fa-external-link-alt" aria-hidden="true" style={{ marginLeft: 5 }}></i>
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
                            GitHub <i className="fab fa-github" aria-hidden="true" style={{ marginLeft: 5 }}></i>
                          </a>
                        )}
                      </div>
                    </div>
                  </div>
                </article>
              );
            })}
          </div>

          {/* Projects Pagination Component (on Home view) */}
          {!showAll && totalPages > 1 && (
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
                  <i className="fas fa-chevron-left" aria-hidden="true"></i> Previous
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
                  Next <i className="fas fa-chevron-right" aria-hidden="true"></i>
                </button>
              </div>
            </div>
          )}

          {/* View Profile on GitHub & Internal Links */}
          <div className="portfolio-more text-center" style={{ display: 'flex', flexWrap: 'wrap', gap: '12px', justifyContent: 'center' }}>
            {!isStandalonePage && (
              <a
                href="/projects"
                className="btn btn-secondary"
                onClick={(e) => handleLinkClick(e, '/projects')}
              >
                View All Projects
              </a>
            )}
            <a
              href="https://github.com/VALENTINE-it"
              target="_blank"
              rel="noopener noreferrer"
              className="btn btn-primary"
              aria-label="Visit my GitHub profile"
            >
              View GitHub Profile <i className="fab fa-github" aria-hidden="true" style={{ marginLeft: 8 }}></i>
            </a>
            {isStandalonePage && (
              <a
                href="/contact"
                className="btn btn-secondary"
                onClick={(e) => handleLinkClick(e, '/contact')}
              >
                Discuss a Project
              </a>
            )}
          </div>
        </div>
      </section>
    </>
  );
}

export default ProjectTimeline;
