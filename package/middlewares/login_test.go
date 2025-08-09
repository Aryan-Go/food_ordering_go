package middlewares

import "testing"

func TestEmailver(t *testing.T) {
	got := EmailVerification(`goyal.aryan@gmail.com`)
	want := true
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
