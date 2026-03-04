module.exports = {
  content: [
    './_includes/**/*.html',
    './_layouts/**/*.html',
    './_posts/*.md',
    './*.md',
    './*.html',
    './docker/**/*.md',
    './kubernetes/**/*.md',
    './terraform/**/*.md',
    './helm/**/*.md',
  ],
  theme: {
    extend: {
      fontFamily: {
        sans: ['Instrument Sans', 'system-ui', 'sans-serif'],
        mono: ['JetBrains Mono', 'ui-monospace', 'monospace'],
      },
      colors: {
        primary: {
          50: '#f0f9ff',
          100: '#e0f2fe',
          500: '#0ea5e9',
          600: '#0284c7',
          700: '#0369a1',
        },
        surface: {
          50: '#f8fafc',
          100: '#f1f5f9',
          200: '#e2e8f0',
          800: '#1e293b',
          900: '#0f172a',
        },
      },
      borderRadius: {
        '2xl': '1rem',
      },
      boxShadow: {
        'soft': '0 2px 8px rgba(0,0,0,0.04)',
        'card': '0 1px 3px rgba(0,0,0,0.06)',
      },
      typography: {
        DEFAULT: {
          css: {
            maxWidth: '65ch',
            fontFamily: 'Instrument Sans, system-ui, sans-serif',
            color: 'inherit',
            a: { color: '#0284c7', '&:hover': { color: '#0369a1' } },
            'h1,h2,h3,h4': { fontFamily: 'Instrument Sans, system-ui, sans-serif', fontWeight: '600' },
            code: { fontFamily: 'JetBrains Mono, monospace', fontWeight: '400', backgroundColor: '#f1f5f9', padding: '0.2em 0.4em', borderRadius: '0.25rem' },
            'code::before, code::after': { content: '""' },
          },
        },
      },
    },
  },
  plugins: [
    require('@tailwindcss/typography'),
  ],
} 