package log2_test

import (
	"fmt"
	"testing"

	"github.com/msample/log2"
)

func TestInfo(t *testing.T) {
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

func TestError(t *testing.T) {
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
