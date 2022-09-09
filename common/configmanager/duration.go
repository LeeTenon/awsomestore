package configmanager

import (
    "log"
    "math"
    "strconv"
    "strings"
    "time"
)

const defaultDuration = time.Second

type Duration string

func (d Duration) AsDuration() time.Duration {
    s := string(d)
    index := strings.Index(s, "s")
    if index > 0 {
        numStr := s[:index]
        num, err := strconv.ParseFloat(numStr, 64)
        if err != nil {
            log.Println("parse time.duration fail, check your config file")
            return defaultDuration
        }
        return time.Duration(num * math.Pow(10, 9))
    }
    log.Println("parse time.duration fail, check your config file")
    return defaultDuration
}
