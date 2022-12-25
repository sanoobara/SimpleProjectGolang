package file

import (
	"fmt"
	"io"
	"log"
	"os"

	"path/filepath"
)

func OpenFile(pathfile string) []byte {
	file, err := os.Open(pathfile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	data := make([]byte, 256)

	for {
		n, err := file.Read(data)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		fmt.Print(string(data[:n]))
	}

	return data
}

type s struct {
	data []byte
	name string
}

func Dir(pathdir string) {

	var rows = make([]s, 2, 300)
	i := 0

	err := filepath.Walk(pathdir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			extension := filepath.Ext(path)

			if !(info.IsDir()) && extension == ".dst" {
				rows[i].data = OpenFile(path)
				rows[i].name = info.Name()
				i++
			}
			//fmt.Println(path, info.Size())

			return nil
		})

	for _, f := range rows {
		fmt.Println(f.name)
	}

	if err != nil {
		log.Println(err)
	}
}
