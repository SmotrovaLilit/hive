// Code generated for package driver by go-bindata DO NOT EDIT. (@generated)
// sources:
// ../contrib/sql/migrations/postgres/1.sql
// ../contrib/sql/migrations/tests/1_test.sql
package driver

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _ContribSqlMigrationsPostgres1Sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x56\x5b\x6f\x9b\x4c\x10\x7d\xe7\x57\xcc\x1b\xa0\x2f\x91\x3e\xb5\x4a\x1f\x12\xa9\x12\x81\x75\x83\x82\xb1\xcb\xa5\x49\xfa\x82\xb6\xb0\xb6\xb7\xb6\x81\xee\xae\x73\x51\xd5\xff\x5e\x71\x33\x18\x03\xc6\xa9\xdf\xb0\xcf\x39\x9e\x39\x73\x66\xec\xcb\x4b\xf8\x6f\x4b\x97\x0c\x0b\x02\x7e\x2a\xe9\x0e\xd2\x3c\x04\xde\xd3\x1c\x41\xc8\x48\x44\x62\x41\xf1\x86\x07\xe2\x2d\x25\xa0\xb9\x80\x6c\x7f\x0a\x8a\x9c\xd0\x28\x94\x2f\x40\x4e\x31\xe7\x2f\x09\x8b\x64\xf5\xe6\x80\xca\xc9\x66\x11\x70\xc2\x9e\x69\x48\x02\x46\x7e\xed\x08\x17\x6d\x8d\x4d\xb2\xa4\x71\x26\xc2\xc8\x92\x72\xc1\xb0\xa0\x49\x9c\x09\xed\x95\xb4\x5b\x0b\x81\x39\x01\x7b\xe6\x01\x7a\x34\x5d\xcf\x05\x9a\x57\x24\xde\x24\x45\x02\x00\x48\xd7\xd0\x7a\xdd\x9a\x5f\x5c\xe4\x98\x9a\x05\x73\xc7\x9c\x6a\xce\x13\xdc\xa3\xa7\x8b\x1c\x4c\xa3\x36\xf8\x9b\xe6\xe8\x77\x9a\xa3\x7c\xb8\xba\x52\xf3\x6f\xb1\x7d\xcb\x02\xdf\x36\xbf\xfa\xa8\xe0\x08\x86\xa9\xe0\x4d\xce\x4f\x9e\xc4\x3f\xaa\x87\x3d\xc7\x40\x13\xcd\xb7\x3c\x90\x7f\xff\x91\xaf\xaf\x73\x4c\x53\x20\xe0\xe1\x8a\x6c\x71\xb0\x63\x1b\x10\xe4\x55\x40\x5b\x40\x1a\xd9\x77\x50\x0f\xa5\xc3\x82\xc1\xe6\x4b\x81\x74\x9d\xc1\x4c\xdb\x03\x07\x4d\x90\x83\x6c\x1d\xd5\xf2\xa0\xa4\x6b\x15\x66\x36\x18\xc8\x42\x1e\x02\x5d\x73\x75\xcd\x28\xbd\xd8\x12\xb1\x4a\x4a\x0f\x8f\xb2\x51\x35\x52\x40\xc3\x24\x5e\xd0\x65\x87\x61\xa7\x4c\x7b\x87\x0f\x41\xf1\xde\x82\x12\xd6\x9b\x8a\xd1\xee\x34\x64\x07\x8d\x6a\x00\x07\x3d\xab\x6b\x3b\x33\x78\xfa\x1d\xd2\xef\x41\xd9\x90\x78\x29\x56\x4a\x2d\xa3\xc2\x67\xf8\x5f\x3d\x65\x53\xd7\xfe\xf5\x98\x73\xc6\xc2\x54\x55\x7f\xfc\xa4\xb6\xac\xed\x6c\x81\xbc\xa6\x94\x11\x1e\xe0\x32\xee\x9e\x39\x45\xae\xa7\x4d\xe7\xf0\x60\x7a\x77\xf9\x23\x7c\x9f\xd9\x08\x5a\xe1\xa1\x9c\xef\x48\xb4\xa7\x8d\xe0\xed\x93\x64\xcf\x1e\x14\xb5\x50\xa9\xae\x4e\xb6\x72\xf9\x22\x36\xd7\xae\xb3\xf8\x43\xde\x8a\xe0\x88\x30\x7e\x1c\xdf\x51\x39\x2e\xb4\x70\x28\xe8\x33\xa9\x19\x47\x5b\xd3\x53\x43\xb1\x69\xf5\xd9\xf9\x97\x1a\xd6\x34\x3e\x18\x63\xff\x69\x1e\x7b\x8b\x38\xe1\x9c\x26\x71\xdf\xb2\x0d\x04\x8a\x8f\x4e\xd4\xa8\x40\xf5\x27\x63\x38\x50\xa7\x79\x5d\x81\xc2\x3b\xb1\xca\x86\x17\x62\x51\x88\xbd\x4b\xa5\x79\x87\x2b\xbb\xce\x3b\xc6\xcd\xf1\x14\xee\x80\x69\x1b\xe8\x11\x62\xbc\x25\x19\xbc\x6b\xc2\xa0\xd0\xe8\x22\xcf\x42\xc6\x6e\xfe\xe8\x1b\xc9\x4b\x2c\x19\xce\x6c\x5e\xce\x7a\xf8\xc2\xde\x9c\x80\x1e\x7c\x5e\x06\xa5\x93\xd3\x02\x1e\x57\x5c\x01\xb2\xbf\x13\xe6\xa4\xca\x5e\x7b\x83\xba\x51\xbd\x19\xbf\x91\xfe\x06\x00\x00\xff\xff\xc4\xae\x43\x62\xf0\x08\x00\x00")

func ContribSqlMigrationsPostgres1SqlBytes() ([]byte, error) {
	return bindataRead(
		_ContribSqlMigrationsPostgres1Sql,
		"../contrib/sql/migrations/postgres/1.sql",
	)
}

func ContribSqlMigrationsPostgres1Sql() (*asset, error) {
	bytes, err := ContribSqlMigrationsPostgres1SqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../contrib/sql/migrations/postgres/1.sql", size: 2288, mode: os.FileMode(420), modTime: time.Unix(1565962017, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _ContribSqlMigrationsTests1_testSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x52\xc1\x4e\xc2\x40\x10\xbd\xf7\x2b\x26\xbd\xb4\xc4\x05\xc2\x15\x2f\x9a\xc8\x81\xc4\x40\x22\xa0\xc7\x66\xdd\x9d\xb6\x13\xda\x6e\xdd\xd9\x52\x88\xe1\xdf\x4d\x4b\x95\xa2\xc4\x78\xf2\x34\xd3\xd7\xc9\x7b\x33\xfb\xde\x70\x08\x37\x39\x25\x56\x3a\x84\x4d\xe9\xcd\x17\xab\xd9\xd3\x1a\xe6\x8b\xf5\x12\x48\x63\xe1\xc8\x1d\x20\x24\x2d\xc0\x59\x49\x8e\x23\x56\x29\xe6\x32\xaa\x6c\x36\xf0\x9e\xef\x1f\x37\xb3\x15\x84\x81\x96\x4e\x0e\x27\x81\x80\x20\x36\x26\x18\xdc\x7a\x57\x79\x22\x65\xb1\x6d\x65\xd6\x50\x76\x60\xb9\x15\x90\xa3\x4b\x8d\x16\xa0\x4c\x11\x53\x72\x26\x9e\x08\x08\x4a\xc9\x5c\x1b\xab\x1b\xf6\x77\x3f\x36\xc6\x9f\xfa\xaf\xd2\xfa\xc7\xbf\xe8\x44\x27\x2c\x26\xb4\x3d\xc9\xde\xff\x46\xfd\x3c\x73\xa9\x7c\xba\xea\x0e\xf7\x32\x2f\x33\x1c\x19\x9b\xfc\xb7\xa4\x32\xf9\x0f\x49\xc6\x2c\x8e\x18\xed\x8e\x14\x46\x16\xdf\x2a\x64\x77\x72\x08\xf7\x25\x59\xe4\x48\x3a\x01\xc4\x5c\xa1\x6e\xdb\x6e\xa6\xb1\xec\xfc\x91\xa2\xd4\x68\x59\x80\x54\x8e\x76\xf8\x69\x01\x0b\xd8\x52\xa1\x2f\x96\x5a\x2c\x5f\xc2\xc1\x57\x09\x52\xe7\x4a\x9e\x8e\xc7\x75\x5d\x8f\x8c\x3d\x8c\x38\x1d\xb7\xd6\x1c\x83\xef\x66\xb5\x48\x66\x12\x2a\xae\x1c\xc1\x4c\xa6\x80\x90\x7f\x59\x5c\x56\x2e\x6d\x9e\x49\x49\xd7\x21\xbd\xd0\xf4\xd2\xc7\xa4\xdb\xf0\x5d\x2e\xda\x95\x49\x23\xdd\x0f\xf9\x83\xa9\x0b\xef\x23\x00\x00\xff\xff\x3c\xee\x8d\x2f\xf6\x02\x00\x00")

func ContribSqlMigrationsTests1_testSqlBytes() ([]byte, error) {
	return bindataRead(
		_ContribSqlMigrationsTests1_testSql,
		"../contrib/sql/migrations/tests/1_test.sql",
	)
}

func ContribSqlMigrationsTests1_testSql() (*asset, error) {
	bytes, err := ContribSqlMigrationsTests1_testSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../contrib/sql/migrations/tests/1_test.sql", size: 758, mode: os.FileMode(420), modTime: time.Unix(1565962017, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"../contrib/sql/migrations/postgres/1.sql":   ContribSqlMigrationsPostgres1Sql,
	"../contrib/sql/migrations/tests/1_test.sql": ContribSqlMigrationsTests1_testSql,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"..": &bintree{nil, map[string]*bintree{
		"contrib": &bintree{nil, map[string]*bintree{
			"sql": &bintree{nil, map[string]*bintree{
				"migrations": &bintree{nil, map[string]*bintree{
					"postgres": &bintree{nil, map[string]*bintree{
						"1.sql": &bintree{ContribSqlMigrationsPostgres1Sql, map[string]*bintree{}},
					}},
					"tests": &bintree{nil, map[string]*bintree{
						"1_test.sql": &bintree{ContribSqlMigrationsTests1_testSql, map[string]*bintree{}},
					}},
				}},
			}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
