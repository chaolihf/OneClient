
function doHttpRequest(){
	let client=callMethod("http","newClient");
	//return client.ExecuteTextRequest("http://134.64.116.90:8101/sso/index.html?res=workflow", "Get", "",""); 
	return client;
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

var result=doHttpRequest();
console.Log(result);
result;