package main

import (
	// "fmt"
	// "bufio"
	"fmt"
	"os"
)

func main() {
	// Open the file
	// f, err := os.Open("myFile.txt") 
	// defer f.Close() //Close the file after the function exits
	// if err != nil {
	// 	panic(err)
	// } 
	// fileInfo, err := f.Stat()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("File name: ", fileInfo.Name()) //myFile.txt
	// fmt.Println("is Folder: ", fileInfo.IsDir()) //false
	// fmt.Println("File size: ", fileInfo.Size()) //bytes
	// fmt.Println("File mode: ", fileInfo.Mode()) //permissions
	// fmt.Println("Last modified: ", fileInfo.ModTime()) //time


	// Read the file
	// f, err := os.Open("myFile.txt") 
	// fileInfo, err := f.Stat()
	// buf := make([]byte, fileInfo.Size())
	// d, err := f.Read(buf)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("data", d, buf)
	// for i:=0; i<len(buf); i++ {
	// 	print(string(buf[i]))
	// }


	// Read the file using ReadFile
	// d, err := os.ReadFile("myFile.txt") // Drawback: Read the whole file at once // Not suitable for large files
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(d, string(d))



	// Read folders
	// dir, err := os.Open("../")
	// defer dir.Close()
	// if err != nil {
	// 	panic(err)
	// }
	// fileInfo, err := dir.ReadDir(0) 
	// // n <= 0: read all files
	// // n > 0: read n files
	// for _, file := range fileInfo {
	// 	fmt.Println(file.Name())
	// }
	


	// Create a file
	// f, err := os.Create("newFile1.txt")
	// defer f.Close()
	// if err != nil {
	// 	panic(err)
	// }
	// f.WriteString("Anish is learning Go") // Append data to the file

	// bytes := []byte("Anish is from Pune")
	// f.Write(bytes) // Append data to the file



	// Read and write to a file (Streaming fashion)
	// srcFile, err := os.Open("myFile.txt")
	// defer srcFile.Close()
	// if err != nil {
	// 	panic(err)
	// }

	// destFile, err := os.Create("destFile.txt")
	// defer destFile.Close()
	// if err != nil {
	// 	panic(err)
	// }

	// reader := bufio.NewReader(srcFile)
	// writer := bufio.NewWriter(destFile)

	// for {
	// 	b, err := reader.ReadByte()
	// 	if err != nil {
	// 		if err.Error() != "EOF" {
	// 			panic(err)
	// 		}
	// 		break
	// 	}
	// 	er := writer.WriteByte(b)
	// 	if er != nil {
	// 		panic(er)
	// 	}
	// }
	// writer.Flush()
	// fmt.Println("File copied successfully")



	// Delete a file
	er := os.Remove("newFile1.txt")
	if er != nil {
		panic(er)
	}
	fmt.Println("File deleted successfully")

}