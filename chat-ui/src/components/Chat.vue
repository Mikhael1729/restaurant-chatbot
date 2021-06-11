<template>
  <div class="chat">
    <ChatHistory :messages="messages" :error="error" />
    <InputSection @newMessage="sendMessage" :loading="loading" :error="error" />
  </div>
</template>

<script lang="ts">
import ChatHistory from "./ChatHistory.vue";
import InputSection from "./ChatInput.vue";
import Message, { Sender } from "../models/Message";
import SendMessageResponse from "../models/SendMessageResponse";
import { defineComponent, ref, onMounted } from "vue";

export default defineComponent({
  name: "chat",
  components: { ChatHistory, InputSection },
  setup() {
    const messages = ref([] as Message[]);
    const loading = ref(false);
    const error = ref("");

    const getMessages = async () => {
      try {
        const response = await fetch("http://localhost:9090/messages");
        const data: Message[] = await response.json();

        messages.value = data || [];

        if (error.value) error.value = "";
      } catch (e) {
        error.value =
          "Error: No se pudieron cargar los mensajes. Intente nuevamente recargando la página";
      }
    };

    const sendMessage = async (messageText: string) => {
      loading.value = true;

      try {
        const userMessage: Message = {
          id: 0,
          text: messageText,
          dateTime: "",
          sender: Sender.Customer,
          category: "",
        };

        const response = await fetch("http://localhost:9090/messages", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({ text: messageText }),
        });

        messages.value.push(userMessage);

        const data: SendMessageResponse = await response.json();

        messages.value[messages.value.length - 1] = data.message;
        messages.value.push(data.response);
        if (error.value) error.value = "";
      } catch (e) {
        error.value =
          "Lo sentimos, su mensaje no pudo ser enviado. intente de nuevo más tarde";
      }
      loading.value = false;
    };

    onMounted(getMessages);

    return {
      messages,
      getMessages,
      sendMessage,
      loading,
      error,
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
