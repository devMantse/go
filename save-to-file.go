package main

import (
	"fmt"
	"os"
)

func valueToFile(value string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(value, []byte(valueText), 0644)
}
