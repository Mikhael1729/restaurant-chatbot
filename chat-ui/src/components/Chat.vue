<template>
  <div class="chat">
    <ChatHistory />
    <InputSection />
  </div>
</template>

<script lang="ts">
import ChatHistory from "./ChatHistory.vue";
import InputSection from "./ChatInput.vue";
import { defineComponent } from "vue";

export default defineComponent({
  components: {
    ChatHistory,
    InputSection,
  }
}
</script>

<style lang="scss" scoped>
.chat {
  display: flex;
  flex-direction: column;
  padding: 0.5em;
  border: 1px solid var(--text-color);
  border-radius: 0.8em;
  background: var(--surface-a);
}
</style>
