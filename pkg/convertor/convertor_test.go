package convertor

import "testing"

func TestToString(t *testing.T) {
	var n32 int32 = 32
	strN32 := ToString(n32)
	if strN32 != "32" {
		t.Errorf("Number 32(int32) to string result isn't '32', err result is %s", strN32)
	}

	var n64 int64 = 64
	strN64 := ToString(n64)
	if strN64 != "64" {
		t.Errorf("Number 64(int64) to string result isn't '64', err result is %s", strN64)
	}

	var f32 float32 = 1.12345
	strF32 := ToString(f32)
	if strF32 != "1.12345" {
		t.Errorf("Float Number 32.32(float32) to string result isn't '1.12345', err result is %s", strF32)
	}

	var f64 = 912078294.0
	strF64 := ToString(f64)
	if strF64 != "912078294" {
		t.Errorf("Float Number 32.32(float32) to string result isn't '912078294', err result is %s", strF64)
	}
}
