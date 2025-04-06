<template>
    <nav class="container-menu font-menu-bar">
        <router-link to="/rackets" class="container-icon">
            <img src="../assets/logo.png" alt="Rackets" class="icon-logo"/>
        </router-link>
        <nav class="container-menu-items">

            <template v-if="!isAdmin">
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

                <a @click="navigateTo(profileRoute)" :class="{ 'active': isProfileRoute }" class="container-icon">
                    <img src="../assets/auth.png" alt="Login" class="icon"/>
                    <span class="tooltip font-tooltip">Личный<br>кабинет</span>
                </a>
            </template>

            <template v-else>
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

                <a @click="navigateTo(config.VIEWS.admin.profile)" :class="{ 'active': isProfileRoute }" class="container-icon">
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
    computed: {
        computed: {
            role() {
                return this.authStore.role;
            }
        },
        isAdmin() {
            return this.role === 'Admin';
        },
        profileRoute() {
            return this.isAdmin ? config.VIEWS.admin.profile : config.VIEWS.user.profile;
        },
        isProfileRoute() {
            const profileRoutes = [
                config.VIEWS.auth.login,
                config.VIEWS.auth.register,
                config.VIEWS.user.profile,
                config.VIEWS.admin.profile
            ];
            return profileRoutes.includes(this.activeLink);
        }
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
        checkRole() {
            const token = localStorage.getItem('token');
            if (!token) {
                this.role = '';
                return false;
            }
            
            const decoded = this.parseJwt(token);
            this.role = decoded?.Role || '';
            return !!this.role;
        },
        navigateTo(route) {
            if (this.checkRole() || route === config.VIEWS.auth.login) {
                this.activeLink = route;
                this.$router.push(route);
            } else {
                this.$router.push(config.VIEWS.auth.login);
            }
        }
    },
    mounted() {
        this.activeLink = this.$route.path;
        this.checkRole();
    },
    watch: {
        '$route.path'(newPath) {
            this.activeLink = newPath;
        }
    },
}
</script>