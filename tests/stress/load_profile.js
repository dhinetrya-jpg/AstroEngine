import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
    stages: [
        { duration: '1m', target: 2000 },  
        { duration: '3m', target: 10000 }, 
        { duration: '1m', target: 0 },     
    ],
    thresholds: {
        http_req_duration: ['p(95)<200', 'p(99)<500'], 
        http_req_failed: ['rate<0.0001'],
    },
};

const BASE_URL = __ENV.TARGET_URL || 'http://localhost:8080';

export default function () {
    const payload = JSON.stringify({
        base_payload: {
            birth_date: '1989-11-23',
            birth_time: '04:30:00',
            latitude: 22.572645,
            longitude: 88.363892
        },
        milestones: [
            { event_type: 'CAREER_BREAKTHROUGH', timestamp: '2018-07-15T10:00:00Z' },
            { event_type: 'SURGERY', timestamp: '2022-03-22T06:30:00Z' }
        ],
        search_window_minutes: 15
    });

    const params = {
        headers: {
            'Content-Type': 'application/json',
            'X-User-ID': `STRESS_USER_${__VU}_${__ITER}`
        },
    };

    // Execute High-Precision Birth Time Calibration Step
    const btrResponse = http.post(`${BASE_URL}/api/v1/btr/rectify`, payload, params);
    check(btrResponse, {
        'BTR Status is 200': (r) => r.status === 200,
        'BTR Certified True': (r) => r.json().is_certified_true === true,
    });

    sleep(0.5); 

    // Run Multi-Dimensional Conflict Logic Resolution Matrix
    const resolverResponse = http.get(`${BASE_URL}/api/v1/resolver/evaluate?user_id=USR_KOLKATA_89&category=Wealth`, params);
    check(resolverResponse, {
        'Resolver Status is 200': (r) => r.status === 200,
        'Graph Path Verified': (r) => r.json().is_graph_verified === true,
    });

    sleep(0.5);
}
