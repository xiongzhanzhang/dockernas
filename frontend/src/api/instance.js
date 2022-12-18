import http from '@/utils/request'

export const newInstance = (
    name, summary, appName, imageUrl, version, portParams, envParams, localVolume, dfsVolume, iconUrl
) => {
    return http.post("/api/instance", {
        "name": name,
        "summary":summary,
        "appName": appName,
        "imageUrl": imageUrl,
        "version": version,
        "portParams": portParams,
        "envParams": envParams,
        "localVolume": localVolume,
        "dfsVolume": dfsVolume,
        "iconUrl": iconUrl
    })
};

export const getAllInstance = () => {
    return http.get("/api/instance")
}

export const getInstance = (name) => {
    return http.get("/api/instance/"+name)
}

export const stopInstance = (name) => {
    return http.patch("/api/instance/"+name,{"op":"stop"})
}

export const startInstance = (name) => {
    return http.patch("/api/instance/"+name,{"op":"start"})
}

export const editInstance = (name, dataStr) => {
    return http.patch("/api/instance/"+name,{"op":"edit","data":dataStr})
}

export const deleteInstance = (name) => {
    return http.delete("/api/instance/"+name)
}