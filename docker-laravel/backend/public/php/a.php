<?php
$command="http://192.168.0.80:25565/php/a.py ";
exec($command,$output);
print_r($output);
?>