<template>
    <div class="container-page">
        <h1 class="font-container-header">Заказы</h1>

        <div class="grid-order-column">    
            <div class="form-in-row-add-feedback">
                <div class="form-in-row-left">
            
                    <select class="submit-button-filter" @change="sortBy('total_price', $event.target.value)">
                        <option value="">Сумма</option>
                        <option value="asc">По возрастанию</option>
                        <option value="desc">По убыванию</option>
                    </select>

                    <select class="submit-button-filter" @change="sortBy('creation_date', $event.target.value)">
                        <option value="">Создание</option>
                        <option value="desc">Сначала новые</option>
                        <option value="asc">Сначала старые</option>
                    </select>
                
                    <select class="submit-button-filter" @change="sortBy('delivery_date', $event.target.value)">
                        <option value="">Доставка</option>
                        <option value="desc">Сначала новые</option>
                        <option value="asc">Сначала старые</option>
                    </select>

                    <select class="submit-button-filter" @change="filterByRecipient($event.target.value)">
                        <option value="">Получатель</option>
                        <option v-for="recipient in uniqueRecipients" 
                                :key="recipient" 
                                :value="recipient">
                            {{ recipient }}
                        </option>
                    </select>
                
                    <select class="submit-button-filter" @change="filterByStatus($event.target.value)">
                        <option value="">Статус заказа</option>
                        <option value="Done">Выполнены</option>
                        <option value="InProgress">В процессе</option>
                        <option value="Canceled">Отменен</option>
                    </select>
                    
                </div>
                <button class="submit-button-orange" @click="resetFilters">
                    Сбросить
                </button>
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
                        <a class="container-icon-tool" @click.stop="openStatusModal(order)">
                            <img src="../../assets/tool.png" alt="Tool" class="icon-tool"/>
                        </a>
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
   
    <div v-if="showStatusModal" class="modal-overlay" @click.self="closeStatusModal">
        <div class="modal-content">
            <p class="font-form-header">Изменение статуса заказа №{{ selectedOrderId }}</p>
            <div class="form-input">
                <select v-model="selectedStatus">
                    <option value="InProgress">В процессе</option>
                    <option value="Done">Выполнен</option>
                    <option value="Canceled">Отменен</option>
                </select>
            </div>
            <div class="form-in-row">
                <button 
                    class="submit-button-orange"
                    @click="closeStatusModal">
                    Отмена
                </button>
                <button 
                    class="submit-button-green"
                    @click="updateOrderStatus">
                    Сохранить
                </button>
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
            showStatusModal: false,
            selectedStatus: '',
            selectedOrderId: null,
            orders: [],
            config: config,
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
            
            if (this.recipientFilter) {
                result = result.filter(order => order.recipient_name === this.recipientFilter);
            }
            
            if (this.statusFilter) {
                result = result.filter(order => order.status === this.statusFilter);
            }
            
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
        openStatusModal(order) {
            this.selectedOrderId = order.id;
            this.selectedStatus = order.status;
            this.showStatusModal = true;
        },
        
        closeStatusModal() {
            this.showStatusModal = false;
            this.selectedOrderId = null;
            this.selectedStatus = '';
        },
        
        async updateOrderStatus() {
            try {
                const token = localStorage.getItem('token');
                if (!token) {
                    this.$router.push(this.config.VIEWS.auth.login);
                    return;
                }

                const response = await axios.patch(
                    `${config.BACKEND_URL}${config.API.orders}/${this.selectedOrderId}`,
                    {
                        status: this.selectedStatus
                    },
                    {
                        headers: {
                            'Authorization': `Bearer ${token}`
                        }
                    }
                );
                
                if (response.status === 200) {
                    this.closeStatusModal();
                    this.fetchOrders();
                    alert('Статус заказа успешно обновлен!');
                }
            } catch (error) {
                console.error('Ошибка при обновлении статуса:', error);
                alert('Не удалось обновить статус: ' + 
                    (error.response?.data?.message || error.message));
            }
        },
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
            
            const day = date.getDate().toString().padStart(2, '0');
            const year = date.getFullYear();
            
            const hours = date.getHours().toString().padStart(2, '0');
            const minutes = date.getMinutes().toString().padStart(2, '0');
            
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
</style>