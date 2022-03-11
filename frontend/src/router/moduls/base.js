const baseRouter = [
    {
        path: '/home',
        name: 'home',
        component: () => import(/* webpackChunkName: "about" */ '../../views/Home'),
    },
    {
        path: '/about',
        name: 'about',
        component: () => import(/* webpackChunkName: "about" */ '../../views/About'),
    },

    {
        path: '/',
        name: 'index',
        redirect: '/scheduler_watch'
    },
]

export default baseRouter