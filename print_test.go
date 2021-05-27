package test_test

import (
	"fmt"
	"strings"
	"testing"
)

type printT struct {
	testing.T
}

func (pt *printT) Log(args ...interface{}) {
	fmtArgs := make([]string, len(args))
	for i, arg := range args {
		fmtArgs[i] = fmt.Sprintf("%v", arg)
	}
	fmt.Println(strings.Join(fmtArgs, " "))
}

func (pt *printT) Logf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

func (pt *printT) Error(args ...interface{}) {
	fmtArgs := make([]string, len(args))
	for i, arg := range args {
		fmtArgs[i] = fmt.Sprintf("%v", arg)
	}
	fmt.Println(strings.Join(fmtArgs, " "))
}

func (pt *printT) Errorf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

func (pt *printT) Fatal(args ...interface{}) {
	fmtArgs := make([]string, len(args))
	for i, arg := range args {
		fmtArgs[i] = fmt.Sprintf("%v", arg)
	}
	fmt.Println(strings.Join(fmtArgs, " "))
}

func (pt *printT) Fatalf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
