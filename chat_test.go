package rumblelivestreamlib

import (
	"os"
	"testing"
)

const (
	liveUrlEnvVar = "RUMBLE_LIVE_URL"
)

func TestMain(m *testing.M) {
	exitCode := run(m)
	os.Exit(exitCode)
}

func run(m *testing.M) int {
	return m.Run()
}

func TestChatInfo(t *testing.T) {
	url := os.Getenv(liveUrlEnvVar)
	if url == "" {
		t.Skipf("Set %s to run this test.", liveUrlEnvVar)
	}

	client, err := NewClient(NewClientOptions{LiveStreamUrl: url})
	if err != nil {
		t.Fatalf("Want NewClient err = nil, got: %v", err)
	}
	ci, err := client.ChatInfo(false)
	if err != nil {
		t.Fatalf("Want client.ChatInfo err = nil, got: %v", err)
	}

	if ci.ChatID == "" {
		t.Fatalf("Want non-empty chat ID")
	}
	if ci.Page == "" {
		t.Fatal("Want non-empty page")
	}
	if ci.UrlPrefix == "" {
		t.Fatal("Want non-empty url prefix")
	}
}
