package common

import (
	"testing"
)

func TestCreateGuid(t *testing.T) {
	var guid string
	var err error
	for i := 10; i > 0; i-- {
		guid, err = CreateGuid()
		if err != nil {
			t.Error(err)
		} else {
			t.Log(guid)
		}
	}
}

func TestCreateRandomFromSeed(t *testing.T) {
	str, err := CreateRandomFromSeed(5, RandSeed)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(str)
	}
}

func BenchmarkCreateGuid(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		CreateGuid()
	}
	b.StopTimer()
}
