read -p "Input the source file name: " srcFile

read -p "Input your start point: " strtPnt
read -p "Input your end point: " endPnt

read -p "Input your output file name: " outFile

# ffmpeg -ss 00:00:00 -to 00:00:00 -i "in.mp4" -c copy "out.mp4"
# 	// 	ffmpeg -i "out.mp4" "out.mp3"

cutCommand="ffmpeg -ss ${strtPnt} -to ${endPnt} -i \"${srcFile}\" -c copy \"${outFile}\" "

echo $cutCommand