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

<style lang="scss" scoped>
.container {
  display: flex;
  padding: 3em;
  height: 100vh;

  & > *:not(:last-child) {
    margin-right: 3em;
  }
}
</style>
