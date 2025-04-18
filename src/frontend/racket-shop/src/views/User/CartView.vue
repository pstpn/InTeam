<template>
    <div class="container-page">
        <h1 class="font-container-header">Моя корзина</h1>
        
        <div v-if="cartData.lines.length === 0">
            <p class="font-form-body">
                Ваша корзина пуста <router-link :to="config.VIEWS.rackets" class="submit-button-font">  За покупками</router-link> 
            </p>
        </div>

        <div v-else class="grid-cart">
            <div class="grid-cart-column">
                <div class="grid-item" v-for="(item, index) in cartData.lines" :key="index">
                    <div class="grid-cart-column">
                        <div class="grid-photo">
                            <img 
                                v-if="racketImages[item.racket_id]"
                                :src="racketImages[item.racket_id]" 
                                :alt="`Ракетка ${item.racket_id}`"
                                @error="handleImageError(item.racket_id)"
                            />
                            <p v-else class="font-form-body">Загрузка изображения...</p>
                        </div>
                    </div>
                    <div class="grid-cart-column">
                        <div class="form-in-row">
                            <router-link 
                                :to="`${config.VIEWS.rackets}/${item.racket_id}`" 
                                class="submit-button-font">
                                {{ racketNames[item.racket_id] }}
                            </router-link> 
                            <div class="grid-button">
                                <button class="submit-button-cart" @click="decreaseQuantity(item)">
                                    -
                                </button>
                                <p class="font-form-body-bold">{{ item.quantity }} шт.</p>
                                <button class="submit-button-cart" @click="increaseQuantity(item)">
                                    +
                                </button>
                            </div>
                        </div>
                        <div class="form-in-row">
                            <p class="font-form-body">
                                Цена
                            </p>
                            <p class="font-form-body-bold">
                                {{ item.price }} ₽
                            </p>
                        </div>
                        <div class="form-in-row-right">
                            <button 
                                class="submit-button-orange"
                                @click="openDeleteModal(item.racket_id)">
                                Удалить
                            </button>
                        </div>
                    </div>
                    <hr class="line">
                    <hr class="line">
                </div>
            </div>

            <div v-if="isDeleteModalOpen" class="modal-overlay">
                <div class="modal-content">
                    <p class="font-form-header">Подтверждение удаления</p>
                    <p class="font-form-body">Вы уверены, что хотите удалить эту ракетку из корзины?</p>
                    <div class="form-in-row-1">
                        <button 
                            class="submit-button-green" 
                            @click="closeDeleteModal">
                            Нет
                        </button>
                        <button 
                            class="submit-button-orange" 
                            @click="removeItem">
                            Да
                        </button>
                    </div>
                </div>
            </div>

            <div class="grid-cart-column">
                <div class="grid-card-name">
                    <div class="form-in-row">
                        <p class="font-form-body">
                            Товары
                        </p>
                        <p class="font-form-body-bold">
                            {{ cartData.quantity }} 
                        </p>
                    </div>
                    <div class="form-in-row">
                        <p class="font-form-body">
                            Итог
                        </p>
                        <p class="font-form-body-bold">
                            {{ cartData.total_price }} ₽
                        </p>
                    </div>
                    <button 
                        class="submit-button-green"
                        @click="showModal = true">
                        Оформить заказ
                    </button>
                </div>
            </div>
        </div>

        <div v-if="showModal" class="modal-overlay">
            <div class="modal-content">
                <h2 class="font-form-header">Оформление заказа</h2>
                <div class="form-input">
                    <input
                        type="datetime-local"
                        id="delivery_date"
                        v-model="delivery_date"
                        placeholder="Дата доставки"
                        required
                    />
                </div>
                <div class="form-input">
                    <input
                        type="text"
                        id="recipient_name"
                        v-model="recipient_name"
                        placeholder="Введите получателя"
                        required
                    />
                </div>
                <div class="form-input">
                    <input
                        type="text"
                        id="address"
                        v-model="address"
                        placeholder="Введите адрес"
                        required
                    />
                </div>
                <div class="form-in-row-1">
                    <button 
                        class="submit-button-orange" 
                        @click="showModal = false">
                        Отмена
                    </button>
                    <button 
                        class="submit-button-green" 
                        @click="confirmOrder">
                        Подтвердить
                    </button>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import axios from 'axios';
import config from "../../../config.js"

export default {
    name: 'AppMenuBar',
    data() {
        return {
            config: config,
            cartData: {
                lines: [], 
                quantity: 0, 
                total_price: 0 
            },
            showModal: false,
            address: '',
            delivery_date: '',
            recipient_name: '',
            isDeleteModalOpen: false,
            selectedRacketId: 0,
            racketImages: {}, 
            racketNames: {},  
            loadingRackets: {} 
        };
    },
    methods: {
        async fetchCartData() {
            try {
                const token = localStorage.getItem('token');

                if (!token) {
                    this.$router.push(this.config.VIEWS.auth.login);
                    return;
                }

                const response = await axios.get(`${config.BACKEND_URL}${config.API.user.cart}`, {
                    headers: { Authorization: `Bearer ${token}` }
                });
                
                if (response.data) {
                    const cart = response.data.cart;
                    this.cartData = {   
                        lines: cart.lines, 
                        total_price: cart.total_price, 
                        quantity: cart.quantity
                    };
                    
                    this.cartData.lines.forEach(item => {
                        this.fetchRacketData(item.racket_id);
                    });
                }
            } catch (error) {
                this.$router.push(this.config.VIEWS.auth.login);
            }
        },
        
        async fetchRacketData(racketId) {
            if (this.loadingRackets[racketId]) return;
            
            this.loadingRackets = {...this.loadingRackets, [racketId]: true};
            
            try {
                const response = await axios.get(`${config.BACKEND_URL}${config.API.rackets}/${racketId}`);
                
                if (response.data) {
                    const racket = response.data.racket;
                
                    this.racketNames = {...this.racketNames, [racketId]: `${racket.brand} ${racket.id}`};
                    
                    if (racket.image) {
                        this.racketImages = {...this.racketImages, [racketId]: `data:image/jpeg;base64,${racket.image}`};
                    }
                }
            } catch (error) {
                console.error(`Error fetching racket ${racketId}:`, error);
                this.racketImages = {...this.racketImages, [racketId]: require('../../assets/racket.png')};
            } finally {
                this.loadingRackets = {...this.loadingRackets, [racketId]: false};
            }
        },
        
        handleImageError(racketId) {
            this.racketImages = {...this.racketImages, [racketId]: require('../../assets/racket.png')};
        },
        
        async increaseQuantity(item) {
            try {
                const token = localStorage.getItem('token');
                const response = await axios.put(
                    `${config.BACKEND_URL}${config.API.user.cart}/rackets/${item.racket_id}`,
                    { quantity: 1 },
                    { headers: { Authorization: `Bearer ${token}` } }
                );
                if (response.data) {
                    this.fetchCartData();
                }
            } catch (error) {
                console.error('Error increasing quantity:', error);
            }
        },
        
        async decreaseQuantity(item) {
            try {
                const token = localStorage.getItem('token');
                const response = await axios.put(
                    `${config.BACKEND_URL}${config.API.user.cart}/rackets/${item.racket_id}`,
                    { quantity: -1 },
                    { headers: { Authorization: `Bearer ${token}` } }
                );

                if (response.data) {
                    this.fetchCartData();
                }
            } catch (error) {
                console.error('Error decreasing quantity:', error);
            }
        },
        
        async removeItem() {
            try {
                const token = localStorage.getItem('token');
                const response = await axios.delete(
                    `${config.BACKEND_URL}${config.API.user.cart}/rackets/${this.selectedRacketId}`,
                    { headers: { Authorization: `Bearer ${token}` } }
                );
                if (response.data) {
                    this.fetchCartData();
                    this.closeDeleteModal();
                }
            } catch (error) {
                console.error('Error removing item:', error);
            }
        },
        
        async confirmOrder() {
            try {
                if (!this.address.trim() || !this.delivery_date || !this.recipient_name.trim()) {
                    alert('Пожалуйста, заполните все поля');
                    return;
                }

                const token = localStorage.getItem('token');
                
                const formatDate = (dateString) => {
                    const date = new Date(dateString);
                    const timezoneOffset = -date.getTimezoneOffset();
                    const offsetHours = Math.abs(Math.floor(timezoneOffset / 60));
                    const offsetMinutes = Math.abs(timezoneOffset % 60);
                    const offsetSign = timezoneOffset >= 0 ? '+' : '-';
                    const offset = `${offsetSign}${offsetHours.toString().padStart(2, '0')}:${offsetMinutes.toString().padStart(2, '0')}`;
                    
                    return `${date.getFullYear()}-${(date.getMonth()+1).toString().padStart(2, '0')}-${date.getDate().toString().padStart(2, '0')}T${date.getHours().toString().padStart(2, '0')}:${date.getMinutes().toString().padStart(2, '0')}:00${offset}`;
                };

                const response = await axios.post(
                    `${config.BACKEND_URL}${config.API.orders}`,
                    {
                        address: this.address.trim(),
                        delivery_date: formatDate(this.delivery_date),
                        recipient_name: this.recipient_name.trim()
                    },
                    { headers: { Authorization: `Bearer ${token}` } }
                );
                
                if (response.status == 200) {
                    this.showModal = false;
                    this.$router.push(this.config.VIEWS.user.orders);
                }
            } catch (error) {
                alert('Ошибка при оформлении заказа: ' + (error.response?.data?.message || error.message));
            }
        },
        
        openDeleteModal(racketId) {
            this.selectedRacketId = racketId;
            this.isDeleteModalOpen = true;
        },

        closeDeleteModal() {
            this.isDeleteModalOpen = false;
            this.selectedRacketId = null;
        }
    },
    mounted() {
        this.fetchCartData();
    }
}
</script>

<style>
@import "../../css/Containers.css";
@import "../../css/Fonts.css";
@import "../../css/Grid.css";
</style>