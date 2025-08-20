package main

import (
	"fmt"
	"os"
)

func main() {
	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// fileInfo, err := f.Stat()

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Name of the file: ", fileInfo.Name())
	// fmt.Println("file or folder ", fileInfo.IsDir())
	// fmt.Println("file size ", fileInfo.Size())
	// fmt.Println("file Mode ", fileInfo.Mode())
	// fmt.Println("file Modified at ", fileInfo.ModTime())

	//read file

	// f, err := os.Open("example.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// buf := make([]byte, 44)

	// d, err := f.Read(buf)

	// if err != nil {
	// 	panic(err)
	// }

	// for i := 0; i < len(buf); i++ {
	// 	println("data from file", d, string(buf[i]))
	// }

	//another way for smaller files

	// data, err := os.ReadFile("example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(data))

	//read folders

	// dir, err := os.Open("../")
	// if err != nil {
	// 	panic(err)
	// }
	// defer dir.Close()

	// fileInfo, err := dir.ReadDir(-1)
	// if err != nil {
	// 	panic(err)
	// }

	// for _, v := range fileInfo {
	// 	fmt.Println(v.Name())
	// }

	//create a file
	// f, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// // f.WriteString("Hello go")
	// // f.WriteString(" h")

	// arr := []byte("hello golang")

	// f.Write(arr)

	//read and write to another file (streaming fashion)

	// sourceFile, err := os.Open("example.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// defer sourceFile.Close()

	// destFile, err := os.Create("destinationFile.txt")

	// if err != nil {
	// 	panic(err)
	// }
	// defer destFile.Close()

	// reader := bufio.NewReader(sourceFile)
	// writer := bufio.NewWriter(destFile)

	// for {
	// 	b, err := reader.ReadByte()
	// 	if err != nil {
	// 		if err.Error() != "EOF" {
	// 			panic(err)
	// 		}
	// 		break
	// 	}

	// 	e := writer.WriteByte(b)
	// 	if e != nil {
	// 		panic(e)
	// 	}
	// }

	// writer.Flush()

	// fmt.Println("Files Generated and Transferred successfulyy")

	//delete a file

	// err := os.Remove("destinationFile.txt")
	// if err != nil {
	// 	panic(err)
	// }

	if err := os.Remove("sample.txt"); err != nil { //single line

		panic(err)
	}

	fmt.Println("Files deleted Successfully")

}
