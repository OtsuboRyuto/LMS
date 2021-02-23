<!DOCTYPE html>
<html>
    <head>
        <meta charset='utf-8'>
 <script src="https://code.jquery.com/jquery-3.5.1.min.js"></script><!-- Jqueryをinclude /-->
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.0-beta1/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-giJF6kkoqNQ00vy+HMDP7azOuL0xtbfIcaT9wjKHr8RbDVddVHyTfAAsrekwKmP1" crossorigin="anonymous">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <script src ="{{asset('./js/ButtonEvent.js')}}"></script>
    <meta name="viewport" content="width=device-width, initial-scale=1">
        <title>bbs</title>
    </head>
    <script>
    const id = '{{$id}}';
	const group = '{{$group}}';
	//forループはデータの取得差分式もしくは、全権取得の頻繁にリロードする形
    </script>
    <body style="background-image : url(../../img/graWhite2.gif); " class="container-fluid" id="target">
<div style="position:fixed;" class="col-md-3 d-none d-sm-block">
            <!-- ここに機能群のボタンを入れる-->
            <div class="container col-md-12 row">
                <div class="col-md-8" style="background-color:rgba(130,130,130,0.2);margin-top:15%;padding-bottom:15%;border-radius:20px;">
                    <div class="col-md-12" style="text-align: center" class = "row"><img src="../../img/logo.gif"> </div>
                    <button type="button" class="btn btn-outline-secondary col-md-12" id = "toTop" style="margin:1%">トップページ</button><br>
                    <button type="button" class="btn btn-outline-secondary col-md-12" id = "toLogin" style="margin:1%">ログイン画面</button><br>
                    <button type="button" class="btn btn-outline-secondary col-md-12" id = "toMypage" style="margin:1%">マイページ</button><br>
                    <button type="button" class="btn btn-outline-secondary col-md-12" style="margin:1%" id="toGroup">BBS Laravel</button><br> <!-- グループチャット -->
                    <button type="button" class="btn btn-outline-secondary col-md-12" id = "toGobbs" style="margin:1%">BBS Golang</button><br>
                    <button type="button" class="btn btn-outline-secondary col-md-12" id = "toTest" style="margin:1%">オンラインテスト</button><br><!-- テストのページ   -->
                    <button type="button" class="btn btn-outline-secondary col-md-12" id = "toDocument" style="margin:1%">授業資料</button><br>
                    <button type="button" class="btn btn-outline-secondary col-md-12" id = "toAdmin" style="margin:1%">アドミン</button><br> <!-- アドミン用ページへ  -->
                    <!--ページのリロードなしで授業資料を見れる  -->
                </div>
            </div>
        </div>
    	<div id = "whole_contents" class="col-md-8 offset-md-2 text-left row align-items-start border-start border-end border-secondary">
    	<div style = "height:100vh;">
    	        <h1>あなたのIDは{{$id}}です</h1>
        <p>Group {{ $group}}</p>
<div class = "form-group">
            <form method = "post" name = "form">
             <!-- {{csrf_field()}} cookieへの書き込み防止-->
        <textarea name = "content" id = "content" cols = "100", rows = "4" maxlength="128" class = "col-9" style = "resize:none;;float:left;height:6em;"></textarea>
        <button type="button"class = "col-3" id = "submit" style = "float:right;height:6em;">投稿</button>
       		</form>
            <div style = "clear:left;padding-top:2em">
        	<div id = "contents">
        	      
        	  </div>   
            </div>   	
		</div>
		</div>	
		</div>
    </body>
    <script>
    	getData();
    	var i = 0;
    	$('#toMypage').click(function(){
    		window.location.href = "../../mypage.html?hash=" + {{$hash}} + "&id=" + id;
        })
   	    	document.getElementById("submit").onclick = function(){

    		var content = document.form.content.value;
    		$.ajax({
                    url: "/bbs/id/" + id + "/group/" + group + "/content/" + content,
                    type: "get",
                    async:true,
                }).done(function(state) {
                       		getData();
                }).fail(function(){
                    alert("送信失敗");
                });	
           document.form.content.value = "";
    }
    function getData(){
    $.ajax({
                    url: "../group/" + group + "/flag/false",
                    dataType: "json",
                    type: "get",
                    async:true,
                }).done(function(state) {

                	for(;i < state["contents_id"].length;){
           				if(state["contents_id"][i] == id){//自身のidと取得したIDが一致するケースに枠の色を変更する
           					$("#contents").prepend("<div id = \"content\" class = \"border border-success rounded bg-success bg-gradient text-light \">" +  state["contents_id"][i] + "さんの投稿<br>" + state["contents"][i] + "</div>");	
                			$("#contents").prepend("<div style =\"margin:1%;\"></div>");
           				}
           				else{
                			$("#contents").prepend("<div id = \"content\" class = \"border border-secondary bg-gradient rounded bg-secondary text-light\">" +  state["contents_id"][i] + "さんの投稿<br>" + state["contents"][i] + "</div>");	
                			$("#contents").prepend("<div style =\"margin:1%;\"></div>");
                		}
                		i++;
                	}
                		
                }).fail(function(){
                    alert("送信失敗");
                });		
    }	
     setInterval("getData()", 6000);
    </script>

</html>