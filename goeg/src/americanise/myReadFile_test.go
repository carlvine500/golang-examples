package main

import (
	"bufio"
	"fmt"
	regexp2 "github.com/dlclark/regexp2"
	"io"
	"os"
	"regexp"
	"testing"
)

func TestName(t *testing.T) {
	file, err := os.Open("input.txt")
	if err!=nil{
		panic(any(err))
	}
	reader := bufio.NewReader(file)
	for{
		lineWithNext, err := reader.ReadString('\n')
		if err==io.EOF{
			break
		}
		if err!=nil{
			panic(any(err))
		}
		fmt.Print(lineWithNext)// 每行包含着换行符
	}
}

func Test2(t *testing.T) {
	compile := regexp.MustCompile("[A-Za-z]+")
	submatch := compile.FindAllString("hello world",-1)
	for _,v:=range submatch{
		fmt.Println(v)
	}
}

func Test12(t *testing.T) {
	compile := regexp.MustCompile(`\[([A-Za-z]+)\]`)
	submatch := compile.FindAllString("[hello]1[world]a",-1)
	for _,v:=range submatch{
		fmt.Println(v)
	}
}

func Test3(t *testing.T) {
	str := "example123text456"
	// 正向预查需要使用开源包regexp2才行:定义正则表达式，包含一个正向预查来匹配数字
	regex := regexp2.MustCompile(`\d+(?=\D)`,0)
	// 查找所有匹配
	matches, _ := regex.FindStringMatch(str)

	for _, match := range matches.Groups() {
		fmt.Println(match.String())
	}
}