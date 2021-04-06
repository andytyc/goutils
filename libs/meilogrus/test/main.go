package main

import (
	"fmt"
	"time"

	"github.com/andytyc/goutils/libs/meilogrus"
	"github.com/sirupsen/logrus"
)

func main() {
	meilog := &meilogrus.MeiFileLog{LogLevel: logrus.DebugLevel}
	logger, err := meilog.GetLogger()
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println(meilog.LogFileName)
	fmt.Println(meilog.LogFilePath)
	// fmt.Println(meilog.LogFileName)
	// fmt.Println(meilog.LogFileName)

	for i := 0; i < 30; i++ {
		logger.Info("====>", i)
		time.Sleep(time.Second)
	}
}
