package snippet_test

import (
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/lighttiger2505/lipet/internal/path"
	"github.com/lighttiger2505/lipet/internal/snippet"
)

type ValidateSnippetHashTest struct {
	Hash string
	Want bool
}

var ValidateSnippetHashTests = []*ValidateSnippetHashTest{
	&ValidateSnippetHashTest{
		Hash: "9bb9339",
		Want: true,
	},
	&ValidateSnippetHashTest{
		Hash: "9bb9339d32344e3147c3792024a1530e8de10494",
		Want: true,
	},
}

func TestValidateSnippetHash(t *testing.T) {
	for i, test := range ValidateSnippetHashTests {
		got, err := snippet.ValidateSnippetHash(test.Hash)
		if err != nil {
			t.Fatalf("%d: Don't want error. %s", i, err)
		}
		want := test.Want
		if got != want {
			t.Fatalf("%d: Invalid return value \ngot: %v\nwant:%v", i, got, want)
		}
	}
}

func TestCreateAndGet(t *testing.T) {
	createdAt, _ := time.Parse("2006-01-02", "2018-04-01")
	updatedAt, _ := time.Parse("2006-01-02", "2018-04-02")

	snip := &snippet.Snippet{
		Hash:      "idhoge",
		Title:     "titlehoge",
		FileType:  "filetypehoge",
		Content:   "contenthoge",
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
	if err := snippet.Create(snip); err != nil {
		t.Fatal(err)
	}
	snippetPath := path.SnippetPath(snip.Hash)
	defer os.Remove(snippetPath)

	got, err := snippet.GetFull(snip.Hash)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, snip) {
		t.Fatalf("Invalid return value. \ngot :%s\nwant%s", got, snip)
	}
}
