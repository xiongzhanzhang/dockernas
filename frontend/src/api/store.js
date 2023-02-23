import http from '@/utils/request'

export const getApps = () => {
    return http.get("/api/app")
};


export const getAppsByName = (name) => {
    return http.get("/api/app/"+name)
};