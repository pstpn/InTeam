<template>
    <div class="container-page">
        <h1 class="font-container-header">Мои отзывы</h1>

        <div v-if="feedbacks.length === 0" class="font-form-body">
            <p>Вы не оставили еще ни один отзыв</p>
        </div>

        <div v-else class="grid-order-column">
            <div 
                class="grid-card-feedback" 
                v-for="(feedback, index) in feedbacks" 
                :key="index"
            >
                <div class="grid-feedback">
                    <p class="font-form-header">
                        Ракетка {{ feedback.racket_id }}
                    </p>
                    
                    <div class="form-in-row-racket">
                        <div class="grid-photo-ball">
                            <p class="font-form-body">
                                {{ formatDate(feedback.date) }} 
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
            feedbacks: [],
            isDeleteModalOpen: false,
            selectedRacketId: null 
        };
    },
    methods: {
        async fetchFeedbacks() {
            try {
                const token = localStorage.getItem('token');

                if (!token) {
                    this.$router.push(this.config.API.auth.login);
                    return;
                }

                const response = await axios.get(
                    `${config.BACKEND_URL}${config.API.user.feedbacks}`,
                    {
                        headers: {
                            Authorization: `Bearer ${token}`
                        }
                    }
                );

                if (response.data) {
                    console.log(response.data)
                    this.feedbacks = response.data.feedbacks;
                }
            } catch (error) {
                console.error('Ошибка при получении отзывов:', error);
            }
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

/* .modal-overlay {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.5);
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 1000;
}

.modal-content {
    background-color: white;
    padding: 20px;
    border-radius: 30px;
    width: 300px;
    text-align: center;
}

.modal-content h2 {
    margin-bottom: 16px;
}

.modal-buttons {
    display: flex;
    justify-content: space-around;
    margin-top: 20px;
}

.modal-buttons button {
    padding: 8px 16px;
    border: none;
    border-radius: 4px;
    cursor: pointer;
}

.modal-buttons .submit-button-orange {
    background-color: #ffa500;
    color: white;
}

.modal-buttons .submit-button-green {
    background-color: #008000;
    color: white;
} */
</style>