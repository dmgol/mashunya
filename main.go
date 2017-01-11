package main

import (
	"fmt"

	"github.com/dmgol/mashunya/config"
	_ "github.com/dmgol/mashunya/db/migrations"
)

func main() {
	fmt.Println(config.Config)
}
