package main

import (
	"bytes"
	"errors"
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

func TestTateHalfwidthVoicedKana(t *testing.T) {
	var buf bytes.Buffer
	err := tate(&buf, strings.NewReader("ｳﾞﾜﾞｦﾞ\n"), option{})
	if err != nil {
		t.Fatal(err)
	}
	got := buf.String()
	want := "ヴ\nヷ\nヺ\n"
	if got != want {
		t.Fatalf("want %v, but %v", want, got)
	}
}

func TestTateComposedVoicedKana(t *testing.T) {
	var buf bytes.Buffer
	err := tate(&buf, strings.NewReader("か\u3099\n"), option{})
	if err != nil {
		t.Fatal(err)
	}
	got := buf.String()
	want := "が\n"
	if got != want {
		t.Fatalf("want %v, but %v", want, got)
	}
}

func TestTateAsciiBars(t *testing.T) {
	var buf bytes.Buffer
	err := tate(&buf, strings.NewReader("-_\n"), option{})
	if err != nil {
		t.Fatal(err)
	}
	got := buf.String()
	want := "｜\n｜\n"
	if got != want {
		t.Fatalf("want %v, but %v", want, got)
	}
}

func TestTateAsciiSymbols(t *testing.T) {
	var buf bytes.Buffer
	err := tate(&buf, strings.NewReader("/~\n"), option{})
	if err != nil {
		t.Fatal(err)
	}
	got := buf.String()
	want := "＼\n ∫\n"
	if got != want {
		t.Fatalf("want %v, but %v", want, got)
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

type errWriter struct {
	err error
}

func (w *errWriter) Write(b []byte) (int, error) {
	return 0, w.err
}

func TestWriteFail(t *testing.T) {
	want := errors.New("write failed")
	err := tate(&errWriter{err: want}, strings.NewReader("x"), option{})
	if !errors.Is(err, want) {
		t.Fatalf("want %v, but %v", want, err)
	}
}

type shortWriter struct {
}

func (w *shortWriter) Write(b []byte) (int, error) {
	return len(b) - 1, nil
}

func TestShortWriteFail(t *testing.T) {
	err := tate(&shortWriter{}, strings.NewReader("x"), option{})
	if !errors.Is(err, io.ErrShortWrite) {
		t.Fatalf("want %v, but %v", io.ErrShortWrite, err)
	}
}
