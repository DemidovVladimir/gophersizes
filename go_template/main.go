package main

import (
	"log"
	"os"
)

// "github.com/gorilla/websocket"

// // We'll need to define an Upgrader
// // this will require a Read and Write buffer size
// var upgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// }

// type Product struct {
// 	name        string
// 	description string
// 	created     Time
// }

// // define a reader which will listen for
// // new messages being sent to our WebSocket
// // endpoint
// func reader(conn *websocket.Conn) {
// 	for {
// 		// read in a message
// 		messageType, p, err := conn.ReadMessage()
// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}
// 		// print out that message for clarity
// 		fmt.Println(string(p))

// 		if err := conn.WriteMessage(messageType, p); err != nil {
// 			log.Println(err)
// 			return
// 		}

// 	}
// }

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Home Page")
// }

// func wsEndpoint(w http.ResponseWriter, r *http.Request) {
// 	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
// 	fmt.Println(r.Host)

// 	// upgrade this connection to a WebSocket
// 	// connection
// 	ws, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	// listen indefinitely for new messages coming
// 	// through on our WebSocket connection
// 	reader(ws)
// }

// func setupRoutes() {
// 	http.HandleFunc("/", homePage)
// 	http.HandleFunc("/ws", wsEndpoint)
// }

type readOp struct {
	key  int
	read chan int
}

type writeOp struct {
	key   int
	value int
	write chan bool
}

type Task struct {
	Command string
	*log.Logger
}

func NewTask(command string, logger *log.Logger) Task {
	return Task{command, logger}
}

func main() {

	var logger = log.New(os.Stdout, "", 5)

	task := NewTask("validate", logger)

	task.Println("Shit...")

	// var readOps uint64
	// var writeOps uint64

	// reads := make(chan readOp)
	// writes := make(chan writeOp)

	// go func() {
	// 	state := make(map[int]int)

	// 	for {
	// 		select {
	// 		case read := <-reads:
	// 			read.read <- state[read.key]
	// 		case write := <-writes:
	// 			state[write.key] = write.value
	// 			write.write <- true
	// 		}
	// 	}
	// }()

	// for i := 0; i < 100; i++ {
	// 	go func() {
	// 		for {
	// 			read := readOp{
	// 				key:  rand.Intn(10),
	// 				read: make(chan int),
	// 			}
	// 			reads <- read
	// 			<-read.read
	// 			atomic.AddUint64(&readOps, 1)
	// 			time.Sleep(time.Millisecond)
	// 		}
	// 	}()
	// }

	// for i := 0; i < 10; i++ {
	// 	go func() {
	// 		for {
	// 			write := writeOp{
	// 				key:   rand.Intn(10),
	// 				value: rand.Intn(100),
	// 				write: make(chan bool),
	// 			}
	// 			writes <- write
	// 			<-write.write
	// 			atomic.AddUint64(&writeOps, 1)
	// 			time.Sleep(time.Millisecond)
	// 		}
	// 	}()
	// }

	// time.Sleep(time.Second)
	// readOpsFinal := atomic.LoadUint64(&readOps)
	// fmt.Println("readOps:", readOpsFinal)
	// writeOpsFinal := atomic.LoadUint64(&writeOps)
	// fmt.Println("writeOps:", writeOpsFinal)
	// fmt.Println("Hello World")
	// setupRoutes()
	// log.Fatal(http.ListenAndServe(":8080", nil))

}
