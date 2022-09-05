package main

import (
	"fmt"
	"github.com/Bpazy/behappy/images"
	"time"
)

func main() {
	year, week := time.Now().ISOWeek()
	path, err := images.HonorTemplate("南帅", year, week, int(time.Now().Month()), 23)
	if err != nil {
		panic(err)
	}
	fmt.Printf("生成的图像路径: %s", path)
}
