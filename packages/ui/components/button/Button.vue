<script setup lang="ts">
import { cva, type VariantProps } from "cva";
import Circular from "../circular/Circular.vue";

const button = cva("button", {
  variants: {
    type: {
      flat: "btn-flat",
      filled: "btn-filled",
      tonal: "btn-tonal",
      elevated: "btn-elevated",
      outlined: "btn-outlined",
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
  defineProps<{
    type?: ButtonProps["type"];
    size?: ButtonProps["size"];
    color?: ButtonProps["color"];
    disabled?: boolean;
    loading?: boolean;
    icon?: string;
    iconRight?: string;
  }>(),
  {
    type: "elevated",
    size: "md",
    color: "primary",
  }
);
</script>

<template>
  <button :class="button({ type, size, color })" :disabled="disabled">
    <div v-if="loading" class="btn-loading-state-container">
      <Circular size="md"/>
    </div>
    <article>
      <span class="material-icons">{{icon}}</span>
      <slot />
      <span class="material-icons">{{iconRight}}</span>
    </article>
      
    <!-- <slot v-else name="loading">
      
      <div style="display: flex; gap: 0.5rem">
        <Circular size="sm"/>
        Loading...
      </div>
    </slot> -->
  </button>
</template>

<style>
.button {
  border: none;
  cursor: pointer;
  border-radius: 100px;
  width: max-content;
  height: max-content;
  position: relative;

  display: flex;
  gap: .5rem;
  transition: 0.3s;
}

.button::before {
  content: "";
  position: absolute;
  border-radius: 100px;
  background-color: transparent;
  top: 0;
  right: 0;
  width: 100%;
  height: 100%;
  opacity: 8%;
  transition: 0.3s;
}

.btn-filled {
  background-color: var(--bg-color);
  color: var(--white);
}

.btn-filled:hover::before {
  background-color: rgba(255, 255, 255);
}
.btn-filled:is(:focus, :active)::before {
  opacity: 12%;
}

.btn-flat {
  background-color: transparent;
  color: var(--text-color);
}

.btn-elevated {
  box-shadow:
    0px 1px 2px 0px rgba(0, 0, 0, 0.3),
    1px 0px 3px 1px rgba(0, 0, 0, 0.15);
  background-color: var(--surface-color);
  color: var(--text-color);
}

.btn-outlined {
  background-color: transparent;
  border: 1px solid var(--border-color);
  color: var(--text-color);
}

:is(.btn-outlined, .btn-flat, .btn-elevated):hover::before {
  background-color: var(--bg-color);
}

:is(.btn-outlined, .btn-flat, .btn-elevated):is(:focus, :active)::before {
  opacity: 12%;
  background-color: var(--bg-color);
}

.btn-tonal {
  background-color: transparent;
}

.btn-tonal::before {
  background-color: var(--bg-color);
}

.btn-loading-state-container{
  display: flex;
  align-items: center;
  position: absolute;
  left: 5%;
  top: 30%;
  height: 40%;
  width: 20%;

  svg {
    width: 100% !important;
    height: 100% !important;
  }
}

.button:disabled {
  color: var(--disbaled--text-color);
  opacity: 20%;
  cursor: not-allowed;
}

.button:disabled::before {
  background-color: var(--disbaled--surface-color);
  opacity: 25%;
}

.button:has(.btn-loading-state-container){
  color: var(--disbaled--text-color);
  opacity: 20%;
  cursor: not-allowed;
  
  &.sm {
    padding-left: 2.2rem;
    padding-right: 1.8rem;
  }

  &.md {
    padding-left: 2.6rem; 
    padding-right: 2.2rem;

  }

  &.lg {
    padding-left: 3rem; 
    padding-right: 2.6rem;
  }

  &.xl {
    padding-left: 3.4rem; 
    padding-right: 3rem;
  }
}
</style>
