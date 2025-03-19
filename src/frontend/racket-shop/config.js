const BACKEND_URL = "http://192.168.1.199:8081/api";

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