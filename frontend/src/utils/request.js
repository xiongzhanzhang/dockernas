import axios from 'axios'
import storage from "@/utils/storage"
import {ElMessage} from "element-plus"

const http = axios.create({
    timeout: 1000 * 60
})

// request interceptor
http.interceptors.request.use(config => {
    const token = storage.get("USER_TOKEN");
    if (token) {
        config.headers["USER_TOKEN"] = token;
    }
    return config;
}, error => {
    return Promise.reject(error)
})

http.interceptors.response.use(response => {
    return response
  }, error => {
    ElMessage.error(error.response.data.msg)
    // console.log(error)
    return Promise.reject(error)
  })

export default {
    get(url, data = {}) {
        return http.get(url, data)
    },
    post(url, data = {}) {
        return http.post(url, data)
    },
    put(url, data = {}) {
        return http.put(url, data)
    },
    delete(url, data = {}) {
        return http.delete(url, { data: data })
    },
    patch(url, data = {}) {
        return http.patch(url, data)
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
