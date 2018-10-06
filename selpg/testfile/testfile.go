package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	fout1, _ := os.Create("C:\\Users\\o\\Desktop\\test1.txt")
	writer1 := bufio.NewWriter(fout1)
	for i := 0; i < 20; i++ {
		writer1.WriteString("line" + strconv.Itoa(i) + "\f")
	}
	fout2, _ := os.Create("C:\\Users\\o\\Desktop\\test2.txt")
	writer2 := bufio.NewWriter(fout2)
	for i := 0; i < 200; i++ {
		writer2.WriteString("line" + strconv.Itoa(i) + "\n")
	}
	writer1.Flush()
	writer2.Flush()
}
