import http from '@/utils/request'

export const getHostInfo = () => {
    return http.get("/api/host")
}

export const getStorageInfo = () => {
    return http.get("/api/storage")
}