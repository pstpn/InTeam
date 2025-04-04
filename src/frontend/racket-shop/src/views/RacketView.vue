<template>
    <div class="container-page">
        <h1 class="font-container-header">{{ 'Ракетка '+ racket.brand + ' ' + racket.id || 'Загрузка...' }}</h1>

        <div v-if="loading" class="loading-indicator">
            Загрузка данных...
        </div>

        <div v-else-if="error" class="error-message">
            {{ error }}
        </div>

        <div v-else class="grid-racket">
            <div class="grid-cart-column">
                <div class="grid-photo">
                    <img src='../assets/racket.png'>
                </div>
            </div>
            <div class="grid-cart-column">
                <div class="form-in-row">
                    <p class="font-form-body">Цена</p>
                    <p class="font-form-body-bold">{{ racket.price }} ₽</p>
                </div>
                
                <div class="form-in-row">
                    <p class="font-form-body">Бренд</p>
                    <p class="font-form-body-bold">{{ racket.brand }}</p>
                </div>

                <div class="form-in-row">
                    <p class="font-form-body">Баланс</p>
                    <p class="font-form-body-bold">{{ racket.balance }} мм</p>
                </div>

                <div class="form-in-row">
                    <p class="font-form-body">Размер головы</p>
                    <p class="font-form-body-bold">{{ racket.head_size }} кв.см</p>
                </div>

                <div class="form-in-row">
                    <p class="font-form-body">Вес</p>
                    <p class="font-form-body-bold">{{ racket.weight }} г</p>
                </div>
                <div class="form-in-row-right">
                    <button class="submit-button-green" @click="addToCart(racket.id)">
                        Добавить в корзину
                    </button>
                </div>
            </div>
        </div>

        <div class="grid-order-column" v-if="!loading && !error">
            <div v-if="feedbacks.length === 0" class="no-feedbacks">
                <p>Пока нет отзывов</p>
            </div>
            
            <div class="grid-card-feedback" v-for="feedback in feedbacks" :key="feedback.id">
                <hr class="line">
                <div class="grid-feedback">
                    <p class="font-form-header">Пользователь {{ feedback.user_id }}</p>
                    
                    <div class="form-in-row-racket">
                        <div class="grid-photo-ball">
                            <p class="font-form-body">
                                {{ feedback.date }} 
                            </p>
                            <img 
                                v-for="ball in 5" 
                                :key="ball" 
                                src="../assets/ball.png"
                                :style="{ opacity: ball <= feedback.rating ? 1 : 0.3 }"
                            >
                        </div>
                    </div>
                </div>
                <p class="font-form-body">
                    {{ feedback.feedback }}
                </p>
            </div>
        </div>
    </div>
</template>

<script>
import axios from 'axios';
import config from '../../config.js';

export default {
    data() {
        return {
            config: config,
            racket: {
                name: '',
                price: 0,
                brand: '',
                balance: 0,
                head_size: 0,
                weight: 0,
                image_url: ''
            },
            feedbacks: [],
            loading: true,
            error: null
        };
    },
    mounted() {
        this.fetchRacketData();
        this.fetchFeedbacks();
    },
    methods: {
        async fetchRacketData() {
            try {
                const racketId = this.$route.params.id;
                const response = await axios.get(`${config.BACKEND_URL}${config.API.rackets}/${racketId}`);
                this.racket = response.data.racket;

                console.log(this.racket)
            } catch (error) {
                this.error = 'Ошибка при загрузке данных ракетки';
                console.error('Ошибка загрузки ракетки:', error);
            } finally {
                this.loading = false;
            }
        },
        async fetchFeedbacks() {
            try {
                const racketId = this.$route.params.id;
                const response = await axios.get(`${config.BACKEND_URL}${config.API.feedbacks}/${racketId}`);
                this.feedbacks = response.data.feedbacks;

                console.log(this.feedbacks)
            } catch (error) {
                console.error('Ошибка при загрузке отзывов:', error);
            }
        },
        formatDate(dateString) {
            if (!dateString) return '';
            const date = new Date(dateString);
            return date.toLocaleDateString('ru-RU');
        },
        async addToCart(racketId) {
            try {
                const token = localStorage.getItem('token');
                if (!token) {
                    this.$router.push(`${config.API.auth.login}`);
                    return;
                }

                await axios.post(
                    `${config.BACKEND_URL}${config.API.user.cart}`,
                    { 
                        racket_id: racketId,
                        quantity: 1 
                    },
                    {
                        headers: {
                            'Authorization': `Bearer ${token}`
                        }
                    }
                );
                
                alert('Ракетка добавлена в корзину!');
            } catch (error) {
                console.error('Ошибка при добавлении в корзину:', error);
                alert('Не удалось добавить ракетку в корзину');
            }
        }
    }
};
</script>

<style>
@import "../css/Forms.css";
@import "../css/Containers.css";
@import "../css/Fonts.css";
@import "../css/Grid.css";
</style>