import http from '@/utils/request'

export const getHostInfo = () => {
    return http.get("/api/host")
}

export const getStorageInfo = () => {
    return http.get("/api/storage")
}


export const getNetworkInfo = () => {
    return http.get("/api/network")
}

export const getHttpProxyConfig = () => {
    return http.get("/api/httpproxyconfig")
}

export const delHttpProxyConfig = (hostName) => {
    return http.delete("/api/httpproxyconfig/"+hostName)
}

export const postHttpProxyConfig = (hostName, instanceName, port) => {
    return http.post("/api/httpproxyconfig",{
        "hostName":hostName,
        "instanceName":instanceName,
        "port":port
    })
}

export const postDomain = (domain) => {
    return http.post("/api/domain",{
        "domain":domain
    })
}