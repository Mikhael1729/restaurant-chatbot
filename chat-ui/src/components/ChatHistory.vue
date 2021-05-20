<template>
  <div class="chat-history">
    <template v-for="message in messages" :key="message.id">
      <ChatMessage
        :date="message.date"
        :text="message.text"
        :type="message.type"
      />
    </template>
  </div>
</template>

<script lang="ts">
import ChatMessage from "./ChatMessage.vue";
import Message from "../models/Message";
import { defineComponent, PropType } from "vue";

export default defineComponent({
  components: {
    ChatMessage,
  },
  data() {
    return {
      messages: Message.mockData(),
    };
  },

  mounted() {
    console.log(this.$data.messages);
  },
});
</script>

<style lang="scss" scoped>
.chat-history {
  background: inherit;
  display: flex;
  flex-direction: column;
  flex: 1;
  width: 100%;

  & > *:not(:last-child) {
    margin-bottom: 1em;
  }
}
</style>

