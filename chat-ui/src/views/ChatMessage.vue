<template>
  <div class="chat-message">
    <Avatar label="P" shape="circle" size="small" class="avatar" />
    <div class="message">
      <p class="text">
        {{ text }}
      </p>
      <span class="date">{{ stringDate }}</span>
    </div>
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
      stringDate: "hace 5 minutos",
    };
  },
});
</script>

<style lang="scss" scoped>
.chat-message {
  display: flex;

  & > *:not(:last-child) {
    margin-right: 1em;
  }

  .avatar {
    align-self: flex-end;
  }

  .message {
    display: flex;
    flex-direction: column;
    padding: 0.2em 0.5em 0.2em 0.5em;
    border-radius: 0.4em;
    border: 1px solid black;
    border-color: var(--text-color-secondary);
    color: var(--text-color);

    .text {
      margin-bottom: 0.5em;
    }
    .date {
      text-align: right;
      font-size: 0.8em;
      color: var(--surface-900);
    }
  }
}
</style>
