<script setup lang="ts">
const LINKS = [
  {
    title: "Teste CI - Bea 1234",
    href: "/example/1",
    description: "Lorem ipsum...",
  },
  {
    title: "Teste CI - sem sucesso :c",
    href: "/example/1",
    description: "Lorem ipsum...",
  },
];
</script>

<template>
  <main class="main">
    <div class="grid">
      <smexy-card
        v-for="{ title, href, description } in LINKS"
        class="card"
        :href="href"
        :key="title"
        :title="title"
      >
        {{ description }}
      </smexy-card>
    </div>
  </main>
</template>

<style>
@import url("https://fonts.googleapis.com/css2?family=Inter:wght@100;200;300;400;500;600;700;800;900&display=swap");

:root {
  --max-width: 1100px;
  --border-radius: 12px;
  --font-mono: ui-monospace, Menlo, Monaco, "Cascadia Mono", "Segoe UI Mono",
    "Roboto Mono", "Oxygen Mono", "Ubuntu Monospace", "Source Code Pro",
    "Fira Mono", "Droid Sans Mono", "Courier New", monospace;

  --foreground-rgb: black;
  --background-start-rgb: white;
  --background-end-rgb: white;

  --callout-rgb: 20, 20, 20;
  --callout-border-rgb: 108, 108, 108;
  --card-rgb: 100, 100, 100;
  --card-border-rgb: 200, 200, 200;

  --glow-conic: conic-gradient(
    from 180deg at 50% 50%,
    #2a8af6 0deg,
    #a853ba 180deg,
    #e92a67 360deg
  );
}

* {
  box-sizing: border-box;
  padding: 0;
  margin: 0;
}

html,
body {
  max-width: 100vw;
  overflow-x: hidden;
  font-family: "Inter", sans-serif;
}

body {
  color: rgb(var(--foreground-rgb));
  background: linear-gradient(
      to bottom,
      transparent,
      rgb(var(--background-end-rgb))
    )
    rgb(var(--background-start-rgb));
}

a {
  color: inherit;
  text-decoration: none;
}

.main {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  align-items: center;
  padding: 6rem;
  min-height: 100vh;
}

.grid {
  display: grid;
  grid-template-columns: repeat(4, minmax(25%, auto));
  max-width: 100%;
  width: var(--max-width);
}

.card {
  padding: 1rem 1.2rem;
  border-radius: var(--border-radius);
  background: rgba(var(--card-rgb), 0);
  border: 1px solid rgba(var(--card-border-rgb), 0);
  transition:
    background 200ms,
    border 200ms;
}

.card span {
  display: inline-block;
  transition: transform 200ms;
}

.card h2 {
  font-weight: 600;
  margin-bottom: 0.7rem;
}

.card p {
  margin: 0;
  opacity: 0.6;
  font-size: 0.9rem;
  line-height: 1.5;
  max-width: 30ch;
}

@media (prefers-reduced-motion) {
  .card:hover span {
    transform: none;
  }
}

/* Enable hover only on non-touch devices */
@media (hover: hover) and (pointer: fine) {
  .card:hover {
    background: rgba(var(--card-rgb), 0.1);
    border: 1px solid rgba(var(--card-border-rgb), 0.15);
  }

  .card:hover span {
    transform: translateX(4px);
  }
}
</style>
