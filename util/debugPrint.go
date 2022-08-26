package util

import (
	"log"
	"os"
)

func DebugPrint(message ...any) {
	if os.Getenv("UTILDEBUG") != "" {
		log.Println("[DEBUG]", message)
	}
}
