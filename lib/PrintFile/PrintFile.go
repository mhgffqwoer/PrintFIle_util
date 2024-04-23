package PrintFile

import (
	"fmt"
	"io"
	"os"
)

func PrintFile(path string, tail bool, delimiter string, lines int64) {
    if path == "" {
        fmt.Println("Error: specify the file!")
        return
    }
    file, err := os.Open(path)
    if err != nil {
        fmt.Println("Error: file not found!")
        return
    }
    defer file.Close()

	readAll := false
	if lines == 0 {
		readAll = true
	}

    if !tail {
        ch := make([]byte, 1)
        for lines > 0 || readAll {
            _, err := file.Read(ch)
            if err != nil {
                break
            }
            if ch[0] == []byte(delimiter)[0] {
                lines--
            }
            fmt.Print(string(ch[0]))
        }
    } else {
        start := int64(0)
        pos, _ := file.Seek(-1, io.SeekEnd)
        i := int64(1)
		ch := make([]byte, 1)
        for (i < pos && lines > 0) || readAll {
            file.Seek(-i, io.SeekEnd)
            _, err := file.Read(ch)
            if err != nil {
                break
            }
            if ch[0] == []byte(delimiter)[0] {
                start = -i
                lines--
            }
            i++
        }

        file.Seek(start, io.SeekEnd)
        io.Copy(os.Stdout, file)
    }
	fmt.Print("\n")
}