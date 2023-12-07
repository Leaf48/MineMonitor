package utils

import (
	"net"
	"time"
)

var IsConnected bool

func CheckInternetConnection(){
  for {
    conn, err := net.DialTimeout("tcp", "google.com:80", 2500 * time.Millisecond)

    if err != nil{
      IsConnected = false
    }else {
      IsConnected = true
      conn.Close()
    }

    time.Sleep(1 * time.Second)
  }
}
