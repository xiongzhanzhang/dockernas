import http from '@/utils/request'

export const getApps = () => {
    return http.get("/api/app")
};


export const getAppsByName = (name) => {
    if(name.indexOf("/")>0){
        return http.get("/api/extra/app/"+name)
    }
    return http.get("/api/app/"+name)
};