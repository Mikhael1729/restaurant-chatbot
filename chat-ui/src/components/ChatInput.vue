<template>
  <div class="input-section">
    <InputText v-model="message" class="input-message" />
    <Button
      label="Enviar"
      class="send-button p-button-outlined"
      icon="pi pi-send"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  data() {
    return {
      message: ":3",
    };
  },
});
</script>

<style lang="scss" scoped>
.input-section {
  width: 100%;
  display: flex;
  background: inherit;

  .input-message {
    width: 100%;
  }
  .send-button {
    margin-left: 1em;
  }
}
</style>
