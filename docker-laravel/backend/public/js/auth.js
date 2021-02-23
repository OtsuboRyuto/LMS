        var tag = true;
/*
    function QUERY_START() {//テストが始まっているか確認する関数
        return $.ajax({
            url: "./data/ref1.txt",
            dataType: "json",
            type: "get",
        }).done(function(data) {
        alert(data);
            if (data == 0) { //dataがfalseの場合の処理{
                alert("テストはまだ始まっていません");
                return 1;
            } else{
                alert("テストを開始します。");
                return 0;
            }
        }).fail(function() {
            alert('failed');
        });
    }
*/
    function GET_PASS() { //学籍番号を返す関数
        var query = document.cookie;
        localStorage.setItem(query, query); //ロード時の学籍番号が消えるのを防ぐ
        var i;
        var ID = ""; //文字列連結用の箱
        // = から&までの文字を抜き出すプログラム(STUDENT_NUMBERが入っている)
        var tag = 0;
        for (i = 0; query.length != i; i++) {
            if (tag == 2) //タグが1の時にコピー
                ID = ID + query[i];
            if (query[i] == '=')
                tag++;
            //document.write(query[i]);
        }
        return ID;

    }
    function GET_HASH() { //ハッシュ値を返す関数
        var query = document.cookie;
        if( query.match(/&/)){

        localStorage.setItem(query, query); //ロード時の学籍番号が消えるのを防ぐ
        var i;
        var ID = ""; //文字列連結用の箱
        // = から&までの文字を抜き出すプログラム(STUDENT_NUMBERが入っている)
        var tag = 0;
        for (i = 0; query[i] != '&'; i++) {
            if (tag == 1) //タグが1の時にコピー
                ID = ID + query[i];
            if (query[i] == '=')
                tag = 1;
            //document.write(query[i]);
        }
        return ID;
}else{
    alert("値が不正です")
    window.location.href = "index.html";
    }
}
const url = new URL(location.href);
const params = new URLSearchParams(url.search);
const hash = GET_HASH();
const id = params.get("id");
const pass = GET_PASS();

    function AUTH() {
        $.ajax({
            url: "http://192.168.0.80:25565/golang/c_auth",
            dataType: "json",
            type: "POST",
            async:false,
            data: { "data": hash, "id": id },
            async: false,
        }).done(function(target) {
            if (target == 1) {
                //現在のIPアドレスが一致するかどうかをチェック
                $.getJSON("https://api.ipify.org/?format=json", function(connection) { //クライアントのipを取得できるapiに接続し、データを取得。
                    $.ajax({
                        url: "http://192.168.0.80:25565/golang/checkIp",
                        dataType: "json",
                        type: "post",
                        data:{ "id": id, "ip": connection.ip },
                        async: false,
                    }).done(function(state) {
                        if (state == 0) {
                            alert("ipの値が不正です");
                            window.location.href = "index.html";
                        }
                        if (state == 1){
                        $.ajax({
                        url: "http://192.168.0.80:25565/golang/checkPass",
                        dataType: "json",
                        type: "post",
                        data:{ "id": id, "pass":pass},
                        async: false,
                    }).done(function(state) {
                        if (state == 0) {
                            alert("passの値が不正です");
                            window.location.href = "index.html";
                        }else{
   						var page1 = document.getElementById("target");
   						page1.style.display = "block";
                        var loader = document.getElementById("load");
                        if (loader != null){
                        loader.style.display = "none";
                        }
                        //var loader = document.getElementById("load");
                        //loader.style.display = "none";
   						}
                    }).fail(function() {
                        alert("error -check");
                    });
                          }
   
                    }).fail(function() {
                        alert("error -check");
                    });
                });
            } else { //エラー、urlを書いて、疑似的に遷移したパターン。どちらにせよ不正な検出
                window.location.href = "index.html";
            }
        }).fail(function() {
            alert("エラー");
        });
    }
    document.addEventListener('DOMContentLoaded', (event) => {
       AUTH();
    });

//AUTH();
