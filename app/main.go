package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"
	// Available if you need it!
	// "github.com/pingcap/parser"
	// "github.com/pingcap/parser/ast"
)

// Usage: your_sqlite3.sh sample.db .dbinfo
func main() {
	databaseFilePath := os.Args[1]
	command := os.Args[2]

	switch command {
	case ".dbinfo":
		databaseFile, err := os.Open(databaseFilePath)
		if err != nil {
			log.Fatal(err)
		}

		header := make([]byte, 100)

		_, err = databaseFile.Read(header)
		if err != nil {
			log.Fatal(err)
		}

		var pageSize uint16
		if err := binary.Read(bytes.NewReader(header[16:18]), binary.BigEndian, &pageSize); err != nil {
			fmt.Println("Failed to read integer:", err)
			return
		}
		// You can use print statements as follows for debugging, they'll be visible when running tests.
		fmt.Println("Logs from your program will appear here!")

		// Uncomment this to pass the first stage
		fmt.Printf("database page size: %v\n", pageSize)

		bTreePage := make([]byte, 4096)
		_, err = databaseFile.ReadAt(bTreePage, 100)
		if err != nil {
			log.Fatal(err)
		}

		var numTables uint16
		if err := binary.Read(bytes.NewReader(bTreePage[3:5]), binary.BigEndian, &numTables); err != nil {
			fmt.Println("Failed to read integer:", err)
			return
		}

		fmt.Printf("number of tables: %v\n", numTables)
	default:
		fmt.Println("Unknown command", command)
		os.Exit(1)
	}
}
