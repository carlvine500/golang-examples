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
    "bufio"
    "fmt"
    "github.com/carlvine500/golang-examples/goeg/src/safemap"
    "io"
    "log"
    "os"
    "path/filepath"
    "regexp"
    "runtime"
)

var ncpu = runtime.NumCPU()

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU()) // Use all the machine's cores
    if len(os.Args) != 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
        fmt.Printf("usage: %s <file.log>\n", filepath.Base(os.Args[0]))
        os.Exit(1)
    }
    lines := make(chan string, ncpu*4)
    done := make(chan struct{}, ncpu)
    pageMap := safemap.New()
    // 1协程读文件到lines
    go readLines(os.Args[1], lines)
    // 4协程正则提取每行内容到map
    processLines(done, pageMap, lines)
    // 4协程对应的4chan关闭
    waitUntil(done)
    showResults(pageMap)
}



func readLines(filename string, lines chan<- string) {
    file, err := os.Open(filename)
    if err != nil {
        log.Fatal("failed to open the file:", err)
    }
    defer file.Close()
    reader := bufio.NewReader(file)
    for {
        line, err := reader.ReadString('\n')
        if line != "" {
            lines <- line
        }
        if err != nil {
            if err != io.EOF {
                log.Println("failed to finish reading the file:", err)
            }
            break
        }
    }
    close(lines)
}

func processLines(done chan<- struct{}, pageMap safemap.SafeMap,
    lines <-chan string) {
    getRx := regexp.MustCompile(`GET[ \t]+([^ \t\n]+[.]html?)`)
    incrementer := func(value interface{}, found bool) interface{} {
        if found {
            return value.(int) + 1
        }
        return 1
    }
    for i := 0; i < ncpu; i++ {
        go func() {
            for line := range lines {
                if matches := getRx.FindStringSubmatch(line);
                    matches != nil {
                    pageMap.Update(matches[1], incrementer)
                }
            }
            done <- struct{}{}
        }()
    }
}

func waitUntil(done <-chan struct{}) {
    for i := 0; i < ncpu; i++ {
        <-done
    }
}

func showResults(pageMap safemap.SafeMap) {
    pages := pageMap.Close()
    for page, count := range pages {
        fmt.Printf("%8d %s\n", count, page)
    }
}
