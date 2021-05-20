<template>
  <div class="container">
    <div class="title-description">
      <h1 class="title">Restaurante Determinado</h1>

      <br />
      <p class="description">
        Puedes escribirnos cualquier duda acerca de nuestro menú y servicios por
        medio de este chat.
      </p>
    </div>
    <div class="chat">
      <ChatHistory />
      <InputSection />
    </div>
  </div>
</template>

<script setup>
import InputSection from "./InputSection";
import ChatHistory from "./ChatHistory";
</script>

<style lang="scss" scoped>
.container {
  display: flex;
  padding: 3em;
  height: 100vh;

  & > *:not(:last-child) {
    margin-right: 3em;
  }
}

.title-description {
  display: flex;
  flex-direction: column;
  border-radius: 0.5em;
  padding: 2em;
  color: red;
  width: 20%;

  .title {
    color: var(--primary-color);
  }

  .description {
  }
}

.chat {
  width: 100%;
  display: flex;
  flex-direction: column;
  padding: 2em;
  border: 1px solid var(--text-color);
  border-radius: 2em;
  background: var(--surface-a);
}
</style>

