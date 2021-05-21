<template>
  <div
    :class="{
      'chat-message': true,
      'chat-message-sender': isCustomer(),
    }"
  >
    <Avatar
      shape="circle"
      size="small"
      :class="{ avatar: true, 'avatar-sender': isCustomer() }"
    >
      <i
        v-if="!isCustomer()"
        class="pi pi-android"
        style="background: inherit"
      />
      <i v-else class="pi pi-user" style="background: inherit" />
    </Avatar>
    <div class="message">
      <p>
        {{ text }}
      </p>
      <span>{{ messageDate() }}</span>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { Sender } from "../models/Message";

export default defineComponent({
  props: {
    text: String,
    date: String,
    sender: Number,
  },
  methods: {
    isCustomer: function () {
      return this.sender === Sender.Customer;
    },
    messageDate: function () {
      const parts = this.date?.split("T");

      if (parts?.length !== 2) return "";

      const date = parts[0];
      const timeInfo = parts[1]?.split(".");
      const timeParts = timeInfo[0].split(":");
      const time = `${timeParts[0]}:${timeParts[1]}`;

      return `${date}  ${time}`;
    },
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

/* It's used to align the message at the end of the row in the case the sender is a customer */
.chat-message-sender {
  flex-direction: row-reverse;
}
</style>
