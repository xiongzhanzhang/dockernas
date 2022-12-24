import http from '@/utils/request'

export const login = (user, passwd) => {
    return http.post("/api/login",{
        "user":user,
        "passwd":passwd
    })
}