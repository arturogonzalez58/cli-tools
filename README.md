# Command line tools
This repository contains a set of useful command line applications based on the book [Powerful Command-Line Applications in go](https://pragprog.com/titles/rggo/powerful-command-line-applications-in-go/).
Tools:
- WordCounter
## Build
Run the command
```shell
make
```
### WordCounter

#### Usages
- Count all the words in the input
```shell
echo "This is a test" | wc
```
Expected Output:
```shell
4
```
- Count the lines in the input
```shell
echo "This is a test" | wc -l
```
Expected Output:
```shell
1
```
- Count the number of times a regular expression matches in the input
```shell
echo "This is a test" | wc -r test
```
Expected Output:
```shell
1
```