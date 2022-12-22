import http from '@/utils/request'

export const getDfsDirs = (path) => {
    return http.get("/api/filesystem?path="+encodeURI(path))
}
