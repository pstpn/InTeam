<template>
    <div class="container-page">
        <h1 class="font-container-header">Личный кабинет</h1>
        
        <div class="grid-profile">
            <div class="grid-card-name">
                <p class="font-grid-1">{{ userData.name }} {{ userData.surname }}</p>
                <p class="font-grid-2">{{ userData.email }}</p>
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
import axios from 'axios';
import config from "../../../config.js"

export default {
    name: 'AppMenuBar',
    data() {
        return {
            config: config,
            userData: {
                name: '',
                surname: '',
                email: ''
            }
        };
    },
    methods: {
        async fetchUserData() {
            try {
                const token = localStorage.getItem('token');

                if (!token) {
                    this.$router.push(this.config.API.auth.login);
                    return;
                }

                const cur_url = config.BACKEND_URL + config.API.user.profile;
                const response = await axios.get(cur_url, {
                    headers: {
                        Authorization: `Bearer ${token}`
                    }
                });
                
                if (response.data) {
                    this.userData = {   
                        name: response.data.name,
                        surname: response.data.surname,
                        email: response.data.email
                    };
                }
            } catch (error) {
                console.error('Error fetching user data:', error);
                this.$router.push(this.config.API.auth.login);
            }
        },
        navigateTo(route) {
            this.$router.push(route);
        },
        logout() {
            localStorage.removeItem('token');
            this.$router.push(this.config.API.auth.login);
        }
    },
    mounted() {
        this.fetchUserData();
    }
}
</script>

<style>
@import "../../css/Forms.css";
@import "../../css/Containers.css";
@import "../../css/Fonts.css";
@import "../../css/Grid.css";
</style>