import { createResolver, defineNuxtModule } from "@nuxt/kit";
import { join } from "path";

export default defineNuxtModule({
  setup(_, nuxt) {
    nuxt.hook("components:dirs", (dirs) => {
      const { resolve } = createResolver(import.meta.url)

      dirs.push({
        path: resolve('./components'),
        prefix: 'smexy',
        // enabled: true,
        // global: true
      });
    });
  },
});
