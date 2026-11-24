import PageHeader from '../components/PageHeader';
import About from '../components/About';

export default function AboutPage({ onNavigate }) {
  return (
    <>
      <PageHeader
        title="About Valentine Omondi Awili"
        subtitle="Full-stack software developer and ICT student at Maseno University building React, Go, blockchain, and resilient web applications in Kenya."
        breadcrumbLabel="About"
        onNavigate={onNavigate}
      />
      <About isStandalonePage onNavigate={onNavigate} />
    </>
  );
}
