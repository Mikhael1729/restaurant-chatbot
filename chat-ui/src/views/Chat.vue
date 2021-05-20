<template>
  <div class="p-grid">
    <div class="container">
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
  margin: 0 auto;
  height: 100vh;
  max-width: 80%;
  display: flex;
  flex-direction: column;
  padding: 3em;
}
</style>

