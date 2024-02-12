// https://nuxt.com/docs/api/configuration/nuxt-config

import {browserslistToTargets} from 'lightningcss';
import browserslist from 'browserslist'

export default defineNuxtConfig({
  devtools: { enabled: true },
  app: {
    head: {
      title: "App",
    },
  },
  modules: [
    [
      "@nuxt-modules/compression",
      {
        algorithm: "brotliCompress",
      },
    ],
    "ui",
  ],
  vite: {
    css: {
      transformer: 'lightningcss',
      lightningcss: {
        targets: browserslistToTargets(browserslist('>= 0.25%'))
      }
    },
    build: {
      cssMinify: 'lightningcss'
    },
    plugins: [],
    vue: {
      customElement: true,
    },
    vueJsx: {
      mergeProps: true,
    },
  },
  css: ["./assets/css/utilities.css"],
});
