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
                    <img :src="'data:image/jpeg;base64,' + imageData" alt="Динамическое изображение">
                </div>
            </div>
            
            <div class="grid-cart-column">
                <div class="grid-card-feedback-1">
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
                        <button
                             class="submit-button-green" 
                            @click="addToCart(racket.id)"
                            :disabled="!racket.available">
                            Добавить в корзину
                        </button>
                    </div>
                </div>
            </div>
        </div>

        <div class="grid-order-column" v-if="!loading && !error">
            <div v-if="feedbacks.length === 0" class="grid-card-name-1">
                <hr class="line">
                
                <p class="font-form-body">
                    Отзывов нет <router-link :to="config.VIEWS.user.feedbacks" class="submit-button-font"> Будьте первым</router-link> 
                </p>
            </div>
            
            <div class="grid-card-feedback-1" v-for="feedback in feedbacks" :key="feedback.id">
                <hr class="line">
                <div class="grid-feedback">
                    <p class="font-form-header">{{ feedback.username }}</p>
                    
                    <div class="form-in-row-right">
                        <p class="font-form-body feedback-date">
                            {{ formatDate(feedback.date) }} 
                        </p>
                        
                        <div class="form-in-row-right">
                            <img 
                                v-for="(ball, ballIndex) in feedback.rating" 
                                :key="ballIndex"
                                src="../assets/ball.png"
                                class="rating-ball"
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
            imageData: null,
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

                this.imageData = this.racket.image;

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

            } catch (error) {
                console.error('Ошибка при загрузке отзывов:', error);
            }
        },
        formatDate(dateString) {
            const date = new Date(dateString);
            
            const day = date.getDate().toString().padStart(2, '0');
            const year = date.getFullYear();
         
            const monthNames = [
                'января', 'февраля', 'марта', 'апреля', 'мая', 'июня',
                'июля', 'августа', 'сентября', 'октября', 'ноября', 'декабря'
            ];
            
            return `${day} ${monthNames[date.getMonth()]} ${year} г.`;
        },
        async addToCart(racketId) {
            try {
                const token = localStorage.getItem('token');
                if (!token) {
                    this.$router.push(`${config.VIEWS.auth.login}`);
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
                this.$router.push(`${config.VIEWS.user.cart}`);
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

button:disabled {
    background-color: #d9d9d9 !important;
    cursor: not-allowed;
}
</style>