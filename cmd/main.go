package main

import (
	"fmt"

	"github.com/ichi-pg/golang-middleware/util"
	"github.com/ichi-pg/golang-server/internal/presentation/echo"
)

func main() {
	if err := util.InitRand(); err != nil {
		fmt.Println(err.Error())
		return
	}
	if err := echo.Start(); err != nil {
		fmt.Println(err.Error())
	}
}
