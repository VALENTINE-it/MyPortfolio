import PageHeader from '../components/PageHeader';
import ProjectTimeline from '../components/ProjectTimeline';

export default function ProjectsPage({ onNavigate }) {
  return (
    <>
      <PageHeader
        title="Software Development Projects &amp; Accomplishments"
        subtitle="Production web applications, IoT telemetry systems, and Web3 architectures engineered with React, Go, SQLite, and REST APIs."
        breadcrumbLabel="Projects"
        onNavigate={onNavigate}
      />
      <ProjectTimeline isStandalonePage onNavigate={onNavigate} />
    </>
  );
}
