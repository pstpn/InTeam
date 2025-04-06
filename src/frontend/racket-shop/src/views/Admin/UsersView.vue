<template>
    <div class="container-page">
        <h1 class="font-container-header">Пользователи</h1>

        <!-- Поле поиска -->
        <div class="grid-order-column">
            <div class="form-in-row">
                <input
                    type="text"
                    v-model="searchQuery"
                    placeholder="Поиск по имени, email или роли"
                    class="input"
                    @input="filterUsers"
                >
            </div>
        </div>

        <div class="grid-order-column">
            <div class="grid-card-name" v-for="(user) in filteredUsers" :key="user.id">
                <div class="form-in-row">
                    <div class="form-in-row-left">
                        <p class="font-form-body-bold">
                            Пользователь {{ user.user_id }}
                        </p>
                        <button :class="{
                            'submit-button-orange': user.role === 'Admin',
                            'submit-button-green': user.role !== 'Admin'
                        }">
                            {{ getRoleName(user.role) }}
                        </button>
                        <a class="container-icon-tool" @click.stop="openRoleModal(user)">
                            <img src="../../assets/tool.png" alt="Tool" class="icon-tool"/>
                        </a>
                    </div>
                </div>
                <div class="form-in-row">
                    <p class="font-form-body">
                        Имя
                    </p>
                    <p class="font-form-body-bold">
                        {{ user.name }}
                    </p>
                </div>
                <div class="form-in-row">
                    <p class="font-form-body">
                        Фамилия
                    </p>
                    <p class="font-form-body-bold">
                        {{ user.surname }}
                    </p>
                </div>
                <div class="form-in-row">
                    <p class="font-form-body">
                        Email
                    </p>
                    <p class="font-form-body-bold">
                        {{ user.email }} 
                    </p>
                </div>
            </div>
        </div>
    </div>

    <div v-if="showRoleModal" class="modal-overlay" @click.self="closeRoleModal">
        <div class="modal-content">
            <p class="font-form-header">Изменение роли пользователя</p>
            <div class="form-in-row">
                <p class="font-form-body-bold">{{ selectedUser.name }} {{ selectedUser.surname }}</p>
            </div>
            <div class="form-input">
                <select v-model="selectedRole">
                    <option value="User">Покупатель</option>
                    <option value="Admin">Админ</option>
                </select>
            </div>
            <div class="form-in-row">
                <button 
                    class="submit-button-orange"
                    @click="closeRoleModal">
                    Отмена
                </button>
                <button
                    class="submit-button-green"
                    @click="updateUserRole">
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
            users: [],
            filteredUsers: [],
            searchQuery: '',
            config: config,
            loading: true,
            showRoleModal: false,
            selectedRole: '',
            selectedUser: null
        };
    },
    methods: {
        openRoleModal(user) {
            this.selectedUser = user;
            this.selectedRole = user.role;
            this.showRoleModal = true;
        },
        
        closeRoleModal() {
            this.showRoleModal = false;
            this.selectedUser = null;
            this.selectedRole = '';
        },
        
        async updateUserRole() {
            try {
                const token = localStorage.getItem('token');
                if (!token) {
                    this.$router.push(this.config.VIEWS.auth.login);
                    return;
                }
                
                const response = await axios.patch(
                    `${config.BACKEND_URL}${config.API.admin.users}/${this.selectedUser.user_id}`,
                    {
                        role: this.selectedRole
                    },
                    {
                        headers: {
                            'Authorization': `Bearer ${token}`
                        }
                    }
                );
                
                if (response.status === 200) {
                    this.closeRoleModal();
                    this.fetchUsers(); // Обновляем список пользователей
                    alert('Роль пользователя успешно обновлена!');
                }
            } catch (error) {
                console.error('Ошибка при обновлении роли:', error);
                alert('Не удалось обновить роль: ' + 
                    (error.response?.data?.message || error.message));
            }
        },
        async fetchUsers() {
            try {
                this.loading = true;
                const token = localStorage.getItem('token');

                if (!token) {
                    this.$router.push(this.config.VIEWS.auth.login);
                    return;
                }

                const response = await axios.get(`${config.BACKEND_URL}${config.API.admin.users}`, {
                    headers: {
                        Authorization: `Bearer ${token}`
                    }
                });

                if (response.data) {
                    this.users = response.data.users || [];
                    this.filteredUsers = [...this.users];
                }
            } catch (error) {
                console.error('Ошибка при получении пользователей:', error);
                this.users = [];
                this.filteredUsers = [];
            } finally {
                this.loading = false;
            }
        },
        
        getRoleName(role) {
            const roleNames = {
                'Admin': 'Админ',
                'User': 'Покупатель',
            };
            return roleNames[role] || role;
        },
        
        filterUsers() {
            if (!this.searchQuery) {
                this.filteredUsers = [...this.users];
                return;
            }
            
            const query = this.searchQuery.toLowerCase();
            this.filteredUsers = this.users.filter(user => {
                return (
                    user.name.toLowerCase().includes(query) ||
                    (user.surname.toLowerCase().includes(query)) ||
                    (user.email.toLowerCase().includes(query)) ||
                    (this.getRoleName(user.role).toLowerCase().includes(query)
                ))
            });
        }
    },
    mounted() {
        this.fetchUsers();
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