package commandUtils

import (
	"github.com/ggg17226/aghost-go-base/pkg/utils/fileUtils"
	"github.com/ggg17226/aghost-go-base/pkg/utils/randomUtils"
	"io"
	"os"
	"testing"
)

func TestExecCommand(t *testing.T) {
	randomFileName := randomUtils.RandStringWithLength(16)
	randomFileContent := randomUtils.RandStringWithLength(16)
	cmd := "echo '" + randomFileContent + "' > /tmp/test_" + randomFileName
	exitCode, err := ExecCommand(cmd)
	if err != nil {
		t.Fatalf("exec command error:  %s", err.Error())
	}
	if exitCode != 0 {
		t.Fatalf("exec command get non zero exit code")
	}

	if !fileUtils.FileExists("/tmp/test_" + randomFileName) {
		t.Fatalf("output file does not exist")
	}

	file, err := os.Open("/tmp/test_" + randomFileName)
	if err != nil {
		t.Fatalf("open output file error:  %s", err.Error())
	}
	defer file.Close()
	defer os.Remove("/tmp/test_" + randomFileName)

	content, err := io.ReadAll(file)
	if err != nil {
		t.Fatalf("read output file error:  %s", err.Error())
	}

	if string(content) != randomFileContent+"\n" {
		t.Fatalf("output file content does not match")
	}

}

func TestExecCommandAndGetStdout(t *testing.T) {
	cmd := "/usr/bin/free -m"
	exitCode, stdout, err := ExecCommandAndGetOutput(cmd)
	if err != nil {
		t.Fatalf("exec cmd error:  %s", err.Error())
	}
	t.Log(exitCode)
	t.Log(*stdout)
}
