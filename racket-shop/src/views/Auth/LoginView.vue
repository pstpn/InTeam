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
                    type="password"
                    id="password"
                    v-model="password"
                    placeholder="Пароль"
                    :class="{ 'input-error': error }"
                    @input="resetError"
                    required
                />
            </div>

            <div class="form-in-row" v-if="showError">
                <p class="font-form-body-error">{{ errorMessage }}</p>
            </div>

            <div class="form-in-row">
                <p class="font-form-body">
                    Нет аккаунт? <router-link to="/auth/register" class="submit-button-font">Зарегистрироваться</router-link> 
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
import backend_url from "../../../config.js"

export default {
    data() {
        return {
            email: '',
            password: '',
            error: false,
            showError: false,
            errorMessage: 'Некорректные данные для ввода!'
        };
    },
    methods: {
        async login() {

            // const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
            // if (!this.email || !emailRegex.test(this.email)) {
            //     this.showError = true;
            //     this.error = true;

            //     this.errorMessage = 'Некорректный формат электронной почты!';
            //     return;
            // }

            // if (!this.password || this.password.length == 0) {
            //     this.showError = true;
            //     this.error = true;

            //     this.errorMessage = 'Пароль не должен быть пустым!';
            //     return;
            // }

            try {
                const cur_url = backend_url + 'api/auth/login'
                const response = await axios.post(cur_url, {
                    email: this.email,
                    password: this.password
                });
                
                if (response.data) {
                    console.log('Login successful:', response.data);
                    localStorage.setItem('token', response.data.access_token);

                    this.error = false;
                    this.$router.push('/profile');
                }
            } catch (error) {
                console.error('Error logging in:', error);
                this.error = true;

                this.showError = true;
            }
        },
        resetError() {
            this.error = false;
            this.showError = false;
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
</style>