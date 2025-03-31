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
    const nav = document.querySelector('nav');
    const navLinks = document.querySelector('.nav-links');
    const menuButton = document.querySelector('.mobile-menu-button');
    
    if (!menuButton) return;
    
    // Toggle mobile menu
    menuButton.addEventListener('click', (e) => {
        e.stopPropagation();
        navLinks.classList.toggle('show');
        
        // Toggle icon
        const icon = menuButton.querySelector('i');
        if (navLinks.classList.contains('show')) {
            icon.classList.remove('fa-bars');
            icon.classList.add('fa-times');
        } else {
            icon.classList.remove('fa-times');
            icon.classList.add('fa-bars');
        }
    });
    
    // Close mobile menu when clicking outside
    document.addEventListener('click', (e) => {
        if (!nav.contains(e.target) && navLinks.classList.contains('show')) {
            navLinks.classList.remove('show');
            const icon = menuButton.querySelector('i');
            icon.classList.remove('fa-times');
            icon.classList.add('fa-bars');
        }
    });
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