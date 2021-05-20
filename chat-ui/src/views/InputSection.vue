<template>
  <div class="input-section">
    <InputText v-model="message" class="input-message" />
    <Button>Enviar</Button>
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
  background: yellow;
  width: 100%;
  display: flex;

  .input-message {
    width: 100%;
  }
}
</style>
