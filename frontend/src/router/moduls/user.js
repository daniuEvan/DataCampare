const userRouter = [
  {
    path: '/login',
    name: 'login',
    component: () => import(/* webpackChunkName: "about" */ '../../views/user/Login'),
  },
  {
    path: '/register',
    name: 'register',
    component: () => import(/* webpackChunkName: "about" */ '../../views/user/Register'),
  },
  {
    path: '/profile',
    name: 'profile',
    meta: {
      // auth: true, // 需要验证
    },
    component: () => import(/* webpackChunkName: "about" */ '../../views/user/Profile'),
  },
];

export default userRouter;
