import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

const app = createApp(App)
localStorage.setItem('token', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDM4NDgxNTAsImlhdCI6MTc0MzgwNDk1MCwiVXNlcklEIjoyLCJSb2xlIjoiQWRtaW4ifQ.EjA7WY8VUb9gm_G5j58hNX_-u6-Br8AtY56QXsRXho0');

app.use(router)

app.mount('#app')