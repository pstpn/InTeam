<template>
    <div class="container-page">
        <h1 class="font-container-header">Ракетки</h1>
        <div class="grid-order-column">
            <div class="form-in-row">
                <button 
                    class="submit-button-green"
                    @click="showAddRacketModal = true">
                    Добавить ракетку
                </button>
            </div>
        </div>
        <div v-if="showAddRacketModal" class="modal-overlay" @click.self="closeAddRacketModal">
            <div class="modal-content">
                <p class="font-form-header">Добавление ракетки</p>
                <div class="form-input">
                    <input
                        type="text"
                        v-model="newRacket.balance"
                        placeholder="Баланс"
                        required
                    />
                </div>

                <div class="form-input">
                    <input
                        type="text"
                        v-model="newRacket.weight"
                        placeholder="Вес (г)"
                        required
                    />
                </div>

                <div class="form-input">
                    <input
                        type="text"
                        v-model="newRacket.head_size"
                        placeholder="Размер головы (кв.см)"
                        required
                    />
                </div>

                <div class="form-input">
                    <input
                        type="text"
                        v-model="newRacket.brand"
                        placeholder="Бренд"
                        required
                    />
                </div>

                <div class="form-input">
                    <input
                        type="text"
                        v-model="newRacket.price"
                        placeholder="Цена"
                        required
                    />
                </div>

                <div class="form-input">
                    <input
                        type="text"
                        v-model="newRacket.quantity"
                        placeholder="Количество"
                        required
                    />
                </div>

                <div class="form-in-row">
                    <input
                        type="file"
                        ref="fileInput"
                        accept="image/*"
                        style="display: none"
                        @change="handleImageUpload"
                    >
                    <button
                        class="submit-button-grey"
                        :class="{ 'image-selected': selectedImage }"
                        @click="triggerFileInput"
                        >
                        {{ selectedImage ? 'Изменить изображение' : 'Добавить изображение' }}
                    </button>
                    <button
                        v-if="selectedImage"
                        class="submit-button-orange"
                        @click="resetImage"
                        >
                        Сбросить
                    </button>
                    <button 
                        class="submit-button-green"
                        @click="addRacket">
                            Добавить
                    </button>
                </div>
            </div>
        </div>
        <div class="grid-order-column">
            <div class="form-in-row-add-feedback">
                <div class="form-in-row-left">
                   
                    <select class="submit-button-filter" @change="sortBy('weight', $event.target.value)">
                        <option value="">Вес</option>
                        <option value="asc">По возрастанию</option>
                        <option value="desc">По убыванию</option>
                    </select>

                    <select class="submit-button-filter" @change="sortBy('balance', $event.target.value)">
                        <option value="">Баланс</option>
                        <option value="asc">По возрастанию</option>
                        <option value="desc">По убыванию</option>
                    </select>

                    <select class="submit-button-filter" @change="sortBy('head_size', $event.target.value)">
                        <option value="">Размер</option>
                        <option value="asc">По возрастанию</option>
                        <option value="desc">По убыванию</option>
                    </select>

                    <select class="submit-button-filter" @change="filterByBrand($event.target.value)">
                        <option value="">Бренд</option>
                        <option v-for="brand in uniqueBrands" :key="brand" :value="brand">
                            {{ brand }}
                        </option>
                    </select>

                    <select class="submit-button-filter" @change="sortBy('price', $event.target.value)">
                        <option value="">Цена</option>
                        <option value="asc">По возрастанию</option>
                        <option value="desc">По убыванию</option>
                    </select>

                    <select class="submit-button-filter" @change="filterByAvailability($event.target.value)">
                        <option value="">Доступность</option>
                        <option value="true">Доступные</option>
                        <option value="false">Недоступные</option>
                    </select>
                </div>
                <button class="submit-button-orange" @click="resetFilters">
                    Сбросить
                </button>
            </div>
        </div>

        <div class="grid-order-column">
            <div class="grid-card-name" v-for="racket in filteredRackets" :key="racket.id">
               
                <div class="form-in-row">
                    <div class="form-in-row-left">
                        <p class="font-form-body-bold">
                            {{ racket.brand }} {{ racket.id }}
                        </p>
                        <button class="submit-button-green" :class="{ 'submit-button-orange': !racket.available }">
                            {{ racket.available ? 'Доступна' : 'Недоступна' }}
                        </button>
                        <a class="container-icon-tool" @click.stop="openUpdateModal(racket)">
                            <img src="../../assets/tool.png" alt="Tool" class="icon-tool"/>
                        </a>
                    </div>
                    <p class="font-form-body-bold">
                        Цена: {{ racket.price }} ₽
                    </p>
                </div>
                <div class="form-in-row">
                    <p class="font-form-body">
                        Количество
                    </p>
                    <p class="font-form-body-bold">
                        {{ racket.quantity }} шт
                    </p>
                </div>
                <div class="form-in-row">
                    <p class="font-form-body">
                        Вес
                    </p>
                    <p class="font-form-body-bold">
                        {{ racket.weight }} г
                    </p>
                </div>
                <div class="form-in-row">
                    <p class="font-form-body">
                        Баланс
                    </p>
                    <p class="font-form-body-bold">
                        {{ racket.balance }} мм
                    </p>
                </div>
                <div class="form-in-row">
                    <p class="font-form-body">
                        Размер головы
                    </p>
                    <p class="font-form-body-bold">
                        {{ racket.head_size }} мм
                    </p>
                </div>
            </div>
        </div>
    </div>

    <div v-if="showUpdateModal" class="modal-overlay" @click.self="closeUpdateModal">
        <div class="modal-content">
            <p class="font-form-header">Обновление количества</p>
            <div class="form-input">
                <input
                    type="number"
                    v-model="updatedQuantity"
                    placeholder="Новое количество"
                    min="0"
                    required
                />
            </div>
            <div class="form-in-row">
                <button 
                    class="submit-button-green"
                    @click="updateRacketQuantity">
                    Обновить
                </button>
                <button 
                    class="submit-button-orange"
                    @click="closeUpdateModal">
                    Отмена
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
            showUpdateModal: false,
            updatedQuantity: 0,
            selectedRacketId: null,
            selectedImage: null,
            selectedImagePreview: null,
            showAddRacketModal: false,
            newRacket: {
                price: '',
                brand: '',
                balance: '',
                head_size: '',
                weight: '',
                quantity: ''
            },
            rackets: [],
            config: config,
            loading: true,
            activeDropdown: null,
            currentSort: {
                field: null,
                direction: null
            },
            brandFilter: null,
            availabilityFilter: null
        };
    },
    computed: {
        uniqueBrands() {
            const brands = new Set();
            this.rackets.forEach(racket => brands.add(racket.brand));
            return Array.from(brands);
        },
        filteredRackets() {
            let result = [...this.rackets];
            
            if (this.brandFilter) {
                result = result.filter(racket => racket.brand === this.brandFilter);
            }
            
            if (this.availabilityFilter !== null) {
                result = result.filter(racket => racket.available === this.availabilityFilter);
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
        openUpdateModal(racket) {
            this.selectedRacketId = racket.id;
            this.updatedQuantity = racket.quantity;
            this.showUpdateModal = true;
        },
        
        closeUpdateModal() {
            this.showUpdateModal = false;
            this.selectedRacketId = null;
            this.updatedQuantity = 0;
        },
        
        async updateRacketQuantity() {
            try {
                const token = localStorage.getItem('token');
                if (!token) {
                    this.$router.push(this.config.VIEWS.auth.login);
                    return;
                }

                const response = await axios.patch(
                    `${config.BACKEND_URL}${config.API.rackets}/${this.selectedRacketId}`,
                    {
                        quantity: this.updatedQuantity
                    },
                    {
                        headers: {
                            'Authorization': `Bearer ${token}`
                        }
                    }
                );
                
                if (response.status === 200) {
                    this.closeUpdateModal();
                    this.fetchRackets();
                    alert('Количество успешно обновлено!');
                }
            } catch (error) {
                console.error('Ошибка при обновлении количества:', error);
                alert('Не удалось обновить количество: ' + 
                    (error.response?.data?.message || error.message));
            }
        },
        resetImage() {
            this.selectedImage = null;
            this.selectedImagePreview = null;
            this.$refs.fileInput.value = '';
            this.$emit('image-reset');
        },
        triggerFileInput() {
            this.$refs.fileInput.click()
        },

        handleImageUpload(event) {
            const file = event.target.files[0]
            if (file) {
                this.selectedImage = file
                
                const reader = new FileReader()
                reader.onload = (e) => {
                this.selectedImagePreview = e.target.result
                }
                reader.readAsDataURL(file)
            }
        },
        removeImage() {
            this.selectedImage = null
            this.selectedImagePreview = null
            this.$refs.fileInput.value = ''
        },
        openAddRacketModal() {
            this.showAddRacketModal = true;
        },
        closeAddRacketModal() {
            this.showAddRacketModal = false;
            this.resetNewRacketForm();
        },
        resetNewRacketForm() {
            this.newRacket = {
                name: '',
                price: 0,
                brand: '',
                balance: 0,
                head_size: 0,
                weight: 0
            };
        },
        async addRacket() {
            try {
                const token = localStorage.getItem('token');
                if (!token) {
                    this.$router.push(this.config.VIEWS.auth.login);
                    return;
                }

                const racketData = {
                    price: this.newRacket.price,
                    brand: this.newRacket.brand,
                    balance: this.newRacket.balance,
                    head_size: this.newRacket.head_size,
                    weight: this.newRacket.weight,
                    quantity: this.newRacket.quantity,
                    image: this.selectedImage
                };

                const response = await axios.post(
                    `${config.BACKEND_URL}${config.API.rackets}`,
                        racketData,
                    {
                        headers: {
                            'Authorization': `Bearer ${token}`,
                            'Content-Type': 'multipart/form-data'
                        }
                    }
                );
                
                if (response.status === 200) {
                    this.closeAddRacketModal();
                    this.fetchRackets();
                    alert('Ракетка успешно добавлена!');
                }
            } catch (error) {
                console.error('Ошибка при добавлении ракетки:', error);
                alert('Не удалось добавить ракетку: ' + 
                    (error.response?.data?.message || error.message));
            }
        },
        async fetchRackets() {
            try {
                this.loading = true;
                const token = localStorage.getItem('token');

                if (!token) {
                    this.$router.push(this.config.VIEWS.auth.login);
                    return;
                }

                const response = await axios.get(`${config.BACKEND_URL}${config.API.rackets}`, {
                    headers: {
                        Authorization: `Bearer ${token}`
                    }
                });

                if (response.data) {
                    this.rackets = response.data.rackets || [];
                }
            } catch (error) {
                console.error('Ошибка при получении ракеток:', error);
                this.rackets = [];
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
        filterByBrand(brand) {
            this.brandFilter = brand;
            this.activeDropdown = null;
        },
        filterByAvailability(available) {
            this.availabilityFilter = available;
            this.activeDropdown = null;
        },
        resetFilters() {
            this.currentSort = { field: null, direction: null };
            this.brandFilter = null;
            this.availabilityFilter = null;
        },
        closeDropdownsOnClickOutside(e) {
            if (!e.target.closest('.filter-dropdown')) {
                this.activeDropdown = null;
            }
        }
    },
    mounted() {
        this.fetchRackets();
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