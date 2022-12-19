package main

import (
	"fmt"
	"gitlab.boquar.tech/galileosky/pkg/crc"
	"gitlab.boquar.tech/galileosky/pkg/logger"
	"test/test1"
	"test/test2"
)

func main() {
	var dataArray []uint8
	var polynom uint32

	result := crc.CalcCRC32(dataArray, polynom)

	fmt.Println(result)

	logger.Init(
		"cfg.App.Name",
		"cfg.App.Environment",
		"cfg.Log.Level",
	)

	logger.I.Fatal("app - Run - postgres.New: %w")

	test := test1.Test{
		Name: "super",
	}

	test2 := test2.Test(test)

	fmt.Print(test2)
}
