function httpGet(path)
{
    var xmlHttp = new XMLHttpRequest();
    xmlHttp.open( "POST", "http://127.0.0.1:4000/post", false ); // false for synchronous request
    xmlHttp.send("page="+path);
    return xmlHttp.responseText;
}
var page;
page = encodeURIComponent(document.getElementById("side-content").innerHTML);
setInterval(function(){httpGet(page);}, 5000);
setTimeout("location.reload(true);",10000);




