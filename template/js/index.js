$('#sendBtn').click(function () {
    var question = $('#question').val();
    //发送后清空输入内容
    $('#question').val("");
    appendDialog("Koala.jpg", "Me", question);
    $.ajax({
        type: 'post',
        url: '/chart/',
        data: {
            question: question
        },
        success: (function (data) {
            obj = JSON.parse(data);
            appendDialog("Jellyfish.jpg", "Robot", obj.text + obj.url);
        }),
        error: function (xhr, textStatus) {
            alert(textStatus);
        }
    });
});
function appendDialog(imageUrl, title, content){
    $('#dialog').append('<a href=\"javascript:void(0);\" class=\"weui-media-box weui-media-box_appmsg\">\n' +
        '        <div class=\"weui-media-box__hd\">\n' +
        '            <img class=\"weui-media-box__thumb\" src=\"./../img/' +imageUrl + '" alt=\"\">\n' +
        '        </div>\n' +
        '        <div class=\"weui-media-box__bd\">\n' +
        '            <h4 class=\"weui-media-box__title\">'+title+'</h4>\n' +
        '            <p class=\"weui-media-box__desc\">'+content+'</p>\n' +
        '        </div>\n' +
        '    </a>');
    var div = document.getElementById('dialog');
    div.scrollTop = div.scrollHeight;
}
//监听回车事件，回车即发送
$('#question').bind('keyup', function(event) {
    if (event.keyCode == "13") {
        //回车执行查询
        $('#sendBtn').click();
    }
});