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
                    required
                />
            </div>

            <div class="form-input">
                <input
                    type="text"
                    id="surname"
                    v-model="surname"
                    placeholder="Фамилия"
                    required
                />
            </div>

            <div class="form-input">
                <input
                    type="text"
                    id="email"
                    v-model="email"
                    placeholder="Электронная почта"
                    required
                />
            </div>

            <div class="form-input">
                <input
                    type="password"
                    id="password"
                    v-model="password"
                    placeholder="Пароль"
                    required
                />
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
import backend_url from "../../../config.js"

export default {
    data() {
        return {
            email: '',
            password: ''
        };
    },
    methods: {
        async register() {
            try {
                const cur_url = backend_url + 'api/auth/register'
                const response = await axios.post(cur_url, {
                    name: this.name,
                    surname: this.surname,
                    email: this.email,
                    password: this.password
                });

                if (response.data) {
                    console.log('Register successful:', response.data);
                }
            } catch (error) {
                console.error('Error logging in:', error);
            }
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