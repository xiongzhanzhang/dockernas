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


