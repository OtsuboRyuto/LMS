//ボタンのページ遷移共通

$("#toTop").click(function() {
    window.location.href = "http://angelos.info.kanagawa-u.ac.jp/";

})

$("#toLogin").click(function() {
    window.location.href = "http://192.168.0.80:25565";

})
$("#toDocument").click(function() {
    //window.location.href = "http://192.168.0.80:25565";
alert("未実装です。");
})

$("#toMypage").click(function() {
    window.location.href = "http://192.168.0.80:25565/mypage.html?hash =" + hash + "&id=" + id;
})

$("#toGroup").click(function() {
    $.ajax({
        url: "http://192.168.0.80:25565/cgi-bin/cgi-bin/attach_unique_id.py",
        dataType: "json",
        type: "post",
        data: JSON.stringify({ "id": id }),
        async: false,
    }).done(function(state) {
        if (state == 0) {
            alert("ipの値が不正です");
            window.location.href = "index.html";
        }
    }).fail(function() {
        alert("エラー1");
    });



    window.location.href = "http://192.168.0.80:25565/bbs/" + hash + "/" + id;

})


$("#toTest").click(function() {
    $.ajax({
        url: "./data/ref2.txt",
        type: "get",
        async: false,
    }).done(function(state) {
       // alert(state);
        $.ajax({
            url: "./data/ref1.txt",
            type: "get",
            async:false,
        }).done(function(data) {
           // alert(data);
            if (String(data) == String(false)) { //dataがfalseの場合の処理{
                alert("テストはまだ始まっていません");
            } else {
                alert("テストを開始します。");
                window.location.href = "http://192.168.0.80:25565/online_test_2.html?id=" + id + "&test_title=" + state;
            }
        }).fail(function() {
            alert('failed');
        });
    }).fail(function() {
        alert("エラー1");
    });



    // window.location.href = "http://192.168.0.80:25565/bbs/" + hash + "/" + id;

})



$("#toAdmin").click(function() {
    window.location.href = "http://192.168.0.80:25565/online_test_admin.html";

})
$("#toGobbs").click(function() {
    //alert("this");
    $.ajax({
        url: "http://192.168.0.80:25565/cgi-bin/cgi-bin/attach_unique_id.py",
        dataType: "json",
        type: "post",
        data: JSON.stringify({ "id": id }),
        async: false,
    }).done(function(state) {
        //alert("this");
        if (state == 0) {
            alert("ipの値が不正です");
            window.location.href = "index.html";
        }
    }).fail(function() {
        alert("エラー1");
    });
    //alert(hash + "  " + id);
    window.location.href = "http://192.168.0.80:25565/Go_bbs.html?hash=" + hash + "&id=" + id;
})
