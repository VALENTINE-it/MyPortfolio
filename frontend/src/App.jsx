import { useState, useEffect, useCallback } from 'react';
import Navbar from './components/Navbar';
import Footer from './components/Footer';
import PagePagination from './components/PagePagination';
import HomePage from './pages/HomePage';
import AboutPage from './pages/AboutPage';
import ProjectsPage from './pages/ProjectsPage';
import SkillsPage from './pages/SkillsPage';
import ContactPage from './pages/ContactPage';
import NotFoundPage from './pages/NotFoundPage';
import { getRouteConfig, navigateTo, normalizePathname } from './utils/routing';
import { useSEO } from './hooks/useSEO';
import './App.css';

function App({ initialPath }) {
  const [currentPath, setCurrentPath] = useState(() => {
    if (initialPath !== undefined) {
      return normalizePathname(initialPath);
    }
    if (typeof window !== 'undefined') {
      return normalizePathname(window.location.pathname);
    }
    return '/';
  });

  const routeConfig = getRouteConfig(currentPath);
  useSEO(routeConfig);

  // Normalize trailing slashes in browser URL bar cleanly without duplicating history entries
  useEffect(() => {
    if (typeof window === 'undefined') return;
    const rawPath = window.location.pathname;
    const normalized = normalizePathname(rawPath);
    if (rawPath !== normalized) {
      window.history.replaceState(null, '', `${normalized}${window.location.hash || ''}`);
      setCurrentPath(normalized);
    }
  }, []);

  // Synchronize back/forward browser navigation
  useEffect(() => {
    const handlePopState = () => {
      setCurrentPath(normalizePathname(window.location.pathname));
    };

    window.addEventListener('popstate', handlePopState);
    return () => window.removeEventListener('popstate', handlePopState);
  }, []);

  const handleNavigate = useCallback((targetPath) => {
    navigateTo(targetPath, (nextPath) => {
      setCurrentPath(nextPath);
    });
  }, []);

  const renderPage = () => {
    switch (routeConfig.key) {
      case 'home':
        return <HomePage onNavigate={handleNavigate} />;
      case 'about':
        return <AboutPage onNavigate={handleNavigate} />;
      case 'projects':
        return <ProjectsPage onNavigate={handleNavigate} />;
      case 'skills':
        return <SkillsPage onNavigate={handleNavigate} />;
      case 'contact':
        return <ContactPage onNavigate={handleNavigate} />;
      default:
        return <NotFoundPage onNavigate={handleNavigate} />;
    }
  };

  return (
    <div className="app-container">
      <Navbar currentPath={currentPath} onNavigate={handleNavigate} />
      {routeConfig.key === 'home' && <PagePagination />}
      <main id="main-content">{renderPage()}</main>
      <Footer onNavigate={handleNavigate} />
    </div>
  );
}

export default App;
