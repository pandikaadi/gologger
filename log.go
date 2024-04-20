package logger

import (
	"fmt"
)

func Log(mess string) {
	fmt.Println("[LOG]" + mess)
}