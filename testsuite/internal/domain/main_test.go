package domain_test

import (
	"log"
	"os"
	"testing"
)

const VALID_UUIDV4 = "f47ac10b-58cc-4372-a567-0e02b2c3d479"

func TestMain(m *testing.M) {
	log.Println("TestMain from domain")
	os.Exit(m.Run())
}
