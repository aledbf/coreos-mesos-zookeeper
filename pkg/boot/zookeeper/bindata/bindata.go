package bindata

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _pkg_boot_zookeeper_bash_add_node_bash = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x52\xdf\x6f\xda\x30\x10\x7e\xcf\x5f\x71\x72\xa3\x16\x24\x82\x45\xf7\x82\x60\x3c\xb0\x0e\xad\xd2\x26\x6d\x2a\x74\x93\x56\xb5\x95\x49\x2e\xe0\x35\xd8\xae\xed\xa0\xd2\x6d\xff\xfb\xce\x49\xe8\x0f\xc6\xa6\xf6\x25\x8e\x7d\x77\xdf\xdd\xf7\x7d\xe7\xd0\x43\x82\x1a\x8c\x34\x98\x0b\x59\x44\xd1\x01\x84\xb7\x0c\xe7\xe5\x02\xe6\xc2\x61\x06\x5a\x01\xaa\xf5\x5a\xd8\xe8\xe2\x02\xe2\xf7\x93\x77\xe7\x1f\xe0\xf2\x12\x0e\x0f\xab\xcc\xe4\x2e\x8a\x56\x42\xaa\x56\x1b\x7e\x46\x00\x78\x67\xb4\xf5\xf0\x65\x3c\x3b\x1d\xc5\xe1\x3b\xe0\x3f\x2c\xf2\xb9\x54\x11\x45\x53\x03\x5c\x18\xc3\xef\xb5\xee\xa6\xf9\x02\xb8\x36\x3e\x5c\x6e\x10\x0d\xda\x24\x13\x5e\x6c\x63\x94\x5d\x28\x48\xdc\xff\x72\x76\x62\x3c\xd5\x2a\x7f\xa8\x27\x00\xaf\xcb\x74\xf9\x2f\x80\x6b\x4a\xea\x66\x1b\x25\x56\x32\x0d\xc9\x32\x07\xe2\x47\x0d\xd9\x0b\x0a\x58\x50\x60\x08\x7e\x89\x8a\x4a\x89\x76\xba\xd4\x90\x28\x60\xa2\xb0\x28\xb2\x0d\xe9\x20\x9d\x77\xf0\x80\x02\x69\x51\x3a\x4f\xa7\x54\xb9\xb6\x2b\xe1\xa5\x56\x2c\xe8\x55\x38\xac\x10\x82\x34\x7b\xd8\xe4\x05\xa2\x4f\x76\xba\xbf\x88\x12\x40\x2e\x03\xaf\x03\xf8\x86\xa0\x90\x8c\xf4\x1a\x44\x46\xc7\x52\x3a\x50\x3a\xc3\xf0\x40\x0c\x1e\x47\xcb\xa1\x8a\x78\x08\xad\xe5\xa2\xb4\x54\x24\xd5\xd3\x1c\xc2\xfb\xfe\xf1\xf4\xf3\x74\x36\x8a\x5b\x61\x37\x12\x04\xe6\x78\x1c\x5e\x06\x6f\xfa\xfd\x3e\xe7\x43\xc7\xaf\x3a\xd5\xd1\x89\x39\x67\xf0\x76\x2f\x2d\x87\x76\x8d\xb6\x5b\x90\x48\xf0\x0b\xd2\x92\xf6\x28\x83\xa3\xce\x11\x24\x39\xf4\xda\xd4\x65\x7c\x32\x3b\x1f\x7f\xba\x9e\x4e\xce\xbe\x4e\xce\xa6\xd4\x6d\x07\x85\x16\x8a\xdf\xdf\x9c\x14\xb2\xeb\x96\x64\x5a\x05\x07\x2c\xae\x87\x63\x0d\x01\x82\x5e\x58\x34\xc0\xae\x9a\x7e\xac\xdd\x38\x5d\xf9\xc5\xe2\xe7\x5d\xd8\x36\x3f\xb9\xa5\x58\x05\xb4\x6b\x31\x53\x1a\x2c\x6e\xe5\xa9\x5c\xa4\xfb\x6d\x29\x49\xaa\x67\x76\xd6\xd9\x24\xb7\x54\x0b\xa8\xb0\x40\x38\x68\xe6\x6c\x74\xff\x6b\x39\xd8\x93\xd2\xfa\xff\x75\xac\xb7\x93\x41\x12\x7c\x66\x0d\xe9\xb8\x95\x0a\xbf\x77\x65\xaa\xcf\x6a\x23\xb3\xf6\xa8\xb6\xf0\xb8\xd7\xef\x0d\x8e\xc9\xc7\x81\x11\xd6\xcb\x54\x1a\xa1\xfc\xf0\xd1\x5e\x56\x6f\xd5\xef\xe8\x4f\x00\x00\x00\xff\xff\x39\xc8\x85\x07\x38\x04\x00\x00")

func pkg_boot_zookeeper_bash_add_node_bash_bytes() ([]byte, error) {
	return bindata_read(
		_pkg_boot_zookeeper_bash_add_node_bash,
		"pkg/boot/zookeeper/bash/add-node.bash",
	)
}

func pkg_boot_zookeeper_bash_add_node_bash() (*asset, error) {
	bytes, err := pkg_boot_zookeeper_bash_add_node_bash_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "pkg/boot/zookeeper/bash/add-node.bash", size: 1080, mode: os.FileMode(420), modTime: time.Unix(1432183300, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pkg_boot_zookeeper_bash_remove_node_bash = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x91\x5f\x4f\xdb\x30\x14\xc5\xdf\xfd\x29\x8e\x42\x04\xad\xd4\x60\x4d\x7b\xa9\x60\x3c\x74\xac\x02\xb4\x3f\xa0\xb6\x30\x69\x08\x26\x37\xb9\x6e\x3d\x12\x3b\xb3\xdd\x8a\xb1\xed\xbb\xef\x3a\x2d\xd3\xe8\xdb\x5e\x1c\xe7\xfe\x39\xbf\x7b\xae\x03\x45\x14\xe4\xd0\x9a\x96\xb4\x32\xb5\x10\x7b\x48\xb1\x8a\xe6\xab\x05\xe6\x2a\x50\x05\x67\x41\x76\xbd\x56\x5e\xdc\xde\x22\x7f\x37\x7e\x7b\x7d\x86\xbb\x3b\xec\xef\x77\x95\xc5\xa3\x10\x8d\x32\xb6\xd7\xc7\x4f\x01\xd0\x63\xeb\x7c\xc4\xd5\x68\x76\x7e\x92\xa7\xf3\x48\x7e\xf3\x24\xe7\xc6\x0a\xce\xee\xe1\x33\xa1\x54\xd6\xba\x88\x55\x20\xc4\x25\xe1\xe2\x0a\x4e\xf3\xcd\x04\x58\x57\x71\x8c\xa7\x21\xaf\x9d\x6f\x36\x79\x4f\x8d\x5b\xab\xfa\x65\x51\xf7\xc3\x52\xf5\x2a\x44\xf2\xac\xfc\xe5\xfd\xf9\xe5\x74\x76\x92\xf7\xd2\xc4\x05\x21\x0b\x32\x4f\x91\xa3\xd7\xc3\xe1\x50\xca\xe3\x20\xef\x07\xdd\x67\x90\x4b\x99\xe1\x0d\xa4\x6b\xa3\x7c\x72\xee\x81\x88\x71\xb2\x74\x56\xcb\x40\x7e\x4d\xfe\xb0\x36\x21\xe2\x17\xca\x15\xbb\xab\x70\x30\x38\x40\xa1\xf1\xaa\xcf\x94\xd1\xe9\xec\x7a\xf4\xe1\xeb\x74\x3c\xb9\x19\x4f\xa6\x4c\xdb\x51\x61\x9b\xf2\xe9\xe1\xb4\x36\x87\x61\x89\x62\x23\x87\x2c\xdf\x0c\x97\x21\x41\xcc\x82\xa5\x17\x9e\x5a\x64\xf7\x5b\x5e\xd6\x4f\xbb\x31\x1a\x54\x2e\x1d\x97\xbf\xa4\x64\xcf\xf5\xc5\x77\xce\x75\x42\xc7\xc9\xbc\xe5\x1e\x6c\x5b\x26\xe3\x8f\x97\x37\x17\x9f\xce\xd0\xe5\xb1\x05\x6b\xef\x1a\xfc\x1d\xee\x79\x59\xd9\x3f\x7d\x9b\xfb\xff\x99\xf0\xb4\xb5\x51\x74\x2f\xc3\xab\xce\x7b\xa5\x8a\x3b\x2a\x45\xa5\xa2\x92\xdd\xd1\xfc\x30\x55\x3f\x91\xb4\x11\xbf\xc5\x9f\x00\x00\x00\xff\xff\x9f\xa9\x8f\x0e\x72\x02\x00\x00")

func pkg_boot_zookeeper_bash_remove_node_bash_bytes() ([]byte, error) {
	return bindata_read(
		_pkg_boot_zookeeper_bash_remove_node_bash,
		"pkg/boot/zookeeper/bash/remove-node.bash",
	)
}

func pkg_boot_zookeeper_bash_remove_node_bash() (*asset, error) {
	bytes, err := pkg_boot_zookeeper_bash_remove_node_bash_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "pkg/boot/zookeeper/bash/remove-node.bash", size: 626, mode: os.FileMode(420), modTime: time.Unix(1432183381, 0)}
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
	"pkg/boot/zookeeper/bash/add-node.bash":    pkg_boot_zookeeper_bash_add_node_bash,
	"pkg/boot/zookeeper/bash/remove-node.bash": pkg_boot_zookeeper_bash_remove_node_bash,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func     func() (*asset, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"pkg": &_bintree_t{nil, map[string]*_bintree_t{
		"boot": &_bintree_t{nil, map[string]*_bintree_t{
			"zookeeper": &_bintree_t{nil, map[string]*_bintree_t{
				"bash": &_bintree_t{nil, map[string]*_bintree_t{
					"add-node.bash":    &_bintree_t{pkg_boot_zookeeper_bash_add_node_bash, map[string]*_bintree_t{}},
					"remove-node.bash": &_bintree_t{pkg_boot_zookeeper_bash_remove_node_bash, map[string]*_bintree_t{}},
				}},
			}},
		}},
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	if err != nil { // File
		return RestoreAsset(dir, name)
	} else { // Dir
		for _, child := range children {
			err = RestoreAssets(dir, path.Join(name, child))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
