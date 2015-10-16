package main

func main() {
	go AddHandlers(true, ":8080", "server.pem", "server.key")
  go AddHandlers(false,":8280","","")
  for {

  }
}
