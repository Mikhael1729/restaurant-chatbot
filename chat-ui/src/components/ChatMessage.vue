<template>
  <div
    :class="{ 'chat-message': true, 'chat-message-sender': type === 'sender' }"
  >
    <Avatar
      shape="circle"
      size="small"
      :class="{ avatar: true, 'avatar-sender': type === 'sender' }"
    >
      <i
        v-if="type === 'reciver'"
        class="pi pi-android"
        style="background: inherit"
      />
      <i v-else class="pi pi-user" style="background: inherit" />
    </Avatar>
    <div class="message">
      <p>
        {{ text }}
      </p>
      <span>{{ stringDate }}</span>
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
  background: inherit;
  display: flex;

  .avatar {
    align-self: flex-end;
    margin-right: 0.5em;
  }

  .avatar-sender {
    margin-left: 0.5em;
    margin-right: 0px;
  }

  .message {
    display: flex;
    flex-direction: column;
    padding: 0.2em 0.5em 0.2em 0.5em;
    border-radius: 0.4em;
    border: 1px solid black;
    border-color: var(--text-color-secondary);
    color: var(--text-color);
    background: inherit;

    p {
      margin-bottom: 0.5em;
      background: inherit;
    }
    span {
      text-align: right;
      font-size: 0.8em;
      color: var(--surface-900);
      background: inherit;
    }
  }
}

.chat-message-sender {
  flex-direction: row-reverse;
}
</style>
