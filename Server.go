package main

import (
	"fmt"
	"os"
)

var tgfile []byte
var cmfile []byte

func main() {
	files := os.Args[1:]
	if len(files) != 2 {
		panic("檔案數量不對")
	}
	dat, err := os.ReadFile(files[0])
	if err != nil {
		panic(err)
	}
	tgfile = dat
	dat, err = os.ReadFile(files[1])
	if err != nil {
		panic(err)
	}
	cmfile = dat
	var cm, tg, lenght, tgoff int
	cm = 0
	tg = 0
	tgoff = 0
	for cm < len(cmfile) {
		for tg < len(tgfile) {
			if cmfile[cm] == tgfile[tg] {
				lenght = 1
				break
			}
			tg++
		}
		if lenght == 0 {
			tg = tgoff
			cm++
		}
		if lenght != 0 && cm+lenght < len(cmfile) && tg+lenght < len(tgfile) {
			for cmfile[cm+lenght] == tgfile[tg+lenght] {
				lenght++
				if cm+lenght >= len(cmfile) || tg+lenght >= len(tgfile) {
					break
				}
			}
			if lenght > 10 {
				fmt.Printf("開始位置 %d\n結束位置 %d\n結果 %s\n", cm+1, cm+lenght, string(cmfile[cm:cm+lenght]))
				cm = cm + lenght
				tgoff = tg + lenght
				tg = tg + lenght
				lenght = 0
			} else if lenght == 1 {
				lenght = 0
				cm++
			} else {
				cm = cm + lenght
				lenght = 0
			}

		}
	}
}
