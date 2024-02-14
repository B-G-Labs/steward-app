<script setup lang="ts">
import { cva, type VariantProps } from "cva";

const button = cva("button", {
  variants: {
    type: {
      flat: "flat",
      filled: "filled",
      tonal: "tonal",
      elevated: "elevated",
      outlined: "outlined",
    },
    size: {
      sm: "sm",
      md: "md",
      lg: "lg",
      xl: "xl",
    },
  },
});

type ButtonProps = VariantProps<typeof button>;

withDefaults(
  defineProps<{ type?: ButtonProps["type"]; size?: ButtonProps["size"] }>(),
  {
    type: "elevated",
    size: "md",
  }
);
</script>

<template>
  <button :class="button({ type, size })">
    <slot />
  </button>
</template>

<style scoped></style>
