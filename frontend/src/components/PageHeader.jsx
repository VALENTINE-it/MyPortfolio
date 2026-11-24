import './PageHeader.css';

export default function PageHeader({ title, subtitle, breadcrumbLabel, onNavigate }) {
  const handleHomeClick = (e) => {
    if (typeof onNavigate === 'function') {
      e.preventDefault();
      onNavigate('/');
    }
  };

  return (
    <header className="page-header">
      <div className="container">
        <nav className="page-breadcrumb" aria-label="Breadcrumb">
          <ol>
            <li>
              <a href="/" onClick={handleHomeClick}>
                Home
              </a>
            </li>
            <li aria-current="page">{breadcrumbLabel}</li>
          </ol>
        </nav>
        <h1 className="page-header-title">{title}</h1>
        {subtitle && <p className="page-header-subtitle">{subtitle}</p>}
      </div>
    </header>
  );
}
