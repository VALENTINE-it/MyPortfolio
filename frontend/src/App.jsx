import Navbar from './components/Navbar'
import Footer from './components/Footer'
import './App.css'

function App() {
  return (
    <div className="app-container">
      <Navbar />
      <main>
        <section id="home" className="section container" style={{ paddingTop: '8rem', minHeight: '80vh' }}>
          <span className="label-editorial">Introduction</span>
          <h1 className="heading-editorial">VALENTINE OMONDI AWILI</h1>
          <p>Full-Stack Developer</p>
        </section>

        <section id="about" className="section container" style={{ minHeight: '60vh' }}>
          <span className="label-editorial">Background</span>
          <h2 className="heading-editorial">ABOUT</h2>
        </section>

        <section id="projects" className="section container" style={{ minHeight: '60vh' }}>
          <span className="label-editorial">Selected Works</span>
          <h2 className="heading-editorial">PROJECTS</h2>
        </section>

        <section id="skills" className="section container" style={{ minHeight: '60vh' }}>
          <span className="label-editorial">Expertise</span>
          <h2 className="heading-editorial">SKILLS</h2>
        </section>

        <section id="contact" className="section container" style={{ minHeight: '60vh' }}>
          <span className="label-editorial">Get In Touch</span>
          <h2 className="heading-editorial">CONTACT</h2>
        </section>
      </main>
      <Footer />
    </div>
  )
}

export default App
