import http from '@/utils/request'
import storage from "@/utils/storage"

export const newInstance = (instanceParam) => {
    return http.post("/api/instance", instanceParam)
};

export const getAllInstance = () => {
    return http.get("/api/instance")
}

export const getInstance = (name) => {
    return http.get("/api/instance/" + name)
}

export const stopInstance = (name) => {
    return http.patch("/api/instance/" + name, { "op": "stop" })
}

export const startInstance = (name) => {
    return http.patch("/api/instance/" + name, { "op": "start" })
}

export const editInstance = (name, dataStr) => {
    return http.patch("/api/instance/" + name, { "op": "edit", "data": dataStr })
}

export const deleteInstance = (name) => {
    return http.delete("/api/instance/" + name)
}

export const getInstanceLog = (name) => {
    return http.get("/api/instance/" + name + "/log")
}

export const getInstanceEvent = (name) => {
    return http.get("/api/instance/" + name + "/event")
}

export const getInstanceStats = (start, end) => {
    return http.get("/api/instancestats?start=" + start + "&end=" + end)
}

export const getInstanceStatsByName = (name, start, end) => {
    return http.get("/api/instancestats/" + name + "?start=" + start + "&end=" + end)
}

export const getInstanceHttpPort = () => {
    return http.get("/api/instancehttpport")
}

export const getWebTerminalWebsocket = (containerId, columns) => {
    var url=window.location.host+"/terminal?containerid=" + containerId + "&columns=" + columns + "&token=" + storage.get("token")
    if(window.location.protocol=='http:'){
        url="ws://"+url
    }else{
        url="wss://"+url
    }
    var ws = new WebSocket(url);
    return ws
}
