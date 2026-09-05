import http from 'k6/http';
import { check, sleep } from 'k6';
import { randomIntBetween, randomItem } from 'https://jslib.k6.io/k6-utils/1.2.0/index.js';

const BASE_URL = __ENV.BASE_URL || 'http://localhost:8080';

export function setup() {
  const res = http.get(`${BASE_URL}/api/tickets`);
  check(res, { 'tickets loaded': (r) => r.status === 200 });
  return { ticketIds: res.json('data').map((t) => t.id) };
}

export const options = {
  scenarios: {
    ramp_and_hold: {
      executor: 'ramping-vus',
      startVUs: 0,
      stages: [
        { duration: '1m', target: 100 },
        { duration: '30s', target: 500 },
        { duration: '3m', target: 500 },
        { duration: '30s', target: 0 },
      ],
      gracefulRampDown: '15s',
    },
  },
  thresholds: {
    http_req_duration: [
      { threshold: 'p(50)<500', abortOnFail: false },
      { threshold: 'p(95)<2000', abortOnFail: false },
      { threshold: 'p(99)<5000', abortOnFail: false },
    ],
    http_req_failed: [{ threshold: 'rate<0.1', abortOnFail: false }],
    checks: [{ threshold: 'rate>0.95', abortOnFail: false }],
    'http_req_duration{endpoint:register}': [{ threshold: 'p(95)<1500' }],
    'http_req_duration{endpoint:webhook}': [{ threshold: 'p(95)<1500' }],
  },
};

const NAMES = [
  'Budi Santoso', 'Siti Rahayu', 'Ahmad Hidayat', 'Dewi Lestari',
  'Rizki Pratama', 'Anisa Putri', 'Fajar Nugroho', 'Maya Sari',
  'Andi Wijaya', 'Rina Wulandari', 'Dimas Saputra', 'Putri Amelia',
];

const PHONES = [
  '081234567890', '085678901234', '087890123456', '081345678901',
  '085789012345', '087901234567', '081456789012', '085890123456',
];

export default function (data) {
  // 1. List tickets
  const ticketsRes = http.get(`${BASE_URL}/api/tickets`, { tags: { endpoint: 'tickets' } });
  check(ticketsRes, {
    'GET /api/tickets status 200': (r) => r.status === 200,
    'GET /api/tickets has data': (r) => r.json('success') === true,
  });

  // 2. Register participant with random ticket
  const ticketId = randomItem(data.ticketIds);
  const timestamp = Date.now();
  const payload = JSON.stringify({
    name: randomItem(NAMES),
    email: `user${timestamp}${randomIntBetween(1000, 9999)}@loadtest.com`,
    phone: randomItem(PHONES),
    ticket_id: ticketId,
  });

  const registerRes = http.post(`${BASE_URL}/api/register`, payload, {
    headers: { 'Content-Type': 'application/json' },
    tags: { endpoint: 'register' },
  });

  const registerOk = check(registerRes, {
    'POST /api/register status 201': (r) => r.status === 201,
    'POST /api/register success': (r) => r.json('success') === true,
  });

  if (!registerOk) {
    return;
  }

  const orderId = registerRes.json('data.payment.order_id');

  // 3. Random delay simulating payment (1-5 seconds)
  sleep(randomIntBetween(1, 5));

  // 4. Hit payment webhook
  const webhookRes = http.post(
    `${BASE_URL}/api/webhook/payment`,
    JSON.stringify({ order_id: orderId }),
    {
      headers: { 'Content-Type': 'application/json' },
      tags: { endpoint: 'webhook' },
    }
  );

  check(webhookRes, {
    'POST /webhook/payment status 200': (r) => r.status === 200,
    'POST /webhook/payment success': (r) => r.json('success') === true,
    'POST /webhook/payment bib generated': (r) =>
      r.json('data.bib_number') !== null && r.json('data.bib_number') !== undefined,
  });
}
