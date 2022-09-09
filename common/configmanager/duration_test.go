package configmanager

import (
    "github.com/go-playground/assert/v2"
    "testing"
    "time"
)

const testTime Duration = "3s"

func TestAsDuration(t *testing.T) {
    assert.Equal(t, 3*time.Second, testTime.AsDuration())
}
