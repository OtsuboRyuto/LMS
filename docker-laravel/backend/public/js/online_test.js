var params = (new URL(document.location)).searchParams;
var id = params.get("id");
var flag = QUERY_START();
if (flag == false){
    alert("テストは開始していません");
    window.location.href = "http://192.168.0.80:25565/";
}
/*
function QUERY_INTEGRATION(target) {//ファイルに追加されたものが送ったものと同じか確認するための関数
    const change_dir = dir;
    //var URL = change_dir.substr(1) + "/result/" + STUDENT_NUMBER() + ".csv";
    var URL = "../../infra/docker/python/data/" + params.get("test_title") + "/result/" + params.get("id") + ".csv";//userに情報を見せすぎているため本当はあまりよろしくない
    var result = false;
    alert(URL);
    $.ajax({
            // リクエスト送信先URLの設定
            url: URL,
            type: "get",
            dataType: "text",
            async: false,
        }).done(function(data) {
            alert(data);
            var dummy1 = data.replace(/\s+/g, "");
            target = String(target);
            var dummy2 = target.replace(/\s+/g, "");
                                                                            //alert(dummy1 + "    " + dummy2);
            if (dummy1 == dummy2 + ",") {                                   //空白文字が邪魔になるケースがあるため空白文字を除外
                                                                            // if(dummy1 === URL){
                alert("整合性が取れました。");
                result = true;                                               //ファイルの中のデータと送信したデータが同一のものの場合trueを返却
            } else {
                alert("整合性が取れませんでした。もう一度提出します。")
                result = false;
            }

        }).fail(function() {
            alert("通信失敗");
            result = false;
        });
    return result;
}
*/
function QUERY_INTEGRATION(ans) {//引数のデータはサーバーに送ったデータと同じもの
    var result = false;
    $.ajax({
        url: "http://192.168.0.80:25565/golang/integration" + location.search,                                                  //起動するPython Script
        detaType: "json",
        type: "POST",
        data: {"data":ans},                                                          //HTMLから受け取った値
        async: false,                                                             //同期処理を行う(データの整合性チェックのためUIをロック)
    }).done(function(data) {
        if(data == 1){
        	alert("送信成功しました。");
            result = true;
        }
    }).fail(function() {
        alert('送信エラー');

    });
return result;


}

function GET_SERVER_DATA() {
    var URL = "./cgi/result/" + id + ".csv";
    var result = $.ajax({
        type: 'GET',
        url: URL,
        async: false
    }).done(function(data) {
        return result;
    });
}


function CHECK_FORM() {
    var i;
    var len_form = number_of_ques;
    var document_list = "";
    var a_document = "";
    for (i = 0;; i++) {
        if (len_form == i)
            break;
        a_document = document.forms[0].elements[i].value;
        if (a_document == "")
            a_document = "なし";
        if (i % 2 == 0)
            document_list = document_list + (i + 1) + "番の回答      " + a_document + "   ";
        else
            document_list = document_list + (i + 1) + "番の回答      " + a_document + "\n";
    }
    var result = window.confirm(document_list + "\nこの内容で送信します。よろしいですか？");
    if (result) {
                                                                        //「true」の処理
        MAKE_ANSWERS();
    } else {
                                                                        //「false」の処理

    }

}

function QUERY_START() {
    var i = false;
    $.ajax({
        url: "./data/ref1.txt",
        async: false,
        type: "get",
    }).done(function(data) {
        if (String(data) == "true")
        i = true;

    });
    return i;
}
/*
function STUDENT_NUMBER() {                                             //学籍番号を返す関数
    var query = location.search;
    localStorage.setItem(query, query)                                  //ロード時の学籍番号が消えるのを防ぐ
    var i;
    var student_number = "";                                            //文字列連結用の箱
                                                                        // = から&までの文字を抜き出すプログラム(STUDENT_NUMBERが入っている)
    var tag = 0;
    for (i = 0; query[i] != '&'; i++) {
        if (tag == 1)                                                   //タグが1の時にコピー
            student_number = student_number + query[i];
        if (query[i] == '=')
            tag = 1;
                                                                        //document.write(query[i]);
    }
    return student_number;
}

function SECOND_QUERY() {                                             //学籍番号を返す関数
    var query = location.search;
    localStorage.setItem(query, query)                                  //ロード時の学籍番号が消えるのを防ぐ
    var i;
    var data = "";                                            //文字列連結用の箱
                                                                        // = から&までの文字を抜き出すプログラム(STUDENT_NUMBERが入っている)
    var tag = 0;
    for (i = 0; query[i] < query.length; i++) {                                                 //タグが1の時にコピー
            data = data + query[i];                                                    //document.write(query[i]);
            alert(query);
        }
    alert(data);
    return data;
}
*/

var flag = false;

function QUERY_TERMINATE() {                                            //テストが終わってるか確認するための関数
                                                                        // Ajax通信処理
    $.ajax({
                                                                        // リクエスト送信先URLの設定
        url: "./data/ref1.txt",
                                                                        // 非同期通信フラグ
        async: true,
        type: "get",
    }).done(function(data) {
        if (String(data) == "false") {
            //alert("test at 156");
            stopTimer();
            NEXT();
        }
    }).fail(function() {
        alert("サーバーとの通信に失敗");


    });
}
var testTimer = setInterval("QUERY_TERMINATE()", 3000)
function stopTimer() {
    clearInterval(testTimer);
}


function NEXT() {
    alert("制限時刻になりました。強制的に提出して、テストを終了します。");
    MAKE_ANSWERS();
}





function SEND_ANSWER(ans) {                                                   //データ送信用の関数
    var send_data = {"data": ans};
    $.ajax({
        url: "http://192.168.0.80:25565/golang/test" + location.search,                                                  //起動するPython Script
        detaType: "json",
        type: "POST",
        data: send_data,                                                          //HTMLから受け取った値
        async: false,                                                             //同期処理を行う(データの整合性チェックのためUIをロック)
    }).done(function(data) {
    }).fail(function() {
        alert('送信エラー');

    });

}

function MAKE_COMMA_SEPARATED_VALUES(target) {
    var i;
    var separated = "";
    const array_length = target.length;
    for (i = 0; i < array_length; i++)
        separated = separated + "," + target[i];
    return separated;

}

function MAKE_ANSWERS() {                                                              //配列の中にデータを格納0の要素に学籍番号)
    var i = 0;
    var ans = [];
    for (i = 1; i < number_of_ques + 1; i++) {
        data = document.getElementById("apply" + i).value;
        var dummy1 = data.replace(/\s+/g, "");
        if (dummy1 == "")
            ans.push(0);
        else
            ans.push(dummy1);
    }
    const a = MAKE_COMMA_SEPARATED_VALUES(ans);
    SEND_INTEGRATION(a);
}

function IS_FULLSCREEN() {
    if ((document.FullscreenElement !== undefined && document.FullscreenElement !== null) || // Firefox
        (document.webkitFullscreenElement !== undefined && document.webkitFullscreenElement !== null) || // Chrome・Safari など
        (document.msFullscreenElement !== undefined && document.msFullscreenElement !== null)) { // IE
        return true; // フルスクリーン中
    } else {
        return false; // フルスクリーンでないか、API非対応端末
    }
}

function SEND_INTEGRATION(ans) {
    var i = false;
    var count = 0;
    for (;count!=3;) {
        if (count == 3) {
            window.alert("整合性チェックが3回通りませんでした。\n原因不明のため管理者に連絡お願いいたします。\n 手動で採点します。ユーザー画面のスクリーンショットを行います。");
            /*
            html2canvas(document.getElementById("target"), {
                onrendered: function(canvas) {
                    //imgタグのsrcの中に、html2canvasがレンダリングした画像を指定する。
                    var imgData = canvas.toDataURL();
                    document.getElementById("result").src = imgData;
                }
            });*/
        }
        //alert();
        count += 1;
        SEND_ANSWER(ans);
        //i = GET_SERVER_DATA().responseText;
        i = QUERY_INTEGRATION(ans);
        if (i == true)
            break;
}
    alert("テストを終了します。お疲れ様でした。");
   window.location.href = "http://192.168.0.80:25565/mypage.html?id=" + id;
}