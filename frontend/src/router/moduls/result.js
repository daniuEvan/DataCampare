const resultRouter = [
    {
        path: '/result/table',
        name: 'resultTableQuery',
        component: () => import(/* webpackChunkName: "about" */ '../../views/resultTableMsg/Index'),
    },
]

export default resultRouter