import { createRouter, createWebHashHistory, RouteRecordRaw } from "vue-router";
import Chat from "../views/Chat.vue";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    name: "Chat del restaurante",
    component: Chat,
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;
