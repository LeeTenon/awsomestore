package logx

import (
    "github.com/go-kratos/kratos/v2/log"
    "github.com/go-kratos/kratos/v2/middleware/tracing"
    "io"
)

func New(svcName string, w io.Writer) log.Logger {
    return log.With(
        log.NewStdLogger(w),
        "ts", log.DefaultTimestamp,
        "caller", log.DefaultCaller,
        "service.name", svcName,
        "trace.id", tracing.TraceID(),
        "span.id", tracing.SpanID(),
    )
}
