export const getInstanceWebUrl = (instanceParam, portParamItem) => {
    if(instanceParam.networkMode === "host"){
        return "http://"+window.location.hostname+":"+portParamItem.key;
    }
    if(portParamItem.value==="" || instanceParam.networkMode === "nobund"){
        return null;
    }
    // return window.location.protocol + "//"+window.location.hostname+":"+port
    return "http://"+window.location.hostname+":"+portParamItem.value
}

export const getInstancePortText = (instanceParam, portParamItem) => {
    if(instanceParam.networkMode === "host"){
        return portParamItem.key;
    }

    if(portParamItem.value==="" || instanceParam.networkMode === "nobund"){
        return portParamItem.key + " -> ";
    }

    return portParamItem.key + " -> " + portParamItem.value;
}

export const splitRouterPathByIndex = (router, index) =>{
    return router.split("/").slice(0,index).join("/")
}