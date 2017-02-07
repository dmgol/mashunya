package main

import _ "github.com/dmgol/mashunya/db/migrations"
import "github.com/dmgol/mashunya/app/server"

func main() {
	go server.ServeAdmins()
	server.ServeAPI()
}
