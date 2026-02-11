const CACHE_NAME = 'cafe';
const urlsToCache = [
    '/',
    '/static/css/style.css',
    '/static/js/htmx.min.js',
    '/static/js/app.js',
    '/static/manifest.json'
];

self.addEventListener('install', (event) => {
    event.waitUntil(
        caches.open(CACHE_NAME)
            .then((cache) => cache.addAll(urlsToCache))
    );
});

self.addEventListener('fetch', (event) => {
    event.respondWith(
        caches.match(event.request)
            .then((response) => response || fetch(event.request))
    );
});