// package main

// import (
// 	"fmt"
// 	"monke/repl"
// 	"os"
// 	"os/user"
// )

// 	func main() {
// 		user, err := user.Current()
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Printf("Hello %s! This is the U++ (Utsav Programming Language) xDDD !\n",
// 			user.Username)
// 		fmt.Printf("Koi bhi commands do mujhe , mai lexing karke deta hu ;)\n")
// 		repl.Start(os.Stdin, os.Stdout)
// 	}
// // package main

// // import (
// // 	"bufio"
// // 	"fmt"
// // 	"monke/repl"
// // 	"os"
// // 	"os/user"
// // )

// // func main() {
// // 	user, err := user.Current()
// // 	if err != nil {
// // 		panic(err)
// // 	}

// // 	fmt.Printf("Hello %s! This is the U++ (Utsav Programming Language) xDDD !\n", user.Username)
// // 	fmt.Printf("Koi bhi commands do mujhe, mai lexing karke deta hu ;)\n")

// // 	if len(os.Args) < 3 {
// // 		fmt.Println("Usage: <executable> <input_file_path> <output_file_path>")
// // 		os.Exit(1)
// // 	}

// // 	inputFilePath := os.Args[1]
// // 	outputFilePath := os.Args[2]

// // 	// Open the input file
// // 	inputFile, err := os.Open(inputFilePath)
// // 	if err != nil {
// // 		fmt.Printf("Failed to open input file: %v\n", err)
// // 		os.Exit(1)
// // 	}
// // 	defer inputFile.Close()

// // 	// Create or open the output file
// // 	outputFile, err := os.Create(outputFilePath)
// // 	if err != nil {
// // 		fmt.Printf("Failed to create or open output file: %v\n", err)
// // 		os.Exit(1)
// // 	}
// // 	defer outputFile.Close()

// // 	// Start processing the file and redirect output to the outputFile
// // 	repl.Start(bufio.NewReader(inputFile), outputFile)
// // }
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"monke/repl"
// 	"os"
// 	"os/user"
// 	"strings"
// )

// func main() {
// 	user, err := user.Current()
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Printf("Hello %s! This is the U++ (Utsav Programming Language) xDDD !\n", user.Username)

// 	if len(os.Args) < 2 {
// 		fmt.Println("Please provide a file path")
// 		os.Exit(1)
// 	}

// 	filePath := os.Args[1]
// 	// Check if the file has a .upp extension
// 	if !strings.HasSuffix(filePath, ".upp") {
// 		fmt.Println("Please provide a file with a .upp extension")
// 		os.Exit(1)
// 	}

// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		fmt.Printf("Failed to open file: %v\n", err)
// 		os.Exit(1)
// 	}
// 	defer file.Close()

//		// Start processing the file
//		repl.Start(bufio.NewReader(file), os.Stdout)
//	}
package main

import (
	"bufio"
	"fmt"
	"monke/repl"
	"os"
	"strings"
)

func main() {
	// user, err := user.Current()
	// if err != nil {
	// 	panic(err)
	// }

	//fmt.Printf("Hello %s! This is the U++ (Utsav Programming Language) xDDD !\n", user.Username)

	if len(os.Args) < 2 {
		fmt.Println("Please provide a file path")
		os.Exit(1)
	}

	filePath := os.Args[1]
	if !strings.HasSuffix(filePath, ".bottle") {
		fmt.Println("Please provide a file with a .bottle extension")
		os.Exit(1)
	}

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Create a buffered reader
	reader := bufio.NewReader(file)

	// Optionally skip the first line if it's a shebang
	peek, _ := reader.Peek(2)
	if string(peek) == "#!" {
		_, _ = reader.ReadString('\n')  // Read and discard the shebang line
	}

	// Start processing the file
	repl.Start(reader, os.Stdout)
}
