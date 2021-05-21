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

export default defineComponent({
  data() {
    return {
      message: "",
    };
  },
  methods: {
    sendMessage: async function () {
      console.log(this.message);
      const response = await fetch("http://localhost:9090/messages", {
        method: "POST",
        mode: "no-cors", // no-cors, *cors, same-origin
        cache: "no-cache",
        headers: {
          "Content-Type": "application/json",
        },
        referrerPolicy: "no-referrer",
        body: JSON.stringify({ text: this.message }),
      });

      console.log(response);
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
