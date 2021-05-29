<template>
  <div class="chat">
    <ChatHistory :messages="messages" />
    <InputSection @newMessage="sendMessage" :loading="loading" />
  </div>
</template>

<script lang="ts">
import ChatHistory from "./ChatHistory.vue";
import InputSection from "./ChatInput.vue";
import Message from "../models/Message";
import { defineComponent, ref, onMounted } from "vue";

interface SendMessageResponse {
  response: Message;
  message: Message;
}

export default defineComponent({
  name: "chat",
  components: { ChatHistory, InputSection },
  setup() {
    const messages = ref([] as Message[]);
    const loading = ref(false);

    const getMessages = async () => {
      fetch("http://localhost:9090/messages")
        .then((res) => res.json())
        .then((data: Message[]) => {
          messages.value = data || [];
        });
    };

    const sendMessage = async (messageText: string) => {
      loading.value = true;
      const response = await fetch("http://localhost:9090/messages", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ text: messageText }),
      });

      loading.value = false;
      const responseMessage: Message = await response.json();
      messages.value.push(responseMessage);
    };

    onMounted(getMessages);

    return {
      messages,
      getMessages,
      sendMessage,
      loading,
    };
  },
});
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
