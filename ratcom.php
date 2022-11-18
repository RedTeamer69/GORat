<?php

if(isset($_GET['hostname']))
{
    // send a PowerShell command to client
   echo "Write-Host 'Installing updates'";
}

if ($_SERVER['REQUEST_METHOD'] === 'POST') {
  $cac1 = htmlspecialchars($_POST["cmd"]);
  $cac2 = htmlspecialchars($_POST["displ"]);
  $cac3 = htmlspecialchars($_POST["hostname"]);

  $stamp = date('Ymd_His');

  $hn = fopen("tmp/access.log", "a+");
  $r = fwrite($hn, "New client request: " . $stamp . "\n");
  $r = fwrite($hn, "Hostname: " . $cac3 . "\n");
  $r = fwrite($hn, "Command: " . $cac1 . "\n");
  $r = fwrite($hn, "Response: " . $cac2 . "\n");
  $r = fwrite($hn, "Client request end: " . $stamp. "\n\n");
  fclose($hn);
}
?>
