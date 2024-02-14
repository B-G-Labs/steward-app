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
    color: {
      primary: "primary",
      secondary: "secondary", 
      tertiary: "tertiary",
    },
  },
});

type ButtonProps = VariantProps<typeof button>;

withDefaults(
  defineProps<{ type?: ButtonProps["type"]; size?: ButtonProps["size"], color?: ButtonProps["color"] }>(),
  {
    type: "elevated",
    size: "md",
    color: "primary"
  }
);

</script> 

<template>
  <button :class="button({ type, size, color })">
    <slot />
  </button>
</template>

<style scoped>
.button {
  border: none;
  cursor: pointer;
  border-radius: 100px;
}

.button.primary{
  --bg-color: var(--primary-40);
  --text-color: var(--primary-40);
  --surface-color: var(--primary-99);
}

.button.secondary{
  --bg-color: var(--secondary-40);
  --text-color: var(--secondary-40);
  --surface-color: var(--secondary-99);
}

.button.tertiary{
  --bg-color: var(--tertiary-40);
  --text-color: var(--tertiary-40);
  --surface-color: var(--tertiary-99);
}

.flat{
  background-color: transparent;
  color: var(--text-color);
}


.filled{
  background-color: var(--bg-color);
  color: var(--white)
}

.tonal {
  background-color: transparent;
  position: relative;
  z-index: 20;
  overflow: hidden;
}
.tonal::before{
  position: absolute;
  height: 100%;
  width: 100%;
  content: '';
  top: 0;
  left: 0;
  background-color: var(--bg-color);
  opacity: 8%;
  z-index: 1;  
}

.elevated{
  box-shadow: 0px 1px 2px 0px rgba(0, 0, 0, 0.3), 1px 0px 3px 1px rgba(0, 0, 0, 0.15);
  background-color: var(--surface-color);
  color: var(--text-color)
}
.outlined{
  background-color: transparent;
  border: 1px solid var(--primary-40);
  color: var(--primary-40);
}

.md {
  font-size: 1.4rem;
  padding: 1rem 2.4rem;
}
.lg {
  background-color: cadetblue;
}
</style>
