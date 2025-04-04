<template>
    <div class="container-page container-single-form">
        <div class="form">
            <h1 class="font-form-header">Регистрация</h1>

            <div class="form-input">
                <input
                    type="text"
                    id="name"
                    v-model="name"
                    placeholder="Имя"
                    @input="resetError"
                    required
                />
            </div>

            <div class="form-input">
                <input
                    type="text"
                    id="surname"
                    v-model="surname"
                    placeholder="Фамилия"
                    @input="resetError"
                    required
                />
            </div>

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
                    Есть аккаунт? <router-link to="/auth/login" class="submit-button-font">Войти</router-link> 
                </p>
                <button 
                    class="submit-button-green"
                    @click="register">
                    Зарегестрироваться
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
            name: '',
            surname: '',
            email: '',
            password: '',
            error: false,
            showError: false,
            errorMessage: 'Некорректные данные для ввода!',
            showPassword: false
        };
    },
    methods: {
        async register() {
            try {
                const cur_url = config.BACKEND_URL + config.API.auth.register;
                const response = await axios.post(cur_url, {
                    name: this.name,
                    surname: this.surname,
                    email: this.email,
                    password: this.password
                });

                if (response.data) {
                    console.log('Register successful:', response.data);
                    localStorage.setItem('token', response.data.access_token);

                    this.error = false;
                    this.$router.push(config.VIEWS.user.profile);
                }
            } catch (error) {
                console.error('Error register in:', error);
                
                this.error = true;
                this.showError = true;
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
</style>