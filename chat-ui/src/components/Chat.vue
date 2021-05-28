<template>
  <div class="chat">
    <ChatHistory :messages="messages" />
    <InputSection @newMessage="sendMessage" />
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

    const getMessages = async () => {
      fetch("http://localhost:9090/messages")
        .then((res) => res.json())
        .then((data: Message[]) => {
          messages.value = data || [];
        });
    };

    const sendMessage = async (messageText: string) => {
      const response = await fetch("http://localhost:9090/messages", {
        method: "POST",
        // mode: "no-cors", // no-cors, *cors, same-origin
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ text: messageText }),
      });

      const data: SendMessageResponse = await response.json();
      messages.value.push(data.message);
      messages.value.push(data.response);
    };

    onMounted(getMessages);

    return {
      messages,
      getMessages,
      sendMessage,
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
