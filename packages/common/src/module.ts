import { defineNuxtModule, addPlugin, createResolver, addImports } from '@nuxt/kit'

// Module options TypeScript interface definition
export interface ModuleOptions {}

export default defineNuxtModule<ModuleOptions>({
  meta: {
    name: 'common',
    configKey: 'common'
  },
  // Default configuration options of the Nuxt module
  defaults: {},
  setup (options, nuxt) {
    const {resolve} = createResolver(import.meta.url)

    // Do not add the extension since the `.ts` will be transpiled to `.mjs` after `npm run prepack`
    addPlugin(resolve('./runtime/plugin'))

    addImports({
      name: "useApi",
      as: "useApi",
      from: resolve('runtime/composables/useApi') 
    })

    addImports({
      name: "useAuth",
      as: "useAuth",
      from: resolve('runtime/composables/useAuth') 
    })
  }
})
