import plugin from "tailwindcss/plugin";
import forms from "@tailwindcss/forms";

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./view/*"],
  theme: {
    extend: {},
  },
  plugins: [
    forms,
    plugin(function ({ addUtilities }) {
      addUtilities({
        ".debug": {
          border: "4px solid red",
        },
      });
    }),
  ],
};
