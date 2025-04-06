<template>
    <div class="container-page">
        <h1 class="font-container-header">Заказы</h1>

        <div class="grid-order-column">  
            <div class="form-in-row-add-feedback">
                <div class="form-in-row-left">
                    <div class="filter-dropdown">
                        <button class="submit-button-filter" @click="toggleDropdown('total_price')">
                            Сумма
                            <span class="dropdown-arrow">▼</span>
                        </button>
                        <div class="dropdown-menu" v-show="activeDropdown === 'total_price'">
                            <button @click="sortBy('total_price', 'asc')">По возрастанию</button>
                            <button @click="sortBy('total_price', 'desc')">По убыванию</button>
                        </div>
                    </div>

                    <div class="filter-dropdown">
                        <button class="submit-button-filter" @click="toggleDropdown('creation_date')">
                            Создание
                            <span class="dropdown-arrow">▼</span>
                        </button>
                        <div class="dropdown-menu" v-show="activeDropdown === 'creation_date'">
                            <button @click="sortBy('creation_date', 'desc')">Сначала новые</button>
                            <button @click="sortBy('creation_date', 'asc')">Сначала старые</button>
                        </div>
                    </div>

                    <div class="filter-dropdown">
                        <button class="submit-button-filter" @click="toggleDropdown('delivery_date')">
                            Доставка
                            <span class="dropdown-arrow">▼</span>
                        </button>
                        <div class="dropdown-menu" v-show="activeDropdown === 'delivery_date'">
                            <button @click="sortBy('delivery_date', 'desc')">Сначала новые</button>
                            <button @click="sortBy('delivery_date', 'asc')">Сначала старые</button>
                        </div>
                    </div>

                    <!-- Фильтр по получателю -->
                    <div class="filter-dropdown">
                        <button class="submit-button-filter" @click="toggleDropdown('recipient')">
                            Получатель
                            <span class="dropdown-arrow">▼</span>
                        </button>
                        <div class="dropdown-menu" v-show="activeDropdown === 'recipient'">
                            <button v-for="recipient in uniqueRecipients" :key="recipient" 
                                    @click="filterByRecipient(recipient)">
                                {{ recipient }}
                            </button>
                        </div>
                    </div>

                    <!-- Фильтр по статусу -->
                    <div class="filter-dropdown">
                        <button class="submit-button-filter" @click="toggleDropdown('status')">
                            Статус заказа
                            <span class="dropdown-arrow">▼</span>
                        </button>
                        <div class="dropdown-menu" v-show="activeDropdown === 'status'">
                            <button @click="filterByStatus('Done')">Выполнены</button>
                            <button @click="filterByStatus('InProgress')">В процессе</button>
                            <button @click="filterByStatus('Canceled')">Отменен</button>
                        </div>
                    </div>

                    <button class="submit-button-orange" @click="resetFilters">
                        Сбросить
                    </button>
                </div>
            </div>  
        </div>

        <div class="grid-order-column">
            <div class="grid-card-name" v-for="order in filteredOrders" :key="order.id">
                <div class="form-in-row">
                    <div class="form-in-row-left">
                        <p class="font-form-body-bold">
                            Заказ №{{ order.id }}
                        </p>
                        <button :class="{
                            'submit-button-green': order.status === 'Done',
                            'submit-button-grey': order.status === 'InProgress',
                            'submit-button-orange': order.status === 'Canceled'
                        }">
                            {{ getStatusText(order.status) }}
                        </button>
                    </div>
                    <p class="font-form-body-bold">
                        Сумма: {{ order.total_price }} ₽
                    </p>
                </div>
                    <div class="form-in-row">
                        <p class="font-form-body">
                            Товары
                        </p>
                        <p class="font-form-body-bold">
                            {{ order.lines.length }} шт
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
                        Адрес
                    </p>
                    <p class="font-form-body-bold">
                        {{ order.address }}
                    </p>
                </div>
                <div class="form-in-row">
                    <p class="font-form-body">
                        Дата создания
                    </p>
                    <p class="font-form-body-bold">
                        {{ formatDate(order.creation_date) }}
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
</template>

<script>
import axios from 'axios';
import config from "../../../config.js";

export default {
    data() {
        return {
            config: config,
            orders: [],
            loading: true,
            activeDropdown: null,
            currentSort: {
                field: null,
                direction: null
            },
            recipientFilter: null,
            statusFilter: null
        };
    },
    computed: {
        uniqueRecipients() {
            const recipients = new Set();
            this.orders.forEach(order => {
                if (order.recipient_name) {
                    recipients.add(order.recipient_name);
                }
            });
            return Array.from(recipients);
        },
        
        filteredOrders() {
            let result = [...this.orders];
            
            // Фильтр по получателю
            if (this.recipientFilter) {
                result = result.filter(order => order.recipient_name === this.recipientFilter);
            }
            
            // Фильтр по статусу
            if (this.statusFilter) {
                result = result.filter(order => order.status === this.statusFilter);
            }
            
            // Сортировка
            if (this.currentSort.field) {
                result.sort((a, b) => {
                    let valueA = a[this.currentSort.field];
                    let valueB = b[this.currentSort.field];
                    
                    if (!isNaN(valueA) && !isNaN(valueB)) {
                        valueA = Number(valueA);
                        valueB = Number(valueB);
                    }
                    
                    if (valueA < valueB) return this.currentSort.direction === 'asc' ? -1 : 1;
                    if (valueA > valueB) return this.currentSort.direction === 'asc' ? 1 : -1;
                    return 0;
                });
            }
            
            return result;
        }
    },
    methods: {
        async fetchOrders() {
            try {
                this.loading = true;
                const token = localStorage.getItem('token');

                if (!token) {
                    this.$router.push(this.config.VIEWS.auth.login);
                    return;
                }

                const ordersResponse = await axios.get(`${config.BACKEND_URL}${config.API.orders}`, {
                    headers: { Authorization: `Bearer ${token}` }
                });

                // console.log(ordersResponse.data)
                if (ordersResponse.data) {
                    this.orders = ordersResponse.data.orders || [];
                }
            } catch (error) {
                console.error('Ошибка при получении заказов:', error);
                this.orders = [];
            } finally {
                this.loading = false;
            }
        },
        
        toggleDropdown(dropdownName) {
            this.activeDropdown = this.activeDropdown === dropdownName ? null : dropdownName;
        },
        
        sortBy(field, direction) {
            this.currentSort = { field, direction };
            this.activeDropdown = null;
        },
        
        filterByRecipient(recipient) {
            this.recipientFilter = recipient;
            this.activeDropdown = null;
        },
        
        filterByStatus(status) {
            this.statusFilter = status;
            this.activeDropdown = null;
        },
        
        resetFilters() {
            this.currentSort = { field: null, direction: null };
            this.recipientFilter = null;
            this.statusFilter = null;
        },
        
        closeDropdownsOnClickOutside(e) {
            if (!e.target.closest('.filter-dropdown')) {
                this.activeDropdown = null;
            }
        },
        
        getStatusText(status) {
            const statusMap = {
                'Done': 'Выполнен',
                'InProgress': 'В процессе',
                'Canceled': 'Отменен'
            };
            return statusMap[status] || status;
        },
        
        formatDate(dateString) {
            const date = new Date(dateString);
            
            // Форматируем дату
            const day = date.getDate().toString().padStart(2, '0');
            const year = date.getFullYear();
            
            // Форматируем время
            const hours = date.getHours().toString().padStart(2, '0');
            const minutes = date.getMinutes().toString().padStart(2, '0');
            
            // Названия месяцев на русском
            const monthNames = [
                'января', 'февраля', 'марта', 'апреля', 'мая', 'июня',
                'июля', 'августа', 'сентября', 'октября', 'ноября', 'декабря'
            ];
            
            return `${day} ${monthNames[date.getMonth()]} ${year} г., ${hours}:${minutes}`;
        }
    },
    mounted() {
        this.fetchOrders();
        document.addEventListener('click', this.closeDropdownsOnClickOutside);
    },
    beforeUnmount() {
        document.removeEventListener('click', this.closeDropdownsOnClickOutside);
    }
};
</script>

<style>
@import "../../css/Containers.css";
@import "../../css/Fonts.css";
@import "../../css/Grid.css";
@import "../../css/Forms.css";
@import "../../css/Icons.css";
@import "../../css/Menu.css";
</style>