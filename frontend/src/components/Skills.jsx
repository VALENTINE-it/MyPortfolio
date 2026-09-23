import { useState, useEffect, useCallback } from 'react'
import { skillService } from '../services/skillService'
import './Skills.css'

const SKILL_ICONS = {
  frontend: 'fas fa-laptop-code',
  backend: 'fas fa-server',
  databases: 'fas fa-database',
  tools: 'fas fa-tools',
  other: 'fas fa-brain',
}

const FALLBACK_SKILLS = [
  {
    category: 'Frontend Development',
    icon: 'fas fa-laptop-code',
    description: 'React, JavaScript, HTML5, CSS3, Vite, Component Architecture, Responsive Design',
    skills: ['React', 'JavaScript', 'HTML5', 'CSS3', 'Vite', 'Responsive Design'],
  },
  {
    category: 'Backend Systems',
    icon: 'fas fa-server',
    description: 'Go (Golang), RESTful APIs, HTTP Handlers, Routing, Concurrency, Architecture',
    skills: ['Go', 'REST APIs', 'HTTP Handlers', 'Routing', 'Concurrency'],
  },
  {
    category: 'Databases & Storage',
    icon: 'fas fa-database',
    description: 'SQLite, SQL Schema Design, Relational Normalization, Indexing, Persistence',
    skills: ['SQLite', 'SQL', 'Schema Design', 'Data Normalization'],
  },
  {
    category: 'Tools & DevOps',
    icon: 'fas fa-terminal',
    description: 'Linux Environments, Git, GitHub Workflows, VS Code, CI/CD Fundamentals',
    skills: ['Linux', 'Git', 'GitHub', 'VS Code', 'Debugging'],
  },
]

function Skills() {
  const [categories, setCategories] = useState(FALLBACK_SKILLS)

  const fetchSkills = useCallback(async () => {
    try {
      const data = await skillService.getSkills({ grouped: true })
      if (Array.isArray(data) && data.length > 0) {
        // Map icons to categories
        const enriched = data.map((item) => {
          const lower = (item.category || '').toLowerCase()
          let icon = SKILL_ICONS.other
          if (lower.includes('front')) icon = SKILL_ICONS.frontend
          else if (lower.includes('back')) icon = SKILL_ICONS.backend
          else if (lower.includes('data')) icon = SKILL_ICONS.databases
          else if (lower.includes('tool')) icon = SKILL_ICONS.tools

          return {
            ...item,
            icon,
            description: item.skills ? item.skills.join(', ') : '',
          }
        })
        setCategories(enriched)
      } else {
        setCategories(FALLBACK_SKILLS)
      }
    } catch {
      setCategories(FALLBACK_SKILLS)
    }
  }, [])

  useEffect(() => {
    fetchSkills()
  }, [fetchSkills])

  return (
    <div className="service" id="skills">
      <div className="container">
        <div className="section-header text-center">
          <p>What I Do?</p>
          <h2>Technical Skills</h2>
        </div>

        <div className="service-grid">
          {categories.map((cat, idx) => (
            <div key={cat.category || idx} className="service-item">
              <div className="service-icon">
                <i className={cat.icon || 'fas fa-code'}></i>
              </div>
              <div className="service-text">
                <h3>{cat.category}</h3>
                <p>{cat.description}</p>
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  )
}

export default Skills
