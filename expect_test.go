package expect_test

import (
	"fmt"
	"testing"
)

type fakeT struct {
	testing.T
	LogCalls   []string
	ErrorCalls []string
	FatalCalls []string
}

func (ft *fakeT) Logf(format string, args ...interface{}) {
	if ft.LogCalls == nil {
		ft.LogCalls = []string{}
	}
	call := fmt.Sprintf(format, args...)
	ft.LogCalls = append(ft.LogCalls, call)
}

func (ft *fakeT) Errorf(format string, args ...interface{}) {
	if ft.ErrorCalls == nil {
		ft.ErrorCalls = []string{}
	}
	call := fmt.Sprintf(format, args...)
	ft.ErrorCalls = append(ft.ErrorCalls, call)
}

func (ft *fakeT) Fatalf(format string, args ...interface{}) {
	if ft.FatalCalls == nil {
		ft.FatalCalls = []string{}
	}
	call := fmt.Sprintf(format, args...)
	ft.FatalCalls = append(ft.FatalCalls, call)
}
