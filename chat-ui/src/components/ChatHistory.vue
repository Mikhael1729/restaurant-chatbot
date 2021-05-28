<template>
  <div
    id="history"
    class="chat-history"
    :class="{ 'empty-chat-history': isEmpty }"
    ref="historyRef"
  >
    <div v-if="isEmpty" class="is-empty">Historial vacío</div>
    <template v-for="message in messages" :key="message.id">
      <ChatMessage
        :date="message.dateTime"
        :text="message.text"
        :sender="message.sender"
      />
    </template>
  </div>
</template>

<script lang="ts">
import ChatMessage from "./ChatMessage.vue";
import {
  defineComponent,
  PropType,
  ref,
  computed,
  watch,
  onMounted,
  onRenderTriggered,
  onRenderTracked,
} from "vue";

export default defineComponent({
  name: "chat-history",
  components: { ChatMessage },
  props: { messages: Array as PropType<Array<number>> },
  setup(props, context) {
    const isEmpty = computed(() => props.messages?.length === 0);
    const historyRef = ref(null);

    watch(
      () => props.messages?.length,
      (first, second) => {
        const el: any = historyRef.value;
        if (el) {
          el.scrollTop = el.scrollHeight;
        }
      }
    );

    return {
      isEmpty,
      historyRef,
    };
  },
});
</script>

<style lang="scss" scoped>
.chat-history {
  display: flex;
  flex-direction: column;
  width: 100%;
  overflow: auto;
  padding: 1.5em;

  & > *:not(:last-child) {
    margin-bottom: 1em;
  }
}
.empty-chat-history {
  align-items: center;
  justify-content: center;

  .is-empty {
    font-style: italic;
  }
}
</style>
