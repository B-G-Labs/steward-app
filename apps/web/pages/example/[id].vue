<template>
  <div>
    <div class="components-container">
      <smexy-button size="lg">Big</smexy-button>
      <smexy-button>Medium</smexy-button>
      <smexy-button>Small</smexy-button>
    </div>

    <div class="components-container">
      <smexy-button type="filled">Fill</smexy-button>
      <smexy-button type="outlined">Outline</smexy-button>
      <smexy-button type="flat">Flat</smexy-button>
      <smexy-button type="elevated">Elevated</smexy-button>
      <smexy-button type="tonal">Tonal</smexy-button>
    </div>

    <div>
      <smexy-button>Primary</smexy-button>
      <smexy-button>Primary small</smexy-button>
    </div>

    <div>
      <smexy-button>Secondary</smexy-button>
      <smexy-button>Secondary small</smexy-button>
    </div>

    <div>
      <smexy-button>butt</smexy-button>
      <smexy-button>butter</smexy-button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRoute } from "vue-router";

const route = useRoute();
</script>

<style lang="scss">

.components-container {
  padding: 1rem;
  display: flex;
  align-items: center;
  gap: 1rem;
}

</style>
