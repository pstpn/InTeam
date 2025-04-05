import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

const app = createApp(App)
localStorage.setItem('token', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDM5MTg5ODIsImlhdCI6MTc0Mzg3NTc4MiwiVXNlcklEIjoyLCJSb2xlIjoiQWRtaW4ifQ.UMgpUSErCKGS9wSO2-CY0LN5Wd50zn-WnL0p4Kwhvlw');

app.use(router)

app.mount('#app')