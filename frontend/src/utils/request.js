import axios from 'axios';
// import storageService from '../service/storageService';

// 引入封装axios(推荐)
const service = axios.create({
    baseURL: process.env.VUE_APP_BASE_URL,
    timeout: 1000 * 10,
    // timeout: process.env.REQUEST_TIME_OUT,
});

// 添加请求拦截器,动态更新headers
service.interceptors.request.use(function (config) {
    // 在发送请求之前做些什么
    // Object.assign(config.headers, {Authorization: `Bearer ${storageService.get(storageService.USER_TOKEN)}`});
    return config;
}, function (error) {
    // 对请求错误做些什么
    console.log(error);
    // return Promise.reject(error);
});

export default service;
