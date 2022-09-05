package gate

type ErrorInfo struct {
    Code   int    `json:"code"`
    Reason string `json:"reason"`
    Msg    string `json:"msg"`
}
