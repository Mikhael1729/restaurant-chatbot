<template>
  <div class="container">
    <TitleDescription />
    <div class="chat">
      <ChatHistory />
      <InputSection />
    </div>
  </div>
</template>

<script setup>
import InputSection from "./views/InputSection";
import ChatHistory from "./views/ChatHistory";
import TitleDescription from "./views/TitleDescription";
</script>

<style>
* {
  margin: 0;
  padding: 0;
  background-color: var(--surface-b);
  font-family: var(--font-family);
  color: var(--text-color);
}
</style>

<style lang="scss" scoped>
.container {
  display: flex;
  padding: 3em;
  height: 100vh;

  & > *:not(:last-child) {
    margin-right: 3em;
  }
}

.chat {
  width: 100%;
  display: flex;
  flex-direction: column;
  padding: 2em;
  border: 1px solid var(--text-color);
  border-radius: 0.8em;
  background: var(--surface-a);
}
</style>

