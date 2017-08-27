# Ruler
Ruler shows table from csv and ltsv and so on.  
Please look at the Example and Usage.

## Example
```
$ cat hoge.csv
column1,column2,column3,column4
1,2,3,4
5,6,,7
8,,,9

$ cat hoge.csv | ruler
+---------+---------+---------+---------+
| column1 | column2 | column3 | column4 |
+---------+---------+---------+---------+
| 1       | 2       | 3       | 4       |
| 5       | 6       |         | 7       |
| 8       |         |         | 9       |
+---------+---------+---------+---------+

```

## Usage
```
Usage of ruler:
  -a string
        Specify the text align. [right/left] (default "left")
  -f string
        Specify the format. [csv/tsv/ltsv] (default "csv")
  -n    Specify when there is no header.
  -v    Output version number.
```

### Specify the text align
Default "left".

```
$ cat hoge.csv | ruler -a right
+---------+---------+---------+---------+
| column1 | column2 | column3 | column4 |
+---------+---------+---------+---------+
|      1  |       2 |       3 |       4 |
|      5  |       6 |         |       7 |
|      8  |         |         |       9 |
+---------+---------+---------+---------+
```

### Specify then format
The choise is:
* CSV
* TSV
* LTSV

Default "csv".

```
$ cat hoge.log 
time:30/Nov/2016:00:55:08 +0900 host:xxx.xxx.xxx.xxx    forwardedfor:-  req:GET /v1/xxx HTTP/1.1        status:200
time:30/Nov/2016:00:56:37 +0900 host:xxx.xxx.xxx.xxx    forwardedfor:-  req:GET /v1/yyy HTTP/1.1        status:200	size:123

$ cat hoge.log | ruler -f ltsv
+----------------------------+-----------------+--------------+----------------------+--------+------+
| time                       | host            | forwardedfor | req                  | status | size |
+----------------------------+-----------------+--------------+----------------------+--------+------+
| 30/Nov/2016:00:55:08 +0900 | xxx.xxx.xxx.xxx | -            | GET /v1/xxx HTTP/1.1 | 200    |      |
| 30/Nov/2016:00:56:37 +0900 | xxx.xxx.xxx.xxx | -            | GET /v1/yyy HTTP/1.1 | 200    | 123  |
+----------------------------+-----------------+--------------+----------------------+--------+------+
```

### Headerless
header skip mode.

Default "false"

```
$ cat hoge.csv
1,2,3,4
11,22,33,44
111,222,333,444

$ cat hoge.csv | ruler -n
+-----+-----+-----+-----+
| 1   | 2   | 3   | 4   |
| 11  | 22  | 33  | 44  |
| 111 | 222 | 333 | 444 |
+-----+-----+-----+-----+
```

## Installation
Executable binaries are available at [releases](https://github.com/morix1500/go-ruler/releases).

```
$ wget https://github.com/morix1500/go-ruler/releases/download/v0.1.1/ruler_linux_amd64 -O ruler
$ chmod a+x ruler
```

## License
Please see the [LICENSE](./LICENSE) file for details.  

## Author
Shota Omori(Morix)  
https://github.com/morix1500
