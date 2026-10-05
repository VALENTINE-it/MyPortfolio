import './Projects.css'

function ProjectItem({ item, isLeft }) {
  const technologies = Array.isArray(item.technologies)
    ? item.technologies
    : []

  const title = item.name || item.title
  const role = item.role || item.category || 'Full-Stack Project'
  const date = item.period || item.year || '2026'

  return (
    <div className={`timeline-item ${isLeft ? 'left' : 'right'}`} id={`timeline-${item.id}`}>
      <div className="timeline-date">{date}</div>
      <div className="timeline-text">
        {item.logo && (
          <div className="timeline-header-org">
            <div className="timeline-circular-logo">
              <img src={item.logo} alt={`${title} Logo`} />
            </div>
          </div>
        )}
        <h2>{title}</h2>
        <h4>{role}</h4>
        <p>{item.description}</p>

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
          {item.githubUrl && (
            <a
              href={item.githubUrl}
              target="_blank"
              rel="noopener noreferrer"
              className="btn btn-sm"
              aria-label={`View ${title} on GitHub`}
            >
              GitHub <i className="fab fa-github" style={{ marginLeft: 6 }}></i>
            </a>
          )}
          {item.liveUrl && (
            <a
              href={item.liveUrl === '#' ? '#!' : item.liveUrl}
              target={item.liveUrl && item.liveUrl !== '#' && item.liveUrl !== '#!' ? '_blank' : undefined}
              rel={item.liveUrl && item.liveUrl !== '#' && item.liveUrl !== '#!' ? 'noopener noreferrer' : undefined}
              className="btn btn-secondary btn-sm"
              aria-label={`Visit live site for ${title}`}
              onClick={(e) => {
                if (item.liveUrl === '#' || item.liveUrl === '#!') {
                  e.preventDefault()
                }
              }}
            >
              Live Site <i className="fas fa-external-link-alt" style={{ marginLeft: 6 }}></i>
            </a>
          )}
        </div>

        {(item.aboutOrg || item.orgDetail) && (
          <div className="timeline-org-detail">
            <p className="timeline-org-detail-text">
              <strong>About {title}:</strong> {item.aboutOrg || item.orgDetail}
            </p>
          </div>
        )}
      </div>
    </div>
  )
}

export default ProjectItem
