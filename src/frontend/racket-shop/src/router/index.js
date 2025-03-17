import { createRouter, createWebHistory } from 'vue-router'
import config from "../../config";

import LoginView from '../views/Auth/LoginView.vue'
import RegisterView from '../views/Auth/RegisterView.vue'
import ProfileView from '../views/User/ProfileView.vue'
import CartView from '../views/User/CartView.vue'
import FeedbacksView from '../views/User/FeedbacksView.vue'
import OrdersView from '../views/User/OrdersView.vue'

const router = createRouter({
    history: createWebHistory(),
    routes: [
      {
        path: config.API.auth.login,
        component: LoginView,
      },
      {
        path: config.API.auth.register,
        component: RegisterView,
      },
      {
        path: config.API.user.profile,
        component: ProfileView,
      },
      {
        path: config.API.user.cart,
        component: CartView,
      },
      {
        path: config.API.user.feedbacks,
        component: FeedbacksView,
      },
      {
        path: config.API.user.orders,
        component: OrdersView,
      },
    ],
})

export default router