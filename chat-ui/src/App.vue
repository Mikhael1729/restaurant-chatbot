<template>
  <div class="p-grid">
    <div class="p-col-12 p-lg-3">
      <TitleDescription />
    </div>
    <div class="p-col-12 p-lg-9">
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

body {
  margin: 0px;
  background-color: var(--surface-b);
  padding: 3em;
  height: 100vh;
  width: 100%;
}
</style>
