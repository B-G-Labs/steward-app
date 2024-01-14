// https://nuxt.com/docs/api/configuration/nuxt-config

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
