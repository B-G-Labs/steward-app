// https://nuxt.com/docs/api/configuration/nuxt-config

import { browserslistToTargets } from "lightningcss";
import browserslist from "browserslist";

export default defineNuxtConfig({
  devtools: { enabled: true },
  app: {
    head: {
      title: "App",
      link: [
        {
          href: "https://fonts.googleapis.com/css2?family=Material+Icons", 
          rel: "stylesheet" 
        }
      ]
    },
  },
  modules: [
    "@pinia/nuxt",
    [
      "@nuxt-modules/compression",
      {
        algorithm: "brotliCompress",
      },
    ],
    "ui",
    "@nuxtjs/color-mode"
  ],
  vite: {
    css: {
      transformer: "lightningcss",
      lightningcss: {
        targets: browserslistToTargets(browserslist(">= 0.25%")),
      },
    },
    build: {
      cssMinify: "lightningcss",
    },
    plugins: [],
    vueJsx: {
      mergeProps: true,
    },
  },  
  css: ["./assets/css/utilities.css"],
});
