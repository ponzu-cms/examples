package content

import (
	"fmt"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"

	"github.com/bosssauce/reference"
)

type Article struct {
	item.Item

	Title    string   `json:"title"`
	Photo    string   `json:"photo"`
	Body     string   `json:"body"`
	Category []string `json:"category"`
	Author   string   `json:"author"`
}

// MarshalEditor writes a buffer of html to edit a Article within the CMS
// and implements editor.Editable
func (a *Article) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(a,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Article field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Title", a, map[string]string{
				"label":       "Title",
				"type":        "text",
				"placeholder": "Enter the Title here",
			}),
		},
		editor.Field{
			View: editor.File("Photo", a, map[string]string{
				"label":       "Photo",
				"placeholder": "Upload the Photo here",
			}),
		},
		editor.Field{
			View: editor.Richtext("Body", a, map[string]string{
				"label":       "Body",
				"placeholder": "Enter the Body here",
			}),
		},
		editor.Field{
			View: editor.Tags("Category", a, map[string]string{
				"label":       "Category",
				"placeholder": "+Category",
			}),
		},
		editor.Field{
			View: reference.Select("Author", a, map[string]string{
				"label": "Author",
			}, "Author", `{{.name}}`),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Article editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Article"] = func() interface{} { return new(Article) }
}

func (a *Article) String() string { return a.Title }

func (a *Article) IndexContent() bool { return true }

func (a *Article) Push() []string {
	return []string{
		"author",
		"photo",
	}
}
