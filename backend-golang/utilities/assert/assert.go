package assert

import "testing"

//True ..
func True(t *testing.T, condition bool) {
	t.Helper()
	if !condition {
		t.Fail()
	}
}

//NoError ..
func NoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fail()
	}
}

//ErrorRaised ..
func ErrorRaised(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Fail()
	}
}
