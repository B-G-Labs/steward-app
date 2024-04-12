<template>
  <div class="content-container">
    <smexy-circular size="sm"></smexy-circular>
    <smexy-circular></smexy-circular>
    <smexy-circular size="lg"></smexy-circular>
    <smexy-circular size="xl"></smexy-circular>

    <smexy-button :loading="loading" size="sm">Button</smexy-button>
    <smexy-button :loading="loading">Button</smexy-button>
    <smexy-button :loading="loading" size="lg">Button</smexy-button>
    <smexy-button :loading="loading" size="xl">Button</smexy-button>
    <smexy-button @click="toggle">Button</smexy-button>
  </div>
</template>

<script setup lang="ts">
import {ref} from 'vue'

const loading = ref(true)
function toggle () {
  loading.value = !loading.value
}
</script>

<style lang="scss">
.content-container {
  display: flex;
  gap: 2.5rem;
}

.components-container {
  background-color: var(--white);
  border-radius: 50px;
  width: fit-content;
  padding: 2.5rem;
  display: flex;
  flex-direction: column;
  gap: 2.5rem;
  flex-wrap: wrap;

  .btn-color,
  .btn-types {
    display: flex;
    gap: 1.5rem;
    align-items: flex-end;
  }
}
</style>
