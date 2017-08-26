package main

import (
	"bytes"
	"fmt"
	"testing"
	"strings"
	"github.com/nak1114/goutil/assert"
)

func ExampleCsv() {
	test_record := `column1,column2,column3,column4
1,2,3,4
5,6,,7
8,,,9
`
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}

	args := strings.Split("ruler", " ")

	assert.StubIO(test_record, func() {
		cli.Run(args)
	})
	fmt.Println(outStream.String())
	// Output: +---------+---------+---------+---------+
	// | column1 | column2 | column3 | column4 |
	// +---------+---------+---------+---------+
	// | 1       | 2       | 3       | 4       |
	// | 5       | 6       |         | 7       |
	// | 8       |         |         | 9       |
	// +---------+---------+---------+---------+
}

func ExampleTsv() {
	test_record := `column1	column2	column3	column4
1	2	3	4
5	6		7
8			9
`
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}

	args := strings.Split("ruler -f tsv", " ")

	assert.StubIO(test_record, func() {
		cli.Run(args)
	})
	fmt.Println(outStream.String())
	// Output: +---------+---------+---------+---------+
	// | column1 | column2 | column3 | column4 |
	// +---------+---------+---------+---------+
	// | 1       | 2       | 3       | 4       |
	// | 5       | 6       |         | 7       |
	// | 8       |         |         | 9       |
	// +---------+---------+---------+---------+
}

func ExampleLtsv() {
	test_record := `column1:1	column2:2	column3:3	column4:4
column1:5	column2:6	column4:7
column1:8	column4:9
column5:10
`
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}

	args := strings.Split("ruler -f ltsv", " ")

	assert.StubIO(test_record, func() {
		cli.Run(args)
	})
	fmt.Println(outStream.String())
	// Output: +---------+---------+---------+---------+---------+
	// | column1 | column2 | column3 | column4 | column5 |
	// +---------+---------+---------+---------+---------+
	// | 1       | 2       | 3       | 4       |         |
	// | 5       | 6       |         | 7       |         |
	// | 8       |         |         | 9       |         |
	// |         |         |         |         | 10      |
	// +---------+---------+---------+---------+---------+
}

func TestTextAlign(t *testing.T) {
	test_record := `column1,column2,column3,column4
1,2,3,4
5,6,,7
8,,,9
`
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}

	args := strings.Split("ruler -a left", " ")

	var status int

	assert.StubIO(test_record, func() {
		status = cli.Run(args)
	})

	if status != 0 {
		t.Errorf("status should be %d, but %d", 0, status)
	}

	want :=  `+---------+---------+---------+---------+
| column1 | column2 | column3 | column4 |
+---------+---------+---------+---------+
|       1 |       2 |       3 |       4 |
|       5 |       6 |         |       7 |
|       8 |         |         |       9 |
+---------+---------+---------+---------+
`
	result := outStream.String()
	if result != want {
		t.Errorf("output should be \n%s, but \n%s", want, result)
	}
}

func TestHeaderless(t *testing.T) {
	test_record := `1,2,3,4
5,6,,7
8,,,9
`
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}

	args := strings.Split("ruler -n", " ")

	var status int

	assert.StubIO(test_record, func() {
		status = cli.Run(args)
	})

	if status != 0 {
		t.Errorf("status should be %d, but %d", 0, status)
	}

	want :=  `+---+---+---+---+
| 1 | 2 | 3 | 4 |
| 5 | 6 |   | 7 |
| 8 |   |   | 9 |
+---+---+---+---+
`
	result := outStream.String()
	if result != want {
		t.Errorf("output should be \n%s, but \n%s", want, result)
	}
}

func TestMultiByte(t *testing.T) {
	test_record := `カラム1,カラム2,カラム3,カラム4
りんご,ゴリラ,ラッパ,パンダ
だるま,マーマレード,,ドリル
ルーラ,,,ラーメン
`
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}

	args := strings.Split("ruler", " ")

	var status int

	assert.StubIO(test_record, func() {
		status = cli.Run(args)
	})

	if status != 0 {
		t.Errorf("status should be %d, but %d", 0, status)
	}

	want :=  `+---------+--------------+---------+----------+
| カラム1 | カラム2      | カラム3 | カラム4  |
+---------+--------------+---------+----------+
| りんご  | ゴリラ       | ラッパ  | パンダ   |
| だるま  | マーマレード |         | ドリル   |
| ルーラ  |              |         | ラーメン |
+---------+--------------+---------+----------+
`
	result := outStream.String()
	if result != want {
		t.Errorf("output should be \n%s, but \n%s", want, result)
	}
}
