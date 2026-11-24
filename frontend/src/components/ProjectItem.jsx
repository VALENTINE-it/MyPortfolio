import OptimizedImage from './OptimizedImage';
import './Projects.css';

const LOGO_DIMENSIONS = {
  '/images/hopereach.jpeg': { width: 800, height: 791 },
  '/images/afrinit.jpeg': { width: 400, height: 400 },
  '/images/zone01.png': { width: 181, height: 107 },
};

function ProjectItem({ item, isLeft }) {
  const technologies = Array.isArray(item.technologies)
    ? item.technologies
    : [];

  const title = item.name || item.title;
  const role = item.role || item.category || 'Full-Stack Project';
  const date = item.period || item.year || '2026';
  const hasValidLiveUrl = Boolean(item.liveUrl && item.liveUrl !== '#' && item.liveUrl !== '#!');
  const logoDims = LOGO_DIMENSIONS[item.logo] || { width: 88, height: 88 };

  return (
    <article className={`timeline-item ${isLeft ? 'left' : 'right'}`} id={`timeline-${item.id}`}>
      <div className="timeline-date">{date}</div>
      <div className="timeline-text">
        {item.logo && (
          <div className="timeline-header-org">
            <div className="timeline-circular-logo">
              <OptimizedImage
                src={item.logo}
                alt={`${title} Logo`}
                width={logoDims.width}
                height={logoDims.height}
                loading="lazy"
                decoding="async"
              />
            </div>
          </div>
        )}
        <h3>{title}</h3>
        <p className="timeline-role">{role}</p>
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

        {(item.githubUrl || hasValidLiveUrl) && (
          <div className="timeline-actions">
            {item.githubUrl && (
              <a
                href={item.githubUrl}
                target="_blank"
                rel="noopener noreferrer"
                className="btn btn-sm"
                aria-label={`View ${title} on GitHub`}
              >
                GitHub <i className="fab fa-github" aria-hidden="true" style={{ marginLeft: 6 }}></i>
              </a>
            )}
            {hasValidLiveUrl && (
              <a
                href={item.liveUrl}
                target="_blank"
                rel="noopener noreferrer"
                className="btn btn-secondary btn-sm"
                aria-label={`Visit live site for ${title}`}
              >
                Live Site <i className="fas fa-external-link-alt" aria-hidden="true" style={{ marginLeft: 6 }}></i>
              </a>
            )}
          </div>
        )}

        {(item.aboutOrg || item.orgDetail) && (
          <div className="timeline-org-detail">
            <p className="timeline-org-detail-text">
              <strong>About {title}:</strong> {item.aboutOrg || item.orgDetail}
            </p>
          </div>
        )}
      </div>
    </article>
  );
}

export default ProjectItem;
