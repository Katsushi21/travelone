const defaultTheme = require('tailwindcss/defaultTheme');

module.exports = {
  mode: 'jit',
  content: [
    './src/pages/**/*.{js,ts,jsx,tsx}',
    './src/components/**/*.{js,ts,jsx,tsx}',
  ],
  content: [
    './src/pages/**/*.{js,ts,jsx,tsx}',
    './src/components/**/*.{js,ts,jsx,tsx}',
  ],
  theme: {
    fontFamily: {
      ja: [...defaultTheme.fontFamily.sans],
      en: [...defaultTheme.fontFamily.sans],
    },
    extend: {
      colors: {
        theme: {
          light: '#ffffff',
          medium: '#cccccc',
          DEFAULT: '#242424',
          dark: '#111111',
        },
        primary: {
          // light: '',
          // medium: '',
          DEFAULT: '#242424',
          // dark: '',
        },
      },
    },
  },
  plugins: [],
};
