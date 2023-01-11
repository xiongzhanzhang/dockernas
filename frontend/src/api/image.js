import http from '@/utils/request'

export const getImages = () => {
    return http.get("/api/image")
};

export const delImages = (imageInfo) => {
    return http.delete("/api/image", imageInfo)
};
