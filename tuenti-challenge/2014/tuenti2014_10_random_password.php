#! /usr/bin/env php
<?php
/*

Tuenti Challenge 4, Apr 2014, Challenge 10 - Random password

Challenge not submitted to the contest.


Source code for the challenge page can be got from index.phps or index.php~
We have to find the parent PID to learn the seed.

This program tries PID until the password is got (current one is 1336).

*/

date_default_timezone_set("GMT");

$pid = 1;
$hour = 12;
$minute = 0;
$password = "0";

$found = false;
do {
    srand(mktime($hour, $minute, 0) * $pid);
    $password = rand();
    echo "TRYING PID: $pid, PASSWORD: $password, hour: $hour, minute: $minute\n";

    $page = file_get_contents("http://random.contest.tuenti.net/?input=f175ff28be&password=$password");
    if ($page == "wrong!") {
        $date_parser = date_parse(substr($http_response_header[5],6 ));
        $header_hour = $date_parser['hour'];
        $header_minute = $date_parser['minute'];
        if ($hour != $header_hour or $minute != $header_minute) {
            echo "TIME DOES NOT MATCH, TRYING AGAIN\n";
            $hour = $header_hour;
            $minute = $header_minute;
        } else {
            $pid++;
        }
    } else {
        echo "FOUND, PID: $pid, PASSWORD: $page\n";
        $found = true;
    }
} while (!$found);
