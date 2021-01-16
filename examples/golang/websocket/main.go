package main

import (
	`log`
	`net/http`

	`github.com/gorilla/websocket`
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main()  {
	log.Println("start a server")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "home.html")
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received a request")

		_, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}
		w.Write([]byte("OK"))
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
