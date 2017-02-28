# log2
--
    import "github.com/msample/log2"

Log2 defines Debug, Info, Warn, Error and Audit structured logging functions
whose implementations may be swaapped atomically.

Main purpose: provide a standard way for all go packages to link to logging
functions without restricting what actual logging implementation an app chooses
to use (other than imposing specific levels & structured logging). If all
packages used something like this, it would be easy for an App to configure
logging, including changing during runtime (e.g. in response to SIG_USR1 &
SIG_USR2).

Does not support pkg-specific logging level based muting like log4j unless a
"pkg" param convention is used to logging calls.

Debug, Info, Error, etc logging functions take alternating key and value pairs
as arguments as per go-kit/kit/log. Keys will be converted string with
fmt.Sprint(). If given an uneven number of keyvals, the first keyval will be
assumed to be a value and will be given the key "msg" implicitly (simplify
migration at cost of alloc)

## Usage

#### func  Audit

```go
func Audit(keyvals ...interface{}) error
```

#### func  Debug

```go
func Debug(keyvals ...interface{}) error
```

#### func  Error

```go
func Error(keyvals ...interface{}) error
```

#### func  Info

```go
func Info(keyvals ...interface{}) error
```

#### func  Warn

```go
func Warn(keyvals ...interface{}) error
```

#### type LogFunc

```go
type LogFunc func(...interface{}) error
```


#### func  SwapAudit

```go
func SwapAudit(f LogFunc) LogFunc
```

#### func  SwapDebug

```go
func SwapDebug(f LogFunc) LogFunc
```

#### func  SwapError

```go
func SwapError(f LogFunc) LogFunc
```

#### func  SwapInfo

```go
func SwapInfo(f LogFunc) LogFunc
```

#### func  SwapWarn

```go
func SwapWarn(f LogFunc) LogFunc
```
