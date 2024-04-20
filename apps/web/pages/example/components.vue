<template>
  <div class="content-container">
    <div class="components-container">
      <div class="btn-color">
        <smexy-button size="xl">Button xl</smexy-button>
        <smexy-button size="lg">Button lg</smexy-button>
        <smexy-button size="md">Button md</smexy-button>
        <smexy-button size="sm">Button sm</smexy-button>
      </div>

      <div class="btn-color">
        <smexy-button color="secondary" size="xl">Button xl</smexy-button>
        <smexy-button color="secondary" size="lg">Button lg</smexy-button>
        <smexy-button color="secondary" size="md">Button md</smexy-button>
        <smexy-button color="secondary" size="sm">Button sm</smexy-button>
      </div>

      <div class="btn-color">
        <smexy-button color="tertiary" size="xl">Button xl</smexy-button>
        <smexy-button color="tertiary" size="lg">Button lg</smexy-button>
        <smexy-button color="tertiary" size="md">Button md</smexy-button>
        <smexy-button color="tertiary" size="sm">Button sm</smexy-button>
      </div>
    </div>

    <div class="components-container">
      <div class="btn-types">
        <smexy-button type="filled" icon="person" icon-right="notifications" loading>Fill</smexy-button>
        <smexy-button type="outlined">Outline</smexy-button>
        <smexy-button type="flat">Flat</smexy-button>
        <smexy-button type="elevated">Elevated</smexy-button>
        <smexy-button type="tonal">Tonal</smexy-button>
        <smexy-button type="tonal" disabled>Tonal</smexy-button>
      </div>

      <div class="btn-types">
        <smexy-button color="secondary" type="filled">Fill</smexy-button>
        <smexy-button color="secondary" type="outlined">Outline</smexy-button>
        <smexy-button color="secondary" type="flat">Flat</smexy-button>
        <smexy-button color="secondary" type="elevated">Elevated</smexy-button>
        <smexy-button color="secondary" type="tonal">Tonal</smexy-button>
        <smexy-button color="secondary" type="tonal" disabled>Tonal</smexy-button>

      </div>

      <div class="btn-types">
        <smexy-button color="tertiary" type="filled">Fill</smexy-button>
        <smexy-button color="tertiary" type="outlined">Outline</smexy-button>
        <smexy-button color="tertiary" type="flat">Flat</smexy-button>
        <smexy-button color="tertiary" type="elevated">Elevated</smexy-button>
        <smexy-button color="tertiary" type="tonal">Tonal</smexy-button>
        <smexy-button color="tertiary" type="tonal" disabled>Tonal</smexy-button>
      </div>
    </div>


    <div class="components-container">
      <smexy-navigation-drawer />
    </div>


    <div class="components-container">
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
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const loading = ref(true)

function toggle() {
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
