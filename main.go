package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("This is a simple application to cut an .mp4 video file")
	fmt.Println("This just cuts the start and end position and then based on decision converts it to an .mp3")
	fmt.Println("And then it deletes the source video file")
	fmt.Println("")

	fmt.Println("Input the source file name: ")
	scanner.Scan()
	sourceFile := scanner.Text()

	fmt.Println("Input your start point: ")
	scanner.Scan()
	startPoint := scanner.Text()

	fmt.Println("Input your end point: ")
	scanner.Scan()
	endPoint := scanner.Text()

	cutCommand := ""

	fmt.Printf("Your source file is: %v and your start point: %v and your end point %v", sourceFile, startPoint, endPoint)

}
