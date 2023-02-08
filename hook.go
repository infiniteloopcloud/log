package log

import "context"

func PreHookSetLogLevel(_ context.Context) {
	SetLevel(LevelToUint())
}
