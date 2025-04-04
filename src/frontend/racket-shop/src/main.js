import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

const app = createApp(App)
// localStorage.setItem('token', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDMyOTU3MTYsImlhdCI6MTc0MzI1MjUxNiwiVXNlcklEIjoyLCJSb2xlIjoiQWRtaW4ifQ.TgtAoBygjtpeFNmu4oDZQHKUbxmSWKa3G3yG_NAs6S4');

app.use(router)

app.mount('#app')