import axios from 'axios'
import storage from "@/utils/storage"
import { ElMessage } from "element-plus"
import router from '@/router'

const http = axios.create({
    timeout: 1000 * 60
})

// request interceptor
http.interceptors.request.use(config => {
    const token = storage.get("token");
    if (token) {
        config.headers["token"] = token;
    }
    return config;
}, error => {
    return Promise.reject(error);
})

http.interceptors.response.use(response => {
    return response;
}, error => {
    console.log(error);
    if (error.response != null) {
        if (error.response.status == 555) {
            storage.rm("token");
            router.replace({ path: "/login" });
        } else if (error.response.status == 556) {
            router.replace({ path: "/basepath" });
        } else {
            var msg = error.response.data.msg;
            if (msg == null || msg == "") {
                msg = error.response.statusText;
            }
            ElMessage.error(msg);
        }
    }
    else {
        msg = error.message;
        ElMessage.error(msg);
    }

    return Promise.reject(error);
})

export default {
    get(url, data = {}) {
        return http.get(url, data);
    },
    post(url, data = {}) {
        return http.post(url, data);
    },
    put(url, data = {}) {
        return http.put(url, data);
    },
    delete(url, data = {}) {
        return http.delete(url, { data: data });
    },
    patch(url, data = {}) {
        return http.patch(url, data);
    },
    upload(url, params) {
        return http.post(url, params, {
            headers: {
                'Content-Type': 'multipart/form-data'
            },
            timeout: 1000 * 60 // 上传文件超时1分钟
        })
    }
}
