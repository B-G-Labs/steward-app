<template>
  <div>
    <div class="components-container">
      <smexy-button>Big</smexy-button>
      <smexy-button>Medium</smexy-button>
      <smexy-button>Small</smexy-button>
    </div>

    <div class="components-container">
      <smexy-button btnStyle="fill">Fill</smexy-button>
      <smexy-button>Outline</smexy-button>
      <smexy-button>Flat</smexy-button>
    </div>

    <div>
      <smexy-button style="margin: 10px">Primary</smexy-button>
      <smexy-button>Primary small</smexy-button>
    </div>

    <div>
      <smexy-button style="margin: 10px">Secondary</smexy-button>
      <smexy-button>Secondary small</smexy-button>
    </div>

    <div>
      <smexy-button style="margin: 10px">butt</smexy-button>
      <smexy-button class="primaryMedium">butter</smexy-button>
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
