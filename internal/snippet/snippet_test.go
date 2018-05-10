package snippet_test

import (
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/lighttiger2505/lipet/internal/path"
	"github.com/lighttiger2505/lipet/internal/snippet"
)

func TestCreateAndGet(t *testing.T) {
	createdAt, _ := time.Parse("2006-01-02", "2018-04-01")
	updatedAt, _ := time.Parse("2006-01-02", "2018-04-02")

	snip := &snippet.Snippet{
		ID:        "idhoge",
		Title:     "titlehoge",
		FileType:  "filetypehoge",
		Content:   "contenthoge",
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
	if err := snippet.Create(snip); err != nil {
		t.Fatal(err)
	}
	snippetPath := path.SnippetPath(snip.ID)
	defer os.Remove(snippetPath)

	got, err := snippet.Get(snip.ID)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, snip) {
		t.Fatalf("Invalid return value. \ngot :%s\nwant%s", got, snip)
	}
}
