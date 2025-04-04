const BACKEND_URL = "http://localhost:8081";

const API = {
    auth: {
        login: '/api/auth/login',
        register: '/api/auth/register',
    },
    user: {
        profile: '/api/profile',
        cart: '/api/cart',
        feedbacks: '/api/feedbacks'
    },
    admin: {
        profile: '/api/profile',
        users: '/api/users'
    },
    rackets: '/api/rackets',
    feedbacks: '/api/feedbacks/rackets',
    orders: '/api/orders'
};

const VIEWS = {
    auth: {
        login: '/auth/login',
        register: '/auth/register',
    },
    user: {
        profile: '/user/profile',
        orders: '/user/orders',
        cart: '/user/cart',
        feedbacks: '/user/feedbacks'
    },
    admin: {
        profile: '/admin/profile',
        orders: '/admin/orders',
        cart: '/admin/rackets',
        rackets: '/admin/rackets',
        users: '/admin/users'
    },
    rackets: '/rackets'
};

export default { BACKEND_URL, API, VIEWS };