package test

import (
	"net/http"
	"testing"

	"github.com/c12s/whirlpool-go"
)

func TestLibraryOnRestNode(t *testing.T) {
	var probe []whirlpool.Probe

	probe = append(probe, whirlpool.NewLivenessProbe())

	restProbe := whirlpool.NewRestProbe("http://localhost", "metrics", 8080, probe)

	var configs []whirlpool.ProbeMethod

	probingClient := whirlpool.New(append(configs, &restProbe))

	response, err := probingClient.StartProbing()

	if err != nil {
		t.Fatalf("Failed to reach the rest node: %v", err)
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %d", response.StatusCode)
	}
}
