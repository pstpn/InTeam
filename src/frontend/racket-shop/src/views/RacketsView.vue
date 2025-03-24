<template>
    <div class="container-page">
        <h1 class="font-container-header">Каталог</h1>
        
        <div class="grid-rackets">
            <div v-for="racket in rackets" :key="racket.id" class="grid-racket-card">
                <div class="grid-photo">
                    <img 
                        :src="racketImage(racket)" 
                        @error="handleImageError(racket)"
                    />
                </div>
                <p class="font-form-body-bold">{{ racket.brand }} {{ racket.id }}</p>
                <div class="form-in-row">
                    <router-link :to="`${config.API.rackets}/${racket.id}`" class="submit-button-font-thin">Описание</router-link> 
                    <button 
                        class="submit-button-green"
                        @click="addToCart(racket.id)">
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
            error: null,
            imageCache: {}
        };
    },
    methods: {
        async fetchRackets() {
            try {
                const cur_url = config.BACKEND_URL + config.API.rackets;
                const response = await axios.get(cur_url);

                if (response.data) {
                    this.rackets = response.data.rackets;
                    this.rackets.forEach(racket => {
                        this.getRacketPhoto(racket);
                    });
                }
            } catch (err) {
                this.error = err.response?.data?.message || 'Не удалось загрузить данные';
            }
        },
        
        async getRacketPhoto(racket) {
            try {
                if (racket.image && !this.imageCache[racket.id]) {
                    const cur_url = 'http://localhost:5000/convert';

                    const response = await axios.post(cur_url, {
                        image_data: racket.image
                    });
                    
                    this.imageCache[racket.id] = URL.createObjectURL(
                        new Blob([response.data], { type: 'image/jpeg' })
                    );
                    
                    this.$forceUpdate();
                }
            } catch (err) {
                console.error('Ошибка загрузки изображения:', err);
            }
        },     
        racketImage(racket) {
            return 'data:image/jpg;base64,' + racket.image;
        },
        handleImageError(racket) {
            return 'data:image/jpg;base64,' + racket.image;
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
                        "racket_id": racketId,
                        "quantity": 1 
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
    },
    beforeUnmount() {
        Object.values(this.imageCache).forEach(url => {
            URL.revokeObjectURL(url);
        });
    }
};
</script>

<style>
@import "../css/Forms.css";
@import "../css/Containers.css";
@import "../css/Fonts.css";
@import "../css/Grid.css";
</style>