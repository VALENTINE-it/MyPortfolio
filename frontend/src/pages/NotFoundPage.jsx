import PageHeader from '../components/PageHeader';

export default function NotFoundPage({ onNavigate }) {
  const handleNav = (e, path) => {
    if (typeof onNavigate === 'function') {
      e.preventDefault();
      onNavigate(path);
    }
  };

  return (
    <>
      <PageHeader
        title="404 — Page Not Found"
        subtitle="The page you requested does not exist or may have moved."
        breadcrumbLabel="404"
        onNavigate={onNavigate}
      />
      <section className="section" aria-label="Not found navigation options">
        <div className="container" style={{ textAlign: 'center', padding: '50px 15px 80px' }}>
          <h2 style={{ marginBottom: '14px' }}>Explore My Portfolio</h2>
          <p style={{ maxWidth: '580px', margin: '0 auto 28px' }}>
            Use the links below to return to the homepage or browse my projects, technical skills, and contact details.
          </p>
          <div style={{ display: 'flex', flexWrap: 'wrap', gap: '12px', justifyContent: 'center' }}>
            <a href="/" className="btn" onClick={(e) => handleNav(e, '/')}>
              Back to Home
            </a>
            <a href="/projects" className="btn btn-secondary" onClick={(e) => handleNav(e, '/projects')}>
              View Projects
            </a>
            <a href="/skills" className="btn btn-secondary" onClick={(e) => handleNav(e, '/skills')}>
              Technical Skills
            </a>
            <a href="/contact" className="btn btn-secondary" onClick={(e) => handleNav(e, '/contact')}>
              Contact Me
            </a>
          </div>
        </div>
      </section>
    </>
  );
}
