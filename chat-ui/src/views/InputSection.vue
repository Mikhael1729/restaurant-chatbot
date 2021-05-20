<template>
  <div class="input-section">
    <InputText v-model="message" class="input-message" />
    <Button class="send-button">Enviar</Button>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  data() {
    return {
      message: ":3"
    }
  }
})
</script>

<style lang="scss" scoped>
.input-section {
  width: 100%;
  display: flex;

  .input-message {
    width: 100%;
  }
  .send-button {
    margin-left: 1em;
  }
}
</style>
