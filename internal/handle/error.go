package handle

import (
	"fmt"
	"os"
)

func Error(s string) {
	fmt.Print("\033[1;33m" + s + "\033[0m\n\n")
	os.Exit(1)
}
