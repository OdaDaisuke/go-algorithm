package process

import (
	"os"
	"fmt"
	"runtime"
)

func Run() {
	fmt.Println("Stdout Fd %v", os.Stdout.Fd())
	fmt.Println("Num CPU %v", runtime.NumCPU())
}
