// Code generated by go-bindata.
// sources:
// web/index.html
// DO NOT EDIT!

package main

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

var _webIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x56\x6d\x6f\xdb\x46\x0c\xfe\xee\x5f\xc1\x68\x1b\x22\x23\xf5\x49\x4e\x51\x2c\xb0\x25\x03\x5d\x1b\x6c\x19\xb2\xb6\xc0\x82\x01\xc3\xb0\x0f\x67\x89\x96\xe8\x9e\x74\xda\x1d\xe5\xda\x18\xfc\xdf\x0b\x59\x96\xad\xc4\x6f\x0a\xbf\xc8\x3a\x93\x7c\xc8\x87\x2f\xba\xe0\xea\xe3\xe7\x0f\x4f\x7f\x7f\xb9\x87\x94\x33\x35\xe9\x05\xf5\xa3\x17\xa4\x28\xe3\x49\x0f\xb6\x12\x64\xc8\x12\xa2\x54\x1a\x8b\x1c\x3a\x25\xcf\x06\x77\x4e\xeb\x6f\x26\x56\x38\xf9\x4b\x1b\x46\x52\xf0\x11\x33\x1d\x78\xf5\xd9\x5e\xc7\x46\x86\x0a\x06\x6b\xa2\xd0\x49\x99\x0b\x3b\xf2\x3c\x39\x97\x4b\x91\x68\x9d\x28\x94\x05\x59\x11\xe9\x6c\x73\xe6\x29\x9a\x5a\x6f\xfe\x5f\x89\x66\xe5\xbd\x15\xb7\x62\xb8\x7d\x11\x19\xe5\x62\x6e\x9d\x49\xe0\xd5\xfe\x26\xbd\xc0\xab\x63\xed\x05\x53\x1d\xaf\x5a\x80\x31\x2d\x20\x52\xd2\xda\xd0\x71\xc0\xf2\x4a\x61\xe8\x14\x32\x8e\x29\x4f\x46\xe0\x17\x4b\xb8\xf5\x7f\x6a\x9e\xe3\x56\x36\x3b\x07\xe9\xb0\x31\x63\x5c\xf2\x40\x2a\x4a\xf2\x51\x84\x39\xa3\x19\x3b\xcf\xb2\x85\xf7\x45\x11\x78\xe9\xf0\x88\x93\xe2\x9c\x8f\x03\xf5\x46\x9e\x52\xb2\x20\x8b\x02\x58\x7e\x45\x0b\x9c\x22\x14\x46\x2f\x28\xc6\x18\x28\x93\x09\x5a\x90\x79\x5c\x9d\x45\x68\x6d\xad\x91\x01\xe5\xc0\x1a\xe4\xd4\xb2\x91\x11\xc3\x02\x23\xd6\x66\xab\x2f\xe0\x24\xd8\x97\xda\x0b\xe5\x09\xe4\x65\x86\x46\x97\xb6\x36\x82\x19\x29\xb4\x90\xc9\xd5\x26\x0e\xb0\x3a\x43\x60\xca\x50\x1c\xe6\xe9\x15\x47\x92\x2f\x15\x50\x1c\x3a\x94\x25\x03\x45\x96\x77\x65\xc8\xe4\x72\x90\x22\x25\x29\x8f\x6e\xdf\xf9\xc5\x72\x0c\x7a\x81\x66\xa6\xf4\xb7\xc1\x6a\x24\x4b\xd6\x15\x37\x47\x30\x4a\x75\xac\x4c\xa6\x55\x74\x2f\xa6\xc5\x8b\x1e\xa8\x22\x60\x6c\xa1\x77\xaa\x44\x40\x59\xd2\x04\xef\x80\x54\xdc\xea\xa2\x56\xf8\xf0\xce\x5f\xa4\x63\xc7\x3b\x1d\xc2\xb6\xef\x79\x55\x6c\xa1\xbd\xb9\x5c\xc8\xfa\xf4\x08\xee\x8f\x6e\xac\xa3\x32\xc3\x9c\xfb\xc2\xa0\x8c\x57\xee\xac\xcc\x23\x26\x9d\xbb\x7d\xf8\xff\x90\x94\x46\x1a\x2d\x30\x38\x33\x68\xd3\x47\xb2\xbc\xb1\x38\x65\xb0\x03\x14\xd5\xc4\xb9\x97\x15\x1b\xa9\x32\x19\xc1\xf5\xaf\xf7\x4f\xd7\x6f\x3a\x1b\x95\x46\x8d\xe0\xba\x9a\xfb\x91\xe7\x0d\x6f\x7f\x16\xbe\xf0\xc5\x70\x74\xe7\xdf\xf9\xde\xf5\x45\x2f\xeb\xbe\x88\x75\x8e\x7b\x2a\x62\xc9\xb2\xdf\x3d\xe4\xa2\x5a\x5c\x0f\x55\x4b\x6f\x68\xf9\xfd\xcf\xcf\x9f\xc4\xe6\xac\x76\xd4\x1f\x77\x88\xe0\xa4\xca\xfa\x74\x51\x9e\xd5\xe2\x34\x88\x45\x7e\xa8\xfa\x70\x21\x95\xdb\x32\x79\x03\x6f\x7d\xdf\x3f\x62\xb7\xee\x1f\x42\xee\xea\xff\x22\xd7\x4d\x82\xe7\x1a\x47\x21\x43\xa9\x20\x84\xa6\xef\x44\x82\x7c\xaf\xb0\xfa\xf9\xcb\xea\x21\x76\xf7\xe3\x7b\x9a\x83\x6f\x29\x29\x74\x4b\x25\x52\x69\x3f\xa4\xa4\xe2\x4f\x3a\x46\xeb\xf6\xbb\x34\x60\xa9\x84\xc1\x4c\x2f\x70\x63\x58\x39\x51\xd2\xf2\xe6\xe5\x0c\x67\x67\x58\x9f\x69\x03\x2e\x41\x08\xfe\x18\x08\x02\xa8\x28\x10\x0a\xf3\x84\xd3\x31\xd0\xcd\x4d\x87\xc6\xa9\x48\xd1\xd3\x79\xc5\x8a\x64\xf9\x0f\xfd\xdb\xc9\x42\x51\x9b\xc6\xc8\xa0\x64\xdc\x32\xe9\x3a\x8f\x0f\x67\xf8\x6b\x84\x66\xe0\xea\xe9\xbc\x9e\x7c\xb8\x0a\x81\x4d\x89\x5d\x48\xdc\x85\x41\x82\xf2\x1c\xcd\x6f\x4f\x7f\x3c\x42\x58\xe5\x20\x72\x99\x21\xdc\x80\x03\x6e\xb1\x5b\xf4\x42\x88\xbe\x73\xb9\xe9\x01\x95\xc5\xd7\xa0\x23\x83\x3c\xc3\xc1\xfb\x0e\x14\x34\x22\x45\x6a\x70\x16\x3a\x3f\x5c\x5d\x0e\x74\x6f\x73\x2c\xf9\xd7\x90\xa7\xf3\x48\x51\xf4\x15\x42\x78\xb6\x76\xbb\x7a\x68\x24\xd2\xb9\xd5\x0a\x85\xd2\x89\x5b\x7f\x7c\x3a\xec\x98\x97\x72\x6e\x20\x9d\xbe\xb0\x26\x82\x10\x9c\xe3\x3b\xd5\x81\x9b\x5d\xfe\xdd\x91\xd7\xdd\x55\x15\x09\x59\x14\x98\xc7\xf5\xd4\xca\xcb\x85\x5d\x77\xd9\x04\x6d\x9f\x8a\xce\x2d\xdd\xc3\x93\xd6\x27\x78\x7f\x39\xac\x2f\x85\xd5\x2d\x71\x73\xb3\xfd\x1e\x00\x00\xff\xff\x21\xdb\x42\x0b\xf1\x0a\x00\x00")

func webIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_webIndexHtml,
		"web/index.html",
	)
}

func webIndexHtml() (*asset, error) {
	bytes, err := webIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/index.html", size: 2801, mode: os.FileMode(436), modTime: time.Unix(1513050686, 0)}
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
	"web/index.html": webIndexHtml,
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
	"web": &bintree{nil, map[string]*bintree{
		"index.html": &bintree{webIndexHtml, map[string]*bintree{}},
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

