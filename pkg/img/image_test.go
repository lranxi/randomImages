package img

import (
	"testing"
)

func TestComputePHash(t *testing.T) {
	phash, err := ComputePHash("/Users/wenlys/1.jpg")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(phash)
}

func BenchmarkComputePHash(b *testing.B) {
	phash, err := ComputePHash("/Users/wenlys/1.jpg")
	if err != nil {
		b.Fatal(err)
	}
	b.Log(phash)
}

func TestCompression(t *testing.T) {
	file, err := Compression("/Users/wenlys/code/images/api/static/images/uncheck/P8JrqNWzwZ3nvI5O.jpg", uint(1920), uint(1080))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(file)
}

func BenchmarkCompressImg(b *testing.B) {
	b.ResetTimer()
	file, err := Compression("/Users/wenlys/code/images/api/static/images/uncheck/P8JrqNWzwZ3nvI5O.jpg", uint(1920), uint(1080))
	if err != nil {
		b.Fatal(err)
	}
	b.Log(file)
}
