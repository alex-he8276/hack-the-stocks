/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/**/*.{js,jsx,ts,tsx}",
  ],
  
  theme: {
    extend: {
      backgroundImage: {
        'image1': "url('/public/images/image1.jpg')",
        'image2': "url('/public/images/image2.jpg')",
      }
    },
  },
  plugins: [],
}
