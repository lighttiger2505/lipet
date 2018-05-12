package snippet_test

import (
	"testing"

	"github.com/lighttiger2505/lipet/internal/snippet"
)

type GetFileExtensionTest struct {
	Arg  string
	Want string
}

var GetFileExtensionTests = []*GetFileExtensionTest{
	&GetFileExtensionTest{Arg: "Vim", Want: "vim"},
	&GetFileExtensionTest{Arg: "vim", Want: "vim"},
	&GetFileExtensionTest{Arg: "vimscript", Want: "vim"},
	&GetFileExtensionTest{Arg: "Vim script", Want: "vim"},
	&GetFileExtensionTest{Arg: "hoge", Want: "hoge"},
}

func TestGetFileExtension(t *testing.T) {
	for i, test := range GetFileExtensionTests {
		got := snippet.GetFileExtension(test.Arg)
		want := test.Want
		if got != want {
			t.Fatalf("%d: Invalid return value.\ngot :%s\nwant:%s", i, got, want)
		}
	}
}

type GetFiletypeTest struct {
	Arg  string
	Want string
}

var GetFiletypeTests = []*GetFiletypeTest{
	&GetFiletypeTest{Arg: "Vim", Want: "Vimscript"},
	&GetFiletypeTest{Arg: "vim", Want: "Vimscript"},
	&GetFiletypeTest{Arg: "vimscript", Want: "Vimscript"},
	&GetFiletypeTest{Arg: "Vim script", Want: "Vimscript"},
	&GetFiletypeTest{Arg: "hoge", Want: ""},
}

func TestGetFiletype(t *testing.T) {
	for i, test := range GetFiletypeTests {
		got := snippet.GetFiletype(test.Arg)
		want := test.Want
		if got != want {
			t.Fatalf("%d: Invalid return value.\ngot :%s\nwant:%s", i, got, want)
		}
	}
}
