package logx

import (
    "context"
    "fmt"
    "github.com/go-kratos/kratos/v2/errors"
    "github.com/go-kratos/kratos/v2/log"
    "github.com/go-kratos/kratos/v2/middleware"
    "github.com/go-kratos/kratos/v2/transport"
    "time"
)

func Server(logger log.Logger) middleware.Middleware {
    return func(handler middleware.Handler) middleware.Handler {
        return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
            var (
                code      int32
                reason    string
                kind      string
                operation string
            )
            startTime := time.Now()
            if info, ok := transport.FromServerContext(ctx); ok {
                kind = info.Kind().String()
                operation = info.Operation()
            }
            reply, err = handler(ctx, req)
            if se := errors.FromError(err); se != nil {
                code = se.Code
                reason = se.Reason
                _ = log.WithContext(ctx, logger).Log(
                    log.LevelError,
                    "kind", "server",
                    "component", kind,
                    "operation", operation,
                    "args", extractArgs(req),
                    "code", code,
                    "reason", reason,
                    "stack", err.Error(),
                    "duration", fmt.Sprintf("%dms", time.Since(startTime).Milliseconds()),
                )
            } else {
                _ = log.WithContext(ctx, logger).Log(
                    log.LevelInfo,
                    "kind", "server",
                    "component", kind,
                    "operation", operation,
                    "args", extractArgs(req),
                    "duration", fmt.Sprintf("%dms", time.Since(startTime).Milliseconds()),
                )
            }
            return
        }
    }
}

func extractArgs(req interface{}) string {
    if stringer, ok := req.(fmt.Stringer); ok {
        return stringer.String()
    }
    return fmt.Sprintf("%+v", req)
}
