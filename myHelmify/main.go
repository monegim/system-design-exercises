package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"path/filepath"

	"helm.sh/helm/v3/pkg/chartutil"
)

func main() {
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()
	// text := scanner.Text()
	// fmt.Println(text)
	// scanner.Scan()
	// text = scanner.Text()
	// fmt.Println(text)
	// for {
	// 	scanner.Scan()
	// 	text := scanner.Text()
	// 	if len(text) != 0 {
	// 		fmt.Println(text)
	// 	} else {
	// 		break
	// 	}
	// }
	// if scanner.Err() != nil {
	// 	fmt.Println("Error: ", scanner.Err())
	// }
	name := "test1"
	createHelmChart(name)
	files := listDir(name)
	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
	}
}

func createHelmChart(name string) {
	_, err := chartutil.Create(name, ".")
	if err != nil {
		log.Fatal(err)
	}
}

func listDir(dir string) []fs.FileInfo {
	abs := filepath.Join(dir, "templates")
	files, err := ioutil.ReadDir(abs)
	if err != nil {
		log.Fatal(err)
	}
	return files
}
