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

func TestCreate(t *testing.T) {
	createdAt, _ := time.Parse("2006-01-02", "2018-04-01")
	updatedAt, _ := time.Parse("2006-01-02", "2018-04-02")

	// Create snippet
	snip := &snippet.Snippet{
		Hash:      "040b59d",
		Title:     "titlehoge",
		FileType:  "filetypehoge",
		Content:   "contenthoge",
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
	if err := snippet.Create(snip); err != nil {
		t.Fatal(err)
	}

	// Remove created snippet
	snippetPath := path.SnippetPath(snip.Hash)
	defer os.Remove(snippetPath)

	// Assert updated snippet
	got, err := snippet.Get(snip.Hash)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, snip) {
		t.Fatalf("Invalid return value. \ngot :%s\nwant%s", got, snip)
	}
}

func TestUpdate(t *testing.T) {
	createdAt, _ := time.Parse("2006-01-02", "2018-04-01")
	updatedAt, _ := time.Parse("2006-01-02", "2018-04-02")

	// Create snippet
	snip := &snippet.Snippet{
		Hash:          "040b59d",
		Title:         "titlehoge",
		FileType:      "filetypehoge",
		FileExtension: "fileextensionhoge",
		Content:       "contenthoge",
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
	}
	if err := snippet.Create(snip); err != nil {
		t.Fatal(err)
	}

	// Remove created snippet
	snippetPath := path.SnippetPath(snip.Hash)
	defer os.Remove(snippetPath)

	// Update snippet
	snip.Title = "titlechange"
	snip.Content = "contentchange"
	newUpdatedAt, _ := time.Parse("2006-01-02", "2018-04-03")
	snip.UpdatedAt = newUpdatedAt
	if err := snippet.Update(snip); err != nil {
		t.Fatal(err)
	}

	// Assert updated snippet
	got, err := snippet.Get(snip.Hash)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, snip) {
		t.Fatalf("Invalid return value. \ngot :%s\nwant%s", got, snip)
	}
}
