// Code generated by go-bindata.
// sources:
// static/README.md
// static/html/devel/api.html
// static/html/devel/index.html
// static/html/devel/js/opentracing-browser.min.js
// static/html/devel/js/ws.js
// static/html/devel/openapi/api.swagger.json
// static/html/devel/openapi/index.html
// static/html/devel/stylesheet.css
// static/html/devel/ws.html
// static/html/favicon.ico
// static/html/index.html
// static/sql/crebas.sql
// DO NOT EDIT!

package staticgen

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// bindataRead reads the given file from disk. It returns an error on failure.
func bindataRead(path, name string) ([]byte, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset %s at %s: %v", name, path, err)
	}
	return buf, err
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

// readmeMd reads file data from disk. It returns an error on failure.
func readmeMd() (*asset, error) {
	path := "/home/jean/Private/tpro/GitHub/rpc-sample-app/static/README.md"
	name := "README.md"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// htmlDevelApiHtml reads file data from disk. It returns an error on failure.
func htmlDevelApiHtml() (*asset, error) {
	path := "/home/jean/Private/tpro/GitHub/rpc-sample-app/static/html/devel/api.html"
	name := "html/devel/api.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// htmlDevelIndexHtml reads file data from disk. It returns an error on failure.
func htmlDevelIndexHtml() (*asset, error) {
	path := "/home/jean/Private/tpro/GitHub/rpc-sample-app/static/html/devel/index.html"
	name := "html/devel/index.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// htmlDevelJsOpentracingBrowserMinJs reads file data from disk. It returns an error on failure.
func htmlDevelJsOpentracingBrowserMinJs() (*asset, error) {
	path := "/home/jean/Private/tpro/GitHub/rpc-sample-app/static/html/devel/js/opentracing-browser.min.js"
	name := "html/devel/js/opentracing-browser.min.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// htmlDevelJsWsJs reads file data from disk. It returns an error on failure.
func htmlDevelJsWsJs() (*asset, error) {
	path := "/home/jean/Private/tpro/GitHub/rpc-sample-app/static/html/devel/js/ws.js"
	name := "html/devel/js/ws.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// htmlDevelOpenapiApiSwaggerJson reads file data from disk. It returns an error on failure.
func htmlDevelOpenapiApiSwaggerJson() (*asset, error) {
	path := "/home/jean/Private/tpro/GitHub/rpc-sample-app/static/html/devel/openapi/api.swagger.json"
	name := "html/devel/openapi/api.swagger.json"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// htmlDevelOpenapiIndexHtml reads file data from disk. It returns an error on failure.
func htmlDevelOpenapiIndexHtml() (*asset, error) {
	path := "/home/jean/Private/tpro/GitHub/rpc-sample-app/static/html/devel/openapi/index.html"
	name := "html/devel/openapi/index.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// htmlDevelStylesheetCss reads file data from disk. It returns an error on failure.
func htmlDevelStylesheetCss() (*asset, error) {
	path := "/home/jean/Private/tpro/GitHub/rpc-sample-app/static/html/devel/stylesheet.css"
	name := "html/devel/stylesheet.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// htmlDevelWsHtml reads file data from disk. It returns an error on failure.
func htmlDevelWsHtml() (*asset, error) {
	path := "/home/jean/Private/tpro/GitHub/rpc-sample-app/static/html/devel/ws.html"
	name := "html/devel/ws.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// htmlFaviconIco reads file data from disk. It returns an error on failure.
func htmlFaviconIco() (*asset, error) {
	path := "/home/jean/Private/tpro/GitHub/rpc-sample-app/static/html/favicon.ico"
	name := "html/favicon.ico"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// htmlIndexHtml reads file data from disk. It returns an error on failure.
func htmlIndexHtml() (*asset, error) {
	path := "/home/jean/Private/tpro/GitHub/rpc-sample-app/static/html/index.html"
	name := "html/index.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// sqlCrebasSql reads file data from disk. It returns an error on failure.
func sqlCrebasSql() (*asset, error) {
	path := "/home/jean/Private/tpro/GitHub/rpc-sample-app/static/sql/crebas.sql"
	name := "sql/crebas.sql"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
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
	"README.md": readmeMd,
	"html/devel/api.html": htmlDevelApiHtml,
	"html/devel/index.html": htmlDevelIndexHtml,
	"html/devel/js/opentracing-browser.min.js": htmlDevelJsOpentracingBrowserMinJs,
	"html/devel/js/ws.js": htmlDevelJsWsJs,
	"html/devel/openapi/api.swagger.json": htmlDevelOpenapiApiSwaggerJson,
	"html/devel/openapi/index.html": htmlDevelOpenapiIndexHtml,
	"html/devel/stylesheet.css": htmlDevelStylesheetCss,
	"html/devel/ws.html": htmlDevelWsHtml,
	"html/favicon.ico": htmlFaviconIco,
	"html/index.html": htmlIndexHtml,
	"sql/crebas.sql": sqlCrebasSql,
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
	"README.md": &bintree{readmeMd, map[string]*bintree{}},
	"html": &bintree{nil, map[string]*bintree{
		"devel": &bintree{nil, map[string]*bintree{
			"api.html": &bintree{htmlDevelApiHtml, map[string]*bintree{}},
			"index.html": &bintree{htmlDevelIndexHtml, map[string]*bintree{}},
			"js": &bintree{nil, map[string]*bintree{
				"opentracing-browser.min.js": &bintree{htmlDevelJsOpentracingBrowserMinJs, map[string]*bintree{}},
				"ws.js": &bintree{htmlDevelJsWsJs, map[string]*bintree{}},
			}},
			"openapi": &bintree{nil, map[string]*bintree{
				"api.swagger.json": &bintree{htmlDevelOpenapiApiSwaggerJson, map[string]*bintree{}},
				"index.html": &bintree{htmlDevelOpenapiIndexHtml, map[string]*bintree{}},
			}},
			"stylesheet.css": &bintree{htmlDevelStylesheetCss, map[string]*bintree{}},
			"ws.html": &bintree{htmlDevelWsHtml, map[string]*bintree{}},
		}},
		"favicon.ico": &bintree{htmlFaviconIco, map[string]*bintree{}},
		"index.html": &bintree{htmlIndexHtml, map[string]*bintree{}},
	}},
	"sql": &bintree{nil, map[string]*bintree{
		"crebas.sql": &bintree{sqlCrebasSql, map[string]*bintree{}},
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

