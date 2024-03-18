package main

import (
	"fmt"
	"os"
	"path"
)

func maintr(strrr string) string {
	newstr := ""
	for _, cchar := range strrr {
		ichar := rune(cchar)
		if ichar < 100000 || string(cchar) == "\\" {
			newstr = newstr + string(cchar)
		}

	}
	return newstr
}

func mainRX(dirarg string) {
	//argz := os.Args
	//dirarg := argz[len(argz)-1]

	files, err := os.ReadDir(dirarg)
	if err != nil {
		panic(err)
	}

	for _, ffile := range files {
		if ffile.IsDir() {
			mainRX(path.Join(dirarg, ffile.Name()))
			continue
		}
		oldname := ffile.Name()
		newname := maintr(oldname)
		if len(oldname) != len(newname) {
			oldpath := path.Join(dirarg, oldname)
			newpath := path.Join(dirarg, newname)
			err = os.Rename(oldpath, newpath)
			if err != nil {
				panic(err)
			}
			fmt.Printf("%v -> %v\n\n", oldpath, newpath)
		}
	}
}

func main() {
	mainRX(os.Args[len(os.Args)-1])
}
