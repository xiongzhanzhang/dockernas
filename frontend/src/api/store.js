import http from '@/utils/request'

export const getApps = () => {
    return http.get("/api/app")
};