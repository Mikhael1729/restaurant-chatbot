<template>
  <div class="container">
    <ChatHistory />
    <InputSection />
  </div>
</template>

<script setup>
import InputSection from "./InputSection";
import ChatHistory from "./ChatHistory";
</script>

<style lang="scss" scoped>
.container {
  height: 100vh;
  width: 100%;
  display: flex;
  background: red;
  flex-direction: column;
}
</style>

