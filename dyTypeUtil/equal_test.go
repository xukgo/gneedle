package dyTypeUtil

import (
	"sync"
	"testing"
)

func TestEqualPoint(t *testing.T) {
	var actual bool
	p := &sync.Map{}
	p2 := &sync.Map{}
	actual = Equal(p, p)
	if !actual {
		t.Fail()
	}
	actual = Equal(p, p2)
	if actual {
		t.Fail()
	}
}

func TestEqualBool(t *testing.T) {
	var actual bool
	actual = Equal(true, true)
	if !actual {
		t.Fail()
	}
	actual = Equal(false, false)
	if !actual {
		t.Fail()
	}
	actual = Equal(true, false)
	if actual {
		t.Fail()
	}
	actual = Equal(false, true)
	if actual {
		t.Fail()
	}
}

func TestEqualBoolString(t *testing.T) {
	var actual bool
	actual = Equal(true, "1")
	if !actual {
		t.Fail()
	}
	actual = Equal(true, "TRUE")
	if !actual {
		t.Fail()
	}
	actual = Equal(true, "true")
	if !actual {
		t.Fail()
	}
	actual = Equal("true", true)
	if !actual {
		t.Fail()
	}
	actual = Equal(false, "0")
	if !actual {
		t.Fail()
	}
	actual = Equal(false, "FALSE")
	if !actual {
		t.Fail()
	}
	actual = Equal(false, "false")
	if !actual {
		t.Fail()
	}
	actual = Equal(false, false)
	if !actual {
		t.Fail()
	}
	actual = Equal(true, "0")
	if actual {
		t.Fail()
	}
	actual = Equal(false, "true")
	if actual {
		t.Fail()
	}
	actual = Equal("true", false)
	if actual {
		t.Fail()
	}
}

func TestEqualBytesString(t *testing.T) {
	var actual bool
	actual = Equal([]byte("123abc"), "123abc")
	if !actual {
		t.Fail()
	}
	actual = Equal("123abc", []byte("123abc"))
	if !actual {
		t.Fail()
	}
	actual = Equal([]byte("464"), "123abc")
	if actual {
		t.Fail()
	}
	actual = Equal("fdsf", []byte("123abc"))
	if actual {
		t.Fail()
	}
}

func TestEqualBoolNumber(t *testing.T) {
	var actual bool
	actual = Equal(true, int8(1))
	if !actual {
		t.Fail()
	}
	actual = Equal(true, int64(1))
	if !actual {
		t.Fail()
	}
	actual = Equal(false, int8(0))
	if !actual {
		t.Fail()
	}
	actual = Equal(false, int8(0))
	if !actual {
		t.Fail()
	}
	actual = Equal(int8(1), true)
	if !actual {
		t.Fail()
	}
	actual = Equal(int64(1), true)
	if !actual {
		t.Fail()
	}
	actual = Equal(int8(0), false)
	if !actual {
		t.Fail()
	}
	actual = Equal(int8(0), false)
	if !actual {
		t.Fail()
	}

	actual = Equal(true, int8(2))
	if actual {
		t.Fail()
	}
	actual = Equal(int8(3), false)
	if actual {
		t.Fail()
	}
}

func TestEqualNumber(t *testing.T) {
	var actual bool
	actual = Equal(int64(123), int8(123))
	if !actual {
		t.Fail()
	}
	actual = Equal(int64(456), float32(456))
	if !actual {
		t.Fail()
	}
	actual = Equal(uint64(789), float64(789))
	if !actual {
		t.Fail()
	}
	actual = Equal(uint64(789), float64(789.001))
	if !actual {
		t.Fail()
	}

	actual = Equal(int64(456), int8(123))
	if actual {
		t.Fail()
	}
	actual = Equal(int64(456), float32(426.1))
	if actual {
		t.Fail()
	}
	actual = Equal(uint64(789), float64(543.54365))
	if actual {
		t.Fail()
	}
}

func TestEqualNumberString(t *testing.T) {
	var actual bool
	actual = Equal("123", int8(123))
	if !actual {
		t.Fail()
	}
	actual = Equal(float64(123.456), "123.456")
	if !actual {
		t.Fail()
	}
	actual = Equal(uint64(789), "789")
	if !actual {
		t.Fail()
	}

	actual = Equal("123.9", int8(123))
	if actual {
		t.Fail()
	}
	actual = Equal(float64(1234), "123.456")
	if actual {
		t.Fail()
	}
	actual = Equal(uint64(789), "1789")
	if actual {
		t.Fail()
	}
}

func BenchmarkEqualNumberString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Equal(123.456, "123.456")
	}
}

func BenchmarkEqualNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Equal(int64(1789), float64(1789.2))
	}
}
