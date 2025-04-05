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
                        id="email"
                        v-model="email"
                        placeholder="Баланс"
                        required
                    />
                </div>

                <div class="form-input">
                    <input
                        type="text"
                        id="email"
                        v-model="email"
                        placeholder="Вес"
                        required
                    />
                </div>

                <div class="form-input">
                    <input
                        type="text"
                        id="email"
                        v-model="email"
                        placeholder="Размер головы"
                        required
                    />
                </div>

                <div class="form-input">
                    <input
                        type="text"
                        id="email"
                        v-model="email"
                        placeholder="Бренд"
                        required
                    />
                </div>

                <div class="form-input">
                    <input
                        type="text"
                        id="email"
                        v-model="email"
                        placeholder="Цена"
                        required
                    />
                </div>

                <div class="form-input">
                    <input
                        type="text"
                        id="email"
                        v-model="email"
                        placeholder="Количество"
                        required
                    />
                </div>

                <div class="form-in-row">
                    <p class="font-form-body">
                        Изображение 
                    </p>
                    <button 
                        class="submit-button-green">
                        Добавить
                    </button>
                </div>
            </div>
        </div>
        <div class="grid-order-column">
            <div class="form-in-row">
                <div class="filter-dropdown">
                    <button class="submit-button-filter" @click="toggleDropdown('weight')">
                        Вес
                        <span class="dropdown-arrow">▼</span>
                    </button>
                    <div class="dropdown-menu" v-show="activeDropdown === 'weight'">
                        <button @click="sortBy('weight', 'asc')">По возрастанию</button>
                        <button @click="sortBy('weight', 'desc')">По убыванию</button>
                    </div>
                </div>

                <div class="filter-dropdown">
                    <button class="submit-button-filter" @click="toggleDropdown('balance')">
                        Баланс
                        <span class="dropdown-arrow">▼</span>
                    </button>
                    <div class="dropdown-menu" v-show="activeDropdown === 'balance'">
                        <button @click="sortBy('balance', 'asc')">По возрастанию</button>
                        <button @click="sortBy('balance', 'desc')">По убыванию</button>
                    </div>
                </div>

                <div class="filter-dropdown">
                    <button class="submit-button-filter" @click="toggleDropdown('head_size')">
                        Размер
                        <span class="dropdown-arrow">▼</span>
                    </button>
                    <div class="dropdown-menu" v-show="activeDropdown === 'head_size'">
                        <button @click="sortBy('head_size', 'asc')">По возрастанию</button>
                        <button @click="sortBy('head_size', 'desc')">По убыванию</button>
                    </div>
                </div>

                <div class="filter-dropdown">
                    <button class="submit-button-filter" @click="toggleDropdown('brand')">
                        Бренд
                        <span class="dropdown-arrow">▼</span>
                    </button>
                    <div class="dropdown-menu" v-show="activeDropdown === 'brand'">
                        <button v-for="brand in uniqueBrands" :key="brand" 
                                @click="filterByBrand(brand)">
                            {{ brand }}
                        </button>
                    </div>
                </div>

                <div class="filter-dropdown">
                    <button class="submit-button-filter" @click="toggleDropdown('price')">
                        Цена
                        <span class="dropdown-arrow">▼</span>
                    </button>
                    <div class="dropdown-menu" v-show="activeDropdown === 'price'">
                        <button @click="sortBy('price', 'asc')">По возрастанию</button>
                        <button @click="sortBy('price', 'desc')">По убыванию</button>
                    </div>
                </div>

                <div class="filter-dropdown">
                    <button class="submit-button-filter" @click="toggleDropdown('availability')">
                        Доступность
                        <span class="dropdown-arrow">▼</span>
                    </button>
                    <div class="dropdown-menu" v-show="activeDropdown === 'availability'">
                        <button @click="filterByAvailability(true)">Доступные</button>
                        <button @click="filterByAvailability(false)">Недоступные</button>
                    </div>
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
                        <a class="container-icon-tool">
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
                        Голова
                    </p>
                    <p class="font-form-body-bold">
                        {{ racket.head_size }} мм
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
            showAddRacketModal: false,
            newRacket: {
                name: '',
                price: 0,
                brand: '',
                balance: 0,
                head_size: 0,
                weight: 0
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
            
            // Применяем фильтр по бренду
            if (this.brandFilter) {
                result = result.filter(racket => racket.brand === this.brandFilter);
            }
            
            // Применяем фильтр по доступности
            if (this.availabilityFilter !== null) {
                result = result.filter(racket => racket.available === this.availabilityFilter);
            }
            
            // Применяем сортировку
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
            console.log(brand)
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
@import "../../css/Icons.css";
@import "../../css/Menu.css";

</style>