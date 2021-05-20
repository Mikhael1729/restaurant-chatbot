<template>
  <div class="container">
    <div class="title-description">
      <h1 class="title">Restaurante<br />Chat</h1>

      <p class="description">
        <i>
          Puedes escribirnos cualquier duda acerca de nuestros servicios
          por medio de este chat
        </i>
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
  width: 30%;

  .title {
    color: #dbdbdb;
  }

  .description {
    flex: 1;
    margin-top: 1em;
    font-size: 0.8em;
    line-height: 170%;
  }

  .logo {
    width: 80px;
    align-self: center;
  }
}

.chat {
  width: 100%;
  display: flex;
  flex-direction: column;
  padding: 2em;
  border: 1px solid var(--text-color);
  border-radius: 0.8em;
  background: var(--surface-a);
}
</style>

