<template>
    <div class="container-page">
        <h1 class="font-container-header">Мои отзывы</h1>

        <div v-if="feedbacks.length === 0">
            <p class="font-form-body">
                Вы не оставили еще ни один отзыв
                <button 
                    class="submit-button-green" 
                    @click="openAddFeedbackModal">
                    Добавить отзыв
                </button>
            </p>
        </div>

        <div v-else class="grid-order-column">
            <div class="form-in-row-add-feedback">
                <div class="form-in-row-left">
                    <div class="filter-dropdown">
                        <button class="submit-button-filter" @click="toggleDropdown('rating')">
                            Рейтинг
                            <span class="dropdown-arrow">▼</span>
                        </button>
                        <div class="dropdown-menu" v-show="activeDropdown === 'rating'">
                            <button v-for="n in 5" :key="n" @click="filterByRating(n)">
                                {{ n }} звезд{{ n === 1 ? 'а' : n < 5 ? 'ы' : '' }}
                            </button>
                        </div>
                    </div>

                    <div class="filter-dropdown">
                        <button class="submit-button-filter" @click="toggleDropdown('date')">
                            Дата
                            <span class="dropdown-arrow">▼</span>
                        </button>
                        <div class="dropdown-menu" v-show="activeDropdown === 'date'">
                            <button @click="sortBy('date', 'desc')">Сначала новые</button>
                            <button @click="sortBy('date', 'asc')">Сначала старые</button>
                        </div>
                    </div>
                    <button 
                        class="submit-button-orange" 
                        @click="resetFilters"
                        :disabled="!isFilterActive">
                        Сбросить
                    </button>
                </div>
                <button 
                    class="submit-button-green" 
                    @click="openAddFeedbackModal">
                    Добавить отзыв
                </button>
            </div>
            <div class="grid-card-feedback-2" 
                 v-for="(feedback, index) in filteredFeedbacks" 
                :key="index">
                <div class="grid-feedback">
                    <p class="font-form-header" @click="goToRacket(feedback.racket_id)">
                        {{ feedback.racket_id + ' ' +  feedback.racket_brand}}
                    </p>
                    
                    <div class="form-in-row-right">
                        <p class="font-form-body feedback-date">
                            {{ formatDate(feedback.date) }} 
                        </p>
                        
                        <div class="form-in-row-right">
                            <img 
                                v-for="(ball, ballIndex) in feedback.rating" 
                                :key="ballIndex"
                                src="../../assets/ball.png"
                                class="rating-ball"
                                >
                        </div>
                    </div>
                </div>
                <p class="font-form-body">
                    {{ feedback.feedback }} 
                </p>
                <div class="form-in-row-right">
                    <button 
                        class="submit-button-orange"
                        @click="openDeleteModal(feedback.racket_id)">
                        Удалить
                    </button>
                </div>
                <hr class="line">
            </div>
        </div>
    </div>
    <div v-if="showAddFeedbackModal" class="modal-overlay" @click.self="closeAddFeedbackModal">
        <div class="modal-content">
            <h2 class="font-form-header">Добавить отзыв</h2>
            
            <div class="form-in-row">
                <select 
                    v-model="newFeedback.racket_id" 
                    class="feedback-select"
                    required
                    @change="onRacketSelect"
                >
                    <option value="" disabled selected>Выберите ракетку</option>
                    <option 
                        v-for="racket in userRackets" 
                        :key="racket.racket.id" 
                        :value="racket.racket.id"
                        :disabled="hasFeedbackForRacket(racket.id)"
                    >
                        {{ getRacketDisplayName(racket) }}
                        <span v-if="hasFeedbackForRacket(racket.id)">(отзыв уже оставлен)</span>
                    </option>
                </select>
            </div>
            
            <div class="form-in-row">
                <p class="font-form-body">
                    Оценка
                </p>
                <div class="grid-photo-ball-feedback">
                    <img 
                        v-for="n in 5" 
                        :key="n" 
                        src="../../assets/ball.png"
                        :class="{ 'active': n <= newFeedback.rating }"
                        @click="newFeedback.rating = n"
                        :title="`Оценка ${n} из 5`"
                    >
                </div>
            </div>
            
            <div class="form-input">
                <textarea 
                    v-model="newFeedback.text" 
                    placeholder="Ваш отзыв" 
                    required
                    rows="5"
                ></textarea>
            </div>
            
            <div class="form-in-row-left">
                <button 
                    class="submit-button-green" 
                    @click="submitFeedback"
                    :disabled="!isFeedbackValid">
                    Отправить
                </button>
            </div>
        </div>
    </div>

    <div v-if="isDeleteModalOpen" class="modal-overlay">
        <div class="modal-content">
            <p class="font-form-header">Подтверждение удаления</p>
            <p class="font-form-body">Вы уверены, что хотите удалить этот отзыв?</p>
            <div class="form-in-row-1">
                <button 
                    class="submit-button-green" 
                    @click="closeDeleteModal">
                    Нет
                </button>
                <button 
                    class="submit-button-orange" 
                    @click="confirmDelete">
                    Да
                </button>
            </div>
        </div>
    </div>
</template>

<script>
import axios from 'axios';
import config from "../../../config.js";

export default {
    name: 'AppFeedbacks',
    data() {
        return {
            config: config,
            feedbacks: [],
            userRackets: [],
            isDeleteModalOpen: false,
            selectedRacketId: null,
            showAddFeedbackModal: false,
            newFeedback: {
                racket_id: '',
                rating: 0,
                text: ''
            },
            activeDropdown: null,
            currentFilters: {
                rating: null,
                sortField: null,
                sortDirection: null
            }
        };
    },
    methods: {
        async fetchUserRackets() {
            try {
                const token = localStorage.getItem('token');
                if (!token) {
                    this.$router.push(this.config.VIEWS.auth.login);
                    return;
                }

                const ordersResponse = await axios.get(
                    `${config.BACKEND_URL}${config.API.orders}`,
                    { headers: { Authorization: `Bearer ${token}` } }
                );

                const uniqueRacketIds = new Set();

                if (ordersResponse.data?.orders) {
                    ordersResponse.data.orders.forEach(order => {
                        if (order.lines && order.lines.length > 0) {
                            order.lines.forEach(line => {
                                if (line.racket_id) {
                                    uniqueRacketIds.add(line.racket_id);
                                }
                            });
                        }
                    });
                }

                const racketIds = Array.from(uniqueRacketIds);
               
                if (racketIds.length > 0) {
                    const racketsPromises = racketIds.map(racketId => 
                        axios.get(`${config.BACKEND_URL}${config.API.rackets}/${racketId}`, {
                            headers: { Authorization: `Bearer ${token}` }
                        }).then(res => res.data)
                        .catch(() => null)
                    );

                    const racketsData = await Promise.all(racketsPromises);
                    this.userRackets = racketsData.filter(racket => racket !== null);
                } else {
                    this.userRackets = [];
                }

            } catch (error) {
                console.error('Ошибка при загрузке ракеток из заказов:', error);
                this.userRackets = [];
            }
        },
        
        async submitFeedback() {

            try {
                const token = localStorage.getItem('token');
                const response = await axios.post(
                    `${config.BACKEND_URL}${config.API.user.feedbacks}`,
                    {
                        racket_id: this.newFeedback.racket_id,
                        rating: this.newFeedback.rating,
                        feedback: this.newFeedback.text
                    },
                    { headers: { Authorization: `Bearer ${token}` } }
                );
                
                if (response.status === 200) {
                    this.closeAddFeedbackModal();
                    this.fetchFeedbacks();
                }
            } catch (error) {
                console.error('Ошибка при отправке отзыва:', error);
                alert('Не удалось отправить отзыв');
            }
        },

        async fetchFeedbacks() {
            try {
                const token = localStorage.getItem('token');
                if (!token) {
                    this.$router.push(this.config.API.auth.login);
                    return;
                }

                const response = await axios.get(
                    `${config.BACKEND_URL}${config.API.user.feedbacks}`,
                    { headers: { Authorization: `Bearer ${token}` } }
                );

                if (response.data?.feedbacks) {
                    const racketPromises = response.data.feedbacks.map(feedback => 
                        axios.get(`${config.BACKEND_URL}${config.API.rackets}/${feedback.racket_id}`, {
                            headers: { Authorization: `Bearer ${token}` }
                        }).then(res => res.data)
                        .catch(() => null)
                    );

                    const racketsData = await Promise.all(racketPromises);
                    this.feedbacks = response.data.feedbacks.map((feedback, index) => {
                        const racket = racketsData[index];
                        
                        return {
                            ...feedback,
                            racket_brand: racket.racket.brand
                        };
                    });

                    this.filteredFeedbacks = [...this.feedbacks];
                }
            } catch (error) {
                console.error('Ошибка при получении отзывов:', error);
                this.feedbacks = [];
                this.filteredFeedbacks = [];
            }
        },

        async confirmDelete() {
            if (this.selectedRacketId) {
                await this.deleteFeedback(this.selectedRacketId);
                this.closeDeleteModal();
            }
        },

        async deleteFeedback(racketId) {
            try {
                const token = localStorage.getItem('token');
                const response = await axios.delete(
                    `${config.BACKEND_URL}${config.API.user.feedbacks}/${racketId}`,
                    {
                        headers: {
                            Authorization: `Bearer ${token}`
                        }
                    }
                );
                
                if (response.status == 200) {
                    this.fetchFeedbacks();
                }
            } catch (error) {
                console.error('Ошибка при удалении отзыва:', error);
            }
        },

        toggleDropdown(dropdownName) {
            this.activeDropdown = this.activeDropdown === dropdownName ? null : dropdownName;
        },

        filterByRating(rating) {
            this.currentFilters.rating = rating;
            this.activeDropdown = null;
            this.applyFilters();
        },

        sortBy(field, direction) {
            this.currentFilters.sortField = field;
            this.currentFilters.sortDirection = direction;
            this.activeDropdown = null;
            this.applyFilters();
        },

        resetFilters() {
            this.currentFilters = {
                rating: null,
                sortField: null,
                sortDirection: null
            };
            this.filteredFeedbacks = [...this.feedbacks];
        },

        applyFilters() {
            let result = [...this.feedbacks];

            if (this.currentFilters.rating !== null) {
                result = result.filter(f => f.rating === this.currentFilters.rating);
            }

            if (this.currentFilters.sortField) {
                result.sort((a, b) => {
                    const field = this.currentFilters.sortField;
                    const valueA = new Date(a[field]);
                    const valueB = new Date(b[field]);

                    if (this.currentFilters.sortDirection === 'asc') {
                        return valueA - valueB;
                    } else {
                        return valueB - valueA;
                    }
                });
            }

            this.filteredFeedbacks = result;
        },

        goToRacket(racketId) {
            this.$router.push(`${this.config.VIEWS.rackets}/${racketId}`);
        },

        openDeleteModal(racketId) {
            this.selectedRacketId = racketId;
            this.isDeleteModalOpen = true;
        },

        closeDeleteModal() {
            this.isDeleteModalOpen = false;
            this.selectedRacketId = null;
        },

        getRacketDisplayName(racket) {
            return `${racket.racket.brand || 'Ракетка'} ${racket.racket.name || racket.racket.id}`;
        },

        hasFeedbackForRacket(racketId) {
            return this.feedbacks.some(f => f.racket_id === racketId);
        },

        onRacketSelect() {
            console.log('Выбрана ракетка:', this.newFeedback.racket_id);
        },

        openAddFeedbackModal() {
            this.fetchUserRackets();
            this.showAddFeedbackModal = true;
        },

        closeAddFeedbackModal() {
            this.showAddFeedbackModal = false;
            this.resetNewFeedbackForm();
        },

        resetNewFeedbackForm() {
            this.newFeedback = {
                racket_id: '',
                rating: 0,
                text: ''
            };
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
        }
    },
    computed: {
        isFeedbackValid() {
            return this.newFeedback.racket_id && 
                   this.newFeedback.rating > 0 && 
                   this.newFeedback.text.trim().length > 0;
        },
        isFilterActive() {
            return this.currentFilters.rating !== null || 
                   this.currentFilters.sortField !== null;
        }
    },
    mounted() {
        this.fetchFeedbacks();
    }
};
</script>

<style scoped>
@import "../../css/Containers.css";
@import "../../css/Fonts.css";
@import "../../css/Grid.css";
@import "../../css/Modal.css";


</style>