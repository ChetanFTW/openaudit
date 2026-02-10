package main

import (
	"fmt"
	"os"

	"github.com/yourusername/openaudit/internal/scanner"
)

func main() {
	path := "."

	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	result, err := scanner.Scan(path)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("OpenAudit Report")
	fmt.Println("----------------")
	fmt.Println("Go Module:", result.HasGoMod)
	fmt.Println("CI Config:", result.HasCI)
	fmt.Println("README:", result.HasReadme)
}
