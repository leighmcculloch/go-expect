package test_test

import (
	"fmt"
	"strings"
	"testing"
)

type fakeT struct {
	testing.T
	LogCalls   []string
	ErrorCalls []string
	FatalCalls []string
}

func (ft *fakeT) Log(args ...interface{}) {
	fmtArgs := make([]string, len(args))
	for i, arg := range args {
		fmtArgs[i] = fmt.Sprintf("%v", arg)
	}
	ft.Logf(strings.Join(fmtArgs, " "))
}

func (ft *fakeT) Logf(format string, args ...interface{}) {
	if ft.LogCalls == nil {
		ft.LogCalls = []string{}
	}
	call := fmt.Sprintf(format, args...)
	ft.LogCalls = append(ft.LogCalls, call)
}

func (ft *fakeT) Error(args ...interface{}) {
	fmtArgs := make([]string, len(args))
	for i, arg := range args {
		fmtArgs[i] = fmt.Sprintf("%v", arg)
	}
	ft.Errorf(strings.Join(fmtArgs, " "))
}

func (ft *fakeT) Errorf(format string, args ...interface{}) {
	if ft.ErrorCalls == nil {
		ft.ErrorCalls = []string{}
	}
	call := fmt.Sprintf(format, args...)
	ft.ErrorCalls = append(ft.ErrorCalls, call)
}

func (ft *fakeT) Fatal(args ...interface{}) {
	fmtArgs := make([]string, len(args))
	for i, arg := range args {
		fmtArgs[i] = fmt.Sprintf("%v", arg)
	}
	ft.Fatalf(strings.Join(fmtArgs, " "))
}

func (ft *fakeT) Fatalf(format string, args ...interface{}) {
	if ft.FatalCalls == nil {
		ft.FatalCalls = []string{}
	}
	call := fmt.Sprintf(format, args...)
	ft.FatalCalls = append(ft.FatalCalls, call)
}
