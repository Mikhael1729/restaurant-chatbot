<template>
  <div v-if="error" class="error-message">
    <i class="pi pi-times"></i>{{ error }}
  </div>
  <div
    id="history"
    class="chat-history"
    :class="{ 'empty-chat-history': isEmpty }"
  >
    <div v-if="isEmpty" class="is-empty">Historial vacío</div>
    <div v-if="!isEmpty" class="space" />
    <template v-for="message in messages" :key="message.id">
      <ChatMessage
        :date="message.dateTime"
        :text="message.text"
        :sender="message.sender"
      />
    </template>
  </div>
</template>

<script lang="ts">
import ChatMessage from "./ChatMessage.vue";
import Message from "../models/Message";
import { defineComponent, PropType, computed } from "vue";

export default defineComponent({
  name: "chat-history",
  components: { ChatMessage },
  props: { messages: Array as PropType<Array<Message>>, error: String },
  setup(props) {
    const isEmpty = computed(() => props.messages?.length === 0);

    return {
      isEmpty,
    };
  },
});
</script>

<style lang="scss" scoped>
.error-message {
  background: inherit !important;
  color: var(--pink-100) !important;
  border: 1px solid var(--pink-100);
  padding: 1em;
  border-radius: 10px;
}

.pi-times {
  margin-right: 1em;
  border: 2px solid var(--pink-100);
  padding: 0.3em;
  border-radius: 50%;
}

.chat-history {
  display: flex;
  flex-direction: column;
  width: 100%;
  overflow: auto;
  padding: 1.5em;
  flex: 1;
  color: var(--text-color);

  & > *:not(:last-child) {
    margin-bottom: 1em;
  }
}

.space {
  flex: 1;
}

.empty-chat-history {
  align-items: center;
  justify-content: center;

  .is-empty {
    font-style: italic;
  }
}
</style>
