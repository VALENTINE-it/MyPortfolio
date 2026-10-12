import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  stages: [
    { duration: '30s', target: 100 },   // ramp up to 100 VUs
    { duration: '1m', target: 500 },    // ramp up to 500 VUs
    { duration: '1m', target: 1000 },   // ramp up to 1,000 VUs
    { duration: '30s', target: 0 },     // ramp down
  ],
  thresholds: {
    http_req_duration: ['p(95)<200'],   // 95% of requests must complete below 200ms
    http_req_failed: ['rate<0.01'],     // less than 1% failed requests
  },
};

const BASE_URL = __ENV.TARGET_URL || 'http://localhost:8080';

export default function () {
  // 1. Health check
  const healthRes = http.get(`${BASE_URL}/health/live`);
  check(healthRes, {
    'health status 200': (r) => r.status === 200,
  });

  // 2. Fetch Projects (cached public endpoint)
  const projectsRes = http.get(`${BASE_URL}/api/projects`);
  check(projectsRes, {
    'projects status 200': (r) => r.status === 200,
    'projects has data': (r) => r.json('data') !== undefined,
  });

  // 3. Fetch Skills
  const skillsRes = http.get(`${BASE_URL}/api/skills`);
  check(skillsRes, {
    'skills status 200': (r) => r.status === 200,
  });

  sleep(1);
}
