<template>
  <div class="container p-grid">
    <div class="title-description p-col-12 p-lg-3">
      <TitleDescription />
    </div>
    <div class="chat p-col-12 p-lg-9">
      <Chat />
    </div>
  </div>
</template>

<script setup>
import Chat from "./components/Chat";
import TitleDescription from "./components/TitleDescription";
</script>

<style>
* {
  font-family: var(--font-family);
  color: var(--text-color);
}
</style>

<style lang="scss" scoped>
.container {
  background-color: var(--surface-b);
  padding: 3em;
  height: 100vh;

  .title-description {
    color: red;
  }

  .chat {
  }
}
</style>
