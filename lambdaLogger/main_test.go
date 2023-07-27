package lambdaLogger

import (
	"bytes"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer

	l := New("test", true, &buf)
	l.Print(Info, "Hello, Info!")

	logOutput := buf.String()
	expectedOutput := "INFO testHello, Info!"

	if !strings.Contains(logOutput, expectedOutput) {
		t.Errorf("expected '%s' but got '%s'", expectedOutput, logOutput)
	}
}

func TestPrint(t *testing.T) {
	var buf bytes.Buffer

	_ = New("test", true, &buf)
	l := GetLogger()
	l.Print(Warning, "Hello, Warning!")

	logOutput := buf.String()
	expectedOutput := "WARNING testHello, Warning!"

	if !strings.Contains(logOutput, expectedOutput) {
		t.Errorf("expected '%s' but got '%s'", expectedOutput, logOutput)
	}
}

func TestPrintln(t *testing.T) {
	var buf bytes.Buffer

	l := New("test", true, &buf)
	l.Println(Error, "Hello, Error!")

	logOutput := buf.String()
	expectedOutput := "ERROR testHello, Error!"

	if !strings.Contains(logOutput, expectedOutput) {
		t.Errorf("expected '%s' but got '%s'", expectedOutput, logOutput)
	}
}

func TestPrintJson(t *testing.T) {
	var buf bytes.Buffer

	l := New("test", true, &buf)
	l.PrintJson(Debug, []byte(`{"message": "Hello, Debug!"}`))

	logOutput := buf.String()
	expectedOutput := "DEBUG test{\"message\": \"Hello, Debug!\"}"

	if !strings.Contains(logOutput, expectedOutput) {
		t.Errorf("expected '%s' but got '%s'", expectedOutput, logOutput)
	}
}

func TestPrintStruct(t *testing.T) {
	type testStruct struct {
		Message string `json:"message"`
	}

	var buf bytes.Buffer

	l := New("test", true, &buf)
	l.PrintStruct(Info, testStruct{Message: "Hello, Struct!"})

	logOutput := buf.String()
	expectedOutput := "INFO test{\"message\":\"Hello, Struct!\"}"

	if !strings.Contains(logOutput, expectedOutput) {
		t.Errorf("expected '%s' but got '%s'", expectedOutput, logOutput)
	}
}
