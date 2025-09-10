package gogs

import (
	"log"
	"os"
	"testing"
)

var hook *Webhook

func TestMain(m *testing.M) {
	// setup
	var err error
	hook, err = New(Options.Secret("sampleToken!"))
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
	//teardown
}
