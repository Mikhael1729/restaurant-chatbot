<template>
  <div class="chat-message">
    <span>{{ text }}</span>
    <span>{{ stringDate }}</span>
  </div>
</template>

<script lang="ts">
import { defineComponent, PropType } from "vue";

export default defineComponent({
  props: {
    text: String,
    date: Date,
    type: Object as PropType<"reciver" | "sender">,
  },
  data() {
    return {
      stringDate: this.date?.toString(),
    }
  }
});
</script>

<style lang="scss" scoped>
.chat-message {
  display: flex;
  background: red;

  & > *:not(:last-child) {
    margin-right: 1em;
  }
}
</style>
