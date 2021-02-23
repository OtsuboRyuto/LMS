<?php
	/*
	apiにてLAN内別サーバーのapiよりデータを取得する
	引数
	*/
$url = "192.168.0.80:25566/". $_POST["url"];
// $url = "localhost:25566/data/ref1.txt";
//cURLセッションを初期化する

$ch = curl_init();
 
//送信データを指定
//$data = array('samurai' => 1);
 
//URLとオプションを指定する
curl_setopt($ch, CURLOPT_URL, $url);
curl_setopt($ch, CURLOPT_POST, false);
//curl_setopt($ch, CURLOPT_POSTFIELDS, $data);
 
//URLの情報を取得する
$res =  curl_exec($ch);
//echo 1;
//結果を表示する

//セッションを終了する
curl_close($ch);
    ?>