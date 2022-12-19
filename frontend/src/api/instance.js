import http from '@/utils/request'

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