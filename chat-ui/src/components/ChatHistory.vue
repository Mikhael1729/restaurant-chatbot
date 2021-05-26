<template>
  <div class="chat-history" :class="{ 'empty-chat-history': isEmpty() }">
    <div v-if="isEmpty()" class="is-empty">Historial vacío</div>
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
import { defineComponent } from "vue";

export default defineComponent({
  components: {
    ChatMessage,
  },
  data() {
    return {
      messages: [] as Message[],
    };
  },
  mounted() {
    fetch("http://localhost:9090/messages")
      .then((res) => res.json())
      .then((data: Message[]) => {
        this.messages = data;
        console.log(this.messages)
      });
  },
  methods: {
    isEmpty: function () {
      return this.messages.length === 0;
    },
  },
});
</script>

<style lang="scss" scoped>
.chat-history {
  background: inherit;
  display: flex;
  flex-direction: column; 
  justify-content: flex-end;
  flex: 1;
  width: 100%;
  overflow: auto;
  padding: 1.5em;

  & > *:not(:last-child) {
    margin-bottom: 1em;
  }
}
.empty-chat-history {
  align-items: center;
  justify-content: center;

  .is-empty {
    font-style: italic;
  }
}
</style>
