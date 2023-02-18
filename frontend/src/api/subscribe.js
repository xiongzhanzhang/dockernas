import http from '@/utils/request'

export const getSubscribes = () => {
    return http.get("/api/subscribe")
};

export const createSubscribe = (name, url) => {
    return http.post("/api/subscribe", {
        name: name,
        url: url
    })
};

export  const deleteSubscribe = (name) => {
    return http.delete("/api/subscribe/"+name)
}

export  const flushSubscribe = () => {
    return http.patch("/api/subscribe")
}