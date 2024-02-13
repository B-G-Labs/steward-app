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

<style scoped>
.button {
  border: none;
  cursor: pointer;
  border-radius: 100px;
}
.flat{
  background-color: transparent;
  color: var(--primary);
}

.filled{
  background-color: inherit;
  color: var(--white)
}

.tonal{
  background-color: #F0F8DE;  
}

.elevated{
  box-shadow: 0px 1px 2px 0px rgba(0, 0, 0, 0.3), 1px 0px 3px 1px rgba(0, 0, 0, 0.15);
}
.outlined{
  background-color: transparent;
  border: 1px solid var(--primary);
  color: var(--primary);
}

.md {
  font-size: 1.4rem;
  padding: 1rem 2.4rem;
}
.lg {
  background-color: cadetblue;
}
</style>
