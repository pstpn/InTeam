<template>
    <nav class="container-menu font-menu-bar">
        <router-link to="/rackets" class="container-icon">
            <img src="../assets/logo.png" alt="Rackets" class="icon-logo"/>
        </router-link>
        <nav class="container-menu-items">
            <router-link to="/orders" class="container-icon">
                <img src="../assets/order.png" alt="Order" class="icon"/>
                <span class="tooltip font-tooltip">Заказы</span>
            </router-link>

            <router-link to="/feedbacks" class="container-icon">
                <img src="../assets/feedback.png" alt="Feedback" class="icon"/>
                <span class="tooltip font-tooltip">Отзывы</span>
            </router-link>

            <router-link to="/cart" class="container-icon">
                <img src="../assets/cart.png" alt="Cart" class="icon"/>
                <span class="tooltip font-tooltip">Корзина</span>
            </router-link>

            <router-link to="/auth/login" class="container-icon">
                <img src="../assets/auth.png" alt="Login" class="icon"/>
                <span class="tooltip font-tooltip">Личный<br>кабинет</span>
            </router-link>
        </nav>
    </nav>  
</template>

<script>
import '../css/Containers.css'
import '../css/Icons.css'
import '../css/Fonts.css'

export default {
    name: 'AppMenuBar',
}

</script>