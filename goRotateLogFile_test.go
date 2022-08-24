package gorotatelogfile

import (
	"testing"
	"time"
)

func Test01(t *testing.T) {
	LogInit("/tmp/testfile-01.log", 10, true, 0, 0)
	fixedFieldsString := make(map[string]string)
	fixedFieldsString["TID"] = "TID012345"
	fixedFieldsInt := make(map[string]int)
	fixedFieldsInt["TInt"] = 1234567890
	logs := Logs{FixedFieldsString: fixedFieldsString, FixedFieldsInt: fixedFieldsInt}
	for i := 0; i < 1000; i++ {
		time.Sleep(10 * time.Millisecond)
		logs.Info("HELLO INFO")
	}
}
