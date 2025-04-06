<template>
    <div class="container-page container-single-form">
        <div class="form">
            <h1 class="font-form-header">Авторизация</h1>

            <div class="form-input">
                <input
                    type="text"
                    id="email"
                    v-model="email"
                    placeholder="Электронная почта"
                    :class="{ 'input-error': error }"
                    @input="resetError"
                    required
                />
            </div>

            <div class="form-input">
                <input
                    :type="showPassword ? 'text' : 'password'"
                    id="password"
                    v-model="password"
                    placeholder="Пароль"
                    :class="{ 'input-error': error }"
                    @input="resetError"
                    required
                />
                <img 
                    :src="showPassword ? require('@/assets/noview.png') : require('@/assets/view.png')" 
                    class="icon" 
                    @click="togglePasswordVisibility"
                />
            </div>
        
            <div class="form-in-row" v-if="showError">
                <p class="font-form-body-error">{{ errorMessage }}</p>
            </div>

            <div class="form-in-row">
                <p class="font-form-body">
                    Нет аккаунта? <router-link :to="config.VIEWS.auth.register" class="submit-button-font">Зарегистрироваться</router-link> 
                </p>
                <button 
                    class="submit-button-green"
                    @click="login">
                    Войти
                </button>
            </div>
        </div>
    </div>
</template>

<script>
import axios from 'axios';
import config from "../../../config.js"

export default {
    data() {
        return {
            config: config,
            email: '',
            password: '',
            error: false,
            showError: false,
            errorMessage: 'Некорректные данные для ввода!',
            showPassword: false, 
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
        async login() {
            try {
                const cur_url = config.BACKEND_URL + config.API.auth.login;
                const response = await axios.post(cur_url, {
                    email: this.email,
                    password: this.password
                });
                
                if (response.data?.access_token) {
                    localStorage.setItem('token', response.data.access_token);
                    
                    const token = response.data.access_token;
                    const decodedToken = this.parseJwt(token);
                    
                    this.error = false;
                    if (decodedToken?.Role === 'Admin') {
                        this.$router.push(config.VIEWS.admin.profile).then(() => {
                            window.location.reload()
                        });
                    } else {
                        this.$router.push(config.VIEWS.user.profile);
                    }
                } else {
                    throw new Error('No access token in response');
                }

            } catch (error) {
                console.error('Error logging in:', error);
                this.error = true;
                this.showError = true;
                this.errorMessage = error.response?.data?.message || 
                                'Ошибка входа. Проверьте email и пароль';
            }
        },
        resetError() {
            this.error = false;
            this.showError = false;
        },
        togglePasswordVisibility() {
            this.showPassword = !this.showPassword;
        }
    }
};
</script>

<style>
@import "../../css/Containers.css";
@import "../../css/Forms.css";
@import "../../css/Fonts.css";
@import "../../css/Input.css";
@import "../../css/Buttons.css";

.password-toggle-icon {
    position: absolute;
    right: 10px;
    top: 50%;
    transform: translateY(-50%);
    cursor: pointer;
    width: 20px;
    height: 20px;
}
</style>