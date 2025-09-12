package bitbucket_server

import (
	"log"
	"os"
	"testing"
)

const path = "/webhooks"

var hook *Webhook

func TestMain(m *testing.M) {
	// setup
	var err error
	hook, err = New(Options.Secret("secret"))
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
	// teardown
}
