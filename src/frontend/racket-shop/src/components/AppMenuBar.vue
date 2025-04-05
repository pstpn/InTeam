<template>
    <nav class="container-menu font-menu-bar">
        <router-link to="/rackets" class="container-icon">
            <img src="../assets/logo.png" alt="Rackets" class="icon-logo"/>
        </router-link>
        <nav class="container-menu-items">
            <!-- Иконки для пользователя -->
            <template v-if="role === 'User'">
                <a @click="navigateTo(config.VIEWS.user.orders)" :class="{ 'active': activeLink === config.VIEWS.user.orders }" class="container-icon">
                    <img src="../assets/order.png" alt="Order" class="icon"/>
                    <span class="tooltip font-tooltip">Заказы</span>
                </a>

                <a @click="navigateTo(config.VIEWS.user.feedbacks)" :class="{ 'active': activeLink === config.VIEWS.user.feedbacks }" class="container-icon">
                    <img src="../assets/feedback.png" alt="Feedback" class="icon"/>
                    <span class="tooltip font-tooltip">Отзывы</span>
                </a>

                <a @click="navigateTo(config.VIEWS.user.cart)" :class="{ 'active': activeLink === config.VIEWS.user.cart }" class="container-icon">
                    <img src="../assets/cart.png" alt="Cart" class="icon"/>
                    <span class="tooltip font-tooltip">Корзина</span>
                </a>

                <a @click="navigateTo(config.VIEWS.user.profile)" :class="{ 'active': activeLink === config.VIEWS.auth.login || activeLink === config.VIEWS.auth.register || activeLink === config.VIEWS.user.profile}" class="container-icon">
                    <img src="../assets/auth.png" alt="Login" class="icon"/>
                    <span class="tooltip font-tooltip">Личный<br>кабинет</span>
                </a>
            </template>

            <!-- Иконки для администратора -->
            <template v-else-if="role === 'Admin'">
                <a @click="navigateTo(config.VIEWS.admin.rackets)" :class="{ 'active': activeLink === config.VIEWS.admin.rackets }" class="container-icon">
                    <img src="../assets/rackets.png" alt="Rackets" class="icon"/>
                    <span class="tooltip font-tooltip">Ракетки</span>
                </a>

                <a @click="navigateTo(config.VIEWS.admin.orders)" :class="{ 'active': activeLink === config.VIEWS.admin.orders }" class="container-icon">
                    <img src="../assets/order.png" alt="Orders" class="icon"/>
                    <span class="tooltip font-tooltip">Заказы</span>
                </a>

                <a @click="navigateTo(config.VIEWS.admin.users)" :class="{ 'active': activeLink === config.VIEWS.admin.users }" class="container-icon">
                    <img src="../assets/users.png" alt="Users" class="icon"/>
                    <span class="tooltip font-tooltip">Пользователи</span>
                </a>

                <a @click="navigateTo(config.VIEWS.admin.profile)" :class="{ 'active': activeLink === config.VIEWS.auth.login || activeLink === config.VIEWS.auth.register || activeLink === config.VIEWS.admin.profile}" class="container-icon">
                    <img src="../assets/auth.png" alt="Login" class="icon"/>
                    <span class="tooltip font-tooltip">Личный<br>кабинет</span>
                </a>
            </template>
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
            role: '',
            config: config, 
        };
    },
    methods: {
        parseJwt(token) {
            try {
                const base64Url = token.split('.')[1];
                const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
                const jsonPayload = decodeURIComponent(atob(base64).split('').map(function(c) {
                    return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
                }).join(''));

                return JSON.parse(jsonPayload);
            } catch (e) {
                return null;
            }
        },
        hasToken() {
            const token = localStorage.getItem('token');
            
            if (!token) {
                return false;
            }
            
            const decoded = this.parseJwt(token);

            console.log(decoded.Role)

            if (decoded && decoded.Role) {
                this.role = decoded.Role;
                return true;
            }
            return false;
        },
        navigateTo(route) {
            if (this.hasToken()) {
                this.activeLink = route;
                this.$router.push(route);
            } else {
                this.$router.push(config.VIEWS.auth.login);
            }
        }
    },
    mounted() {
        this.activeLink = this.$route.path;
        this.hasToken();
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