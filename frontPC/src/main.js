import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import PrimeVue from "primevue/config";

import "primevue/resources/themes/mdc-light-indigo/theme.css";
import "primevue/resources/primevue.min.css";
import "primeicons/primeicons.css";
import "primeflex/primeflex.css";

import InputText from "primevue/inputtext";
import Button from "primevue/button";

const app = createApp(App)
app.use(PrimeVue);

app.use(router)
app.component("pv-input-text", InputText);
app.component("pv-button", Button);
app.mount('#app')
