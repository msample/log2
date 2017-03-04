package log2_test

import (
	"errors"
	"fmt"
	"math"
	"testing"

	"github.com/msample/log2"
)

func ExampleError() {
	result, err := doIt()
	if err != nil {
		log2.Error("msg", err)
		return
	}
	fmt.Printf(result)
}

func doIt() (string, error) {
	return "", errors.New("I'm sorry, Dave. I'm afraid I can't do that.")
}

// check each level's log & swap funcs to make sure they're using the
// same func var
func TestDebug(t *testing.T) {
	testLevel(t, log2.Debug, log2.SwapDebug)
}

func TestInfo(t *testing.T) {
	testLevel(t, log2.Info, log2.SwapInfo)
}

func TestWarn(t *testing.T) {
	testLevel(t, log2.Warn, log2.SwapWarn)
}

func TestError(t *testing.T) {
	testLevel(t, log2.Error, log2.SwapError)
}

func TestAudit(t *testing.T) {
	testLevel(t, log2.Audit, log2.SwapAudit)
}

// swap in new log func under f using swap() to ensure that f()
// results in the swapped in function being called correctly
func testLevel(t *testing.T, f log2.LogFunc, swap log2.SwapFunc) {

	// ensure params passed through unmutated and in original
	// order, also track total times called
	callCount := 0
	prev := swap(func(kvs ...interface{}) error {
		callCount++
		sum := 0
		prev := math.MinInt64
		for _, kv := range kvs {
			if _, ok := kv.(string); ok {
				continue
			}
			v := kv.(int)
			sum += v
			if v < prev {
				t.Errorf("Param ordering check: integers should be ascending: %v", kvs)
			}
			prev = v
		}
		if sum != 100 {
			t.Errorf("Param mutation check: expected sum of %v, got %v from %v\n", 100, sum, kvs)
		}
		return nil
	})
	f(10, 20, 30, 40)
	f(100)
	f(1, 99)
	f(-1, 1, 100)
	f(-1, -1, 2, 100)
	f(-1, -1, 1, 1, 2, 4, 94)
	f(-1, -1, 1, 1, 2, 4, 6, 88)
	f("k1", 100)
	f("k1", 50, "k2", 50)
	f("xyzk3", -400, "k1", 50, "k2", 50, "k", 400)

	switch callCount {
	case 0:
		t.Errorf("given log func not calling swapped-in log func impl. Mismatched log var?")
	case 10:
	default:
		t.Errorf("Unexpected call count to swapped in log func: %v, expected %v", callCount, 10)
	}

	// now verify that "msg" gets prepended for odd length keyvals
	// lists and that errors are returned as is
	swap(func(kvs ...interface{}) error {
		k1, ok := kvs[0].(string)
		if !ok || k1 != "msg" {
			t.Errorf(`expected first key of "msg', got %v\n`, kvs[0])
		}
		return testErr("Oops")
	})
	err := f(1)
	if _, ok := err.(testErr); !ok {
		t.Errorf("expected testErr type back, got %T\n", err)
	}
	f("msg", 1, 2, 3, 4, 5)
	swap(prev)
}

type testErr string

func (o testErr) Error() string {
	return string(o)
}

// Show tests - just to see output and absence of panics. They don't
// test much
func TestShowInfo(t *testing.T) {
	log2.Info("foo")
	log2.Info("foo", 2)

	f := log2.SwapInfo(func(kv ...interface{}) error {
		fmt.Printf("Info: %v\n", kv)
		return nil
	})
	log2.Info("foo")
	log2.Info("foo", 2)
	log2.Info("foo", 2, "msg", 83.83, "level", "info", "baz", true)

	o := log2.SwapInfo(f)
	log2.Info("foo")
	log2.Info("foo", 2)
	log2.Info("foo", 2, "msg", 83.83, "level", "info", "baz", true)

	log2.SwapInfo(o)
	log2.Info("foo")
	log2.Info("foo", 2)
	log2.Info("foo", 2, "msg", 83.83, "level", "info", "baz", true)

	log2.SwapInfo(func(kv ...interface{}) error {
		fmt.Printf("Info: 2nd impl: %v\n", kv)
		return nil
	})
	log2.Info("foo")
	log2.Info("foo", 2)
	log2.Info("foo", 2, "msg", 83.83, "level", "info", "baz", true)
}

func TestShowError(t *testing.T) {
	log2.Error("foo")
	log2.Error("foo", 2)

	f := log2.SwapError(func(kv ...interface{}) error {
		fmt.Printf("Error: %v\n", kv)
		return nil
	})
	log2.Error("foo")
	log2.Error("foo", 2)
	log2.Error("foo", 2, "msg", 83.83, "level", "info", "baz", true)

	o := log2.SwapError(f)
	log2.Error("foo")
	log2.Error("foo", 2)
	log2.Error("foo", 2, "msg", 83.83, "level", "info", "baz", true)

	log2.SwapError(o)
	log2.Error("foo")
	log2.Error("foo", 2)
	log2.Error("foo", 2, "msg", 83.83, "level", "info", "baz", true)

	log2.SwapError(func(kv ...interface{}) error {
		fmt.Printf("Error: 2nd impl: %v\n", kv)
		return nil
	})
	log2.Error("foo")
	log2.Error("foo", 2)
	log2.Error("foo", 2, "msg", 83.83, "level", "info", "baz", true)
}
