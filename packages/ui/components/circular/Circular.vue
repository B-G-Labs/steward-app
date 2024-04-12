<script setup lang="ts">
import { cva, type VariantProps } from "cva";

const circular = cva("circular", {
  variants: {
    type: {
      single: 'single',
      triple: 'triple'
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

type CircularProps = VariantProps<typeof circular>;

withDefaults(
  defineProps<{
    type?: CircularProps['type'];
    size?: CircularProps['size'];
    color?: CircularProps['color'];
  }>(),
  {
    type: 'single',
    size: "md",
    color: "primary",
  }
);

const sizeInPixels = 50
const strokeWidth = 20

</script>

<template>
  <div class="loader-container" :class="circular({size, color})">
    <svg
      :viewBox="`0 0 ${(sizeInPixels + strokeWidth) * 2} ${(sizeInPixels + strokeWidth) * 2}`"
      :height="50 * 2.5"
      :width="50 * 2.5"
      :class="circular({})"
    >
    
      <circle
        v-if="type === 'single'"
        :cx="sizeInPixels + strokeWidth"
        :cy="sizeInPixels + strokeWidth + 1"
        :r="sizeInPixels"
        
        fill="none"
        opacity="1"
        class="loader-circle"
        stroke-linecap="round"
      />
      <circle
      v-if="type === 'triple'"
        :cx="sizeInPixels + strokeWidth"
        :cy="sizeInPixels + strokeWidth + 1"
        :r="sizeInPixels"
        
        fill="none"
        opacity="1"
        class="loader-circle-2"
        stroke-linecap="round"
      />
    </svg>
  </div>
</template>

<style>
.loader-container {
  position: relative;
  padding: 0;
  
  > svg {
    height: 100%;
    width: 100%;

    > circle {
      stroke: var(--bg-color);
    }
  }

  &.sm{
    height: 1.8rem;
    width: 1.8rem;

    circle {
      stroke-width: 1.4rem;
    }
  }

  &.md{
    height: 3.6rem;
    width: 3.6rem;

    circle {
      stroke-width: 1.6rem;
    }
  }

  &.lg{
    height: 5.4rem;
    width: 5.4rem;

    circle {
      stroke-width: 1.8rem;
    }
  }

  &.xl{
    height: 7.2rem;
    width: 7.2rem;

    circle {
      stroke-width: 2rem;
    }
  }
  
}

@keyframes rotate-spin {
  to {
    transform: rotate(360deg);
  }
}

@keyframes spin3 {
  0% {
    stroke-dasharray: 1, 313;
  }
  100% {
    stroke-dasharray: 100, 313;
    stroke-dashoffset: -313;
  }
}

@keyframes spin4 {
  0% {
    stroke-dasharray: 1, 107;

  }
  50% {
    stroke-dasharray: 80, 20;
    stroke-dashoffset: -107;
  }
  100%{
    stroke-dasharray: 1, 107;
  }
}

.loader-circle {
  transform-origin: center;
  animation:
  spin3 1.5s ease-in-out infinite,
  rotate-spin 3.5s linear infinite;
}

.loader-circle-2 {
  transform-origin: center;
  stroke: cadetblue;
  animation:
    spin4 1.5s ease-in-out infinite,
    rotate-spin 3.5s linear infinite;
}


</style>