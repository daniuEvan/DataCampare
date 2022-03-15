import Vue from 'vue';
import VueRouter from 'vue-router';
// import userRouter from './moduls/user';
import baseRouter from "@/router/moduls/base";
import dbLinkRouter from "@/router/moduls/dbLink";
import taskRouter from "@/router/moduls/task";
import schedulerRouter from "@/router/moduls/scheduler";
import resultRouter from "@/router/moduls/result";

Vue.use(VueRouter);

const routes = [
    ...baseRouter,
    ...dbLinkRouter,
    ...taskRouter,
    ...schedulerRouter,
    ...resultRouter,
    // ...userRouter,
];

const router = new VueRouter({
    mode: "history",
    base: process.env.BASE_URL,
    routes,
});


export default router;
