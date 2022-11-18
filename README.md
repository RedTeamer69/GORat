# GORat
Simple RAT tool written in GO for RedTeam test    
This simple tool is designed to simulate an attack on a client PC   
It will help you to check how your user and defense technology react to infection  
Client execution will likely be blocked by local policy. Server side require web server with .php support. It is quick hack used in education phishing campain. Server scipt is buggy and vulnerable.
RedTeam and delete it.

**Usage:**  
*Client side*   
Download and install GO https://go.dev/dl/ <br> 
Set up your domain  
Test<br> 
go run client.go <br> 
Build<br> 
go build client.go <br> 
Distribute client.exe <br> 
 <br> 
*Server side*   
Put ratcom.php to your web server <br> 
Make tmp/access.log writable to php script  <br> 
