// Code generated by go-bindata. DO NOT EDIT.
// sources:
// schemas/jsonschema-draft-04.json (4.357kB)
// schemas/v2/schema.json (40.249kB)

package spec

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _jsonschemaDraft04JSON = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x57\x3d\x6f\xdb\x3c\x10\xde\xf3\x2b\x08\x26\x63\xf2\x2a\x2f\xd0\xc9\x5b\xd1\x2e\x01\x5a\x34\x43\x37\x23\x03\x6d\x9d\x6c\x06\x14\xa9\x50\x54\x60\xc3\xd0\x7f\x2f\x28\x4a\x14\x29\x91\x92\x2d\xa7\x8d\x97\x28\xbc\xaf\xe7\x8e\xf7\xc5\xd3\x0d\x42\x08\x61\x9a\xe2\x15\xc2\x7b\xa5\x8a\x55\x92\xbc\x96\x82\x3f\x94\xdb\x3d\xe4\xe4\x3f\x21\x77\x49\x2a\x49\xa6\x1e\x1e\xbf\x24\xe6\xec\x16\xdf\x1b\xa1\x3b\xf3\xff\x02\xc9\x14\xca\xad\xa4\x85\xa2\x82\x6b\xe9\x6f\x42\x02\x32\x2c\x28\x07\x45\x5a\x15\x3d\x77\x46\x39\xd5\xcc\x25\x5e\x21\x83\xb8\x21\x18\xb6\xaf\x52\x92\xa3\x47\x68\x88\xea\x58\x80\x56\x4e\x1a\xf2\xbd\x4f\xcc\x29\x7f\x52\x90\x6b\x7d\xff\x0f\x48\xb4\x3d\x3f\x21\x7c\x27\x21\xd3\x2a\x6e\x31\xaa\x2d\x53\xdd\xf3\xe3\x42\x94\x54\xd1\x77\x78\xe2\x0a\x76\x20\xe3\x20\x68\xcb\x30\x86\x41\xf3\x2a\xc7\x2b\xf4\x78\x8e\xfe\xef\x90\x91\x8a\xa9\xc7\xb1\x1d\xc2\xd8\x2f\x0d\x75\xed\xc1\x4e\x9c\xc8\x25\x43\xac\xa8\xbe\xd7\xcc\xa9\xd1\xa9\x21\xa0\x1a\xbd\x04\x61\x94\x34\x2f\x18\xfc\x3e\x16\x50\x8e\x4d\x03\x6f\x1c\x58\xdb\x48\x23\xbc\x11\x82\x01\xe1\xfa\xd3\x3a\x8e\x30\xaf\x18\x33\x7f\xf3\x8d\x39\x11\x9b\x57\xd8\x2a\xfd\x55\x2a\x49\xf9\x0e\xc7\xec\x37\xd4\x25\xf7\xec\x5c\x66\xc7\xd7\x99\xaa\xcf\x4f\x89\x8a\xd3\xb7\x0a\x3a\xaa\x92\x15\xf4\x30\x6f\x1c\xb0\xd6\x46\xe7\x98\x39\x2d\xa4\x28\x40\x2a\x3a\x88\x9e\x29\xba\x88\x37\x2d\xca\x60\x38\xfa\xba\x5b\x20\xac\xa8\x62\xb0\x4c\xd4\xaf\xda\x45\x0a\xba\x5c\x3b\xb9\xc7\x79\xc5\x14\x2d\x18\x34\x19\x1c\x51\xdb\x25\x4d\xb4\x7e\x06\x14\x38\x6c\x59\x55\xd2\x77\xf8\x69\x59\xfc\x7b\x73\xed\x93\x43\xcb\x32\x6d\x3c\x28\xdc\x1b\x9a\xd3\x62\xab\xc2\x27\xf7\x41\xc9\x08\x2b\x23\x08\xad\x13\x57\x21\x9c\xd3\x72\x0d\x42\x72\xf8\x01\x7c\xa7\xf6\x83\xce\x39\xd7\x82\x3c\x1f\x2f\xd6\x60\x1b\xa2\xdf\x35\x89\x52\x20\xe7\x73\x74\xe0\x66\x26\x64\x4e\xb4\x97\x58\xc2\x0e\x0e\xe1\x60\x92\x34\x6d\xa0\x10\xd6\xb5\x83\x61\x27\xe6\x47\xd3\x89\xbd\x63\xfd\x3b\x8d\x03\x3d\x6c\x42\x2d\x5b\x70\xee\xe8\xdf\x4b\xf4\x66\x4e\xe1\x01\x45\x17\x80\x74\xad\x4f\xc3\xf3\xae\xc6\x1d\xc6\xd7\xc2\xce\xc9\xe1\x29\x30\x86\x2f\x4a\xa6\x4b\x15\x84\x73\xc9\x6f\xfd\x7f\xa5\x6e\x9e\xbd\xf1\xb0\xd4\xdd\x45\x5a\xc2\x3e\x4b\x78\xab\xa8\x84\x74\x4a\x91\x3b\x92\x23\x05\xf2\x1c\x1e\x7b\xf3\x09\xf8\xcf\xab\x24\xb6\x60\xa2\xe8\x4c\x9f\x75\x77\xaa\x8c\xe6\x01\x45\x36\x86\xcf\xc3\x63\x3a\xea\xd4\x8d\x7e\x06\xac\x14\x0a\xe0\x29\xf0\xed\x07\x22\x1a\x65\xda\x44\xae\xa2\x73\x1a\xe6\x90\x69\xa2\x8c\x46\xb2\x2f\xde\x49\x38\x08\xed\xfe\xfd\x41\xaf\x9f\xa9\x55\xd7\xdd\x22\x8d\xfa\x45\x63\xc5\x0f\x80\xf3\xb4\x08\xd6\x79\x30\x9e\x93\xee\x59\xa6\xd0\x4b\xee\x22\xe3\x33\xc1\x3a\x27\x68\x36\x78\x7e\x87\x0a\x06\xd5\x2e\x20\xd3\xaf\x15\xfb\xd8\x3b\x73\x14\xbb\x92\xed\x05\x5d\x2e\x29\x38\x2c\x94\xe4\x42\x45\x5e\xd3\xb5\x7d\xdf\x47\xca\x38\xb4\x5c\xaf\xfb\x7d\xdd\x6d\xf4\xa1\x2d\x77\xdd\x2f\xce\x6d\xc4\x7b\x8b\x4e\x67\xa9\x6f\xfe\x04\x00\x00\xff\xff\xb1\xd1\x27\x78\x05\x11\x00\x00")

func jsonschemaDraft04JSONBytes() ([]byte, error) {
	return bindataRead(
		_jsonschemaDraft04JSON,
		"jsonschema-draft-04.json",
	)
}

func jsonschemaDraft04JSON() (*asset, error) {
	bytes, err := jsonschemaDraft04JSONBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "jsonschema-draft-04.json", size: 4357, mode: os.FileMode(0644), modTime: time.Unix(1567900649, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe1, 0x48, 0x9d, 0xb, 0x47, 0x55, 0xf0, 0x27, 0x93, 0x30, 0x25, 0x91, 0xd3, 0xfc, 0xb8, 0xf0, 0x7b, 0x68, 0x93, 0xa8, 0x2a, 0x94, 0xf2, 0x48, 0x95, 0xf8, 0xe4, 0xed, 0xf1, 0x1b, 0x82, 0xe2}}
	return a, nil
}

var _v2SchemaJSON = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5d\x4f\x93\xdb\x36\xb2\xbf\xfb\x53\xa0\x14\x57\xd9\xae\xd8\x92\xe3\xf7\x2e\xcf\x97\xd4\xbc\xd8\x49\x66\x37\x5e\x4f\x79\x26\xbb\x87\x78\x5c\x05\x91\x2d\x09\x09\x09\x30\x00\x38\x33\x5a\xef\x7c\xf7\x2d\xf0\x9f\x08\x02\x20\x41\x8a\xd2\xc8\x0e\x0f\xa9\x78\x28\xa0\xd1\xdd\x68\x34\x7e\xdd\xf8\xf7\xf9\x11\x42\x33\x49\x64\x04\xb3\xd7\x68\x76\x86\xfe\x76\xf9\xfe\x1f\xe8\x32\xd8\x40\x8c\xd1\x8a\x71\x74\x79\x8b\xd7\x6b\xe0\xe8\xd5\xfc\x25\x3a\xbb\x38\x9f\xcf\x9e\xab\x0a\x24\x54\xa5\x37\x52\x26\xaf\x17\x0b\x91\x17\x99\x13\xb6\xb8\x79\xb5\x10\x59\xdd\xf9\xef\x82\xd1\x6f\xf2\xc2\x8f\xf3\x4f\xb5\x1a\xea\xc7\x17\x45\x41\xc6\xd7\x8b\x90\xe3\x95\x7c\xf1\xf2\x7f\x8b\xca\x45\x3d\xb9\x4d\x32\xa6\xd8\xf2\x77\x08\x64\xfe\x8d\xc3\x9f\x29\xe1\xa0\x9a\xff\xed\x11\x42\x08\xcd\x8a\xd6\xb3\x9f\x15\x67\x74\xc5\xca\x7f\x27\x58\x6e\xc4\xec\x11\x42\xd7\x59\x5d\x1c\x86\x44\x12\x46\x71\x74\xc1\x59\x02\x5c\x12\x10\xb3\xd7\x68\x85\x23\x01\x59\x81\x04\x4b\x09\x9c\x6a\xbf\x7e\xce\x49\x7d\xba\x7b\x51\xfd\xa1\x44\xe2\xb0\x52\xac\x7d\xb3\x08\x61\x45\x68\x46\x56\x2c\x6e\x80\x86\x8c\xbf\xbd\x93\x40\x05\x61\x74\x96\x95\xbe\x7f\x84\xd0\x7d\x4e\xde\x42\xb7\xe4\xbe\x46\xbb\x14\x5b\x48\x4e\xe8\xba\x90\x05\xa1\x19\xd0\x34\xae\xc4\xce\xbe\xbc\x9a\xbf\x9c\x15\x7f\x5d\x57\xc5\x42\x10\x01\x27\x89\xe2\x48\x51\xb9\xda\x40\xd5\x87\x37\xc0\x15\x5f\x88\xad\x90\xdc\x10\x81\x42\x16\xa4\x31\x50\x39\x2f\x38\xad\xab\xb0\x53\xd8\xac\x94\x56\x6f\xc3\x84\xf4\x11\xa4\x50\xb3\xfa\xe9\xd3\x6f\x9f\x3e\xdf\x2f\xd0\xeb\x8f\x1f\x3f\x7e\xbc\xfe\xf6\xe9\xf7\xaf\x5f\x7f\xfc\x18\x7e\xfb\xec\xfb\xc7\xb3\x36\x79\x54\x43\xe8\x29\xc5\x31\x20\xc6\x11\x49\x9e\xe5\x12\x41\x66\xa0\xe8\xed\x1d\x8e\x93\x08\x5e\xa3\x27\x3b\xc3\x7c\xa2\x73\xba\xc4\x02\x2e\xb0\xdc\xf4\xe5\x76\xd1\xca\x96\xa2\x8a\x94\xcd\x21\xc9\x6c\xec\x2c\x70\x42\x9e\x34\x74\x9d\x19\x7c\xcd\x20\x9c\xea\x2e\x0a\xfe\x42\x84\xd4\x29\x04\x8c\x8a\xb4\x41\xa2\xc1\xdc\x19\x8a\x88\x90\x4a\x49\xef\xce\xdf\xbd\x45\x4a\x52\x81\x70\x10\x40\x22\x21\x44\xcb\x6d\xc5\xec\x4e\x3c\x1c\x45\xef\x57\x9a\xb5\x7d\xae\xfe\xe5\xe4\x31\x86\x90\xe0\xab\x6d\x02\x3b\x2e\xcb\x11\x90\xd9\xa8\xc6\x77\xc2\x59\x98\x06\xfd\xf9\x2e\x78\x45\x01\xa6\xa8\xa0\x71\x5c\xbe\x33\xa7\xd2\xd9\x5f\x95\xef\xd9\xd5\xac\xfd\xdc\x5d\xbf\x5e\xb8\xd1\x3e\xc7\x31\x48\xe0\x5e\x4c\x14\x65\xdf\xb8\xa8\x71\x10\x09\xa3\xc2\xc7\x02\xcb\xa2\x4e\x5a\x02\x82\x94\x13\xb9\xf5\x30\xe6\xb2\xa4\xb5\xfe\x9b\x3e\x7a\xb2\x55\xd2\xa8\x4a\xbc\x16\xb6\x71\x8e\x39\xc7\xdb\x9d\xe1\x10\x09\x71\xbd\x9c\xb3\x41\x89\xd7\xa5\x89\xdc\x57\xb5\x53\x4a\xfe\x4c\xe1\xbc\xa0\x21\x79\x0a\x1a\x0f\x70\xa7\x5c\x08\x8e\xde\xb0\xc0\x43\x24\xad\x74\x63\x0e\xb1\xd9\x90\xe1\xb0\x2d\x13\xa7\x6d\x78\xfd\x04\x14\x38\x8e\x90\xaa\xce\x63\xac\x3e\x23\xbc\x64\xa9\xb4\xf8\x03\x63\xde\xcd\xbe\x16\x13\x4a\x55\xac\x82\x12\xc6\xac\xd4\x35\xf7\x22\xd4\x3a\xff\x22\x73\x0e\x6e\x51\xa0\x75\x1e\xae\x8f\xe8\x5d\xc7\x59\xe6\xe4\x9a\x18\x8d\xd6\x1c\x53\x84\x4d\xb7\x67\x28\x37\x09\x84\x69\x88\x12\x0e\x01\x11\x80\x32\xa2\xf5\xb9\xaa\xc6\xd9\x73\x53\xab\xfb\xb4\x2e\x20\xc6\x54\x92\xa0\x9a\xf3\x69\x1a\x2f\x81\x77\x37\xae\x53\x1a\xce\x40\xc4\xa8\x82\x1c\xb5\xef\xda\x24\x7d\xb9\x61\x69\x14\xa2\x25\xa0\x90\xac\x56\xc0\x81\x4a\xb4\xe2\x2c\xce\x4a\x64\x7a\x9a\x23\xf4\x13\x91\x3f\xa7\x4b\xf4\x63\x84\x6f\x18\x87\x10\xbd\xc3\xfc\x8f\x90\xdd\x52\x44\x04\xc2\x51\xc4\x6e\x21\x74\x48\x21\x81\xc7\xe2\xfd\xea\x12\xf8\x0d\x09\xf6\xe9\x47\x35\xaf\x67\xc4\x14\xf7\x22\x27\x97\xe1\xe2\x76\x2d\x06\x8c\x4a\x1c\x48\x3f\x73\x2d\x0b\x5b\x29\x45\x24\x00\x2a\x0c\x11\xec\x94\xca\xc2\xa6\xc1\x37\x21\x43\x83\x3b\x5f\x97\xf1\x43\x5e\x53\x73\x19\xa5\x36\xd8\x2d\x05\x2e\x34\x0b\xeb\x39\xfc\x1d\x63\x51\x01\xbd\x3d\xbb\x90\x84\x40\x25\x59\x6d\x09\x5d\xa3\x1c\x37\xe6\x5c\x16\x9a\x40\x09\x70\xc1\xe8\x82\xf1\x35\xa6\xe4\xdf\x99\x5c\x8e\x9e\x4d\x79\xb4\x27\x2f\xbf\x7e\xf8\x05\x25\x8c\x50\xa9\x98\x29\x90\x62\x60\xea\x75\xae\x13\xca\xbf\x2b\x1a\x29\x27\x76\xd6\x20\xc6\x64\x5f\xe6\x32\x1a\x08\x87\x21\x07\x21\xbc\xb4\xe4\xe0\x32\x67\xa6\xcd\xf3\x1e\xcd\xd9\x6b\xb6\x6f\x8e\x27\xa7\xed\xdb\xe7\xbc\xcc\x1a\x07\xce\x6f\x87\x33\xf0\xba\x51\x17\x22\x66\x78\x79\x8e\xce\xe5\x13\x81\x80\x06\x2c\xe5\x78\x0d\xa1\xb2\xb8\x54\xa8\x79\x09\xbd\xbf\x3c\x47\x01\x8b\x13\x2c\xc9\x32\xaa\xaa\x1d\xd5\xee\xab\x36\xbd\x6c\xfd\x54\x6c\xc8\x08\x01\x3c\xbd\xe7\x07\x88\xb0\x24\x37\x79\x90\x28\x4a\x1d\x10\x1a\x92\x1b\x12\xa6\x38\x42\x40\xc3\x4c\x43\x62\x8e\xae\x36\xb0\x45\x71\x2a\xa4\x9a\x23\x79\x59\xb1\xa8\xf2\xa4\x0c\x60\x9f\xcc\x8d\x40\xf5\x80\xca\xa8\x99\xc3\xa7\x85\x1f\x31\x25\xa9\x82\xc5\x6d\xbd\xd8\x36\x76\x7c\x02\x28\x97\xf6\x1d\x74\x3b\x11\x7e\x91\xae\x32\xf8\x6c\xf4\xe6\x7b\x9a\xa5\x1f\x62\xc6\x21\xcf\x9a\xe5\xed\x8b\x02\xf3\x2c\x33\x33\xdf\x00\xca\xc9\x09\xb4\x04\xf5\xa5\x08\xd7\xc3\x02\x18\x66\xf1\xab\x1e\x83\x37\x4c\xcd\x12\xc1\x1d\x50\xf6\xaa\xbd\xfe\xe2\x73\x48\x38\x08\xa0\x32\x9b\x18\x44\x86\x0b\x6a\xc1\xaa\x26\x96\x2d\x96\x3c\xa0\x54\x65\x73\x87\x15\xca\x15\xe5\xf5\x94\x46\x9f\x33\x1a\x0c\x9a\xb1\x5a\xd9\x6a\x95\xcd\xcb\x7e\xec\x9a\xc5\x94\x3b\x37\x26\x31\xd7\xfc\xe4\x1f\x13\x8c\x31\x75\x9c\xba\xf7\x87\x3c\xa1\xb7\x4f\x17\x1b\x09\x82\x98\xc4\x70\x95\xd3\xe8\x4c\x48\x5a\xa6\xd6\x2a\x3d\x56\x42\x80\x9f\xaf\xae\x2e\x50\x0c\x42\xe0\x35\x34\x3c\x8a\x62\x03\x37\xba\xb2\x27\x04\xda\x25\x8d\x06\xe2\xa0\x13\x8a\xf3\xf5\xec\x10\x72\x67\x88\x90\x3d\x4b\x64\xeb\xaa\xda\x8f\xf7\x5a\x75\x47\x9a\xa8\x51\x70\x26\xd2\x38\xc6\x7c\xbb\x57\xfc\xbd\xe4\x04\x56\xa8\xa0\x54\x9a\x45\xd5\xf7\x0f\x16\xfc\x57\x1c\x3c\xdf\x23\xba\x77\x38\xda\x16\x4b\x31\x53\x6a\x4d\x9a\x15\x63\xe7\xe1\x18\x69\x9f\x22\xe0\x24\xbb\x94\x4b\x97\xee\x2d\xf9\x70\x87\x72\x7b\xe6\xc4\x33\x2a\x66\x5e\x1c\x35\x72\xe3\x2d\xda\x73\xe4\xc7\x51\x6d\xa4\xa1\x2a\x4f\xde\x94\xcb\xb2\x3e\x31\x48\xae\x82\xce\xc9\xc8\x65\xcd\xc3\xb7\x34\xb6\x2b\xdf\x58\x65\x78\x6e\x73\xac\x5e\x24\x0d\x3f\xdc\x70\x23\xc6\xda\x52\x0b\x2d\x63\x7d\xa9\x49\x2d\x54\x48\x28\xc0\x12\x9c\xe3\x63\xc9\x58\x04\x98\x36\x07\xc8\x0a\xa7\x91\xd4\xf0\xbc\xc1\xa8\xb9\x70\xd0\xc6\xa9\xb6\x78\x80\x5a\xa3\xb4\x2c\xf4\x18\x0b\x8a\x9d\xd0\xb4\x55\x10\xee\x0d\xc5\xd6\xe0\x99\x93\xdc\xa1\x04\xbb\xf1\xa7\x23\xd1\xd1\x97\x8c\x87\x13\x0a\x21\x02\xe9\x99\x25\xed\x20\xc5\x92\x66\x3c\x32\x9c\xd6\x06\xb0\x31\x5c\x86\x29\x0a\xcb\x60\x33\x12\xa5\x91\xfc\x96\x75\xd0\x59\xd7\x13\xbd\xd3\x23\x79\xdd\x2a\x90\xa6\x38\x06\x91\x39\x7f\x20\x72\x03\x1c\x2d\x01\x61\xba\x45\x37\x38\x22\x61\x8e\x71\x85\xc4\x32\x15\x28\x60\x61\x16\xb8\x3d\x29\xdc\x4d\x3d\x2f\x12\x13\x7d\xc8\x7e\x37\xee\xa8\x7f\xfa\xdb\xcb\x17\xff\x77\xfd\xf9\x7f\xee\x9f\x3d\xfe\xcf\xa7\xa7\x45\xfb\xcf\x1e\xf7\xf3\xe0\xff\xc4\x51\x0a\x8e\x4c\xcb\x01\xdc\x0a\x65\xb2\x01\x83\xed\x3d\xe4\xa9\xa3\x4e\x2d\x59\xc5\xe8\x2f\x48\x7d\x5a\x6e\x37\xbf\x5c\x9f\x35\x13\x64\x14\xfa\xef\x0b\x68\xa6\x0d\xb4\x8e\xf1\xa8\xff\xbb\x60\xf4\x03\x64\xab\x5b\x81\x65\x51\xe6\xda\xca\xfa\xf0\xb0\xac\x3e\x9c\xca\x26\x0e\x1d\xdb\x57\x5b\xbb\xb4\x9a\xa6\xb6\x9b\x1a\x6b\xd1\x9a\x9e\x7e\x33\x9a\xec\x41\x69\x45\x22\xb8\xb4\x51\xeb\x04\x77\xca\x6f\x7b\x7b\xc8\xb2\xb0\x95\x92\x25\x5b\xd0\x42\xaa\x2a\xdd\x32\x78\x4f\x0c\xab\x68\x46\x6c\xea\x6d\xf4\x5c\x5e\xde\xc4\xac\xa5\xf9\xd1\x00\x9f\x7d\x98\x65\x24\xbd\xc7\x97\xd4\xb3\x3a\xa8\x2b\xa0\x34\x76\xf9\x65\x5f\x2d\x25\x95\x1b\xcf\xd6\xf4\x9b\x5f\x09\x95\xb0\x36\x3f\xdb\xd0\x39\x2a\x93\x1c\x9d\x03\xa2\x4a\xca\xf5\xf6\x10\xb6\x94\x89\x0b\x6a\x70\x12\x13\x49\x6e\x40\xe4\x29\x12\x2b\xbd\x80\x45\x11\x04\xaa\xc2\x8f\x56\x9e\x5c\x6b\xec\x8d\x5a\x0e\x14\x59\x06\x2b\x1e\x24\xcb\xc2\x56\x4a\x31\xbe\x23\x71\x1a\xfb\x51\x2a\x0b\x3b\x1c\x48\x10\xa5\x82\xdc\xc0\xbb\x3e\x24\x8d\x5a\x76\x2e\x09\xed\xc1\x65\x51\xb8\x83\xcb\x3e\x24\x8d\x5a\x2e\x5d\xfe\x02\x74\x2d\x3d\xf1\xef\xae\xb8\x4b\xe6\x5e\xd4\xaa\xe2\x2e\x5c\x5e\xec\x0e\xf5\x5b\x0c\xcb\x0a\xbb\xa4\x3c\xf7\x1f\x2a\x55\x69\x97\x8c\x7d\x68\x95\xa5\xad\xb4\xf4\x9c\xa5\x07\xb9\x7a\x05\xbb\xad\x50\x6f\xfb\xa0\x4e\x9b\x48\x23\x49\x92\x28\x87\x19\x3e\x32\xee\xca\x3b\x46\x7e\x7f\x18\x64\xcc\xcc\x0f\x34\xe9\x36\x8b\xb7\x6c\xa8\xa5\x5b\x54\x4c\x54\x5b\x15\x3a\xf1\x6c\x2d\xfe\x96\xc8\x0d\xba\x7b\x81\x88\xc8\x23\xab\xee\x7d\x3b\x92\xa7\x60\x29\xe3\xdc\xff\xb8\x64\xe1\xf6\xa2\x5a\x59\xdc\x6f\xeb\x45\x7d\x6a\xd1\x76\x1e\xea\xb8\xf1\xfa\x14\xd3\x36\x63\xe5\xd7\xf3\xe4\xbe\x25\xbd\x5e\x05\xeb\x73\x74\xb5\x21\x2a\x2e\x4e\xa3\x30\xdf\xbf\x43\x28\x2a\xd1\xa5\x2a\x9d\x8a\xfd\x76\xd8\x8d\xbc\x67\x65\xc7\xb8\x03\x45\xec\xa3\xb0\x37\x8a\x70\x4c\x68\x91\x51\x8e\x58\x80\xed\x4a\xf3\x81\x62\xca\x96\xbb\xf1\x52\xcd\x80\xfb\xe4\x4a\x5d\x6c\xdf\x6e\x20\x4b\x80\x30\x8e\x28\x93\xf9\xe9\x8d\x8a\x6d\xd5\x59\x65\x7b\xaa\x44\x9e\xc0\xc2\xd1\x7c\x40\x26\xd6\x1a\xce\xf9\xc5\x69\x7b\x6c\xec\xc8\x71\x7b\xe5\x21\x2e\xd3\xe5\x65\x93\x91\x53\x0b\x7b\x3a\xc7\xfa\x17\x6a\x01\xa7\x33\xd0\xf4\x40\x0f\x39\x87\xda\xe4\x54\x87\x3a\xd5\xe3\xc7\xa6\x8e\x20\xd4\x11\xb2\x4e\xb1\xe9\x14\x9b\x4e\xb1\xe9\x14\x9b\xfe\x15\x63\xd3\x47\xf5\xff\x97\x38\xe9\xcf\x14\xf8\x76\x82\x49\x13\x4c\xaa\x7d\xcd\x6c\x62\x42\x49\x87\x43\x49\x19\x33\x6f\xe3\x44\x6e\x9b\xab\x8a\x3e\x86\xaa\x99\x52\x1b\x5b\x59\x33\x02\x09\xa0\x21\xa1\x6b\x84\x6b\x66\xbb\xdc\x16\x0c\xd3\x68\xab\xec\x36\x4b\xd8\x60\x8a\x40\x31\x85\x6e\x14\x57\x13\xc2\xfb\x92\x10\xde\xbf\x88\xdc\xbc\x53\x5e\x7f\x82\x7a\x13\xd4\x9b\xa0\xde\x04\xf5\x90\x01\xf5\x94\xcb\x7b\x83\x25\x9e\xd0\xde\x84\xf6\x6a\x5f\x4b\xb3\x98\x00\xdf\x04\xf8\x6c\xbc\x7f\x19\x80\xaf\xf1\x71\x45\x22\x98\x40\xe0\x04\x02\x27\x10\xd8\x29\xf5\x04\x02\xff\x4a\x20\x30\xc1\x72\xf3\x65\x02\x40\xd7\xc1\xd1\xe2\x6b\xf1\xa9\x7b\xfb\xe4\x20\xc0\x68\x9d\xd4\xb4\xd3\x96\xb5\xa6\xd1\x41\x20\xe6\x89\xc3\x48\x65\x58\x13\x84\x9c\x56\x56\x3b\x0c\xe0\x6b\x83\x5c\x13\xd2\x9a\x90\xd6\x84\xb4\x26\xa4\x85\x0c\xa4\x45\x19\xfd\xff\x63\x6c\x52\xb5\x1f\x1e\x19\x74\x3a\xcd\xb9\x69\xce\xa6\x3a\x0f\x7a\x2d\x19\xc7\x81\x14\x5d\xcb\xd5\x03\xc9\x39\xd0\xb0\xd1\xb3\xcd\xfb\x7a\x2d\x5d\x3a\x48\xe1\xfa\x2e\xe6\x81\x42\x18\x86\xd6\xc1\xbe\xb1\x23\xd3\xf7\x34\xed\x19\x0a\x0b\xc4\x48\x44\xfd\x22\x50\xb6\x42\x58\xbb\xe5\x3d\xa7\x73\xd4\x8b\xc4\x8c\x70\x61\xec\x73\xee\xc3\x81\x8b\xf5\xe2\xd7\x52\x3e\xcf\xeb\xeb\x17\x3b\x71\x16\xda\x7d\xb8\xde\xf0\x7a\x8f\x06\x2d\xa7\x40\x7b\xc1\x9d\x41\x4d\xb6\x61\xa2\x4e\x9f\x3d\xa0\xc5\xae\xe3\x1c\x1d\x40\x6c\x48\x8b\x63\xa0\xb5\x01\xed\x8e\x02\xe9\x86\xc8\x3b\x06\xee\xdb\x4b\xde\xbd\xc0\xa1\x6f\xcb\xda\xfc\xc2\x44\x16\x87\x9c\x17\x31\xd3\x30\x20\x39\x42\xcb\x6f\xf2\xf1\xf4\x72\x10\xf8\x1c\xa0\xf3\xbd\x10\xea\x21\x35\x7d\xe8\x86\xdb\x15\xed\x81\x81\x07\x28\xbb\x13\x28\xc7\xf8\xce\x7d\x8d\xc2\x31\xb4\x7e\x94\xd6\xdb\x55\xef\x4a\xfb\xed\xc3\x40\x3e\xeb\x9f\xe9\x99\x0f\xdf\x08\x65\x88\x27\x73\x86\x31\x9d\x47\xdf\x55\x19\xba\x3d\xee\x15\x0a\xcd\x8c\xaa\x5e\xb9\xf6\x57\x33\x73\x5a\xa1\x89\x7b\x3b\xa0\xb2\xa4\xc2\xf6\xc1\x53\xb5\x00\xca\x23\xe5\xf4\x60\x6a\xb4\x2d\x74\xea\x4e\xed\x3b\xe3\x47\xfb\xed\x82\x3d\x19\xd4\x3b\x6b\xaf\xae\x2b\x2f\x57\xb3\x82\x68\xcb\xed\x88\x2e\xe1\x5c\xd7\x26\xfa\x0a\x65\xe7\xce\x11\x33\xb4\xdd\x66\xe3\x37\xf6\xfa\x70\xd6\x4f\xa1\x21\x51\xd8\x3c\x26\x14\x4b\xc6\x87\x44\x27\x1c\x70\xf8\x9e\x46\xce\xab\x21\x07\x5f\xc1\x76\x17\x1b\x77\xb4\xda\x75\xa0\x0a\x3a\x30\xe1\xf8\x97\x32\x16\x2b\x00\x75\x85\xee\x62\x46\xef\xd3\x85\xb5\x6b\x60\xbe\xf2\x30\x7a\x8c\x0b\x4b\xa6\xd0\xf9\x64\x42\xe7\x07\x41\x41\xe3\x2c\x5d\xf9\x6d\xe9\x39\x98\x3b\x3b\x5d\x67\xd4\x5c\xed\xf2\xf0\x48\x7b\xbd\x2d\x31\xdd\x3f\x34\xad\x44\x76\x51\x9a\x56\x22\xa7\x95\xc8\x69\x25\xf2\xe1\x56\x22\x1f\x00\x32\x6a\x73\x92\xed\xe1\xc6\x7d\x9f\x49\x2c\x69\x7e\xc8\x31\x4c\x0c\xb4\xf2\x54\x3b\x79\x3b\x9e\x4d\xb4\xd1\x18\x3e\x5f\x9a\x93\xa2\x11\xc3\xda\x27\x0b\xaf\x37\x2e\x5c\x37\xfb\xeb\x9a\xd6\xc3\xac\xc3\xcc\xf8\x1e\x5b\x9d\xac\x22\x64\xb7\xed\x26\xb8\xf3\xb9\x3c\xbb\x1f\xe2\xb0\x22\x77\x43\x6a\x62\x29\x39\x59\xa6\xe6\xe5\xcd\x7b\x83\xc0\x5b\x8e\x93\x64\xac\xeb\xca\x4f\x65\xac\x4a\xbc\x1e\xcd\x82\xfa\x3c\x70\x36\xb6\xb5\xed\x79\xef\xec\x68\x00\xff\x54\xfa\xb5\xe3\xf1\xdb\xe1\xbe\xce\x76\x17\xaf\x57\xb6\x6b\x89\x05\x09\xce\x52\xb9\x01\x2a\x49\xbe\xd9\xf4\xd2\xb8\x7a\xbf\x91\x02\xf3\x22\x8c\x13\xf2\x77\xd8\x8e\x43\x8b\xe1\x54\x6e\x5e\x9d\xc7\x49\x44\x02\x22\xc7\xa4\x79\x81\x85\xb8\x65\x3c\x1c\x93\xe6\x59\xa2\xf8\x1c\x51\x95\x05\xd9\x20\x00\x21\x7e\x60\x21\x58\xa9\x56\xff\xbe\xb6\x5a\x5e\x5b\x3f\x1f\xd6\xd3\x3c\xc4\x4d\xba\x99\xb4\x63\x6e\x7d\x3e\x3d\x57\xd2\x18\x5f\x47\xe8\xc3\x06\x8a\x68\x6c\x7f\x3b\x72\x0f\xe7\xe2\x77\x77\xf1\xd0\x99\xab\xdf\x2e\xfe\xd6\xbb\xcd\x1a\xb9\x90\xd1\xaf\xf2\x38\x3d\xdb\x74\xf8\xeb\xe3\xda\xe8\x2a\x62\xb7\xda\x1b\x07\xa9\xdc\x30\x5e\xbc\x68\xfb\x6b\x9f\x97\xf1\xc6\xb1\xd8\x5c\x29\x1e\x49\x30\xc5\xf7\xde\xad\x91\x42\xf9\xdd\xed\x89\x80\x25\xbe\x37\xd7\xe7\x32\x5c\xe6\x35\xac\xd4\x0c\x2d\xf7\x90\xc4\xe3\xf5\xe3\x2f\x7f\x54\x18\x88\xe3\x61\x47\x85\x64\x7f\xc0\xd7\x3f\x1a\x92\x42\xe9\xc7\x1e\x0d\x95\x76\xa7\x51\xa0\x8f\x02\x1b\x46\x9e\x06\x42\xd1\xf2\x01\x07\x02\xde\xe9\x7d\x1a\x0b\xa7\x32\x16\xcc\xc0\xee\xc4\x90\xd2\x5f\x6f\x98\x54\x5d\xf2\x95\xe1\xa7\x69\x10\x3a\x06\xe1\x65\xb3\x17\x47\x58\x78\xd0\x45\xd6\x5b\xd5\x5f\x25\x1d\x71\x49\xa6\x7a\x64\xda\xd0\x6f\xc7\x3a\x4c\xe3\x09\xc0\x6e\x96\x2c\xa7\xa7\x77\x34\x10\x05\x08\x21\x44\x92\x65\x77\xdf\x20\x5c\xbc\xe7\x97\x3f\xf4\x1a\x45\xd6\xe7\x27\x4a\xde\x74\x27\x66\x11\x7d\x70\xba\xd3\x78\xf9\x1e\x0d\xca\xc8\x39\xde\x7c\xb3\xa6\xe1\xbc\xd7\xc1\x6a\x6f\xb3\x0e\x52\xbe\xe4\x98\x8a\x15\x70\x94\x70\x26\x59\xc0\xa2\xf2\x1c\xfb\xd9\xc5\xf9\xbc\xd5\x92\x9c\xa3\xdf\xe6\x1e\xb3\x0d\x49\xba\x87\x50\x5f\x84\xfe\xe9\xd6\xf8\xbb\xe6\xf0\x7a\xeb\xa6\x65\x3b\x86\x8b\x79\x93\xf5\x59\x20\x6e\xb4\xa7\x44\xf4\x3f\xa5\xfe\x67\x42\x12\xdb\xd3\xe7\xbb\xa5\xa3\x8c\x5c\x2b\x97\xbb\xbb\x7f\x8e\xc5\x6e\xed\x43\x5c\xbf\x74\xc8\x8f\xff\xe6\xd6\xbe\x91\xb6\xf5\x95\xe4\xed\x93\xc4\xa8\x5b\xf9\x76\x4d\x35\xb7\xd8\x8c\xb6\x7d\xaf\x72\xe0\xb6\xbd\x01\x63\x9e\x76\xab\x1a\x32\x76\xe4\x8c\x76\xc2\xad\x6c\xa2\x65\xf7\xcf\xf8\xa7\xda\x2a\xb9\x8c\x3d\x3c\xa3\x9d\x64\x33\xe5\x1a\xb5\x2d\xfb\x86\xa2\x5a\x7f\x19\x5b\x7f\xc6\x3f\xd1\x53\xd3\xe2\x41\x5b\xd3\x4f\xf0\xec\xb0\x42\x73\x43\xd2\x68\x27\xd3\x6a\x6a\x34\xf6\x4e\x1e\x52\x8b\x87\x6c\xcc\xae\x44\xfb\x9e\xa7\x51\x4f\x9d\x55\x03\x81\x8e\x67\xfc\xb4\x69\xf0\x3a\x18\xf2\x40\xd0\xf6\xa8\x34\xe3\xc9\x98\xaf\xf6\xda\x24\xd3\xeb\x60\xb9\x0e\xd3\x1f\xa9\xff\xee\x1f\xfd\x37\x00\x00\xff\xff\x69\x5d\x0a\x6a\x39\x9d\x00\x00")

func v2SchemaJSONBytes() ([]byte, error) {
	return bindataRead(
		_v2SchemaJSON,
		"v2/schema.json",
	)
}

func v2SchemaJSON() (*asset, error) {
	bytes, err := v2SchemaJSONBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v2/schema.json", size: 40249, mode: os.FileMode(0644), modTime: time.Unix(1567900649, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xcb, 0x25, 0x27, 0xe8, 0x46, 0xae, 0x22, 0xc4, 0xf4, 0x8b, 0x1, 0x32, 0x4d, 0x1f, 0xf8, 0xdf, 0x75, 0x15, 0xc8, 0x2d, 0xc7, 0xed, 0xe, 0x7e, 0x0, 0x75, 0xc0, 0xf9, 0xd2, 0x1f, 0x75, 0x57}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"jsonschema-draft-04.json": jsonschemaDraft04JSON,

	"v2/schema.json": v2SchemaJSON,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"jsonschema-draft-04.json": &bintree{jsonschemaDraft04JSON, map[string]*bintree{}},
	"v2": &bintree{nil, map[string]*bintree{
		"schema.json": &bintree{v2SchemaJSON, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
