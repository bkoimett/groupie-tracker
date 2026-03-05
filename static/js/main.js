/**
 * Groupie Tracker - Main JavaScript
 * Handles formatting, animations, and interactive features
 */

// Wait for DOM to load
document.addEventListener('DOMContentLoaded', function() {
    initializeApp();
});

/**
 * Initialize application features
 */
function initializeApp() {
    formatDates();
    formatLocations();
    initializeHoverEffects();
    initializeSearchFilter(); // Only runs on home page if search exists
}

/**
 * Format date strings to readable format
 * Looks for elements with class 'format-date' or dates in tour schedules
 */
function formatDates() {
    // Format dates in tour schedule
    document.querySelectorAll('.tour-date').forEach(el => {
        const originalDate = el.textContent.trim();
        const formattedDate = formatDateString(originalDate);
        if (formattedDate !== originalDate) {
            el.textContent = formattedDate;
            el.setAttribute('data-original', originalDate);
        }
    });

    // Format individual date items
    document.querySelectorAll('.date-item p').forEach(el => {
        const originalDate = el.textContent.trim();
        const formattedDate = formatDateString(originalDate);
        if (formattedDate !== originalDate) {
            el.textContent = formattedDate;
        }
    });
}

/**
 * Format location strings (replace underscores, capitalize)
 */
function formatLocations() {
    // Format location cards
    document.querySelectorAll('.location-card p, .tour-location').forEach(el => {
        const originalLocation = el.textContent.trim();
        const formattedLocation = formatLocationString(originalLocation);
        if (formattedLocation !== originalLocation) {
            el.textContent = formattedLocation;
        }
    });
}

/**
 * Format a date string to "DD MMM YYYY" format
 * @param {string} dateStr - Original date string
 * @returns {string} Formatted date string
 */
function formatDateString(dateStr) {
    if (!dateStr) return '';
    
    // Try to parse the date
    const date = new Date(dateStr);
    if (!isNaN(date.getTime())) {
        return date.toLocaleDateString('en-US', {
            year: 'numeric',
            month: 'short',
            day: 'numeric'
        });
    }
    
    // If it's already in a different format, try to clean it up
    const patterns = [
        /(\d{4})-(\d{2})-(\d{2})/, // YYYY-MM-DD
        /(\d{2})\/(\d{2})\/(\d{4})/, // DD/MM/YYYY
        /(\d{1,2})\s+(\w+)\s+(\d{4})/i // DD Month YYYY
    ];
    
    for (let pattern of patterns) {
        const match = dateStr.match(pattern);
        if (match) {
            return dateStr; // Return original if pattern matches but we can't parse
        }
    }
    
    return dateStr;
}

/**
 * Format location string
 * @param {string} locStr - Original location string (e.g., "london_uk")
 * @returns {string} Formatted location (e.g., "London UK")
 */
function formatLocationString(locStr) {
    if (!locStr) return '';
    
    // Split by underscore or space
    return locStr.split(/[_-\s]+/)
        .map(word => {
            // Handle special cases like "uk" -> "UK", "usa" -> "USA"
            if (word.toLowerCase() === 'uk') return 'UK';
            if (word.toLowerCase() === 'usa') return 'USA';
            if (word.toLowerCase() === 'nyc') return 'NYC';
            if (word.toLowerCase() === 'la') return 'LA';
            
            // Capitalize first letter, rest lowercase
            return word.charAt(0).toUpperCase() + word.slice(1).toLowerCase();
        })
        .join(' ');
}

/**
 * Initialize hover effects for cards
 */
function initializeHoverEffects() {
    // Add smooth transitions for all interactive elements
    const interactiveElements = document.querySelectorAll(
        '.artist-card, .info-card, .location-card, .date-item, .tour-item, .member-item'
    );
    
    interactiveElements.forEach(el => {
        el.addEventListener('mouseenter', function(e) {
            this.style.transition = 'all 0.3s cubic-bezier(0.2, 0, 0, 1)';
        });
    });
}

/**
 * Initialize search and filter functionality (for home page)
 * Only activates if search input exists
 */
function initializeSearchFilter() {
    const searchInput = document.getElementById('artist-search');
    if (!searchInput) return;
    
    const artistCards = document.querySelectorAll('.artist-card');
    const noResults = document.createElement('div');
    noResults.className = 'no-results';
    noResults.style.display = 'none';
    noResults.innerHTML = `
        <div class="error-container" style="margin: 2rem auto; padding: 3rem;">
            <h3 style="font-size: 2rem; margin-bottom: 1rem;">🔍 No artists found</h3>
            <p style="color: var(--text-secondary);">Try adjusting your search terms</p>
        </div>
    `;
    
    const artistGrid = document.querySelector('.artist-grid');
    if (artistGrid) {
        artistGrid.parentNode.insertBefore(noResults, artistGrid.nextSibling);
    }
    
    searchInput.addEventListener('input', function(e) {
        const searchTerm = e.target.value.toLowerCase().trim();
        let visibleCount = 0;
        
        artistCards.forEach(card => {
            const artistName = card.querySelector('h2')?.textContent.toLowerCase() || '';
            const members = card.querySelectorAll('p')?.[2]?.textContent.toLowerCase() || '';
            const year = card.querySelector('p')?.textContent.toLowerCase() || '';
            
            const matches = artistName.includes(searchTerm) || 
                           members.includes(searchTerm) || 
                           year.includes(searchTerm) ||
                           searchTerm === '';
            
            if (matches) {
                card.style.display = 'flex';
                visibleCount++;
            } else {
                card.style.display = 'none';
            }
        });
        
        // Show/hide no results message
        if (visibleCount === 0 && searchTerm !== '') {
            noResults.style.display = 'block';
            artistGrid.style.display = 'none';
        } else {
            noResults.style.display = 'none';
            artistGrid.style.display = 'grid';
        }
    });
}

/**
 * Utility function to add loading animation
 * @param {HTMLElement} element - Element to show loading state
 */
function showLoading(element) {
    if (!element) return;
    
    const originalContent = element.innerHTML;
    element.innerHTML = '<div class="loading-spinner" style="text-align: center; padding: 2rem;"><div class="spinner" style="width: 40px; height: 40px; border: 3px solid var(--border-color); border-top-color: var(--accent-primary); border-radius: 50%; animation: spin 1s linear infinite; margin: 0 auto;"></div><p style="margin-top: 1rem; color: var(--text-secondary);">Loading...</p></div>';
    
    // Add spinner animation if not exists
    if (!document.querySelector('#spinner-style')) {
        const style = document.createElement('style');
        style.id = 'spinner-style';
        style.textContent = `
            @keyframes spin {
                to { transform: rotate(360deg); }
            }
        `;
        document.head.appendChild(style);
    }
    
    return originalContent;
}

/**
 * Smooth scroll to element
 * @param {string} elementId - ID of target element
 */
function smoothScrollTo(elementId) {
    const element = document.getElementById(elementId);
    if (element) {
        element.scrollIntoView({
            behavior: 'smooth',
            block: 'start'
        });
    }
}

// Export functions for use in templates
window.groupieTracker = {
    formatDate: formatDateString,
    formatLocation: formatLocationString,
    smoothScrollTo,
    showLoading
};