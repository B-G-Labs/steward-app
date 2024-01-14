<script setup lang="ts">
import { cva, type VariantProps } from "class-variance-authority";

//types
// - filled
// - outline
// - elevated
// - with icon
// - without icon

const button = cva("button", {
  variants: {
    type: {
      primary: "primary",
      secondary: "secondary",
    },
    size: {
      small: "small",
      medium: "medium",
      big: "big"
    },
    btnStyle: {
      fill: "fill",
      outline: "outline",
      flat: "flat"
    }
  },
  compoundVariants: [
    { type: "primary", size: "medium", class: "primaryMedium" },
  ],
});

type ButtonProps = VariantProps<typeof button>;

withDefaults(
  defineProps<{ type?: ButtonProps["type"]; size?: ButtonProps["size"]; btnStyle?: ButtonProps["btnStyle"]}>(),
  {
    type: "primary",
    size: "medium",
    btnStyle: "fill"
  }
);
</script>

<template>
  <button :class="button({ type, size, btnStyle })">
    <slot />
  </button>
</template>

<style scoped>
.button {
  display: inline-flex;
  border-width: 1px;
  border-radius: 6px;
  border-style: solid;
  cursor: pointer;
  height: fit-content;
}


.primary {
  color: rgb(255 255 255);
  background-color: rgb(33, 33, 34);
  border: transparent;
}

.primary:hover {
  background-color: rgb(93, 93, 93);  
}

.secondary {
  background-color: rgb(255 255 255);
  color: rgb(31 41 55);
  border-color: rgb(156 163 175);
}

.secondary:hover {
  background-color: rgb(243 244 246);
}

.small {
  font-size: 1.4rem /* 14px */;
  line-height: 2rem /* 20px */;
  padding: 0.25rem 0.5rem;
}

.medium {
  font-size: 1.6rem /* 16px */;
  line-height: 2.4rem /* 24px */;
  padding: 0.5rem 1rem;
}

.big {
  font-size: 2.4rem;
  line-height: 2.8rem; 
  padding: 0.75rem 1.5rem;
}

.fill {
  background-color: crimson !important;
}

</style>
