package azure

import (
	"log"
	"os"
	"testing"
)

var hook *Webhook

func TestMain(m *testing.M) {
	// setup
	var err error
	hook, err = New()
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
	// teardown
}
