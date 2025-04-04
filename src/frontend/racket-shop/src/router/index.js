import { createRouter, createWebHistory } from 'vue-router'
import config from "../../config"

import LoginView from '../views/Auth/LoginView.vue'
import RegisterView from '../views/Auth/RegisterView.vue'

import ProfileView from '../views/User/ProfileView.vue'
import CartView from '../views/User/CartView.vue'
import FeedbacksView from '../views/User/FeedbacksView.vue'
import OrdersView from '../views/User/OrdersView.vue'
import RacketsView from '../views/RacketsView.vue'
import RacketView from '../views/RacketView.vue'

import AdminProfileView from '../views/Admin/ProfileView.vue'
import AdminRacketsView from '../views/Admin/RacketsView.vue'
import AdminOrdersView from '../views/Admin/OrdersView.vue'
import AdminUsersView from '../views/Admin/UsersView.vue'

const router = createRouter({
    history: createWebHistory(),
    routes: [
      {
        path: config.VIEWS.auth.login,
        component: LoginView,
      },
      {
        path: config.VIEWS.auth.register,
        component: RegisterView,
      },
      {
        path: config.VIEWS.user.profile,
        component: ProfileView,
      },
      {
        path: config.VIEWS.rackets,
        component: RacketsView,
      },
      {
        path: config.VIEWS.rackets+'/:id',
        component: RacketView,
      },
      {
        path: config.VIEWS.user.cart,
        component: CartView,
      },
      {
        path: config.VIEWS.user.feedbacks,
        component: FeedbacksView,
      },
      {
        path: config.VIEWS.user.orders,
        component: OrdersView,
      },
      {
        path: config.VIEWS.admin.rackets,
        component: AdminRacketsView,
      },
      {
        path: config.VIEWS.admin.orders,
        component: AdminOrdersView,
      },
      {
        path: config.VIEWS.admin.users,
        component: AdminUsersView,
      },
      {
        path: config.VIEWS.admin.profile,
        component: AdminProfileView,
      },
    ],
})

export default router