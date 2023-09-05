<?php

$string = "Gabriel"; //String
$integer = 23; // Integer
$float = 1.93; // Float
$bool = true;

$bool = !is_bool($bool) ? $bool : ( $bool ? "true" : "false");

echo $string, PHP_EOL;
echo $integer, PHP_EOL;
echo $float, PHP_EOL;
echo $bool, PHP_EOL;
