/* Utilities (Helpers) */

function removeAllChildrenFrom(view){
    view.innerHTML = ""
}

function callBackendMicroservice(route, body, method){
    return callURL(backendAddress + microserviceEndpoint + route, body, method)
}

function pingBackendServer(){
    return callURL(backendAddress + "/ping")
}

function callURL(url, body, method){
    var request = new XMLHttpRequest();
    
    const isAsync = false
    request.open(method, url, isAsync);
    const contentType = "application/json"
    request.setRequestHeader("Content-Type", contentType)
    
    if (!body) {
        request.send(null);
    }else{
        request.send(JSON.stringify(body))
    }

    return JSON.parse(request.responseText)
}
