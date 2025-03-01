import { createRouter, createWebHistory } from 'vue-router'

import LoginView from '../views/Auth/LoginView.vue'
import RegisterView from '../views/Auth/RegisterView.vue'

const router = createRouter({
    history: createWebHistory(),
    routes: [
      {
        path: '/auth/login',
        component: LoginView,
      },
      {
        path: '/auth/register',
        component: RegisterView,
      },
    ],
})

export default router