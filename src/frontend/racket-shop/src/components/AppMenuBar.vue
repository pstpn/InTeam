<template>
    <nav class="container-menu font-menu-bar">
        <router-link to="/rackets" class="container-icon">
            <img src="../assets/logo.png" alt="Rackets" class="icon-logo"/>
        </router-link>
        <nav class="container-menu-items">
            <a @click="navigateTo(config.API.user.orders)" :class="{ 'active': activeLink === config.API.user.orders }" class="container-icon">
                <img src="../assets/order.png" alt="Order" class="icon"/>
                <span class="tooltip font-tooltip">Заказы</span>
            </a>

            <a @click="navigateTo(config.API.user.feedbacks)" :class="{ 'active': activeLink === config.API.user.feedbacks }" class="container-icon">
                <img src="../assets/feedback.png" alt="Feedback" class="icon"/>
                <span class="tooltip font-tooltip">Отзывы</span>
            </a>

            <a @click="navigateTo(config.API.user.cart)" :class="{ 'active': activeLink === config.API.user.cart }" class="container-icon">
                <img src="../assets/cart.png" alt="Cart" class="icon"/>
                <span class="tooltip font-tooltip">Корзина</span>
            </a>

            <a @click="navigateTo(config.API.user.profile)" class="container-icon">
                <img src="../assets/auth.png" alt="Login" class="icon"/>
                <span class="tooltip font-tooltip">Личный<br>кабинет</span>
            </a>
        </nav>
    </nav>  
</template>

<script>
import config from "../../config.js";

export default {
    name: 'AppMenuBar',
    data() {
        return {
            activeLink: '',
            config: config, 
        };
    },
    methods: {
        hasToken() {
            return localStorage.getItem('token') !== null;
        },
        navigateTo(route) {
            if (this.hasToken()) {
                this.activeLink = route;
                this.$router.push(route);
            } else {
                this.$router.push(config.API.auth.login);
            }
        }
    },
    mounted() {
        this.activeLink = this.$route.path;
    },
    watch: {
        '$route.path'(newPath) {
            this.activeLink = newPath;
        }
    },
}
</script>

<style>
@import "../css/Containers.css";
@import "../css/Icons.css";
@import "../css/Fonts.css";
</style>