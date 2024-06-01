document.addEventListener('DOMContentLoaded', function() {
    function loadPage(page) {
        fetch(`pages/${page}.html`)
            .then(response => response.text())
            .then(data => {
                document.getElementById('content').innerHTML = data;
            })
            .catch(error => {
                console.error('Error loading page:', error);
                document.getElementById('content').innerHTML = '<p>Страница не найдена.</p>';
            });
    }

    document.querySelectorAll('nav a').forEach(link => {
        link.addEventListener('click', function(e) {
            e.preventDefault();
            const page = e.target.getAttribute('data-page');
            history.pushState(null, '', e.target.href);
            loadPage(page);
        });
    });

    window.addEventListener('popstate', function() {
        const path = window.location.pathname.slice(1) || 'home';
        loadPage(path);
    });

    // Load the initial page
    const initialPage = window.location.pathname.slice(1) || 'home';
    loadPage(initialPage);
});
