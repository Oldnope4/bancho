// Code generated by go-bindata.
// sources:
// data/default.png
// DO NOT EDIT!

package avatarserver

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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _dataDefaultPng = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x54\x7b\x3c\x94\xf9\x1a\x7f\xc7\xbc\xf3\x1a\x99\xe6\xd2\xa0\xb9\xa4\x9d\xb5\x25\x2b\xe4\x12\x61\x56\xe6\x95\xeb\x96\x25\xcc\x56\x2b\x72\x8b\x88\x7c\x90\x6e\x12\x66\xac\xdc\xd2\x48\xa5\x53\x8d\x19\x0e\x1d\x8d\xec\x59\xb9\x15\x4b\xaa\x39\x92\x32\x86\xf4\xd1\xb8\x37\xc8\xd4\x8e\xb1\x24\x1a\x97\xe6\x8c\xce\x39\xfb\x39\xbb\xed\x3f\x7d\x9f\xcf\xfb\xcf\xf3\xbc\xbf\xef\xe7\xfb\x3c\xbf\xe7\xfb\x53\xbd\x50\x8d\x00\xf8\x6f\x5d\x3d\x5c\x01\x04\x02\x00\x10\xea\x00\x54\xfd\xc0\x0e\x00\x02\x41\x10\x05\x42\x28\x14\xa4\x05\xa9\x3f\x6d\x34\x04\xa1\xb5\x31\xab\x56\x69\xaf\x5a\x85\xc1\x63\x3e\x02\x8f\xc1\xe2\xf0\x38\x2c\x86\x48\xd4\xd1\x21\x12\xd7\x51\xa9\xd4\x75\x5f\xfc\x8e\x15\x12\x84\xa6\xa6\xe6\x6a\xcc\x6a\x12\x1e\x4f\xa2\xe8\x12\x74\x29\x5f\x7c\x2e\x54\x2d\x00\x01\x0d\xa4\x02\xa9\x48\x04\x01\xd0\x20\x20\x90\x04\x84\x4a\x08\x50\xd5\x22\x35\x90\x2b\x72\x7f\x07\x0a\x42\x20\x41\x0d\x4d\x75\xd1\x1e\x0f\x20\x90\x00\x88\x41\x6a\x22\x51\xff\xf9\x43\x43\x9d\x00\x51\x10\x61\x0d\x51\xef\x4b\x0b\x6b\xd8\x37\xe1\xc2\xed\xee\xdd\xc1\x71\x06\x96\x4e\x3e\x21\x69\xf7\xf3\x8b\xab\x74\x74\xbf\xda\x10\xbf\x72\x56\xe7\x0f\x9c\x1f\xb1\x92\x25\x7c\x9a\xed\x03\x30\x48\xb5\x08\x02\x92\x00\x38\x02\x8b\x95\xdc\xea\x1a\xe3\xf6\xf4\x04\x51\x7c\x56\xee\x4f\x98\x16\x69\x47\x02\xba\xb9\x6b\xcb\xdf\xc0\x1b\xce\xcf\xfd\x25\x4a\x76\x15\xd2\x20\xaa\x52\xbe\x74\xff\x45\x5a\xbe\xf5\xf7\x73\x67\x82\xe4\x8c\xb6\xab\x15\x49\x8e\xe6\xc7\x68\x61\x9e\xcc\x26\xde\x15\x3a\xde\xd2\x50\x30\xd4\xf9\x54\x3f\x93\x4c\x64\xc6\x46\xbc\xc2\x0a\x31\x59\x82\xfa\xea\xb7\x9a\x21\xd0\xe7\x12\x18\x4c\xea\x72\x8b\x99\x35\xb4\x22\xe4\xc5\xdc\xb2\xa0\x63\xdb\xd6\x1e\xbe\xee\x31\x87\x5e\x2a\xdf\x68\xf2\x9d\xa0\x30\x2f\x00\x60\xb5\xa4\xaa\xc1\xc6\x3a\xf5\x05\x8e\xad\x0f\x3c\x71\x4d\xca\x88\x94\x2a\x39\xa7\x42\xd7\x8a\xe5\x32\xcb\x95\x12\xeb\x5e\x4c\xad\xdf\xfe\x38\xd9\x56\x6d\xa2\xa0\x6c\xb4\x90\x41\xfb\x99\x46\xa3\xbd\xdd\xc0\x5f\x7d\xd2\xf1\xef\xb4\x71\xfb\xdd\x85\x8d\xf2\x7a\x05\xdb\xb7\xa0\xf9\xb8\x28\xed\x14\x4a\x46\x72\xb7\x41\x19\xc2\xfd\xe9\x43\x58\xb6\xc0\x8e\xe5\x49\xe6\xdd\x17\x27\x31\x18\xf0\x81\xba\x3e\x81\x99\xed\x2c\x05\x95\xf5\xdc\x67\xfa\x09\x1f\xbf\x17\x8f\x8f\x2a\xf5\xa1\x18\x36\x7a\x34\x67\xc0\x2d\x3e\x04\x87\xb1\xf9\x06\x68\x0d\x4a\x46\x16\x5b\xf1\x03\xd3\xd2\x53\x5b\xe5\xf3\xd5\x03\x32\xba\x51\xcf\x65\xc7\xa0\xa6\xe5\xf3\x52\x76\x31\x02\x75\xa6\x3f\xaa\xfc\xce\xa8\xf8\x90\xde\x26\xe6\xce\x47\x66\xa9\x70\x10\x2c\x8d\x74\x35\x1a\x18\x0e\xcf\x78\x0b\xeb\xc5\xe7\x17\xe4\x3b\x28\x64\x16\xe9\xa9\x42\xf6\x0d\xdb\x47\x88\x4c\xde\x77\xc2\xf5\xd7\x90\x41\x7c\x7a\x9e\x18\x1e\x1f\xd5\x88\x8d\xf0\x62\x61\xe3\x9b\xb8\x90\x39\x87\x34\x58\x34\xd7\xdd\xb6\x41\xf6\xac\xa6\x6a\x78\xda\x26\x24\x8c\xea\xbd\x67\xb2\x4e\x90\x93\xe7\x00\xa5\x85\x24\xfa\xfd\x12\x2f\x03\xf5\x40\xc2\xa5\xf4\xf7\xbe\x21\x92\x65\xce\xb3\x37\xec\x42\x2d\xa4\x16\x22\xc7\x1e\xcb\x29\xcb\x6c\x7d\x98\xa8\xe8\x0f\xd0\xbc\x50\x56\x98\x0d\x74\x67\x2f\x29\xd8\x91\xc4\xfd\xad\x76\xbf\x78\x9f\x2c\x1d\x90\xff\x4b\x2a\x11\xc7\xbf\xa4\x05\xd3\x06\x9d\x79\x6b\xcd\xef\x3d\xd8\x97\xa3\x60\x05\x7c\x22\x7a\xbb\xf3\x03\x33\xaf\xe2\xab\xf3\x8b\xb4\x7f\x5e\xcc\x73\x1e\x5b\xf0\x98\x4a\x69\xc0\x04\xd3\xfe\x17\x6d\x8d\x8d\xfc\xc9\x2e\x6d\x9b\xde\xa3\xeb\x23\x5d\x58\x1a\xf4\x7d\xd5\xcb\xe2\xb8\x97\xb4\xdf\x7c\x81\xd7\xa7\x3f\xa0\xad\xf5\x8d\x78\xe6\xaf\x62\xf5\xcb\x4d\x8e\x92\xab\x7a\xdc\x4b\x93\xcd\xec\x96\xfd\xad\xf8\x19\x7f\x1a\xee\xae\x67\x4f\xa0\xb3\x56\xcb\x5b\x16\x5f\xc1\x01\x0f\x19\xff\x0d\x58\x4e\x7e\xac\x08\xd3\xaf\x62\x74\xd5\xe9\x5c\x3d\x86\xed\xa1\xa2\x85\xfe\xd4\xcd\x03\xa6\x8a\x6c\x74\x03\x33\x04\x3c\x32\x39\xe9\x28\x67\x73\x6e\x86\xfb\x1f\x89\x85\xe0\x51\x4a\xda\x39\xf8\x60\xf5\xc4\xd9\x3d\xf5\xdd\x61\xc3\xbc\x76\x2c\x35\x46\x50\x32\x7a\x1e\xe6\xef\xe5\x32\xb7\xac\x6b\x40\x80\x4d\x61\xa2\x59\xb2\x42\x19\xa0\x02\x76\xc9\x8c\xdf\x0b\x31\x21\x90\x05\x92\xc4\xc6\x6d\xb6\x29\x0c\x8f\xdd\xba\x46\x87\xe2\x1c\xaa\x73\x61\x5b\x8d\x53\x47\xc2\xfb\x40\xf6\xa6\x5d\x27\xa4\xad\x05\x81\x2e\x8a\x21\xd6\xfa\x4b\x7b\x3d\x2c\xe8\x77\x03\xd9\x3c\x0a\x9b\xd4\x83\x89\x82\xac\x51\x75\x4f\xab\x55\x40\x74\x48\x4a\xe6\xe3\xfe\xf7\x4b\xb4\xb1\x3f\xb5\xe2\xe5\x8e\x89\xe3\x6c\x67\x45\xac\x5d\xac\x6d\x33\x9a\xa1\x0a\xc3\xa0\x79\x3c\x3d\x67\x58\x2a\xe1\x44\x19\x47\x9b\x4e\x1f\x78\xcd\x85\x4a\x6e\xf9\xbf\xa4\x2e\x7b\x1f\x36\x3e\xda\x59\xfc\x5b\x32\x3e\xea\x24\x77\x68\xfe\x78\x8a\x1f\xd1\xda\xd3\x26\xd7\xe2\xf1\xba\x31\xa9\xb0\x83\x25\xef\x9d\x88\x80\xb3\xfb\xb8\xae\x46\xa3\x79\x1a\xf1\xc1\x11\x0f\x8e\x1f\xae\x6e\xe4\x9f\x6d\x55\xf2\x43\x4f\x8b\x2d\x9a\x0f\x91\x86\xb7\x92\x5a\xf9\x0b\xba\x5e\x6e\xa7\x1c\xf6\xb9\x65\x2d\xbf\x82\xe2\xdf\x99\x4c\x2d\xf1\x73\x3e\x59\xdb\xdd\x31\xe7\xc0\xf2\xda\xf3\xef\x0c\x86\x5f\x96\x4e\xdc\x0b\xa5\x26\x75\x28\x61\xc4\x39\x46\x2b\x3f\xd4\xfb\x60\x99\x88\xf5\x43\x45\x85\xe8\x12\xb9\xc8\x9c\x11\xda\x30\xbf\x79\x6e\x81\xd6\x11\xb1\x47\x9f\x42\x2b\x20\x19\x6e\x4a\x16\x6d\x93\xbf\x83\xf1\x33\x13\x70\xae\xba\x6d\x24\x85\xf5\x6d\xa7\xb8\x88\xf7\x83\x8b\x6e\xaf\x10\x97\x7b\xe9\x43\xbe\x64\xfe\x74\xdb\xf6\xbc\xbf\x70\x9b\xf4\xa7\xa8\xb6\x0a\x63\xf0\x62\xe8\x53\xaf\x22\x2b\xa5\x67\x8d\x3d\xee\x72\x77\x49\x51\x8e\xfb\x85\xf3\xb6\x77\x76\x52\x47\x38\x0d\xab\xf7\xaa\x27\x89\xdc\xc9\x45\x65\x94\xdd\x92\xc5\x75\x8d\x67\x70\xc6\x5e\x9c\xb0\x35\x33\x3d\x93\xb3\x24\x0e\xbf\x7b\x4b\xe6\xe6\x92\x4d\xa1\x17\x8c\x04\xdb\xd9\x59\xf7\x1d\xe9\xe2\xa2\x12\xc7\x45\xb2\x95\xe1\x63\xb2\x08\x05\xe8\xea\x7c\xe2\x80\x81\x64\x70\x50\x9c\x6c\xf6\xd4\x8f\xd7\x6f\x64\x7a\x2e\x76\xee\xc6\x9d\x1b\x29\x38\xc4\x22\xed\xca\x5f\xd9\x96\xd1\x77\xcd\x42\x3f\x06\x6b\xaf\xcd\xe2\x51\x93\x40\xe5\x95\xfe\xf1\x2d\x0f\xf9\x4a\xae\xa1\xb1\x8b\x5e\x02\x8b\x83\xd3\x69\x99\xc7\xeb\x45\x84\x51\xaa\xca\x49\x3d\xe4\xc8\xc7\x35\xf6\x48\x91\x57\x99\x38\x74\x65\x9f\xf2\x0d\x2c\x83\xa0\x78\xcb\xa4\x74\x7f\x7a\xe3\x3d\x16\xae\x78\x1f\xd5\x70\xf6\x1b\x4a\xf2\x37\x29\xfd\xa7\xa7\x9d\x95\xd3\xd1\x8b\x07\xde\x7c\xc0\xfb\xe3\xf1\x9f\xba\x11\x54\xfb\x71\x90\x5b\x71\x38\xf3\xe7\xc8\x2f\x8f\x96\x0c\x8f\x5e\xb7\x22\x2d\xcf\xd5\x17\xf0\x95\x93\x09\xc1\xbf\xb6\x37\xb5\x8d\xe0\x7e\x35\x9a\xe4\xe6\x19\xce\x49\x9a\xf9\xcf\x16\x68\x95\xf9\x1c\x62\x56\x75\x96\x4d\xfe\xb1\xbb\xfd\x72\x99\xc5\xc7\xdb\x49\x18\xac\x9d\xb5\xd4\x6a\x2f\x69\xeb\xad\xbd\x96\x05\xe2\xdc\xd2\x66\x3b\xad\x6c\x6b\x87\x3c\xa7\xec\x71\xc5\x7e\x7a\xae\x95\x02\x72\x69\xf8\xa3\x6e\x45\xa7\xbb\x8d\x68\x78\x66\x78\x95\x96\x10\x04\x75\x17\xff\xb8\xf4\xa8\x34\x84\x1a\x9a\x1b\x6d\x6e\x56\x7c\xdd\x91\x68\x96\xa3\xf3\x21\xdc\x64\x6a\xe1\x35\x4c\xb0\x4e\xbd\x7d\xdb\x61\xd2\xf9\x84\x0a\x90\xee\x88\x6d\xc7\xb8\xa9\x80\x22\x49\x50\xd2\x50\xb6\xdb\x98\x0f\xb5\x7e\x86\xe5\xf3\xa3\x77\x4f\x0d\xb5\x28\xcb\x29\xdb\x20\x06\x7a\x20\xda\xdf\x99\xc4\x9f\x62\x30\xfe\xef\x51\x44\x37\xd2\x0f\x95\x33\x63\xd0\x23\xd7\x37\x7e\xf5\xf5\x7e\xc9\xfd\x1f\xfb\x26\x08\x31\xa5\x8e\xc6\x5a\xbb\xea\x04\xf5\xbe\x33\x67\x4f\x41\xfe\x4c\xff\x9b\xca\x1d\x4d\x48\x53\x41\x65\xcf\x52\xf1\x50\xda\xe5\xe7\xdf\xbf\x39\xe3\x3d\xc5\x78\x12\x55\x91\xe8\x68\x74\x9c\x7f\xb0\x96\xd9\xc0\xbb\x48\x17\x5b\xf6\x09\x06\xbd\xda\xf5\x65\xe4\x18\xe6\x91\x7f\x8c\x63\x67\x31\x9f\x4b\xe0\xd2\xb7\xd9\xc5\xe9\x56\x39\x0c\x22\x8c\xae\xb4\xb7\x2a\xa2\x4d\x7a\xcd\x0b\xa2\x0d\x94\x9c\x14\x95\xe4\xdf\x01\x00\x00\xff\xff\x0c\xb6\x31\xfc\x19\x08\x00\x00")

func dataDefaultPngBytes() ([]byte, error) {
	return bindataRead(
		_dataDefaultPng,
		"data/default.png",
	)
}

func dataDefaultPng() (*asset, error) {
	bytes, err := dataDefaultPngBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/default.png", size: 2073, mode: os.FileMode(420), modTime: time.Unix(1395500043, 0)}
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
	"data/default.png": dataDefaultPng,
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
	"data": &bintree{nil, map[string]*bintree{
		"default.png": &bintree{dataDefaultPng, map[string]*bintree{}},
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
