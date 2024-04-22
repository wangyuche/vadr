package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var tgfile []rune
var cmfile []rune
var base []string = make([]string, 0)

func main() {
	readFile, err := os.Open("./base.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		base = append(base, fileScanner.Text())
	}
	readFile.Close()

	if len(os.Args) != 2 {
		panic("檔案數量不對")
	}
	dat, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	str := strings.Replace(string(dat), "\n", "", -1)
	tgfile = []rune(str)
	fmt.Println(len([]rune(base[0])))
	for i := 0; i < len(base); i++ {
		temp := []rune(base[i])
		startindex := 0
		endindex := 0
		//fmt.Println(i)
		/*
			for {
				index := indexAt(string(tgfile), string(temp[:10]), startindex+10)
				if index == -1 {
					break
				}
				startindex = index
			}
			for {
				index := indexAt(string(tgfile), string(temp[len(temp)-10:]), endindex+10)
				if index == -1 {
					break
				}
				endindex = index
			}*/
		//fmt.Printf(string(temp[len(temp)-10:]))
		//endindex = strings.Index(string(tgfile), string(temp[len(temp)-10:]))
		//fmt.Printf("開始位置 %d\n結束位置 %d\n結果 %s\n", startindex+1, endindex+10, "")
		/*
			for {
				startindex = indexAt(string(tgfile), string(temp[:10]), startindex+10)
				if startindex == -1 {
					break
				}
				fmt.Printf("開始位置 %d\n結束位置 %d\n結果 %s\n", startindex+1, startindex+len([]byte(base[i])), "")
			}
		*/

		startindex = strings.Index(string(tgfile), string(temp[:10]))
		endindex = strings.Index(string(tgfile), string(temp[len(temp)-10:]))
		if startindex != -1 {
			fmt.Printf("開始位置 %d\n結束位置 %d\n結果 %s\n", startindex+1, endindex+10, "")
		}
	}
}

func indexAt(s, sep string, n int) int {
	idx := strings.Index(s[n:], sep)
	if idx > -1 {
		idx += n
	}
	return idx
}
