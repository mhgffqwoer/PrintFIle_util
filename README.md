# PrintFIle_util

> _This is my PrintFile utility which can print target file to the command line. It is my first university project written in C++, but rewritten by Golang._

## Commands

`$ PrintFile [commands]`

- _--file [string]_, _-f [string]_ - Path to file. [required parameter]
- _--lines [int64]_, _-l [int64]_ - number of lines to print.
- _--tail_, _-t_ - tail mode.
- _--delimiter [string]_ _-d [string]_ - Character that will be accepted as the end of the line.
- _--help_, _-h_ - Help menu.

## Example

Inside the project directory:
`$ PrintFile --file ./main.go --tail --lines 3`

Will be displayed
```
func main() {
        PrintFile.Execute()
}
```
