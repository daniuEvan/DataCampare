const schedulerRouter = [
    {
        path: '/scheduler_manager',
        name: 'schedulerManager',
        component: () => import(/* webpackChunkName: "about" */ '../../views/scheduler/Index'),
    },
]

export default schedulerRouter