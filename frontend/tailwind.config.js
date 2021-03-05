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
    },
    extend: {},
  },
  variants: {
    extend: {},
  },
  plugins: [],
};
