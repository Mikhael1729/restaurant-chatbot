<template>
  <div class="input-section">
    <InputText v-model="message" class="input-message" />
    <Button
      label="Enviar"
      class="send-button p-button-outlined"
      icon="pi pi-send"
      v-on:click="addNewMessage"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";

export default defineComponent({
  name: "chat-input",
  emits: ["newMessage"],
  setup(_, context) {
    const message = ref("");
    const addNewMessage = () => context.emit("newMessage", message.value);

    return { message, addNewMessage };
  },
});
</script>

<style lang="scss" scoped>
.input-section {
  width: 100%;
  display: flex;
  padding: 0.5em;

  .input-message {
    width: 100%;
  }
  .send-button {
    margin-left: 1em;
  }
}
</style>
