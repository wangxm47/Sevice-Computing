package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/pflag"
)

type selpgArgs struct {
	start       int
	end         int
	fileName    string
	destination string
	lineNum     int
	delimiter   bool
}

func initArgs(args *selpgArgs) {
	pflag.IntVarP(&args.start, "StartPage", "s", 0, "sepcify start page greater than 0")
	pflag.IntVarP(&args.end, "EndPage", "e", 0, "sepcify end page greater than 0")
	pflag.IntVarP(&args.lineNum, "LineNumber", "l", 72, "specify the length of a page greater than 0")
	pflag.StringVarP(&args.destination, "Destination", "d", "", "specify the name of printer")
	pflag.BoolVarP(&args.delimiter, "FileDelimiter", "f", false, "default lines-delimited, 'f' for form-feed-delimited")
}
func Checkargs(args *selpgArgs) {
	pageok := args.start <= args.end && args.start > 0
	lineok := args.lineNum > 0
	if !pageok || !lineok {
		pflag.Usage()
		os.Exit(1)
	}
}
func ExecArgs(args *selpgArgs) {
	fin := os.Stdin
	fout := os.Stdout
	var err error
	if args.fileName != "" {
		fin, err = os.Open(args.fileName)
		if err != nil {
			fmt.Fprintf(os.Stdout, "selpg: can not open the file\"%s\"\n", args.fileName)
			fmt.Println(err)
			os.Exit(1)
		}
		defer fin.Close()
	}
	var inpipe io.WriteCloser
	if args.destination != "" {
		cmd := exec.Command("lp", "-d", args.destination)
		inpipe, err = cmd.StdinPipe()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer inpipe.Close()
		cmd.Stdout = fout
		cmd.Start()
	}
	page, line := 1, 0
	if args.delimiter {
		preader := bufio.NewReader(fin)
		for page <= args.end {
			pdata, perr := preader.ReadString('\f')
			if perr != nil || perr == io.EOF {
				if perr == io.EOF {
					fmt.Fprintf(fout, "%s", pdata)
					if page <= args.end {
						fmt.Println("selpg: it's end of file, \"-e\" is greater than the total pages")
					}
					break
				}
			}
			pdata = strings.Replace(pdata, "\f", "\n", -1)
			if page >= args.start {
				fmt.Fprintf(fout, "%s", pdata)
			}
			page++
		}
	} else {
		lreader := bufio.NewReader(fin)
		for page <= args.end {
			line++
			ldata, lerr := lreader.ReadString('\n')
			if lerr != nil || lerr == io.EOF {
				if lerr == io.EOF {
					fmt.Fprintf(fout, "%s", ldata)
					if page <= args.end {
						fmt.Println("selpg: it's end of file, \"-e\" is greater than the total pages")
					} else if line < args.lineNum-1 {
						fmt.Println("selpg: it's end of file, the final page is not full")
					}
					break
				}
			}
			if page >= args.start {
				fmt.Fprintf(fout, "%s", ldata)
			}
			if line%args.lineNum == 0 {
				line = 0
				page++
			}
		}
	}
}
func main() {
	args := selpgArgs{}
	initArgs(&args)
	pflag.Parse()
	if pflag.NArg() == 1 {
		args.fileName = pflag.Arg(0)
	} else {
		args.fileName = ""
	}
	Checkargs(&args)
	ExecArgs(&args)
}
