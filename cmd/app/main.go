package main

import (
	"fmt"
	"github.com/forabelo/test-task-perl/config"
	"github.com/forabelo/test-task-perl/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}

	app.Run(&cfg)
}
