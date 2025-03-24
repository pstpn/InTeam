const BACKEND_URL = process.env.VUE_APP_BACKEND_URL;

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