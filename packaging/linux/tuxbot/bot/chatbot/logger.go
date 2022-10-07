package chatbot

import (
	"runtime"
)

type Logger interface {
	Debug(format string, args ...any)
	Info(format string, args ...any)
	VDebug(format string, args ...any)
	Alert()
	AlertWith(string)
}

type LogFn func(format string, args ...any)

func GenericTraceVerbose(l Logger, s string, f func() (any, error), okLog LogFn, errLog LogFn) func() {
	pc, _, _, ok := runtime.Caller(2)
	details := runtime.FuncForPC(pc)
	name := "unknown function"
	if ok && details != nil {
		name = details.Name()
	}

	return func() {
		ret, err := f()
		if err != nil {
			errLog("`- %s(%s) -> %s`", name, s, err)
		} else {
			okLog("`- %s(%s) -> OK (%+v)`", name, s, ret)
		}
	}
}

func InfoTraceVerbose(l Logger, s string, f func() (any, error)) func() {
	return GenericTraceVerbose(l, s, f, l.Info, l.Info)
}

func DebugTraceResult(l Logger, f func() (any, error)) func() {
	return GenericTraceVerbose(l, "", f, l.Debug, l.Debug)
}

func DebugTrace(l Logger, f func() error) func() {
	return GenericTraceVerbose(l, "", func() (any, error) { return nil, f() }, l.Debug, l.Debug)
}

func DebugTraceVerbose(l Logger, s string, f func() (any, error)) func() {
	return GenericTraceVerbose(l, s, f, l.Debug, l.Debug)
}

func FunctionTraceVerbose(l Logger, s string, f func() (any, error)) func() {
	return GenericTraceVerbose(l, s, f, l.Debug, l.Info)
}
