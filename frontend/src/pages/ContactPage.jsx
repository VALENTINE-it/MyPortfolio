import PageHeader from '../components/PageHeader';
import ContactForm from '../components/ContactForm';

export default function ContactPage({ onNavigate }) {
  return (
    <>
      <PageHeader
        title="Contact Valentine Omondi Awili"
        subtitle="Reach out for full-stack web application development, software engineering opportunities, or open-source collaboration in Kenya and remotely."
        breadcrumbLabel="Contact"
        onNavigate={onNavigate}
      />
      <ContactForm isStandalonePage />
    </>
  );
}
