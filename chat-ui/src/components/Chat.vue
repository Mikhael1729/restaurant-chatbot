<template>
  <div class="chat">
    <ChatHistory />
    <InputSection />
  </div>
</template>

<script setup>
import InputSection from "./ChatInput";
import ChatHistory from "./ChatHistory";
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

