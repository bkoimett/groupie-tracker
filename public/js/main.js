/**
 * Groupie Tracker - Lightweight Theme Manager
 * Optimized for performance - no blocking operations
 */

// Theme Manager - Immediately invoked to prevent flicker
(function() {
    const theme = localStorage.getItem('theme') || 'light';
    document.documentElement.setAttribute('data-theme', theme);
})();

// Wait for DOM to be ready without blocking
if (document.readyState === 'loading') {
    document.addEventListener('DOMContentLoaded', init);
} else {
    init();
}

function init() {
    initThemeToggle();
    initSearch();
    initFormatting();
}

// Theme Toggle - Lightweight
function initThemeToggle() {
    const headerInner = document.querySelector('.header-inner');
    if (!headerInner || document.querySelector('.theme-toggle')) return;

    const currentTheme = document.documentElement.getAttribute('data-theme') || 'light';
    
    const themeToggle = document.createElement('div');
    themeToggle.className = 'theme-toggle';
    themeToggle.innerHTML = `
        <button class="theme-btn ${currentTheme === 'light' ? 'active' : ''}" data-theme="light">
            <span>☀️</span> Light
        </button>
        <button class="theme-btn ${currentTheme === 'dark' ? 'active' : ''}" data-theme="dark">
            <span>🌙</span> Dark
        </button>
    `;

    const logo = document.querySelector('.logo');
    if (logo) {
        logo.after(themeToggle);
    } else {
        headerInner.appendChild(themeToggle);
    }

    // Event delegation for better performance
    themeToggle.addEventListener('click', (e) => {
        const btn = e.target.closest('.theme-btn');
        if (!btn) return;
        
        const theme = btn.dataset.theme;
        document.documentElement.setAttribute('data-theme', theme);
        localStorage.setItem('theme', theme);
        
        document.querySelectorAll('.theme-btn').forEach(b => {
            b.classList.toggle('active', b.dataset.theme === theme);
        });
    });
}

// Search with debounce for performance
function initSearch() {
    const searchInput = document.getElementById('artist-search');
    if (!searchInput) return;

    const artistCards = document.querySelectorAll('.artist-card');
    const artistGrid = document.querySelector('.artist-grid');
    if (!artistGrid) return;

    let timeout;
    const noResultsDiv = createNoResultsElement();
    artistGrid.parentNode.insertBefore(noResultsDiv, artistGrid.nextSibling);

    searchInput.addEventListener('input', (e) => {
        clearTimeout(timeout);
        timeout = setTimeout(() => filterArtists(e.target.value.trim().toLowerCase()), 200);
    });

    function filterArtists(searchTerm) {
        let visibleCount = 0;
        
        artistCards.forEach(card => {
            const text = card.textContent.toLowerCase();
            const matches = searchTerm === '' || text.includes(searchTerm);
            
            card.style.display = matches ? 'block' : 'none';
            if (matches) visibleCount++;
        });

        noResultsDiv.style.display = (visibleCount === 0 && searchTerm !== '') ? 'block' : 'none';
        artistGrid.style.display = (visibleCount === 0 && searchTerm !== '') ? 'none' : 'grid';
    }

    function createNoResultsElement() {
        const div = document.createElement('div');
        div.className = 'no-results';
        div.style.display = 'none';
        div.innerHTML = `
            <div class="error-container" style="margin: 2rem auto;">
                <h3 style="font-size: 2rem; margin-bottom: 1rem;">🔍 No artists found</h3>
                <p style="color: var(--text-secondary);">Try a different search term</p>
            </div>
        `;
        return div;
    }
}

// Format dates and locations - Optimized with batch processing
function initFormatting() {
    // Use requestIdleCallback for non-critical formatting
    if ('requestIdleCallback' in window) {
        requestIdleCallback(formatDates, { timeout: 1000 });
        requestIdleCallback(formatLocations, { timeout: 1000 });
    } else {
        setTimeout(formatDates, 50);
        setTimeout(formatLocations, 50);
    }
}

function formatDates() {
    const monthMap = {
        '01': 'January', '02': 'February', '03': 'March', '04': 'April',
        '05': 'May', '06': 'June', '07': 'July', '08': 'August',
        '09': 'September', '10': 'October', '11': 'November', '12': 'December'
    };

    document.querySelectorAll('.tour-date, .date-item p, .format-date').forEach(el => {
        let date = el.textContent.trim().replace(/^\*+/, '');
        const parts = date.split('-');
        
        if (parts.length === 3 && monthMap[parts[1]]) {
            el.textContent = `${monthMap[parts[1]]} ${parts[2]}, ${parts[0]}`;
        }
    });
}

function formatLocations() {
    document.querySelectorAll('.location-card p, .tour-location, .format-location').forEach(el => {
        let loc = el.textContent.trim()
            .replace(/_/g, ' ')
            .replace(/-/g, ', ')
            .replace(/\s+/g, ' ')
            .trim();
        
        el.textContent = loc.split(' ').map(word => 
            word.length > 0 ? word[0].toUpperCase() + word.slice(1).toLowerCase() : word
        ).join(' ');
    });
}