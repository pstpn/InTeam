<template>
    <div class="container-page">
        <h1 class="font-container-header">Личный кабинет</h1>
        
        <div class="grid-profile">
            <div class="grid-card-name">
                <p class="font-grid-1">{{ userData.name }} {{ userData.surname }}</p>
                <p class="font-form-body">{{ userData.email }}</p>
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
            <div class="grid-card-icon" @click="navigateTo(config.VIEWS.admin.rackets)">
                <h1 class="font-grid-1">Ракетки</h1>
            </div>
            <div class="grid-card-icon" @click="navigateTo(config.VIEWS.admin.orders)">
                <h1 class="font-grid-1">Заказы</h1>
            </div>
            <div class="grid-card-icon" @click="navigateTo(config.VIEWS.admin.users)">
                <h1 class="font-grid-1">Пользователи</h1>
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
                    this.$router.push(this.config.VIEWS.auth.login);
                    return;
                }

                const response = await axios.get(`${config.BACKEND_URL}${config.API.user.profile}`, 
                {
                    headers: {
                        Authorization: `Bearer ${token}`
                    }
                });
                
                if (response.data) {
                    const user = response.data.user
                    this.userData = {   
                        name: user.name,
                        surname: user.surname,
                        email: user.email
                    };
                }
                this.AppMenuBar
        
            } catch (error) {
                console.error('Error fetching user data:', error);
                this.$router.push(this.config.VIEWS.auth.login);
            }
        },
        navigateTo(route) {
            this.$router.push(route);
        },
        logout() {
            localStorage.removeItem('token');
            this.$router.push(this.config.VIEWS.auth.login);
            window.location.reload();
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