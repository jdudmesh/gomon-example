/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.js", "../views/*.html"],
  darkMode: "class",
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
}

