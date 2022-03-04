const dbLinkRouter = [
    {
    path: '/db_link_manager',
    name: 'dbLinkManager',
    component: () => import(/* webpackChunkName: "about" */ '../../views/dbLink/Index'),
   },
]

export default dbLinkRouter