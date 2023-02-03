package log

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"
	"time"
)

var (
	TestField ContextField = "test_field1"
)

func TestMain(m *testing.M) {
	loggable = []fmt.Stringer{TestField, UserID}
	os.Exit(m.Run())
}

func TestParse(t *testing.T) {
	ctx := context.WithValue(context.Background(), TestField, "test")
	ctx = context.WithValue(ctx, UserID, "test_user")
	result := parse(ctx, "info", "test_message", nil)
	t.Log(result)
}

func TestError(t *testing.T) {
	ctx := context.WithValue(context.Background(), TestField, "test")
	ctx = context.WithValue(ctx, UserID, "test_user")
	Error(ctx, errors.New("test_error"), "")
}

func TestTrace(t *testing.T) {
	ctx := context.WithValue(context.Background(), TestField, "test")
	ctx = context.WithValue(ctx, UserID, "test_user")
	ctx = context.WithValue(ctx, TracingTime, time.Now())
	Trace(ctx, "prepare query")
	time.Sleep(300 * time.Millisecond)
	Trace(ctx, "database query")
}

func BenchmarkParse(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		p = blank{}
		ctx := context.WithValue(context.Background(), TestField, "test")
		ctx = context.WithValue(ctx, UserID, "test_user")
		parse(ctx, "info", "test_message", nil)
	}
}

func BenchmarkError(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		p = blank{}
		ctx := context.WithValue(context.Background(), TestField, "test")
		ctx = context.WithValue(ctx, UserID, "test_user")
		Error(ctx, errors.New("test_error"), "")
	}
}
