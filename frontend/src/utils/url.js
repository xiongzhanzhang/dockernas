export const getInstanceWebUrl = (instanceName, port) => {
    if(port==0){
        return null;
    }
    return window.location.protocol + "//"+window.location.hostname+":"+port
}