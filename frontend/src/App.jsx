import Navbar from './components/Navbar'
import Hero from './components/Hero'
import About from './components/About'
import ProjectTimeline from './components/ProjectTimeline'
import Skills from './components/Skills'
import Quotes from './components/Quotes'
import ContactForm from './components/ContactForm'
import Footer from './components/Footer'
import './App.css'

function App() {
  return (
    <div className="app-container">
      <Navbar />
      <main>
        <Hero />
        <About />
        <ProjectTimeline />
        <Skills />
        <Quotes />
        <ContactForm />
      </main>
      <Footer />
    </div>
  )
}

export default App
