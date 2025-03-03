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
                <button class="submit-button-green">Зарегестрироваться</button>
            </div>
        </div>
    </div>
</template>

<script>
</script>

<style>
@import "../../css/Containers.css";
@import "../../css/Forms.css";
@import "../../css/Fonts.css";
@import "../../css/Input.css";
@import "../../css/Buttons.css";
</style>