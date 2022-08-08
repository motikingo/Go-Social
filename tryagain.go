// package main

// import (
// 	"log"
// 	"net/http"
// 	"strings"
// 	"time"

// 	"github.com/gorilla/websocket"
// )

// const(
// 	wrWait =10* time.Second
// 	powait = 60* time.Second
// 	piperiod = (powait*9)/10
// 	maxMessSize = 512
// )

// var(
// 	newl = []byte{"\n"}
// 	sp = []byte{" "}
// )

// var upgder = websocket.Upgrader{
// 	ReadBufferSize:1024,
// 	WriteBufferSize:1024,
// }

// type Clie struct{
// 	h *Hub
// 	c websocket.Conn
// 	s chan []byte
// }

// func(cl *Clie)read(){
// 	defer func ()  {
// 		cl.h.unregister <- cl
// 		cl.c.Close()
// 	}()
// 	cl.c.SetReadDeadline(time.Now().Add(powait))
// 	cl.c.SetReadLimit(maxMessSize)
// 	cl.c.SetPongHandler(func (string)error {cl.c.SetReadDeadline(time.Now().Add(powait));return nil})

// 	for {
// 		_,message,er:= cl.c.ReadMessage()
// 		if er!=nil{
// 			if websocket.IsUnexpectedCloseError(er,websocket.CloseGoingAway,websocket.CloseAbnormalClosure){
// 				log.Printf("erroe is %v",er)
// 			}
// 			break
// 		}

// 		message = []byte(strings.TrimSpace(strings.Replace(string(message),"\n"," ")))
// 		cl.h.broadcast <-message
		
// 	}
// }

// func(cl *Clie)wr(){
// 	t :=time.NewTicker(piperiod)

// 	defer func ()  {
// 		t.Stop()
// 		cl.c.Close()
// 	}()

// 	for {

// 		select{
// 		case m,ok:= <- cl.send:
// 			cl.c.SetWriteDeadline(time.Now().Add(wrWait))
// 			if !ok{
// 				cl.c.WriteMessage(websocket.TextMessage,[]byte{})
// 				break
// 			}

// 			w,er:= cl.c.NextWriter(websocket.TextMessage)

// 			if er!=nil{
// 				log.Printf(er.Error())
// 				return
// 			}

// 			w.Write(m)

// 			for i:=0;i<len(cl.s);i++{

// 				w.Write(newl)
// 				w.Write(<-cl.s)
// 			}
// 		}
		
// 	}

// }

// func servsw(h *Hub,w http.ResponseWriter,r *http.Request){
// 	c,_:= upgrader.Upgrade(w,r,nil)

// 	cl:=&Clie{c:c,h: h,s: make(chan []byte,256)}
// 	cl.h.registy <- cl
// 	go cl.read()
// 	go cl.wr()

// }