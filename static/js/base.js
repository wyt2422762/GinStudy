var $ = layui.jquery
$(function(){
    // 设置jQuery Ajax全局的参数
    $.ajaxSetup({
        error: function (XMLHttpRequest, textStatus, errorThrown) {
            top.layer.msg(XMLHttpRequest.responseJSON.msg);
        },
        complete: function (XMLHttpRequest, textStatus){
            console.log("complete")
        }
    });
});