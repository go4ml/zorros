package tests

import (
	"bytes"
	"go-ml.dev/pkg/zorros/zlog"
	"gotest.tools/assert"
	"strings"
	"testing"
)

func Test_Init(t *testing.T) {
	defer zlog.Config{Name: "test log", Verbose: true}.Init().Close()
	zlog.Info("hello logger!")
}

func Test_LogWriter(t *testing.T) {
	bf := bytes.Buffer{}
	func() {
		defer zlog.Config{Name: "test log", LogWriter: &bf}.Init().Close()
		zlog.Info("hello logger!")
	}()
	assert.Assert(t, strings.Contains(bf.String(), "hello logger!"))
}
