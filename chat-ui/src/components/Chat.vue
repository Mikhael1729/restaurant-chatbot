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
.chat {
  display: flex;
  flex-direction: column;
  padding: 0.5em;
  border: 1px solid var(--text-color);
  border-radius: 0.8em;
  background: var(--surface-a);
  max-height: calc(100vh - 6em);
}
</style>

