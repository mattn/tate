package main

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestTate(t *testing.T) {
	var buf bytes.Buffer
	r := strings.NewReader("Golangを使い、\ntextを縦書きに\n変換するコマンドを\n書いたので、\n今後活用したい。\n")
	err := tate(&buf, r, option{})
	if err != nil {
		t.Fatal(err)
	}
	got := "今書変ｔＧ\n後い換ｅｏ\n活たすｘｌ\n用のるｔａ\nしでコをｎ\nた︑マ縦ｇ\nい　ン書を\n︒　ドき使\n　　をにい\n　　　　︑\n"
	want := buf.String()
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

type errReader struct {
}

func (r *errReader) Read(b []byte) (int, error) {
	return 0, io.ErrUnexpectedEOF
}

func TestFail(t *testing.T) {
	var buf bytes.Buffer
	err := tate(&buf, &errReader{}, option{})
	if err == nil {
		t.Fatal("should be error")
	}
}
