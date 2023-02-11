export const getInstanceWebUrl = (instanceName, port) => {
    if(port==0){
        return null;
    }
    // return window.location.protocol + "//"+window.location.hostname+":"+port
    return "http://"+window.location.hostname+":"+port
}

export const splitRouterPathByIndex = (router, index) =>{
    return router.split("/").slice(0,index).join("/")
}