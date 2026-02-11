if ('serviceWorker' in navigator) {
    window.addEventListener('load', () => {
        navigator.serviceWorker.register('/worker.js')
            .then(registration => {
                console.log('ServiceWorker registered:', registration);
            })
            .catch(error => {
                console.log('ServiceWorker registration failed:', error);
            });
    });
}

let deferredPrompt;
window.addEventListener('beforeinstallprompt', (e) => {
    e.preventDefault();
    deferredPrompt = e;
});

document.body.addEventListener('htmx:afterSwap', (event) => {
    event.detail.elt.querySelectorAll('.animate-slide-up').forEach((el, i) => {
        el.style.animationDelay = `${i * 0.1}s`;
    });
});