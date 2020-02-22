package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"royal-rumble/models"
	"royal-rumble/utils"
	"sort"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("USAGE : %s <input_filename> \n", os.Args[0])
		os.Exit(0)
	}
	names := models.RoyalNames{}
	fileUtil := utils.NewFileUtil(ioutil.ReadFile)
	names, err := fileUtil.ContentToList(os.Args[1])
	if err != nil {
		panic(err.Error())
	}

	sort.Sort(names)
	fmt.Printf("Sorted names: %+v\n", names)
}
