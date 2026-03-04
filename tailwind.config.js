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
        accent: {
          DEFAULT: '#2563eb',
          hover: '#1d4ed8',
          light: '#eff6ff',
          border: '#bfdbfe',
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
            color: '#3f3f46',
            a: { color: '#2563eb', '&:hover': { color: '#1d4ed8' } },
            'h1,h2,h3,h4': { fontFamily: 'Instrument Sans, system-ui, sans-serif', fontWeight: '600', color: '#18181b' },
            code: { fontFamily: 'JetBrains Mono, monospace', fontWeight: '400', backgroundColor: '#f4f4f5', padding: '0.2em 0.4em', borderRadius: '0.25rem', color: '#18181b' },
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