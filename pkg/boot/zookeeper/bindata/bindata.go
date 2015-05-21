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

var _pkg_boot_zookeeper_bash_add_node_bash = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x52\x4d\x6f\x13\x31\x14\xbc\xef\xaf\x18\xb9\xab\x36\x91\xb2\xb1\x52\x2e\x51\x42\x0e\x7c\x89\x4a\x1c\x40\xa2\x08\x89\xaa\x20\xc7\xfb\x36\x31\xdd\xd8\x96\xed\x44\x4d\x81\xff\xce\xf3\x86\xa6\x28\x84\xaa\x17\x7b\xd7\x6f\x66\xde\x9b\xb1\x23\x25\x54\xe4\xe0\x8d\xa7\x46\x99\xb6\x28\x4e\x90\xcf\x6a\x9a\xaf\x17\x98\xab\x48\x35\x9c\x05\xd9\xcd\x46\x85\xe2\xea\x0a\xe5\xeb\x37\x2f\x3f\xbd\xc5\xf5\x35\x4e\x4f\x3b\x64\x75\x5b\x14\x2b\x65\x6c\xaf\x8f\x1f\x05\x40\xb7\xde\x85\x84\x0f\x2f\x2e\x2f\x66\x65\x5e\x27\xf2\x7b\x20\x39\x37\xb6\xe0\xaa\xf6\x90\xca\x7b\x79\xe7\xdc\x50\x37\x0b\x48\xe7\x53\xfe\xb9\x21\xf2\x14\xaa\x5a\x25\x75\x5f\x63\x74\x6b\x51\xc5\xc7\x30\x07\x35\xa9\x9d\x6d\xf6\x7c\x16\x48\x6e\xad\x97\xff\x13\xf8\xc6\xa0\x61\xbd\xb5\x6a\x65\x74\x06\x9b\x06\xec\x8f\x1b\x8a\x27\x10\x44\x4e\x60\x8a\xb4\x24\xcb\x54\xb6\xad\x97\x0e\x95\x85\x50\x6d\x20\x55\x6f\x39\x07\x13\x53\xc4\x5e\x05\xba\x5d\xc7\xc4\xbb\xb1\x8d\x0b\x2b\x95\x8c\xb3\x22\xe7\xd5\x46\xea\x14\x72\x34\x47\xdc\x34\x2d\x51\xaa\x0e\xba\x3f\xc9\x12\xd0\x98\xec\xeb\x04\x9f\x09\x96\xf8\x22\x93\x83\xaa\x79\x5b\x9a\x08\xeb\x6a\xca\x07\xec\xe0\x61\xb4\x06\x5d\x25\x21\xb7\x36\x8b\x75\x60\x92\xb1\x7f\x63\x58\xef\xcb\xbb\x8b\xf7\x1f\x2f\x67\x65\x2f\xbf\x8d\x8a\x20\xa2\x2c\xf3\xc9\xe4\xd9\x78\x3c\x96\x72\x1a\xe5\xd7\x41\xb7\x0d\x4a\x29\x05\x9e\x1f\xb5\x15\x29\x6c\x28\x0c\x5b\x0e\x09\x3f\xa1\xd7\xfc\x8e\x6a\x9c\x0d\xce\x50\x35\x18\xf5\xf3\xd8\x5d\xa2\x82\xe7\x35\x76\x81\xae\x01\x54\xc4\x8e\x77\x3f\xf8\x3f\xe9\x8a\x3d\x31\x7f\x1d\x34\xe6\x37\x28\xef\x6e\x5e\xb5\x66\x18\x97\x7c\xcf\x3b\x25\x51\xee\xfc\x08\x04\xda\xb9\x46\x95\x43\x12\x7f\x26\x2c\x7b\x5a\xa5\xa3\x79\x77\xcb\x6a\x6b\xea\xfe\x6c\xe7\xff\x7c\x34\x1e\x4d\xce\x39\x84\x89\x57\x21\x19\x6d\xbc\xb2\x69\xfa\x90\x8d\x28\x7e\x15\xbf\x03\x00\x00\xff\xff\x98\xb5\xd3\x50\x70\x03\x00\x00")

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

	info := bindata_file_info{name: "pkg/boot/zookeeper/bash/add-node.bash", size: 880, mode: os.FileMode(420), modTime: time.Unix(1432221550, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pkg_boot_zookeeper_bash_remove_node_bash = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x91\x5f\x6f\xd3\x30\x14\xc5\xdf\xfd\x29\x8e\xb2\x68\x6b\xa5\x7a\x16\xe2\xa5\xda\xd8\x43\x19\x15\x43\x20\x31\xb5\x1d\x48\x4c\x1b\x72\x93\xeb\xd6\x2c\xb1\x83\xed\x56\x63\xc0\x77\xe7\x3a\xed\x10\xdb\x1b\x2f\x8e\x73\xff\x9c\xdf\x3d\xd7\x91\x12\x24\x79\x74\xb6\x23\xa3\x6d\x23\xc4\x01\x72\xac\xa6\xe5\x66\x85\xa5\x8e\x54\xc3\x3b\x90\xdb\x6e\x75\x10\xd7\xd7\x28\xdf\x4c\x5f\x5f\xbd\xc5\xcd\x0d\x0e\x0f\xfb\x4a\x79\x2f\x44\xab\xad\x1b\x0c\xf1\x53\x00\x74\xdf\xf9\x90\x70\x39\x59\x5c\x9c\x95\xf9\x3c\x51\xdf\x02\xa9\xa5\x75\x82\xb3\x07\xf8\x4c\xa8\xb4\x73\x3e\x61\x13\x09\x69\x4d\x78\x77\x09\x6f\xf8\x66\x23\x9c\xaf\x39\xc6\xd3\x50\x30\x3e\xb4\xbb\x7c\xa0\xd6\x6f\x75\xf3\xb4\xa8\xff\x61\xa9\x66\x13\x13\x05\x56\xfe\xf2\xfe\xe2\xe3\x7c\x71\x56\x0e\xf2\xc4\x92\x50\x44\x55\xe6\xc8\xc9\xcb\xf1\x78\xac\xd4\x69\x54\xb7\xa3\xfe\x33\x2a\x95\x2a\xf0\x0a\xca\x77\x49\x3d\x78\x7f\x47\xc4\x38\x55\x79\x67\x54\xa4\xb0\xa5\x70\xdc\xd8\x98\xf0\x0b\xd5\x86\xdd\xd5\x38\x1a\x1d\x41\x1a\xbc\x18\x32\x65\x72\xbe\xb8\x9a\x7c\xf8\x3a\x9f\xce\x3e\x4d\x67\x73\xa6\x3d\x53\x61\x9b\xea\xe1\xee\xbc\xb1\xc7\x71\x0d\xb9\x93\x43\x51\xee\x86\x2b\x90\x21\x76\xc5\xd2\xab\x40\x1d\x8a\xdb\x3d\xaf\x18\xe6\xdd\x58\x03\xaa\xd6\x9e\xcb\x9f\x52\x8a\xc7\x7a\xf9\x9d\x73\xbd\xd0\x69\x36\xef\xb8\x07\xfb\x96\x59\xde\x91\x75\x2b\xf4\x79\xec\xc1\x26\xf8\x16\x7f\x87\x7b\x5c\x56\xf1\x4f\xdf\xee\xfe\x7f\x26\x02\xed\x6d\xc8\xfe\x65\x78\xd5\xe5\xa0\xd2\xe9\x99\x8a\xac\x75\xd2\xaa\x3f\xda\x1f\xb6\x1e\x66\x92\xb1\xe2\xb7\xf8\x13\x00\x00\xff\xff\x33\x7d\xe1\x5a\x72\x02\x00\x00")

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

	info := bindata_file_info{name: "pkg/boot/zookeeper/bash/remove-node.bash", size: 626, mode: os.FileMode(420), modTime: time.Unix(1432220733, 0)}
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
