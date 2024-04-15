
function doHttpRequest(url){
	let client=callMethod("http","newClient");
	return client.ExecuteTextRequest(url, "Get", "",""); 
}

/**
 * 调用插件方法，参见JavaScriptService#callPluginsMethod
 * @param {*} code 
 * @param {*} method 
 * @param {*} params 
 */
function callMethod( code, method, params) {
    return _func_predef_proxy.CallPluginsMethod(code, method,  params);
}

//console.Log(result);

function main(url){
	return doHttpRequest(url);
}