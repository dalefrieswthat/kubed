// Initialize Prism.js syntax highlighting
document.addEventListener('DOMContentLoaded', function() {
    Prism.highlightAll();
});

// Add copy button to code blocks
document.addEventListener('DOMContentLoaded', function() {
    const codeBlocks = document.querySelectorAll('pre code');
    
    codeBlocks.forEach(block => {
        const button = document.createElement('button');
        button.className = 'copy-button';
        button.textContent = 'Copy';
        
        const pre = block.parentNode;
        pre.style.position = 'relative';
        pre.appendChild(button);
        
        button.addEventListener('click', async () => {
            try {
                await navigator.clipboard.writeText(block.textContent);
                button.textContent = 'Copied!';
                setTimeout(() => {
                    button.textContent = 'Copy';
                }, 2000);
            } catch (err) {
                console.error('Failed to copy text: ', err);
                button.textContent = 'Failed';
            }
        });
    });
});

// Add mobile menu toggle
document.addEventListener('DOMContentLoaded', function() {
    const nav = document.querySelector('nav');
    const navLinks = document.querySelector('.nav-links');
    
    // Create mobile menu button
    const menuButton = document.createElement('button');
    menuButton.className = 'mobile-menu-button';
    menuButton.innerHTML = '☰';
    nav.insertBefore(menuButton, navLinks);
    
    // Toggle mobile menu
    menuButton.addEventListener('click', () => {
        navLinks.classList.toggle('show');
    });
    
    // Close mobile menu when clicking outside
    document.addEventListener('click', (e) => {
        if (!nav.contains(e.target) && navLinks.classList.contains('show')) {
            navLinks.classList.remove('show');
        }
    });
}); 