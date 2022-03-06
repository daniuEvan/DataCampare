const dbTaskRouter = [
    {
        path: '/task_manager',
        name: 'taskManager',
        component: () => import(/* webpackChunkName: "about" */ '../../views/task/Index'),
    },
]

export default dbTaskRouter