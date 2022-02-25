import http from 'k6/http';

export const options = {
    discardResponseBodies: true,
    scenarios: {
        contacts: {
            executor: 'per-vu-iterations',
            vus: 100,
            iterations: 200,
            maxDuration: '1h00m',
        },
    },
};

export default function () {
    http.get('http://localhost:12345/users/1');
}