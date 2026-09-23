import './Projects.css'

function ProjectItem({ project, isLeft }) {
  const technologies = Array.isArray(project.technologies)
    ? project.technologies
    : []

  return (
    <div className={`timeline-item ${isLeft ? 'left' : 'right'}`} id={`timeline-${project.id}`}>
      <div className="timeline-date">{project.year || '2026'}</div>
      <div className="timeline-text">
        <h2>{project.title}</h2>
        <h4>{project.category || 'Full-Stack Project'}</h4>
        <p>{project.description}</p>

        {technologies.length > 0 && (
          <div className="timeline-techs">
            {technologies.map((tech) => (
              <span key={tech} className="timeline-tech-badge">
                {tech}
              </span>
            ))}
          </div>
        )}

        <div className="timeline-actions">
          {project.githubUrl && (
            <a
              href={project.githubUrl}
              target="_blank"
              rel="noopener noreferrer"
              className="btn btn-sm"
              aria-label={`View ${project.title} on GitHub`}
            >
              GitHub <i className="fab fa-github" style={{ marginLeft: 6 }}></i>
            </a>
          )}
          {project.liveUrl && (
            <a
              href={project.liveUrl}
              target="_blank"
              rel="noopener noreferrer"
              className="btn btn-secondary btn-sm"
              aria-label={`Visit live site for ${project.title}`}
            >
              Demo <i className="fas fa-external-link-alt" style={{ marginLeft: 6 }}></i>
            </a>
          )}
        </div>
      </div>
    </div>
  )
}

export default ProjectItem
