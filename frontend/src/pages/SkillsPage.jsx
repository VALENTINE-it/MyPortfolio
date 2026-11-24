import PageHeader from '../components/PageHeader';
import Skills from '../components/Skills';

export default function SkillsPage({ onNavigate }) {
  return (
    <>
      <PageHeader
        title="Technical Skills &amp; Engineering Stack"
        subtitle="Core capabilities across React frontend engineering, Go backend APIs, SQLite database design, computer networking, Linux, and Web3."
        breadcrumbLabel="Skills"
        onNavigate={onNavigate}
      />
      <Skills isStandalonePage onNavigate={onNavigate} />
    </>
  );
}
