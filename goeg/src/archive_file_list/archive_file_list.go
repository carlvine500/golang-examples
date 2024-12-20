// Copyright © 2011-12 Qtrac Ltd.
// 
// This program or package and any associated files are licensed under the
// Apache License, Version 2.0 (the "License"); you may not use these files
// except in compliance with the License. You can get a copy of the License
// at: http://www.apache.org/licenses/LICENSE-2.0.
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
    "archive/tar"
    "archive/zip"
    "compress/gzip"
    "errors"
    "fmt"
    "io"
    "os"
    "path/filepath"
    "runtime"
    "strings"
)

var FunctionForSuffix = map[string]func(string) ([]string, error){
    ".gz": GzipFileList, ".tar": TarFileList, ".tar.gz": TarFileList,
    ".tgz": TarFileList, ".zip": ZipFileList}

func main() {
    if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
        fmt.Printf("usage: %s archive1 [archive2 [... archiveN]]\n",
            filepath.Base(os.Args[0]))
        os.Exit(1)

    }
    args := commandLineFiles(os.Args[1:])
    archiveFileList := ArchiveFileList
    if len(args[0]) == 1 && strings.IndexAny(args[0], "12345") != -1 {
        which := args[0][0]
        args = args[1:]
        switch which {
        case '2':
            archiveFileList = ArchiveFileList2
        case '3':
            archiveFileList = ArchiveFileList3
        case '4':
            archiveFileList = ArchiveFileList4
        case '5':
            archiveFileList = ArchiveFileListMap
        }
    }
    for _, filename := range args {
        fmt.Print(filename)
        lines, err := archiveFileList(filename)
        if err != nil {
            fmt.Println(" ERROR:", err)
        } else {
            fmt.Println()
            for _, line := range lines {
                fmt.Println(" ", line)
            }
        }
    }
}

func commandLineFiles(files []string) []string {
    if runtime.GOOS == "windows" {
        args := make([]string, 0, len(files))
        for _, name := range files {
            if matches, err := filepath.Glob(name); err != nil {
                args = append(args, name) // Invalid pattern
            } else if matches != nil { // At least one match
                args = append(args, matches...)
            }
        }
        return args
    }
    return files
}

func ArchiveFileList(file string) ([]string, error) {
    if suffix := Suffix(file); suffix == ".gz" {
        return GzipFileList(file)
    } else if suffix == ".tar" || suffix == ".tar.gz" || suffix == ".tgz" {
        return TarFileList(file)
    } else if suffix == ".zip" {
        return ZipFileList(file)
    }
    return nil, errors.New("unrecognized archive")
}

func ArchiveFileList2(file string) ([]string, error) {
    switch suffix := Suffix(file); suffix { // switch可以带计算表达式
    case ".gz":
        return GzipFileList(file)
    case ".tar":
        fallthrough//往后找方法
    case ".tar.gz":
        fallthrough//往后找方法
    case ".tgz":
        return TarFileList(file)
    case ".zip":
        return ZipFileList(file)
    }
    return nil, errors.New("unrecognized archive")
}

func ArchiveFileList3(file string) ([]string, error) {
    switch Suffix(file) {
    case ".gz":
        return GzipFileList(file)
    case ".tar":
        fallthrough
    case ".tar.gz":
        fallthrough
    case ".tgz":
        return TarFileList(file)
    case ".zip":
        return ZipFileList(file)
    }
    return nil, errors.New("unrecognized archive")
}

func ArchiveFileList4(file string) ([]string, error) {
    switch Suffix(file) { // Canonical
    case ".gz":
        return GzipFileList(file)
    case ".tar", ".tar.gz", ".tgz": // switch可以合并写在一起
        return TarFileList(file)
    case ".zip":
        return ZipFileList(file)
    }
    return nil, errors.New("unrecognized archive")
}

func ArchiveFileListMap(file string) ([]string, error) {
    if function, ok := FunctionForSuffix[Suffix(file)]; ok {
        return function(file)
    }
    return nil, errors.New("unrecognized archive")
}

func Suffix(file string) string {
    file = strings.ToLower(filepath.Base(file))
    if i := strings.LastIndex(file, "."); i > -1 {
        if file[i:] == ".bz2" || file[i:] == ".gz" || file[i:] == ".xz" {
            if j := strings.LastIndex(file[:i], ".");
                j > -1 && strings.HasPrefix(file[j:], ".tar") {
                return file[j:]
            }
        }
        return file[i:]
    }
    return file
}

func ZipFileList(filename string) ([]string, error) {
    zipReader, err := zip.OpenReader(filename)
    if err != nil {
        return nil, err
    }
    defer zipReader.Close()
    var files []string
    for _, file := range zipReader.File {
        files = append(files, file.Name)
    }
    return files, nil
}

func GzipFileList(filename string) ([]string, error) {
    reader, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer reader.Close()
    gzipReader, err := gzip.NewReader(reader)
    if err != nil {
        return nil, err
    }
    return []string{gzipReader.Header.Name}, nil
}

func TarFileList(filename string) ([]string, error) {
    reader, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer reader.Close()
    var tarReader *tar.Reader
    if strings.HasSuffix(filename, ".gz") ||
        strings.HasSuffix(filename, ".tgz") {
        gzipReader, err := gzip.NewReader(reader)
        if err != nil {
            return nil, err
        }
        tarReader = tar.NewReader(gzipReader)
    } else {
        tarReader = tar.NewReader(reader)
    }
    var files []string
    for {
        header, err := tarReader.Next()
        if err != nil {
            if err == io.EOF {
                break
            }
            return files, err
        }
        if header == nil {
            break
        }
        files = append(files, header.Name)
    }
    return files, nil
}
