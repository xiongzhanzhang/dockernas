import http from '@/utils/request'

export const getHostInfo = () => {
    return http.get("/api/host")
}