package middlewares

import "testing"

func TestEmailver(t *testing.T) {
	got := Email_verification("fmnbsadfmn,asbf,ams")
	want := false
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
