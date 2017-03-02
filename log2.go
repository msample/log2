// Log2 defines Debug, Info, Warn, Error and Audit structured logging
// functions whose implementations may be swapped atomically.
//
// Main purpose: provide a standard way for all go packages to link to
// logging functions without restricting what actual logging
// implementation an app chooses to use (other than imposing specific
// levels & structured logging).  If all packages used something like
// this, it would be easy for an app to configure logging, including
// changing during runtime (e.g. in response to SIG_USR1 & SIG_USR2).
//
// Does not support pkg-specific logging level based muting like log4j
// unless a "pkg" param convention is used to logging calls.
//
// Debug, Info, Error, etc logging functions take alternating key and
// value pairs as arguments as per go-kit/kit/log. Keys will be
// converted string with fmt.Sprint(). If given an uneven number of
// keyvals, the first keyval will be assumed to be a value and will be
// given the key "msg" implicitly (simplify migration at cost of
// alloc)
//
// The SwapDebug and other swap functions are intended to be used by
// the app writer (e.g. in main()) to configure how the Debug, Info
// etc functions behave. For example, it might make Error() and Warn()
// go to stdout in a simple text format while at the same time
// Error(), Warn() and Info() got to to a centralized log server over
// syslog/UDP in JSON format.  Audit() could be configured to write
// JSON to a Kafka topic partion selected by hashing the "id" keyval
// key. Swap functions make it easy to respond to SIG_HUP for logrotate
// type behvaiour (e.g. close current logging file and reopen it).
//
package log2

import (
	"sync/atomic"
	"unsafe"
)

var (
	debugFunc *LogFunc
	infoFunc  *LogFunc
	warnFunc  *LogFunc
	errorFunc *LogFunc
	auditFunc *LogFunc
)

type LogFunc func(...interface{}) error

func Debug(keyvals ...interface{}) error {
	return log(&infoFunc, keyvals)
}

func SwapDebug(f LogFunc) LogFunc {
	return swap(&debugFunc, f)
}

func Info(keyvals ...interface{}) error {
	return log(&infoFunc, keyvals)
}

func SwapInfo(f LogFunc) LogFunc {
	return swap(&infoFunc, f)
}

func Warn(keyvals ...interface{}) error {
	return log(&warnFunc, keyvals)
}

func SwapWarn(f LogFunc) LogFunc {
	return swap(&warnFunc, f)
}

func Error(keyvals ...interface{}) error {
	return log(&errorFunc, keyvals)
}

func SwapError(f LogFunc) LogFunc {
	return swap(&errorFunc, f)
}

func Audit(keyvals ...interface{}) error {
	return log(&auditFunc, keyvals)
}

func SwapAudit(f LogFunc) LogFunc {
	return swap(&auditFunc, f)
}

func log(fp **LogFunc, keyvals []interface{}) error {
	// grim
	fptr := (*unsafe.Pointer)(unsafe.Pointer(fp))
	f := (*LogFunc)(atomic.LoadPointer(fptr))
	if f == nil {
		return nil
	}
	if len(keyvals)%2 == 0 {
		return (*f)(keyvals...)
	}
	return (*f)(append([]interface{}{"msg"}, keyvals...)...)
}

func swap(fp **LogFunc, f LogFunc) LogFunc {
	x := (*unsafe.Pointer)(unsafe.Pointer(fp))
	rv := (*LogFunc)(atomic.SwapPointer(x, unsafe.Pointer(&f)))
	if rv == nil {
		return nopLog
	}
	return *rv
}

func nopLog(keyvals ...interface{}) error {
	return nil
}
