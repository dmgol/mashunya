package main

import (
	"fmt"

	_ "github.com/dmgol/mashunya/db/migrations"

	"github.com/dmgol/mashunya/config"
)

func main() {
	fmt.Println(config.Config)
}
