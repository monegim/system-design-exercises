package main

import (
	"errors"
	"fmt"
	"log"
	"path/filepath"

	"helm.sh/helm/v3/pkg/chart"
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
}

func createHelmChart(name string) {
	_, err := chartutil.Create(name, ".")
	if err != nil {
		log.Fatal(err)
	}
}
