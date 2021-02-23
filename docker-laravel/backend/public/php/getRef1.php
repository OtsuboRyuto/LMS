<?php 
$fp = fopen("../data/ref1.txt", 'r');
$txt = fgets($fp);
echo $txt;
fclose($fp);

 ?>