import { StrictMode } from 'react';
import { createRoot, hydrateRoot } from 'react-dom/client';
import App from './App.jsx';
import './index.css';

const rootEl = document.getElementById('root');
const app = (
  <StrictMode>
    <App />
  </StrictMode>
);

if (rootEl && rootEl.hasChildNodes()) {
  hydrateRoot(rootEl, app);
} else if (rootEl) {
  createRoot(rootEl).render(app);
}
