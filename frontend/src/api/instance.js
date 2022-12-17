import http from '@/utils/request'

export const newInstance = (
    name, appName, imageUrl, version, portParams, envParams, localVolume, dfsVolume, iconUrl
) => {
    return http.post("/api/instance", {
        "name": name,
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

export const deleteInstance = (name) => {
    return http.delete("/api/instance/"+name)
}