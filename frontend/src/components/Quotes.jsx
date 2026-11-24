import { useState, useEffect } from 'react';
import './Quotes.css';

const QUOTES = [
  {
    quote: '“When something is important enough, you do it even if the odds are not in your favor.”',
    author: 'Elon Musk',
    role: 'CEO of Tesla & Founder of SpaceX',
    image: '/images/profile.svg',
  },
  {
    quote: '“Simplicity is prerequisite for reliability. Great software begins with clear thinking.”',
    author: 'Edsger W. Dijkstra',
    role: 'Turing Laureate & Pioneer',
    image: '/images/profile.svg',
  },
  {
    quote: '“Any fool can write code that a computer can understand. Good programmers write code that humans can understand.”',
    author: 'Martin Fowler',
    role: 'Software Architecture Author',
    image: '/images/profile.svg',
  },
];

export default function Quotes() {
  const [currentIndex, setCurrentIndex] = useState(0);

  useEffect(() => {
    const timer = setInterval(() => {
      setCurrentIndex((prev) => (prev + 1) % QUOTES.length);
    }, 6000);
    return () => clearInterval(timer);
  }, []);

  const current = QUOTES[currentIndex];

  return (
    <section className="testimonial" id="review" aria-label="Engineering Principles">
      <div className="container">
        <h2 className="sr-only">Engineering Principles</h2>
        <div className="testimonial-icon" aria-hidden="true">
          <i className="fa fa-quote-left"></i>
        </div>

        <div className="testimonial-item">
          <div className="testimonial-img">
            <img
              src={current.image}
              alt=""
              width={110}
              height={110}
              loading="lazy"
              decoding="async"
            />
          </div>
          <div className="testimonial-text">
            <p>{current.quote}</p>
            <h3>{current.author}</h3>
            <p className="testimonial-role">{current.role}</p>
          </div>
        </div>

        {/* Carousel indicator dots */}
        <div className="testimonial-dots">
          {QUOTES.map((_, idx) => (
            <button
              key={idx}
              type="button"
              className={`testimonial-dot ${currentIndex === idx ? 'active' : ''}`}
              onClick={() => setCurrentIndex(idx)}
              aria-label={`Go to quote ${idx + 1}`}
            ></button>
          ))}
        </div>
      </div>
    </section>
  );
}
