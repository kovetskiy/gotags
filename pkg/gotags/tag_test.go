package gotags

import (
	"testing"
)

func TestTagString(t *testing.T) {
	tag := NewTag("tagname", "filename", 2, 3, "x")
	tag.Fields["access"] = "public"
	tag.Fields["type"] = "struct"
	tag.Fields["signature"] = "()"
	tag.Fields["empty"] = ""

	expected := "tagname\tfilename\t2;\"\tx\taccess:public\tcolumn:3\tline:2\tsignature:()\ttype:struct"

	s := tag.String()
	if s != expected {
		t.Errorf("Tag.String()\n  is:%s\nwant:%s", s, expected)
	}
}
