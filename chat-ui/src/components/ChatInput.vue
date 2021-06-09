<template>
  <form class="input-section" v-on:submit.prevent="addNewMessage">
    <InputText v-model="message" class="input-message" />
    <Button
      type="submit"
      label="Enviar"
      class="send-button p-button-outlined"
      :icon="icon"
    />
  </form>
</template>

<script lang="ts">
import { defineComponent, ref, watch } from "vue";

export default defineComponent({
  name: "chat-input",
  emits: ["newMessage"],
  props: {
    loading: Boolean,
  },
  setup(props, context) {
    const message = ref("");
    const icon = ref("pi pi-send");

    const addNewMessage = () => {
      context.emit("newMessage", message.value);
      message.value = "";
    };

    // Show loading icon when the data is being send.
    watch(
      () => props.loading,
      (first, second) => {
        if (props.loading) {
          icon.value = "pi pi-spin pi-spinner";
        } else {
          icon.value = "pi pi-send";
        }
      }
    );

    return { message, addNewMessage, icon };
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
