const BACKEND_URL = "http://localhost:8080/api";

const API = {
    auth: {
        login: '/auth/login',
        register: '/auth/register',
    },
    user: {
        profile: '/profile',
        orders: '/profile/orders',
        cart: '/profile/cart',
        feedbacks: '/profile/feedbacks'
    }
};

export default { BACKEND_URL, API };