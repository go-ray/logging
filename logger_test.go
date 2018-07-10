package logging

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	InitLogger("", "", "", "")
	InitLogger("./", "testlog", "", "")
	InitLogger("./", "testlog", "debug", "")
	InitLogger("./", "testlog", "info", "json")
	InitLogger("./", "testlog", "warn", "json")
	InitLogger("./", "testlog", "error", "json")
	InitLogger("./", "testlog", "fatal", "json")
	InitLogger("./", "testlog", "panic", "json")
	InitLogger("./", "testlog", "debug", "json")

	m.Run()
	c := exec.Command("ls")
	o, _ := c.CombinedOutput()
	l := strings.Split(string(o), "\n")

	for _, i := range l {
		if strings.Contains(i, "testlog") {
			os.Remove(i)
		}
	}
	fmt.Println("All test run over")
}

func TestDebug(t *testing.T) {
	s := "Test " + getCurFuncName() + " for log"
	Debug(s)
	checkexist(s, t)
}

func TestInfo(t *testing.T) {
	s := "Test " + getCurFuncName() + " for log"
	Info(s)
	checkexist(s, t)
}

func TestPrint(t *testing.T) {
	s := "Test " + getCurFuncName() + " for log"
	Print(s)
	checkexist(s, t)
}

func TestWarn(t *testing.T) {
	s := "Test " + getCurFuncName() + " for log"
	Warn(s)
	checkexist(s, t)
}

func TestWarning(t *testing.T) {
	s := "Test " + getCurFuncName() + " for log"
	Warning(s)
	checkexist(s, t)
}

func TestError(t *testing.T) {
	s := "Test " + getCurFuncName() + " for log"
	Error(s)
	checkexist(s, t)
}

func TestPanic(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Log("A recover got:", err)
		}
	}()

	s := "Test " + getCurFuncName() + " for log"
	Panic(s)
	checkexist(s, t)
}

func TestDebugln(t *testing.T) {
	s := "Test " + getCurFuncName() + " for log"
	Debugln(s)
	checkexist(s, t)
}

func TestInfoln(t *testing.T) {
	s := "Test " + getCurFuncName() + " for log"
	Infoln(s)
	checkexist(s, t)
}

func TestPrintln(t *testing.T) {
	s := "Test " + getCurFuncName() + " for log"
	Println(s)
	checkexist(s, t)
}

func TestWarnln(t *testing.T) {
	s := "Test " + getCurFuncName() + " for log"
	Warnln(s)
	checkexist(s, t)
}

func TestWarningln(t *testing.T) {
	s := "Test " + getCurFuncName() + " for log"
	Warningln(s)
	checkexist(s, t)
}

func TestErrorln(t *testing.T) {
	s := "Test " + getCurFuncName() + " for log"
	Errorln(s)
	checkexist(s, t)
}

func TestPanicln(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Log("A recover got:", err)
		}
	}()

	s := "Test " + getCurFuncName() + " for log"
	Panicln(s)
	checkexist(s, t)
}

func TestDebugf(t *testing.T) {
	s := "Test " + getCurFuncName() + " for log:%s"
	pad := "go"
	Debugf(s, pad)
	checkexist(fmt.Sprintf(s, pad), t)
}

func TestInfof(t *testing.T) {
	s := "Test " + getCurFuncName() + " for log:%s"
	pad := "go"
	Infof(s, pad)
	checkexist(fmt.Sprintf(s, pad), t)
}

func TestPrintf(t *testing.T) {
	s := "Test " + getCurFuncName() + " for log:%s"
	pad := "go"
	Printf(s, pad)
	checkexist(fmt.Sprintf(s, pad), t)
}

func TestWarnf(t *testing.T) {
	s := "Test " + getCurFuncName() + " for log:%s"
	pad := "go"
	Warnf(s, pad)
	checkexist(fmt.Sprintf(s, pad), t)
}

func TestWarningf(t *testing.T) {
	s := "Test " + getCurFuncName() + " for log:%s"
	pad := "go"
	Warningf(s, pad)
	checkexist(fmt.Sprintf(s, pad), t)
}

func TestErrorf(t *testing.T) {
	s := "Test " + getCurFuncName() + " for log:%s"
	pad := "go"
	Errorf(s, pad)
	checkexist(fmt.Sprintf(s, pad), t)
}

func TestPanicf(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Log("A recover got:", err)
		}
	}()

	s := "Test " + getCurFuncName() + " for log:%s"
	pad := "go"
	Panicf(s, pad)
	checkexist(fmt.Sprintf(s, pad), t)
}

func getCurFuncName() string {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])

	return f.Name()
}

func checkexist(s string, t *testing.T) {
	d, err := ioutil.ReadFile("./testlog")
	if err != nil {
		if os.IsNotExist(err) {
			t.Error("Expected to create the test log file, but not:")
		}
		t.Error("Unit test got an err:", err)
	}
	if !bytes.Contains(d, []byte(s)) {
		t.Error("The string be logged should be in the log file, but not")
	}
}
