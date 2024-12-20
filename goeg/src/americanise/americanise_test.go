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
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
)


	var inFilename       = "input.txt"
	var expectedFilename = "expected.txt"
	var actualFilename   = "actual.txt"

func init() {
	workDir,_ := os.Getwd()
	fmt.Println("init workDir="+workDir)
	inFilename = filepath.Join(workDir,inFilename)
	expectedFilename = filepath.Join(workDir,expectedFilename)
	actualFilename = filepath.Join(workDir,actualFilename)
}

func list(){
	getwd, err2 := os.Getwd()
	fmt.Println(getwd,err2)
	path := "./"

	// 读取目录
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	// 遍历目录条目
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			fmt.Println("Error getting info for", entry.Name(), ":", err)
			continue
		}

		// 输出文件名和类型
		fmt.Printf("Found file: %s, type: %v\n", entry.Name(), info.Mode())
	}
}
func TestAmericanize(t *testing.T) {
	list()
	log.SetFlags(0)
	log.Println("TEST americanize")

	//path, _ := filepath.Split(os.Args[0])
	var inFile, outFile *os.File
	var err error

	//inFilename := filepath.Join(path, inFilename)
	if inFile, err = os.Open(inFilename); err != nil {
		t.Fatal(err)
	}
	defer inFile.Close()

	//outFilename := filepath.Join(path, actualFilename)
	if outFile, err = os.Create(actualFilename); err != nil {
		t.Fatal(err)
	}
	defer outFile.Close()
	defer os.Remove(actualFilename)

	if err := americanise(inFile, outFile); err != nil {
		t.Fatal(err)
	}

	compare(actualFilename, expectedFilename, t)
}

func compare(actual, expected string, t *testing.T) {

	if actualBytes, err := ioutil.ReadFile(actual); err != nil {
		t.Fatal(err)
	} else if expectedBytes, err := ioutil.ReadFile(expected); err != nil {
		t.Fatal(err)
	} else {
		if bytes.Compare(actualBytes, expectedBytes) != 0 {
			t.Fatal("actual != expected")
		}
	}
}
