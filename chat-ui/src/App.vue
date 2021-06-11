<template>
  <div class="p-grid">
    <div class="title-description p-col-12 p-lg-4">
      <TitleDescription />
    </div>
    <div chat class="p-col-12 p-lg-8">
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
}

html {
  height: 100%;
}

body {
  margin: 0px;
  background-color: var(--surface-b);
  padding: 3em;
  height: 100vh;
  width: 100%;
}
</style>

<style lang="scss" scoped>
.title-description {
  display: flex;
  justify-content: center;
}

.chat {
  display: flex;
  height: calc(100vh - 5.5em);
}
</style>
