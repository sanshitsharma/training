package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	numof := len(os.Args)
	switch numof {
	case 4:
		bdata, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			return
		}
		sdata := string(bdata)
		sdata = strings.ReplaceAll(sdata, os.Args[2], os.Args[3])
		extension := filepath.Ext(os.Args[1])
		name := os.Args[1][0 : len(os.Args[1])-len(extension)]
		name = name + "_revised" + extension
		ioutil.WriteFile(name, []byte(sdata), 644)
	case 3:
		bdata, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			return
		}
		re, _ := regexp.Compile(os.Args[2])
		locations := re.FindAllStringIndex(string(bdata), -1)
		fmt.Println(os.Args[2] + " text at: ")
		for _, item := range locations {
			fmt.Printf(" %d ", item[0])
		}
		fmt.Println()
	default:
	}
}
