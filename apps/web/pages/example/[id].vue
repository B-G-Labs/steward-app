<template>
  <div>
    <div class="components-container">
      <smexy-button size="big">Big</smexy-button>
      <smexy-button size="medium">Medium</smexy-button>
      <smexy-button size="small">Small</smexy-button>
    </div>

    <div class="components-container">
      <smexy-button size="medium" btnStyle="fill">Fill</smexy-button>
      <smexy-button size="medium">Outline</smexy-button>
      <smexy-button size="medium">Flat</smexy-button>
    </div>

    <div>
      <smexy-button style="margin: 10px">Primary</smexy-button>
      <smexy-button size="small">Primary small</smexy-button>
    </div>

    <div>
      <smexy-button style="margin: 10px" type="secondary">Secondary</smexy-button>
      <smexy-button type="secondary" size="small">Secondary small</smexy-button>
    </div>

    <div>
      <smexy-button style="margin: 10px" type="primary" size="medium">butt</smexy-button>
      <smexy-button size="small" class="primaryMedium">butter</smexy-button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRoute } from "vue-router";

const route = useRoute();
</script>

<style scoped>
.components-container {
  margin: 1rem; 
  display: flex; 
  align-items: center;
  gap: 1rem; 
}
</style>
