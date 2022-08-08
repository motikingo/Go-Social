package main

import (
	//"bytes"
	"fmt"
	"log"
	"net/http"

	//"strconv"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	//"github.com/mattn/go-sqlite3/upgrade"
)

const(
	writewait = 10*time.Second
	pongWait = 60 *time.Second
	pingPeriod = (pongWait * 9)/10
	maxMessageSize = 512

)

var(
	newline =[]byte("\n")
	space  =[]byte( " ")
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:1024,
	WriteBufferSize:1024,
}
type Client struct{
	hub *Hub
	conn *websocket.Conn 
	send chan []byte
}

func(c *Client) readPump()  {

	defer func(){
		c.hub.unregister <- c
		c.conn.Close()
	}()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {c.conn.SetReadDeadline(time.Now().Add(pongWait));return nil})

	for{
		_,message,er:=c.conn.ReadMessage()
		if er!=nil{
			if websocket.IsUnexpectedCloseError(er,websocket.CloseGoingAway,websocket.CloseAbnormalClosure){
				log.Printf("error %v",er)
			}
			break
		}

		message = []byte(strings.TrimSpace(strings.Replace(string(message),"\n"," ",-1 )))
		c.hub.broadcast <- message
	}

}

func(c *Client)writePump(){
	fmt.Println("write")
	ticker := time.NewTicker(pingPeriod)
	defer func(){
		ticker.Stop()
		c.conn.Close()
	}()

	for{
		select{
		case message,ok:= <-c.send:
			fmt.Println("sed")

			c.conn.SetWriteDeadline(time.Now().Add(writewait))
			if !ok{
				c.conn.WriteMessage(websocket.CloseMessage,[]byte{})
				return
			}
			w,err:= c.conn.NextWriter(websocket.TextMessage)

			if err!=nil{
				log.Printf("error is %v",err)
			}

			w.Write(message)
			n:= len(c.send)
			for i:=0;i<n;i++{

				w.Write(newline)
				w.Write(<-c.send)
			}

		case <- ticker.C:
			fmt.Println("kill")
			c.conn.SetWriteDeadline(time.Now().Add(writewait))
			if er:= c.conn.WriteMessage(websocket.PingMessage,nil);er!=nil{
				return
			} 
		}
	}
}

func Servews(hub *Hub,w http.ResponseWriter,r *http.Request){
	con,er:= upgrader.Upgrade(w,r,nil)
	
	if er!= nil{
		log.Fatal(er)
		return
	}

	client := &Client{
		hub:hub,
		conn:con,
		send:make(chan []byte,256),
	}

	client.hub.registy <- client

	go client.readPump()
	go client.writePump()
	fmt.Println("server")

}