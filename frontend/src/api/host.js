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

export const startHttpGateway = (domain) => {
    return http.post("/api/httpgateway",{})
}

export const restartHttpGateway = (name) => {
    return http.patch("/api/httpgateway", { "op": "restart" })
}

export const stopHttpGateway = (name) => {
    return http.patch("/api/httpgateway", { "op": "stop" })
}

export const enableHttps = () => {
    return http.post("/api/httpgateway/https",{})
}

export const disableHttps = () => {
    return http.delete("/api/httpgateway/https", {})
}
