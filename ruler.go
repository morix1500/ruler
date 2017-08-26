package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"github.com/mattn/go-runewidth"
	"github.com/morix1500/go-ltsv2tsv"
	"io"
	"os"
)

var (
	file_type  string
	text_align string
	headerless bool
	padding    func(string, int, string) string
)

type CLI struct {
	outStream, errStream io.Writer
}

func (c *CLI) readCsv(fp io.Reader) int {
	reader := csv.NewReader(fp)
	reader.Comma = ','
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintln(c.errStream, "csv parse error.")
		return 1
	}
	err = c.ruler(records)
	if err != nil {
		fmt.Fprintln(c.errStream, err)
		return 1
	}

	return 0
}

func (c *CLI) readTsv(fp io.Reader) int {
	reader := csv.NewReader(fp)
	reader.Comma = '\t'
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintln(c.errStream, "tsv parse error.")
		return 1
	}
	err = c.ruler(records)
	if err != nil {
		fmt.Fprintln(c.errStream, err)
		return 1
	}

	return 0
}

func (c *CLI) readLtsv(fp io.Reader) int {
	reader := ltsv2tsv.NewConverter(fp)
	records, err := reader.Converter()
	if err != nil {
		fmt.Fprintln(c.errStream, "ltsv parse error.")
		return 1
	}
	err = c.ruler(records)
	if err != nil {
		fmt.Fprintln(c.errStream, err)
		return 1
	}

	return 0
}

func times(str string, n int) (out string) {
	for i := 0; i < n; i++ {
		out += str
	}
	return
}

func padRight(str string, length int, pad string) string {
	return str + times(pad, length-runewidth.StringWidth(str))
}

func padLeft(str string, length int, pad string) string {
	return times(pad, length-runewidth.StringWidth(str)) + str
}

func createHeader(header_line []string, len_arr []int) (brank_line string, head_line string) {
	for i, column := range header_line {
		brank_line += "+-" + padding("", len_arr[i], "-")
		head_line += "| " + padding(column, len_arr[i], " ")

		if i == len(header_line)-1 {
			brank_line += "-+"
			head_line += " |"
		} else {
			brank_line += "-"
			head_line += " "
		}
	}

	return
}

func (c *CLI) ruler(lines [][]string) error {
	if len(lines) <= 0 {
		return errors.New("empty lines.")
	}
	len_arr := make([]int, len(lines[0]))

	// count max length
	for _, line := range lines {
		for j, column := range line {
			if len_arr[j] < runewidth.StringWidth(column) {
				len_arr[j] = runewidth.StringWidth(column)
			}
		}
	}

	var brank_line string
	var head_line string

	// create rule
	for i, line := range lines {
		if i == 0 {
			brank_line, head_line = createHeader(line, len_arr)
			if !headerless {
				fmt.Fprintln(c.outStream, brank_line)
				fmt.Fprintln(c.outStream, head_line)
			}
			fmt.Fprintln(c.outStream, brank_line)
			continue
		}

		for j, column := range line {
			cell := "| " + padding(column, len_arr[j], " ")
			if j == len(line)-1 {
				cell += " |"
			} else {
				cell += " "
			}
			fmt.Fprint(c.outStream, cell)
		}
		fmt.Fprintln(c.outStream, "")
		// last line
		if i == len(lines) -1 {
			fmt.Fprintln(c.outStream, brank_line)
		}
	}

	return nil
}

func (c *CLI) Run(args []string) int {
	flags := flag.NewFlagSet("ruler", flag.ContinueOnError)
	flags.SetOutput(c.errStream)
	flags.StringVar(&file_type, "f", "csv", "Specify the format. [csv/tsv/ltsv]")
	flags.StringVar(&text_align, "a", "right", "Specify the text align. [right/left]")
	flags.BoolVar(&headerless, "n", false, "Specify when there is no header.")

	if err := flags.Parse(args[1:]); err != nil {
		return 1
	}

	if text_align == "right" {
		padding = padRight
	} else if text_align == "left" {
		padding = padLeft
	} else {
		return 1
	}

	fp := os.Stdin

	switch file_type {
	case "csv":
		return c.readCsv(fp)
	case "tsv":
		return c.readTsv(fp)
	case "ltsv":
		return c.readLtsv(fp)
	default:
		flag.PrintDefaults()
		return 1
	}
}

func main() {
	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}
