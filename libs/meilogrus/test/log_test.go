package test

import (
	"testing"
	"time"

	"github.com/andytyc/goutils/libs/meilogrus"
)

func TestLog(t *testing.T) {
	hhh := 666
	for i := 0; i < 30; i++ {
		meilogrus.Info("===>", i, "nihao", hhh)
		time.Sleep(time.Second)
	}
}
