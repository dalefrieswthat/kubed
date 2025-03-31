// Initialize Prism.js syntax highlighting
document.addEventListener('DOMContentLoaded', function() {
    Prism.highlightAll();
    initAnimations();
    initMobileMenu();
    initCopyButtons();
    initCardHovers();
});

// Initialize animations
function initAnimations() {
    // Animated title
    const titleEl = document.querySelector('.animate-title span');
    if (titleEl) {
        titleEl.style.animationDelay = '0.3s';
    }

    // Add scroll reveal animations
    const elementsToAnimate = document.querySelectorAll('.feature-card, .cta, .warning, .footer-cta');
    
    const observer = new IntersectionObserver((entries) => {
        entries.forEach(entry => {
            if (entry.isIntersecting) {
                entry.target.classList.add('animated');
                observer.unobserve(entry.target);
            }
        });
    }, {
        threshold: 0.1
    });
    
    elementsToAnimate.forEach(el => {
        el.classList.add('animate-on-scroll');
        observer.observe(el);
    });
}

// Add copy button to code blocks
function initCopyButtons() {
    const codeBlocks = document.querySelectorAll('pre code');
    
    codeBlocks.forEach(block => {
        const button = document.createElement('button');
        button.className = 'copy-button';
        button.innerHTML = '<i class="far fa-clipboard"></i>';
        button.setAttribute('aria-label', 'Copy code to clipboard');
        
        const pre = block.parentNode;
        pre.style.position = 'relative';
        pre.appendChild(button);
        
        button.addEventListener('click', async () => {
            try {
                await navigator.clipboard.writeText(block.textContent);
                button.innerHTML = '<i class="fas fa-check"></i>';
                button.classList.add('success');
                
                setTimeout(() => {
                    button.innerHTML = '<i class="far fa-clipboard"></i>';
                    button.classList.remove('success');
                }, 2000);
            } catch (err) {
                console.error('Failed to copy text: ', err);
                button.innerHTML = '<i class="fas fa-times"></i>';
                button.classList.add('error');
                
                setTimeout(() => {
                    button.innerHTML = '<i class="far fa-clipboard"></i>';
                    button.classList.remove('error');
                }, 2000);
            }
        });
    });
}

// Mobile menu toggle
function initMobileMenu() {
    const menuToggle = document.getElementById('menu-toggle');
    const mobileMenu = document.getElementById('mobile-menu');
    
    if (menuToggle && mobileMenu) {
        menuToggle.addEventListener('click', function() {
            mobileMenu.classList.toggle('hidden');
        });
    }
    
    // Add animation classes to elements when they come into view
    const animateElements = document.querySelectorAll('.feature-card, .cta, .footer-cta');
    
    const observer = new IntersectionObserver((entries) => {
        entries.forEach(entry => {
            if (entry.isIntersecting) {
                entry.target.classList.add('animate-in');
                observer.unobserve(entry.target);
            }
        });
    }, {
        threshold: 0.1
    });
    
    animateElements.forEach(element => {
        element.classList.add('opacity-0', 'translate-y-4', 'transition-all', 'duration-500');
        observer.observe(element);
    });
    
    // Add CSS class for animation when visible
    document.head.insertAdjacentHTML('beforeend', `
        <style>
            .animate-in {
                opacity: 1 !important;
                transform: translateY(0) !important;
            }
        </style>
    `);
}

// Add card hover effects
function initCardHovers() {
    const cards = document.querySelectorAll('.feature-card');
    
    cards.forEach(card => {
        card.addEventListener('mouseenter', () => {
            cards.forEach(otherCard => {
                if (otherCard !== card) {
                    otherCard.classList.add('dimmed');
                }
            });
        });
        
        card.addEventListener('mouseleave', () => {
            cards.forEach(otherCard => {
                otherCard.classList.remove('dimmed');
            });
        });
    });
} 