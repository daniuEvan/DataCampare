const baseRouter = [{
    path: '/home',
    name: 'home',
    component: () => import(/* webpackChunkName: "about" */ '../../views/Home'),
},
    {
        path: '/about',
        name: 'about',
        component: () => import(/* webpackChunkName: "about" */ '../../views/About'),
    },
]

export default baseRouter