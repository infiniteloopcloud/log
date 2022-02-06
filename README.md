# Log

### Usage

```go
package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/infiniteloopcloud/log"
)

const (
	CorrelationID log.ContextField = "correlation_id"
)

func main() {
	// Set this globally
	log.SetLevel(log.LevelToUint())
	log.SetLoggableFields([]fmt.Stringer{CorrelationID})
	 
	ctx := context.Background()
	ctx = context.WithValue(ctx, CorrelationID, "123456")
	log.Error(ctx, errors.New(""), "")
	log.Errorf(ctx, errors.New(""), "test: %s", "test")
	log.Warn(ctx, "")
	log.Info(ctx, "")
	log.Debug(ctx, "")
	// etc...
}
```