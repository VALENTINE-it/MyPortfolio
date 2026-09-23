import { useState } from 'react'
import { contactService } from '../services/contactService'
import './ContactForm.css'

function ContactForm() {
  const [formData, setFormData] = useState({
    name: '',
    email: '',
    subject: '',
    message: '',
  })

  const [errors, setErrors] = useState({})
  const [isSubmitting, setIsSubmitting] = useState(false)
  const [submitStatus, setSubmitStatus] = useState(null) // { success: boolean, message: string }

  const validate = () => {
    const errs = {}
    const trimmedName = formData.name.trim()
    const trimmedEmail = formData.email.trim()
    const trimmedSubject = formData.subject.trim()
    const trimmedMessage = formData.message.trim()

    if (!trimmedName) {
      errs.name = 'Please provide your name.'
    } else if (trimmedName.length > 100) {
      errs.name = 'Name cannot exceed 100 characters.'
    }

    if (!trimmedEmail) {
      errs.email = 'Please provide your email address.'
    } else if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(trimmedEmail)) {
      errs.email = 'Please enter a valid email address.'
    }

    if (!trimmedSubject) {
      errs.subject = 'Please enter a subject.'
    } else if (trimmedSubject.length > 200) {
      errs.subject = 'Subject cannot exceed 200 characters.'
    }

    if (!trimmedMessage) {
      errs.message = 'Please write your message.'
    } else if (trimmedMessage.length < 5) {
      errs.message = 'Message must be at least 5 characters long.'
    } else if (trimmedMessage.length > 5000) {
      errs.message = 'Message cannot exceed 5000 characters.'
    }

    setErrors(errs)
    return Object.keys(errs).length === 0
  }

  const handleChange = (e) => {
    const { name, value } = e.target
    setFormData((prev) => ({ ...prev, [name]: value }))
    if (errors[name]) {
      setErrors((prev) => ({ ...prev, [name]: null }))
    }
  }

  const handleSubmit = async (e) => {
    e.preventDefault()
    setSubmitStatus(null)

    if (!validate()) {
      return
    }

    setIsSubmitting(true)
    try {
      const response = await contactService.submitContact({
        name: formData.name.trim(),
        email: formData.email.trim(),
        subject: formData.subject.trim(),
        message: formData.message.trim(),
      })

      setSubmitStatus({
        success: true,
        message: response?.message || 'Thank you! Your message has been sent successfully.',
      })
      setFormData({ name: '', email: '', subject: '', message: '' })
      setErrors({})
    } catch (err) {
      let errorMsg = 'Failed to send message. Please try again later.'
      if (err.status === 429) {
        errorMsg = 'You have submitted several messages recently. Please wait a minute before trying again.'
      } else if (err.data?.message) {
        errorMsg = err.data.message
      }
      setSubmitStatus({
        success: false,
        message: errorMsg,
      })
    } finally {
      setIsSubmitting(false)
    }
  }

  const handleReset = () => {
    setSubmitStatus(null)
    setFormData({ name: '', email: '', subject: '', message: '' })
    setErrors({})
  }

  return (
    <div className="contact" id="contact">
      <div className="container">
        {/* Section Header */}
        <div className="section-header text-center">
          <p>Get In Touch</p>
          <h2>Contact Me</h2>
        </div>

        <div className="contact-row">
          {/* Left Column: Contact Information */}
          <div className="contact-info-col">
            <div className="contact-info">
              <h3>Let&apos;s Build Something Together</h3>
              <p>
                Have a project in mind, an architectural query, a prospective role,
                or an open-source initiative? Let&apos;s discuss and collaborate.
              </p>

              <div className="contact-detail-items">
                <div className="contact-detail-item">
                  <div className="contact-detail-icon">
                    <i className="fas fa-map-marker-alt"></i>
                  </div>
                  <div className="contact-detail-text">
                    <h4>Location</h4>
                    <p>Kisumu, Kenya</p>
                  </div>
                </div>

                <div className="contact-detail-item">
                  <div className="contact-detail-icon">
                    <i className="fas fa-envelope"></i>
                  </div>
                  <div className="contact-detail-text">
                    <h4>Email</h4>
                    <p>
                      <a href="mailto:valentineawili@gmail.com">valentineawili@gmail.com</a>
                    </p>
                  </div>
                </div>

                <div className="contact-detail-item">
                  <div className="contact-detail-icon">
                    <i className="fab fa-github"></i>
                  </div>
                  <div className="contact-detail-text">
                    <h4>GitHub</h4>
                    <p>
                      <a
                        href="https://github.com/VALENTINE-it"
                        target="_blank"
                        rel="noopener noreferrer"
                      >
                        github.com/VALENTINE-it
                      </a>
                    </p>
                  </div>
                </div>

                <div className="contact-detail-item">
                  <div className="contact-detail-icon">
                    <i className="fab fa-linkedin-in"></i>
                  </div>
                  <div className="contact-detail-text">
                    <h4>LinkedIn</h4>
                    <p>
                      <a
                        href="https://www.linkedin.com/in/valentine-omondi-434aa12ab/"
                        target="_blank"
                        rel="noopener noreferrer"
                      >
                        valentine-omondi
                      </a>
                    </p>
                  </div>
                </div>
              </div>
            </div>
          </div>

          {/* Right Column: Contact Form */}
          <div className="contact-form-col">
            <div className="contact-form-wrapper">
              {submitStatus?.success ? (
                <div className="contact-success" role="alert">
                  <i className="fas fa-check-circle success-icon"></i>
                  <h3>Message Sent Successfully!</h3>
                  <p>{submitStatus.message}</p>
                  <button onClick={handleReset} className="btn">
                    Send Another Message
                  </button>
                </div>
              ) : (
                <form onSubmit={handleSubmit} className="contact-form" noValidate>
                  {submitStatus?.success === false && (
                    <div className="contact-alert-error" role="alert">
                      <i className="fas fa-exclamation-triangle"></i> {submitStatus.message}
                    </div>
                  )}

                  <div className="form-group">
                    <input
                      type="text"
                      id="contact-name"
                      name="name"
                      value={formData.name}
                      onChange={handleChange}
                      placeholder="Your Name"
                      className={`form-control ${errors.name ? 'is-invalid' : ''}`}
                      disabled={isSubmitting}
                      aria-label="Your Name"
                    />
                    {errors.name && <span className="help-block">{errors.name}</span>}
                  </div>

                  <div className="form-group">
                    <input
                      type="email"
                      id="contact-email"
                      name="email"
                      value={formData.email}
                      onChange={handleChange}
                      placeholder="Your Email"
                      className={`form-control ${errors.email ? 'is-invalid' : ''}`}
                      disabled={isSubmitting}
                      aria-label="Your Email"
                    />
                    {errors.email && <span className="help-block">{errors.email}</span>}
                  </div>

                  <div className="form-group">
                    <input
                      type="text"
                      id="contact-subject"
                      name="subject"
                      value={formData.subject}
                      onChange={handleChange}
                      placeholder="Subject"
                      className={`form-control ${errors.subject ? 'is-invalid' : ''}`}
                      disabled={isSubmitting}
                      aria-label="Subject"
                    />
                    {errors.subject && <span className="help-block">{errors.subject}</span>}
                  </div>

                  <div className="form-group">
                    <textarea
                      id="contact-message"
                      name="message"
                      rows={4}
                      value={formData.message}
                      onChange={handleChange}
                      placeholder="Message"
                      className={`form-control ${errors.message ? 'is-invalid' : ''}`}
                      disabled={isSubmitting}
                      aria-label="Message"
                    />
                    {errors.message && <span className="help-block">{errors.message}</span>}
                  </div>

                  <div>
                    <button
                      type="submit"
                      disabled={isSubmitting}
                      className="btn contact-btn"
                    >
                      {isSubmitting ? 'Sending...' : 'Send Message'}
                    </button>
                  </div>
                </form>
              )}
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default ContactForm
