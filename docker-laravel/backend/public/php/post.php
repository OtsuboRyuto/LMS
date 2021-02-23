<?php
	/*
	apiにてLAN内別サーバーのapiよりデータを取得する
	引数 url + 1 + 2 + n...... という形式
	*/
	

$i = 0;
$data = [];
//取得したデータがfalseになるまで取得
/*
for($i = 0;isset($_POST[$i])!= false;$i++){
	$data += $_POST[$i];
	//echo $_POST[$i]; 
}

echo 2;
*/
for($i = 0; isset($_POST[$i]) != false;$i++){
$data =  [$i => $_POST[$i]];
$data = $data + [$i => $_POST[$i]];
}
//$data = json_encode($data);

$url = "192.168.0.80:25566/". $_POST["url"];


$ch = curl_init($url);
curl_setopt($ch, CURLOPT_POST, true);
//curl_setopt($ch, CURLOPT_URL, $url);
curl_setopt($ch, CURLOPT_POSTFIELDS, http_build_query($data));
curl_setopt($ch, CURLOPT_RETURNTRANSFER, TRUE); 
//$res =  curl_exec($ch);
$res = curl_exec($ch); 

    if(curl_errno($ch)){        //curlでエラー発生
        $CURLERR .= 'curl_errno：' . curl_errno($ch) . "\n";
        $CURLERR .= 'curl_error：' . curl_error($ch) . "\n";
        $CURLERR .= '▼curl_getinfo' . "\n";
        foreach(curl_getinfo($ch) as $key => $val){
            $CURLERR .= '■' . $key . '：' . $val . "\n";
        }
        echo nl2br($CURLERR);
    }
curl_close($ch);
echo $res;
    ?>