import Hero from '../components/Hero';
import About from '../components/About';
import ProjectTimeline from '../components/ProjectTimeline';
import Skills from '../components/Skills';
import Quotes from '../components/Quotes';
import ContactForm from '../components/ContactForm';

export default function HomePage({ onNavigate }) {
  return (
    <>
      <Hero />
      <About onNavigate={onNavigate} />
      <ProjectTimeline onNavigate={onNavigate} />
      <Skills onNavigate={onNavigate} />
      <Quotes />
      <ContactForm />
    </>
  );
}
