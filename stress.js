import http from 'k6/http';
import { check } from 'k6';

export let options = {
    stages: [
        { duration: '10s', target: 20 },  // ramp up to 20 users
        { duration: '1m', target: 50 },  // stay at 50 users for 1 minutes
        { duration: '20s', target: 0 },   // scale down. Ramping down can be optional
    ],
};

export default function () {
    const res = http.get('http://127.0.0.1:3000/_sys/poke');

    // Check the response
    check(res, {
        'status was 200': (r) => r.status == 200,
        'transaction time OK': (r) => r.timings.duration < 200,
    });
}
