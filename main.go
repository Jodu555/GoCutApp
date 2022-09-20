package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	var defs = true
	var sourceFile string = "input.mp4"
	var startPoint string = "00:05:00"
	var endPoint string = "00:17:00"
	var outputFile string = "out"
	if !defs {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("This is a simple application to cut an .mp4 video file")
		fmt.Println("This just cuts the start and end position and then based on decision converts it to an .mp3")
		fmt.Println("And then it deletes the source video file")
		fmt.Println("")

		fmt.Println("Input the source file name: ")
		scanner.Scan()
		sourceFile = scanner.Text()

		fmt.Println("Input your start point: ")
		scanner.Scan()
		startPoint = scanner.Text()

		fmt.Println("Input your end point: ")
		scanner.Scan()
		endPoint = scanner.Text()

		fmt.Println("Input your output file name wihtou ext: ")
		scanner.Scan()
		outputFile = scanner.Text()
	}

	// 	ffmpeg -ss 00:00:00 -to 00:00:00 -i "in.mp4" -c copy "out.mp4"
	// 	ffmpeg -i "out.mp4" "out.mp3"

	cutCommand := "ffmepg -ss " + startPoint + " -to " + endPoint + " -i \"" + sourceFile + "\" -c copy \"" + outputFile + ".mp4\""

	convertCommand := "ffmpeg -i \"" + outputFile + ".mp4\" \"" + outputFile + ".mp3\""

	// fmt.Printf("Your source file is: %v and your start point: %v and your end point %v", sourceFile, startPoint, endPoint)

	fmt.Println("Your command: \n" + cutCommand + "\n" + convertCommand)

	// cmd := exec.Command("ffmpeg")
	// // cmd := exec.Command(cutCommand)

	// err := cmd.Run()

	// if err != nil {
	// 	log.Fatal(err)
	// }

	out, err := exec.Command(cutCommand).Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("OUT: ", string(out))

}
