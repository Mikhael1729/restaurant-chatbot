<template>
  <div class="container">
    <TitleDescription />
    <Chat />
  </div>
</template>

<script setup>
import Chat from "./components/Chat"
import TitleDescription from "./components/TitleDescription";
</script>

<style>
* {
  margin: 0;
  padding: 0;
  font-family: var(--font-family);
  color: var(--text-color);
}
</style>

<style lang="scss" scoped>
.container {
  display: flex;
  padding: 3em;
  height: 100vh;
  background-color: var(--surface-b);

  & > *:not(:last-child) {
    margin-right: 3em;
  }
}
</style>
