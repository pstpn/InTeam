<template>
    <div class="container-page">
        <h1 class="font-container-header">Личный кабинет</h1>
        
        <div class="grid-profile">
            <div class="grid-card-name">
                <p class="font-grid-1">Имя Фамилия</p>
                <p class="font-grid-2">Email</p>
                <div class="form-in-row-right">
                    <button 
                        class="submit-button-orange"
                        @click="logout">
                        Выйти
                    </button>
                </div>
            </div>
            <div></div>
            <div></div>
            <div class="grid-card-icon">
                <a @click="navigateTo(config.API.user.orders)">
                    <h1 class="font-grid-1">Мои заказы</h1>
                </a>
            </div>
            <div class="grid-card-icon">
                <a @click="navigateTo(config.API.user.feedbacks)">
                    <h1 class="font-grid-1">Мои отзывы</h1>
                </a>
            </div>
        </div>
    </div>
</template>

<script>
import config from "../../../config.js"

export default {
    name: 'AppMenuBar',
    data() {
        return {
            config: config, 
        };
    },
    methods: {
        navigateTo(route) {
            this.$router.push(route);
        },
        logout() {
            localStorage.removeItem('token');
            this.$router.push(this.config.API.auth.login);
        }
    }
}
</script>

<style>
@import "../../css/Forms.css";
@import "../../css/Containers.css";
@import "../../css/Fonts.css";
@import "../../css/Grid.css";
</style>