import {createApp} from 'vue'
import App from './App.vue'
import './style.css';
import router from './router';
import Toast from "vue-toastification"; // Import Vue Router
import "vue-toastification/dist/index.css";

const app = createApp(App)

app.use(router)
app.use(Toast);

app.mount('#app')
