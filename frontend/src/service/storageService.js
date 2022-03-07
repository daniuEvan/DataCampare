//本地缓存服务

const PREFIX = "data_compare";

// user模块
const USER_PREFIX = `${PREFIX}user_`;
const USER_TOKEN = `${USER_PREFIX}token`;
const USER_INFO = `${USER_PREFIX}info`;

// 数据比对
const DB_LINK_LIST = `db_link_list`
const TASK_INFO_LIST = `task_info_list`
const SCHEDULER_INFO_LIST = `scheduler_info_list`


// 存储
const set = (key, data) => {
    localStorage.setItem(key, data);
}

// 读取

const get = (key) => localStorage.getItem(key)

export default {
    get,
    set,
    USER_TOKEN,
    USER_INFO,
    DB_LINK_LIST,
    TASK_INFO_LIST,
    SCHEDULER_INFO_LIST,
}
