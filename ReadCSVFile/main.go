package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Booker interface {
	ReadLine(strLine string) (Book, *string)
}

type Book struct {
	strTitle  string
	sPrice    float32
	nQuantity int
}

func (b Book) String() string {
	return fmt.Sprintf("Title=%s; Price=%.2f; Quantity=%d\n", b.strTitle, b.sPrice, b.nQuantity)
}

func (book *Book) ReadLine(strLine string) bool {

	arrFields := strings.Split(strLine, ";")

	if len(strLine) >= 3 {
		book.strTitle = arrFields[0]
		sPrice, err := (strconv.ParseFloat(arrFields[1], 32))
		if err == nil {
			book.sPrice = float32(sPrice)
		} else {
			return false
		}
		book.nQuantity, err = strconv.Atoi(strings.TrimSpace(arrFields[2]))
		if err != nil {
			return false
		}
		return true
	} else {
		fmt.Printf("Not enough fields in string %s, fields=%d\n", strLine, len(arrFields))
		return false
	}
}

func readBooksFromFile(strPath string) []Book {

	var arrBooks []Book

	inputFile, err := os.Open(strPath)
	if err == nil {
		fmt.Printf("Opened file at %s\n", strPath)
		defer inputFile.Close()
		inputReader := bufio.NewReader(inputFile)
		for {
			inputString, readerError := inputReader.ReadString('\n')
			if readerError == io.EOF {
				break
			} else {
				newBook := new(Book)
				ok := newBook.ReadLine(inputString)
				if ok {
					arrBooks = append(arrBooks, *newBook)
				}
			}
		}
	} else {
		fmt.Printf("Can't open file at %s\n", strPath)
	}
	return arrBooks
}

func main() {

	arrBooks := readBooksFromFile("./products.txt")

	fmt.Printf("Final array has %d elements\n", len(arrBooks))
	for i, book := range arrBooks {
		fmt.Printf("[%d] %s", i, book)
	}
}
