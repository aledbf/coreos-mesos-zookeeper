package bindata

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
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
	name string
	size int64
	mode os.FileMode
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

var _pkg_boot_mesos_marathon_bash_update_hosts_file_bash = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x1c\xca\xc1\x0a\x82\x40\x14\x05\xd0\xfd\xfb\x8a\x8b\x0d\xa2\x0b\xf1\x0f\x5c\x44\x51\xbb\x16\xd5\x4a\x5c\x3c\xf5\x96\x03\x39\x23\xcd\x24\x41\xf4\xef\x31\x6d\x0f\x27\x30\xa2\xa2\xc7\x62\x17\xde\xd4\x3e\x44\x36\x48\x36\xb2\x7f\xdd\xd1\x6b\xe0\x08\xef\x40\xb7\xae\xfa\x94\xb6\x85\xd9\xed\xb7\xd7\x03\xba\x0e\x79\xfe\x9f\xd5\x5b\x64\x56\xeb\x8a\x12\x1f\x01\x38\x4c\x1e\x99\x39\x9e\xce\x17\x98\x62\xf2\x21\x3a\x9d\x59\x66\x68\x1a\xd4\x8c\x43\x9d\x28\xc8\x57\x7e\x01\x00\x00\xff\xff\xed\xa2\x52\x1e\x7a\x00\x00\x00")

func pkg_boot_mesos_marathon_bash_update_hosts_file_bash_bytes() ([]byte, error) {
	return bindata_read(
		_pkg_boot_mesos_marathon_bash_update_hosts_file_bash,
		"pkg/boot/mesos/marathon/bash/update-hosts-file.bash",
	)
}

func pkg_boot_mesos_marathon_bash_update_hosts_file_bash() (*asset, error) {
	bytes, err := pkg_boot_mesos_marathon_bash_update_hosts_file_bash_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "pkg/boot/mesos/marathon/bash/update-hosts-file.bash", size: 122, mode: os.FileMode(420), modTime: time.Unix(1432350154, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"pkg/boot/mesos/marathon/bash/update-hosts-file.bash": pkg_boot_mesos_marathon_bash_update_hosts_file_bash,
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
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"pkg": &_bintree_t{nil, map[string]*_bintree_t{
		"boot": &_bintree_t{nil, map[string]*_bintree_t{
			"mesos": &_bintree_t{nil, map[string]*_bintree_t{
				"marathon": &_bintree_t{nil, map[string]*_bintree_t{
					"bash": &_bintree_t{nil, map[string]*_bintree_t{
						"update-hosts-file.bash": &_bintree_t{pkg_boot_mesos_marathon_bash_update_hosts_file_bash, map[string]*_bintree_t{
						}},
					}},
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

