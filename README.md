# log2
--
    import "github.com/msample/log2"

Log2 defines Debug, Info, Warn, Error and Audit structured logging functions
whose implementations may be swapped atomically.

Main purpose: provide a standard way for all go packages to link to logging
functions without restricting what actual logging implementation an app chooses
to use (other than imposing specific levels & structured logging). If all
packages used something like this, it would be easy for an app to configure
logging, including changing during runtime (e.g. in response to SIG_USR1 &
SIG_USR2).

Does not support pkg-specific logging level based muting like log4j unless a
"pkg" param convention is used to logging calls.

Debug, Info, Error, etc logging functions take alternating key and value pairs
as arguments as per go-kit/kit/log. Keys will be converted string with
fmt.Sprint(). If given an uneven number of keyvals, the first keyval will be
assumed to be a value and will be given the key "msg" implicitly (simplify
migration at cost of alloc)

The SwapDebug and other swap functions are intended to be used by the app writer
(e.g. in main()) to configure how the Debug, Info etc functions behave. For
example, it might make Error() and Warn() go to stdout in a simple text format
while at the same time Error(), Warn() and Info() got to to a centralized log
server over syslog/UDP in JSON format. Audit() could be configured to write JSON
to a Kafka topic partion selected by hashing the "id" keyval key. Swap functions
make it easy to respond to SIG_HUP for logrotate type behvaiour (e.g. close
current logging file and reopen it).

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
