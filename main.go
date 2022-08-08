package main

import (
	//"fmt"
	//"fmt"
	"log"
	"net/http"
	//"github.com/gorilla/websocket"
)

// var upgrader = websocket.Upgrader{
// 	ReadBufferSize : 1024,
// 	WriteBufferSize : 1024,
// }

func main(){

	hub := NewHub()
	go hub.run()
	http.HandleFunc("/",homepage)
	http.HandleFunc("/chat",chat)
	http.HandleFunc("/ws",func(w http.ResponseWriter, r *http.Request) {

		Servews(hub,w,r)
		//fmt.Println(<- hub.broadcast)
	})

	//fmt.Print("vcdscds")
	log.Fatal(http.ListenAndServe(":8080",nil))
	
}

func homepage(w http.ResponseWriter, r *http.Request){

	http.ServeFile(w,r,"home.html")
}

func chat(w http.ResponseWriter, r *http.Request){


	http.ServeFile(w,r,"chat.html")
}

//func wsSocket(w http.ResponseWriter, r *http.Request){


	


	// upgrader.CheckOrigin = func (r *http.Request)bool{return true}
	// con,er:= upgrader.Upgrade(w,r,nil)
	// if er!=nil{
	// 	log.Fatalf("error is %v",er)
	// 	return
	// }

	// for{
	// 	var x string
	// 	messageType,p,err:=con.ReadMessage()
	// 	if err!=nil{
	// 		log.Fatalf("error is %v",err)
	// 		return
	// 	}
	// 	fmt.Println(string(p))

	// 	fmt.Println(messageType)
	// 	fmt.Scanln(&x)
	// 	if er:=con.WriteMessage(messageType,[]byte(x));er!=nil{
	// 		log.Fatal(er)
	// 		return
	// 	}

	// }


//}