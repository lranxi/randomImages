package logger

import (
	"errors"
	"testing"
)

func TestJSONLogger(t *testing.T) {
	logger, err := NewJSONLogger(
		WithField("defined_key", "defined_value"),
	)
	if err != nil {
		t.Fatal(err)
	}
	defer logger.Sync()

	err = errors.New("pkg error")
	logger.Error("发生错误", WrapMeta(nil, NewMeta("key1", "value1"), NewMeta("key2", "value2"))...)
	logger.Error("发生错误", WrapMeta(err, NewMeta("key1", "value1"), NewMeta("key2", "value2"))...)

}

func BenchmarkJsonLogger(b *testing.B) {
	b.ResetTimer()
	logger, err := NewJSONLogger(
		WithField("defined_key", "defined_value"),
	)
	if err != nil {
		b.Fatal(err)
	}

	defer logger.Sync()

}
