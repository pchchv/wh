package gitlab

import (
	"log"
	"os"
	"testing"
)

var hook *Webhook

func TestMain(m *testing.M) {
	// setup
	var err error
	if hook, err = New(Options.Secret("sampleToken!")); err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
	// teardown
}
