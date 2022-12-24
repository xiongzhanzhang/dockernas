import http from '@/utils/request'

export const getDfsDirs = (path) => {
    return http.get("/api/dfsdir?path="+encodeURI(path))
}

export const getSystemDirs = (path) => {
    return http.get("/api/systemdir?path="+encodeURI(path))
}

export const setBasePath = (path) => {
    return http.post("/api/basepath",{"path":path})
}