const colors = require("tailwindcss/colors");

module.exports = {
  purge: ["./src/**/*.svelte", "./src/App.svelte"],
  darkMode: false, // or 'media' or 'class'
  theme: {
    fontFamily: {
      heading: ['"Roboto Slab"'],
      "sm-heading": ["Mada"],
      body: ["Hind"],
    },
    colors: {
      orange: {
        DEFAULT: "#ff3e00",
      },
      transparent: "transparent",
      black: colors.black,
      white: colors.white,
      gray: colors.trueGray,
      indigo: colors.indigo,
      red: colors.rose,
      green: colors.emerald,
      blue: colors.blue,
      yellow: colors.amber,
    },
    extend: {},
  },
  variants: {
    extend: {},
  },
  plugins: [],
};
