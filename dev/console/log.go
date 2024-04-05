package console

import (
	"os"
	"github.com/k0kubun/pp"
)

func Log(a ...interface{}) {
	pp.Println(a...)
}

func LogAndDie(a ...interface{}) {
	pp.Println(a...)
	os.Exit(1)
}
