package log

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type ContextField string

func (c ContextField) String() string {
	return string(c)
}

const (
	ErrorLevel uint8 = iota
	WarnLevel
	InfoLevel
	TraceLevel
	DebugLevel
)

const (
	debugStr = "debug"
	infoStr  = "info"
	warnStr  = "warn"
	errorStr = "error"
	traceStr = "trace"
)

const (
	TracingTime   ContextField = "tracing_time"
	UserID        ContextField = "user_id"
	ClientHost    ContextField = "client_host"
	CorrelationID ContextField = "correlation_id"
	HTTPPath      ContextField = "http_path"
)

var loggable = []fmt.Stringer{UserID, ClientHost, CorrelationID, HTTPPath}

var level = DebugLevel

type Field struct {
	Key   string
	Value string
}

func SetLoggableFields(custom []fmt.Stringer) {
	loggable = custom
}

func GetLoggableFields() []fmt.Stringer {
	return loggable
}

func AppendLoggableFields(l fmt.Stringer) {
	loggable = append(loggable, l)
}

func SetLevel(l uint8) {
	level = l
}

func Debug(ctx context.Context, msg string) {
	p.send(Parse(ctx, debugStr, msg, nil), DebugLevel)
}

func Debugf(ctx context.Context, format string, args ...interface{}) {
	p.send(Parse(ctx, debugStr, fmt.Sprintf(format, args...), nil), DebugLevel)
}

func Info(ctx context.Context, msg string) {
	p.send(Parse(ctx, infoStr, msg, nil), InfoLevel)
}

func Infof(ctx context.Context, format string, args ...interface{}) {
	p.send(Parse(ctx, infoStr, fmt.Sprintf(format, args...), nil), InfoLevel)
}

func Warn(ctx context.Context, msg string) {
	p.send(Parse(ctx, warnStr, msg, nil), WarnLevel)
}

func Warnf(ctx context.Context, format string, args ...interface{}) {
	p.send(Parse(ctx, warnStr, fmt.Sprintf(format, args...), nil), WarnLevel)
}

func Error(ctx context.Context, err error, msg string) {
	p.send(Parse(ctx, errorStr, msg, err), ErrorLevel)
}

func Errorf(ctx context.Context, err error, format string, args ...interface{}) {
	p.send(Parse(ctx, errorStr, fmt.Sprintf(format, args...), err), ErrorLevel)
}

func Trace(ctx context.Context, msg string) {
	f := []Field{
		{
			Key:   "timestamp",
			Value: fmt.Sprintf("%d", time.Now().UTC().UnixNano()),
		},
	}
	ctxVal := ctx.Value(TracingTime)
	if v, ok := ctxVal.(time.Time); ok {
		f = append(f, Field{Key: "spent", Value: time.Since(v).String()})
	}
	p.send(Parse(ctx, traceStr, msg, nil, f...), TraceLevel)
}

func Parse(ctx context.Context, scopeLevel string, msg string, err error, fields ...Field) string {
	var parsable = make(map[string]interface{})
	for _, loggableField := range loggable {
		var val string
		var ctxVal = ctx.Value(loggableField)
		if ctxVal == nil {
			continue
		}
		switch v := ctxVal.(type) {
		case fmt.Stringer:
			val = v.String()
		case string:
			val = v
		}
		if val != "" {
			parsable[loggableField.String()] = val
		}
	}
	parsable["time"] = time.Now().UTC().Format(time.RFC3339Nano)
	parsable["timestamp"] = time.Now().UTC().UnixNano()

	if err != nil {
		parsable["error"] = err.Error()
	}
	if msg != "" {
		parsable["message"] = msg
	}
	for _, field := range fields {
		parsable[field.Key] = field.Value
	}
	parsable["level"] = scopeLevel

	result, err := json.Marshal(parsable)
	if err != nil && level == DebugLevel {
		fmt.Println("Logger malfunctioning: ", err.Error())
	}
	return string(result)
}
