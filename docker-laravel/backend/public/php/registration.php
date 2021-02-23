<?php
/*
引数  id, name, passward, question, answer
受けとったデータのidがもしデータベースに存在しないならpasswardをハッシュ化してデータを書き込む
そうでない場合は0を返却し存在することを伝える
また、異常終了時は3を返却する。

*/
try {
$pdo = new PDO("mysql:dbname=laravel_local;host=db;charset=utf8mb4", "phper", "secret", 
		[
            PDO::ATTR_ERRMODE => PDO::ERRMODE_EXCEPTION,
            PDO::ATTR_DEFAULT_FETCH_MODE => PDO::FETCH_ASSOC,
        ]
	);
$pass = $_POST["passward"];
#$pass = "1";
$passward = hash("sha256",$pass);
$id = $_POST["id"];
$stmt = $pdo->prepare("select exists(select id from users where id = ? )");
$stmt->bindValue(1, (int)$id, PDO::PARAM_INT);
$stmt->execute();
$a = $stmt->fetch(PDO::PARAM_INT);

if($a[0] == 0){
	$stmt = $pdo-> prepare("insert into users (id, name, passward, question, answer) values (?, ?, ?, ?, ?)");
	$stmt->bindValue(1, (int)$id, PDO::PARAM_INT);
	$stmt->bindValue(2, $_POST["username"], PDO::PARAM_STR);
	$stmt->bindValue(3, $passward, PDO::PARAM_STR);
	$stmt->bindValue(4, $_POST["question"], PDO::PARAM_STR);
	$stmt->bindValue(5, $_POST["answer"], PDO::PARAM_STR);
	$stmt->execute();
	echo 1;
}
else
	echo 0;


	} catch (PDOException $e) {
        //echo $e->getMessage();
        echo -1;
        exit;
    }

    ?>