import Vue from 'vue';
import VueRouter from 'vue-router';
// import userRouter from './moduls/user';
import baseRouter from "@/router/moduls/base";

Vue.use(VueRouter);

const routes = [
    ...baseRouter,
    // ...userRouter,
];

const router = new VueRouter({
    mode: "history",
    base: process.env.BASE_URL,
    routes,
});


export default router;
