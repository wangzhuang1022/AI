window.onload=function () {
    $.ajax({
        type: 'get',
        url: 'http://localhost:1234/get/123456',
        success: (function (data) {
            console.info(data);
        }),
        error: function (xhr, textStatus) {
            alert(textStatus);
        }
    });
};
$('#btn1').click(function () {
    $.ajax({
        type: 'post',
        url: 'http://localhost:1234/save/',
        data: {
            userValue: $('#text1').val(),
            userId: '123456'
        },
        success: (function (data) {
            alert(data);
        }),
        error: function (xhr, textStatus) {
            alert(textStatus);
        }

    });
});
$('#btn2').click(function () {
    $.ajax({
        type: 'get',
        url: 'http://localhost:1234/get/123456',
        success: (function (data) {
            $('#text2').val(data)
        }),
        error: function (xhr, textStatus) {
            alert(textStatus);
        }

    });
});