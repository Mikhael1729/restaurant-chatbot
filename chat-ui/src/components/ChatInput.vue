<template>
  <div class="input-section">
    <InputText v-model="message" class="input-message" />
    <Button
      label="Enviar"
      class="send-button p-button-outlined"
      icon="pi pi-send"
      v-on:click="sendMessage"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import Message from "../models/Message";

export default defineComponent({
  data() {
    return {
      message: "",
    };
  },
  methods: {
    sendMessage: async function () {
      const response = await fetch("http://localhost:9090/messages", {
        method: "POST",
        // mode: "no-cors", // no-cors, *cors, same-origin
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ text: this.message }),
      });

      const message: Message = await response.json();
    },
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
