const BACKEND_URL = "http://158.160.175.99:8080/api";

const API = {
    auth: {
        login: '/auth/login',
        register: '/auth/register',
    },
    user: {
        profile: '/profile',
        orders: '/orders',
        cart: '/cart',
        feedbacks: '/feedbacks'
    },
    rackets: '/rackets'
};

export default { BACKEND_URL, API };