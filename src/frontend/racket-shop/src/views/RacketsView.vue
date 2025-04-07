<template>
    <div class="container-page">
        <h1 class="font-container-header">Каталог</h1>
        
        <div class="grid-rackets">
            <div 
                v-for="racket in filteredRackets" 
                :key="racket.id" 
                class="grid-racket-card"
                :class="{ 'unavailable': !racket.available }"
            >
                <div class="grid-photo">
                    <img 
                        v-if="racket.image"
                        :src="'data:image/jpeg;base64,' + racket.image" 
                        :alt="`${racket.brand} ${racket.id}`"
                        @error="handleImageError(racket)"
                        :class="{ 'grayscale': !racket.available }"
                    />
                    <p v-else class="font-form-body">Загрузка изображения...</p>
                </div>
                <p class="font-form-body-bold">{{ racket.brand }} {{ racket.id }}</p>
                <div class="form-in-row">
                    <router-link 
                        :to="`${config.VIEWS.rackets}/${racket.id}`" 
                        class="submit-button-font-thin"
                        :class="{ 'description-link': !racket.available }"
                    >
                        Описание
                    </router-link> 
                    <button 
                        class="submit-button-green"
                        @click="addToCart(racket.id)"
                        :disabled="!racket.available">
                        В корзину
                    </button>
                </div>
                <div v-if="!racket.available" class="font-form-body-error">
                    Нет в наличии
                </div>
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
            rackets: [],
            loading: true,
            error: null
        };
    },
    computed: {
        filteredRackets() {
            return this.rackets;
        }
    },
    methods: {
        async fetchRackets() {
            try {
                const response = await axios.get(`${config.BACKEND_URL}${config.API.rackets}`);
                
                if (response.data) {
                    this.rackets = response.data.rackets.map(racket => ({
                        ...racket,
                        available: racket.available
                    }));
                }
            } catch (err) {
                this.error = err.response?.data?.message || 'Не удалось загрузить данные';
                console.error('Ошибка загрузки ракеток:', err);
            } finally {
                this.loading = false;
            }
        },
        
        async addToCart(racketId) {
            try {
                const token = localStorage.getItem('token');
                if (!token) {
                    this.$router.push(`${config.VIEWS.auth.login}`);
                    return;
                }

                const isAdmin = await this.isAdmin();
                if (isAdmin) {
                    alert('Администраторы не могут добавлять товары в корзину');
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
                
                this.$router.push(`${config.VIEWS.user.cart}`);
            } catch (error) {
                console.error('Ошибка при добавлении в корзину:', error);
            }
        },

        async isAdmin() {
            const token = localStorage.getItem('token');
            if (!token) return false;
            
            try {
                const payload = JSON.parse(atob(token.split('.')[1]));
                return payload.Role === 'Admin';

            } catch (error) {
                console.error('Ошибка проверки роли:', error);
                return false;
            }
        },
    },
    mounted() {
        this.fetchRackets();
    }
};
</script>

<style>
@import "../css/Forms.css";
@import "../css/Containers.css";
@import "../css/Fonts.css";
@import "../css/Grid.css";

.description-link {
    color: #333 !important;
    text-decoration: underline !important;
    pointer-events: auto !important;
    position: relative;
    z-index: 2;
}

.unavailable .submit-button-font-thin {
    opacity: 1 !important;
    color: #333 !important;
}

.unavailable::after {
    pointer-events: none;
}

.unavailable {
    opacity: 0.7;
    position: relative;
}

button:disabled {
    background-color: #d9d9d9 !important;
    cursor: not-allowed;
}
</style>