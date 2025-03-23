<template>
    <div class="container-page">
        <h1 class="font-container-header">Мои заказы</h1>

        <div v-if="orders.length > 0">
            <div class="grid-order-column" v-for="(order, index) in orders" :key="index">
                <div class="grid-card-name">
                    <div class="form-in-row">
                        <button 
                            :class="getStatusButtonClass(order.status)"
                            @click="handleOrderAction(order)">
                            {{ getStatusButtonText(order.status) }}
                        </button>
                        <p class="font-form-body-bold">
                            Цена: {{ order.total_price }} ₽
                        </p>
                    </div>
                    <div class="form-in-row">
                        <p class="font-form-body">
                            Товары
                        </p>
                        <p class="font-form-body-bold">
                            {{ order.items.length }} шт.
                        </p>
                    </div>
                    <div class="form-in-row">
                        <p class="font-form-body">
                            Получатель
                        </p>
                        <p class="font-form-body-bold">
                            {{ order.recipient_name }}
                        </p>
                    </div>
                    <div class="form-in-row">
                        <p class="font-form-body">
                            Адрес доставки
                        </p>
                        <p class="font-form-body-bold">
                            {{ order.delivery_address }}
                        </p>
                    </div>
                    <div class="form-in-row">
                        <p class="font-form-body">
                            Дата доставки
                        </p>
                        <p class="font-form-body-bold">
                            {{ formatDate(order.delivery_date) }}
                        </p>
                    </div>
                </div>
            </div>
        </div>

        <div v-else class="font-form-body">
            <p>У вас пока нет заказов</p>
        </div>
    </div>
</template>

<script>
import axios from 'axios';
import config from "../../../config.js";

export default {
    name: 'AppOrders',
    data() {
        return {
            orders: []
        };
    },
    methods: {
        async fetchOrders() {
            try {
                const token = localStorage.getItem('token');

                if (!token) {
                    this.$router.push(this.config.API.auth.login);
                    return;
                }

                const response = await axios.get(`${config.BACKEND_URL}${config.API.user.orders}`, {
                    headers: {
                        Authorization: `Bearer ${token}`
                    }
                });

                if (response.data) {
                    this.orders = response.data.orders;
                }
            } catch (error) {
                console.error('Ошибка при получении заказов:', error);
            }
        },

        getStatusButtonClass(status) {
            switch (status) {
                case 'InProgress':
                    return 'status-button-inprogress';
                case 'Canceled':
                    return 'status-button-canceled';
                case 'Done':
                    return 'status-button-done';
                default:
                    return 'status-button-default';
            }
        },
        getStatusButtonText(status) {
            switch (status) {
                case 'InProgress':
                    return 'В обработке';
                case 'Canceled':
                    return 'Отменен';
                case 'Done':
                    return 'Завершен';
                default:
                    return 'Неизвестно';
            }
        },

        async handleOrderAction(order) {
            if (order.status === 'InProgress') {
                try {
                    const token = localStorage.getItem('token');
                    const response = await axios.post(
                        `${config.BACKEND_URL}${config.API.user.cancelOrder}/${order.id}`,
                        {},
                        {
                            headers: { Authorization: `Bearer ${token}` }
                        }
                    );

                    if (response.data) {
                        this.fetchOrders(); 
                    }
                } catch (error) {
                    console.error('Ошибка при отмене заказа:', error);
                }
            }
        },

        formatDate(dateString) {
            const date = new Date(dateString);
            return date.toLocaleDateString('ru-RU');
        }
    },
    mounted() {
        this.fetchOrders(); 
    }
};
</script>

<style>
@import "../../css/Containers.css";
@import "../../css/Fonts.css";
@import "../../css/Grid.css";
</style>