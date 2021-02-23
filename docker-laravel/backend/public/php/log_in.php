<?php
/*
サーバーに送られてきたidが存在するか確認する
y      ↓
サーバーから紐づけられたパスワードを取得
       ↓
クライアントから送られてきたパス(ハッシュに変換)を照合して、正しければ1を返却そうでないなら0を返却
異常終了は-1を返却
*/
try {
$pdo = new PDO("mysql:dbname=laravel_local;host=db;charset=utf8mb4", "phper", "secret", 
		[
            PDO::ATTR_ERRMODE => PDO::ERRMODE_EXCEPTION,
            PDO::ATTR_DEFAULT_FETCH_MODE => PDO::FETCH_ASSOC,
        ]
);
#$_POST = ["id" => 1, "passward" => "3"];
$stmt = $pdo -> prepare("select exists(select id from users where id = ?)"); 
$stmt->bindvalue(1, (int)$_POST["id"], PDO::PARAM_INT);
$stmt->execute();
$a = $stmt->fetch(PDO::PARAM_INT);
if ($a[0] == 1){
	$stmt = $pdo -> prepare("select passward from users where id = ?");
	$stmt->bindvalue(1, (int)$_POST["id"], PDO::PARAM_INT);
	$stmt -> execute();
	$pass = $stmt->fetch(PDO::PARAM_STR);
    if($pass["passward"] == hash("sha256", $_POST["passward"]))
    	echo 1;
    else
    	echo 0;
} 
}
catch (PDOException $e) {
        echo -1;
        exit;
    }

?>