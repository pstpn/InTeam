<template>
    <div class="container-page">
        <h1 class="font-container-header">Мои отзывы</h1>

        
        <div v-if="feedbacks.length === 0" class="font-form-body">
            <p>Вы не оставили еще ни один отзыв</p> 

            <div class="grid-order-column">
                <div class="form-in-row-add-feedback">
                    <button class="submit-button-green" @click="openAddFeedbackModal">
                        Добавить отзыв
                    </button>
                </div>
            </div>
        </div>

        <div v-else class="grid-order-column">
            <div class="form-in-row-add-feedback">
                <button class="submit-button-green" @click="openAddFeedbackModal">
                    Добавить отзыв
                </button>
            </div>
            <div 
                class="grid-card-feedback" 
                v-for="(feedback, index) in feedbacks" 
                :key="index">
                <div class="grid-feedback">
                    <p class="font-form-header" @click="goToRacket(feedback.racket_id)">
                        {{ feedback.racket_id + ' ' +  feedback.racket_brand}}
                    </p>
                    
                    <div class="form-in-row-racket">
                        <div class="grid-photo-ball">
                            <p class="font-form-body">
                                {{ feedback.date }} 
                            </p>
                            <img 
                                v-for="(ball, ballIndex) in feedback.rating" 
                                :key="ballIndex" 
                                src="../../assets/ball.png"
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
            }
        };
    },
    computed: {
        isFeedbackValid() {
            console.log(this.newFeedback.racket_id, this.newFeedback.rating, this.newFeedback.text.trim())
            return this.newFeedback.racket_id && 
                   this.newFeedback.rating > 0 && 
                   this.newFeedback.text.trim().length > 0;
        }
    },
    methods: {
        getRacketDisplayName(racket) {
            console.log('here', racket.racket)
            return `${racket.racket.brand || 'Ракетка'} ${racket.racket.name || racket.racket.id}`;
        },
        
        hasFeedbackForRacket(racketId) {
            return this.feedbacks.some(f => f.racket_id === racketId);
        },
        
        onRacketSelect() {
            // Можно добавить дополнительную логику при выборе ракетки
            console.log('Выбрана ракетка:', this.newFeedback.racket_id);
        },
        openAddFeedbackModal() {
            this.fetchUserRackets(); // Загружаем ракетки пользователя
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
        
        async fetchUserRackets() {
            try {
                const token = localStorage.getItem('token');
                if (!token) {
                    this.$router.push(this.config.VIEWS.auth.login);
                    return;
                }

                // 1. Получаем все заказы пользователя
                const ordersResponse = await axios.get(
                    `${config.BACKEND_URL}${config.API.orders}`,
                    { headers: { Authorization: `Bearer ${token}` } }
                );

                // 2. Собираем все уникальные racket_id из заказов
                const uniqueRacketIds = new Set(); // Используем Set для автоматического удаления дубликатов

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

                // 3. Преобразуем Set обратно в массив
                const racketIds = Array.from(uniqueRacketIds);
               
                // 4. Получаем полные данные по каждой ракетке (опционально)
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

                // console.log(this.userRackets)
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

            // 1. Получаем отзывы
            const response = await axios.get(
                `${config.BACKEND_URL}${config.API.user.feedbacks}`,
                {
                    headers: {
                        Authorization: `Bearer ${token}`
                    }
                }
            );

            if (response.data?.feedbacks) {
                // 2. Создаем массив промисов для получения данных о ракетках
                const racketPromises = response.data.feedbacks.map(feedback => 
                    axios.get(`${config.BACKEND_URL}${config.API.rackets}/${feedback.racket_id}`, {
                        headers: { Authorization: `Bearer ${token}` }
                    }).then(res => res.data)
                    .catch(() => null) // Игнорируем ошибки для отдельных запросов
                );

                // 3. Получаем данные всех ракеток
                const racketsData = await Promise.all(racketPromises);

                // 4. Объединяем отзывы с данными ракеток
                this.feedbacks = response.data.feedbacks.map((feedback, index) => {
                    const racket = racketsData[index];
                    return {
                        ...feedback,
                        racket_brand: racket.racket.brand
                    };
                });

                console.log(this.feedbacks)
            }
        } catch (error) {
            console.error('Ошибка при получении отзывов:', error);
            this.feedbacks = [];
        }
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

        formatDate(dateString) {
            const date = new Date(dateString);
            return date.toLocaleDateString('ru-RU'); 
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