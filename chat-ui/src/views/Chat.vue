<template>
  <div class="container">
    <div class="chat">
    </div>
    <div class="input-section">
      <InputText v-model="message" />
      <Button>Hey</Button>
    </div>
  </div>
</template>


<style lang="scss" scoped>
.container {
  height: 100vh;
  width: 100%;
  display: flex;
  background: red;
  flex-direction: column;

  .input-section {
    background: yellow;
    width: 100%;
  }
}
</style>

<script>
export default {
  data() {
    return {
      message: ":O",
    };
  },
};
</script>
