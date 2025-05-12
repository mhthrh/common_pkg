package text

import (
	env "github.com/mhthrh/common_pkg/util/environment"
	"github.com/mhthrh/common_pkg/util/file"
	"os"
	"path/filepath"
)

var (
	appPath = ""
)

func init() {
	appPath = env.GetAppPath()
}

type File struct {
	path string
	name string
}

func New(path, name string) file.IFile {
	return &File{
		path: path,
		name: name,
	}
}
func (f *File) Read() ([]byte, error) {
	data, e := os.ReadFile(filepath.Join(appPath, f.path, f.name))
	if e != nil {
		return nil, e
	}
	return data, nil
}

func (f *File) Write(bytes []byte) error {
	//TODO implement me
	panic("implement me")
}
