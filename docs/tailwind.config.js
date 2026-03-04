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
      colors: {
        primary: {
          light: '#4A89DC',
          DEFAULT: '#3B78CC',
          dark: '#2D68B2',
        },
        secondary: {
          light: '#6A50A7',
          DEFAULT: '#5A4090',
          dark: '#4A3078',
        },
      },
      typography: {
        DEFAULT: {
          css: {
            maxWidth: '65ch',
            color: 'inherit',
            a: {
              color: '#3B78CC',
              '&:hover': {
                color: '#2D68B2',
              },
            },
            'h1,h2,h3,h4': {
              color: 'inherit',
              fontWeight: '700',
            },
            code: {
              color: '#24292e',
              fontWeight: '400',
              backgroundColor: '#f6f8fa',
              padding: '0.2em 0.4em',
              borderRadius: '0.25rem',
            },
            'code::before': {
              content: '""',
            },
            'code::after': {
              content: '""',
            },
          },
        },
      },
    },
  },
  plugins: [
    require('@tailwindcss/typography'),
  ],
} 