package requestid

// REF: https://github.com/gofiber/fiber/blob/main/middleware/requestid/requestid.go

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

// contextKey is a custom type used for context keys to avoid conflicts with keys from other packages.
// It's unexported to ensure that it does not conflict with context keys defined elsewhere.
type contextKey int

// Constants defining keys for storing and retrieving values from context.
const (
	requestIDKey contextKey = iota // requestIDKey is used to store and retrieve the request ID from the context.
)

// SetContext adds the request ID from the fiber context (c) to the provided context (ctx).
// It returns a new context with the request ID included.
//
// Parameters:
// - ctx: The context to which the request ID will be added.
// - c: The fiber context containing the request ID.
//
// Returns:
// - A new context with the request ID.
func SetContext(ctx context.Context, c *fiber.Ctx) context.Context {
	if rid, ok := c.Locals("requestid").(string); ok {
		return context.WithValue(ctx, requestIDKey, rid)
	}
	return ctx
}

// FromContext retrieves the request ID from the given context.
//
// Parameters:
// - c: The context from which the request ID will be retrieved.
//
// Returns:
// - The request ID as a string if present in the context; otherwise, returns an empty string.
func FromContext(c context.Context) string {
	if rid, ok := c.Value(requestIDKey).(string); ok {
		return rid
	}
	return ""
}
