package models

import (
	"os"
	"testing"
)

func TestPopulatePlayerst(t *testing.T) {
	players, err := PopulatePlayers()

	t.Log(os.Getenv("MONGOLAB_URI"))

	if err != nil {
		t.Fatal("PopulatePlayers returned the following error:", err)
	}

	if len(*players) != 4 {
		t.Error("Playres should have a len of 4 not:", len(*players))
		t.Fail()
	}
}
