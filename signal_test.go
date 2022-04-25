package c2api

import "testing"

func TestSubmitSignal(t *testing.T) {
	client, err := NewDefaultClient("")
	if err != nil {
		t.Error("Unexpected error")
	}

	signal := Signal{}

	conf, resp, _ := client.Signal.SubmitSignal("1234", signal)

	t.Log(conf)
	t.Log(resp.Body)
}
