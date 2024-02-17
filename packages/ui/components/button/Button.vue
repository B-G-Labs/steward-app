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
    }
  },
});

type ButtonProps = VariantProps<typeof button>;

withDefaults(
  defineProps<{ 
    type?: ButtonProps["type"],
    size?: ButtonProps["size"], 
    color?: ButtonProps["color"], 
    disabled?: boolean
  }>(),
  {
    type: "elevated",
    size: "md",
    color: "primary"
  }
);

</script> 

<template>
  <button :class="button({ type, size, color })" :disabled="disabled">
    <slot />
  </button>
</template>

<style scoped>
.button {
  border: none;
  cursor: pointer;
  border-radius: 100px;
  width: max-content;
  height: max-content;
  position: relative;
  /* overflow: hidden; */
}

.button::before{
  content: '';
  position: absolute;
  border-radius: 100px;
  background-color: transparent;
  top: 0;
  right: 0;
  width: 100%;
  height: 100%;
  opacity: 8%;
  transition: .3s;
}

.button.primary{
  --bg-color: var(--primary-40);
  --text-color: var(--primary-40);
  --border-color: var(--primary-40);
  --surface-color: var(--primary-99);

  --disbaled--surface-color: var(--primary-40);
  --disbaled--text-color: var(--primary-10);
}

.button.secondary{
  --bg-color: var(--secondary-40);
  --text-color: var(--secondary-40);
  --border-color: var(--secondary-40);
  --surface-color: var(--secondary-99);

  --disbaled--surface-color: var(--primary-40);
  --disbaled--text-color: var(--secondary-10);
}

.button.tertiary{
  --bg-color: var(--tertiary-40);
  --text-color: var(--tertiary-40);
  --border-color: var(--tertiary-40);
  --surface-color: var(--tertiary-99);

  --disbaled--surface-color: var(--primary-40);
  --disbaled--text-color: var(--tertiary-10);
}

.filled{
  background-color: var(--bg-color);
  color: var(--white);
}

.filled:hover::before{
  background-color: rgba(255, 255, 255);
}
.filled:is(:focus, :active)::before{
  opacity: 12%;
}

.flat{
  background-color: transparent;
  color: var(--text-color);
}

.elevated{
  box-shadow: 0px 1px 2px 0px rgba(0, 0, 0, 0.3), 1px 0px 3px 1px rgba(0, 0, 0, 0.15);
  background-color: var(--surface-color);
  color: var(--text-color)
}
.outlined{
  background-color: transparent;
  border: 1px solid var(--border-color);
  color: var(--text-color);
}

:is(.outlined, .flat, .elevated):hover::before{
  background-color: var(--bg-color);
}

:is(.outlined, .flat, .elevated):is(:focus, :active)::before{
  opacity: 12%;
  background-color: var(--bg-color);
}

.tonal {
  background-color: transparent;
}
.tonal::before{
  background-color: var(--bg-color);
  /* opacity: 8%; */
}
.sm {
  font-size: 1.2rem;
  line-height: 1.6rem;
  padding: .6rem 2rem;
}

.md {
  font-size: 1.4rem;
  line-height: 2rem;
  padding: 1rem 2.4rem;
}
.lg {
  font-size: 1.6rem;
  line-height: 2.4rem;
  padding: 1.4rem 2.8rem;
}

.xl {
  font-size: 1.8rem;
  line-height: 2.8rem;
  padding: 1.8rem 3.2rem;
}

.button:disabled{
  color: var(--disbaled--text-color);
  opacity: 20%;
  cursor: not-allowed;
}

.button:disabled::before{
  background-color: var(--disbaled--surface-color);
  opacity: 25%;
}
</style>
