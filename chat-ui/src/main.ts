import { createApp } from "vue";
import App from "./App.vue";
import store from "./store";
import PrimeVue from "primevue/config";
import Button from "primevue/button";
import InputText from "primevue/inputtext";
import Avatar from "primevue/avatar"
import Card from "primevue/card"

import 'primeflex/primeflex.css';
import "primevue/resources/themes/bootstrap4-dark-blue/theme.css"; //theme
import "primevue/resources/primevue.min.css"; //core css
import "primeicons/primeicons.css"; //icons

createApp(App)
  // Plugins.
  .use(store)
  .use(PrimeVue)
  // PrimeVue components
  .component("Button", Button)
  .component("InputText", InputText)
  .component("Avatar", Avatar)
  .component("Card", Card)
  // Mounting
  .mount("#app");

