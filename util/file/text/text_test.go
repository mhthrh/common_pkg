package text_test

import (
	"github.com/mhthrh/common_pkg/pkg/model/test"
	"github.com/mhthrh/common_pkg/util/file/text"
	"testing"
)

type input struct {
	path       string
	name       string
	content    string
	isFullPath bool
}
type outPut struct {
	content string
	err     error
}

func TestText_Read(t *testing.T) {
	tst := test.Test{
		ID:   1,
		Name: "read file",
		Input: input{
			path:       "",
			name:       "file.txt",
			isFullPath: true,
		},
		Output: outPut{
			content: "file read test",
			err:     nil,
		},
	}
	txt := text.New(tst.Input.(input).path, tst.Input.(input).name, tst.Input.(input).isFullPath)
	byt, err := txt.Read()
	if err != nil {
		t.Errorf("read error-1")
	}
	if string(byt) != tst.Output.(outPut).content {
		t.Error("read error-2")
	}

}
func TestText_Write(t *testing.T) {
	tst := test.Test{
		ID:   1,
		Name: "write file",
		Input: input{
			path:       "",
			content:    "write test",
			name:       "newFile.txt",
			isFullPath: true,
		},
		Output: outPut{
			content: "write test",
			err:     nil,
		},
	}
	txt := text.New(tst.Input.(input).path, tst.Input.(input).name, tst.Input.(input).isFullPath)
	err := txt.Write([]byte(tst.Input.(input).content))
	if err != nil {
		t.Error("write error-1")
	}

	byt, err := txt.Read()
	if err != nil {
		t.Errorf("read error-1")
	}
	if string(byt) != tst.Output.(outPut).content {
		t.Error("read error-2")
	}

}
