package logging

import "fmt"

func Errorf(format string, args ...any) {
	fmt.Printf("[ERROR] "+format, args...)
}

func Infof(format string, args ...any) {
	fmt.Printf("[INFO] "+format, args...)
}

func Plainln(a ...any) {
	fmt.Println(a...)
}
