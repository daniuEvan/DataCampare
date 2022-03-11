const schedulerRouter = [
    {
        path: '/scheduler_manager',
        name: 'schedulerManager',
        component: () => import(/* webpackChunkName: "about" */ '../../views/scheduler/Index'),
    },
    {
        path: '/scheduler_watch',
        name: 'schedulerWatch',
        component: () => import(/* webpackChunkName: "about" */ '../../views/watch/Index'),
    },

]

export default schedulerRouter