<template>
    <div class="container-page">
        <h1 class="font-container-header">Каталог</h1>
        
        <div class="grid-rackets">
            <div v-for="racket in rackets" :key="racket.id" class="grid-racket-card">
                <div class="grid-photo">
                    <img 
                        v-if="racket.imageData"
                        :src="'data:image/jpeg;base64,' + racket.imageData" 
                        :alt="`${racket.brand} ${racket.id}`"
                        @error="handleImageError(racket)"
                    />
                    <p v-else class="loading-text">Загрузка изображения...</p>
                </div>
                <p class="font-form-body-bold">{{ racket.brand }} {{ racket.id }}</p>
                <div class="form-in-row">
                    <router-link 
                        :to="`${config.VIEWS.rackets}/${racket.id}`" 
                        class="submit-button-font-thin"
                    >
                        Описание
                    </router-link> 
                    <button 
                        class="submit-button-green"
                        @click="addToCart(racket.id)"
                        :disabled="!racket.imageLoaded"
                    >
                        В корзину
                    </button>
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
    methods: {
        async fetchRackets() {
            try {
                const response = await axios.get(`${config.BACKEND_URL}${config.API.rackets}`);
                if (response.data) {
                    this.rackets = response.data.rackets.map(racket => ({
                        ...racket,
                        imageData: null,
                        imageLoaded: false,
                        imageError: false
                    }));
                    // Загружаем изображения для каждой ракетки
                    this.rackets.forEach(racket => this.loadRacketImage(racket));
                }
            } catch (err) {
                this.error = err.response?.data?.message || 'Не удалось загрузить данные';
                console.error('Ошибка загрузки ракеток:', err);
            } finally {
                this.loading = false;
            }
        },
        
        async loadRacketImage(racket) {
            try {
                if (racket.image) {
                    // Если изображение уже в Base64
                    racket.imageData = racket.image;
                    racket.imageLoaded = true;
                } else {
                    // Если нужно загружать отдельно
                    const response = await axios.get(
                        `${config.BACKEND_URL}${config.API.rackets}/${racket.id}/image`,
                        { responseType: 'arraybuffer' }
                    );
                    
                    const base64 = btoa(
                        new Uint8Array(response.data)
                            .reduce((data, byte) => data + String.fromCharCode(byte), '')
                    );
                    
                    racket.imageData = base64;
                    racket.imageLoaded = true;
                }
            } catch (err) {
                console.error('Ошибка загрузки изображения:', err);
                racket.imageError = true;
            }
        },
        
        handleImageError(racket) {
            racket.imageError = true;
            racket.imageLoaded = false;
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
            } catch (err) {
                const errorMsg = err.response?.data?.message || 'Не удалось добавить в корзину';
                alert(errorMsg);
            }
        }
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
</style>