package main

import (
	"fmt"
	"gallery/routes"
	"gallery/services"
)

func main() {
	// _ = services.ConnectDB("demo:demo@tcp(127.0.0.1:3306)/Galleries?parseTime=true")
	_ = services.ConnectDB("root:Hieu2951999@tcp(127.0.0.1:3306)/Galleries?parseTime=true")

	fmt.Println("Connected !")
	err := services.CreateLogger()
	if err != nil {
		panic(err)
	}
	defer services.CloseLogger()
	g := routes.Create()
	g.Run("127.0.0.1:3000")

	// var imageSrc = "./services/cat.jpg"
	// sizes, err := services.ResizeAll(imageSrc)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(sizes)
	// }
	// image := services.ImageSize{
	// 	Path: "./services/cat.jpg",
	// }
	// w, h, err := services.GetDimension(image.Path)
	// if err != nil {
	// 	panic(err)
	// }
	// resolutions := []int{1920, 1600, 1280, 1024, 800, 256}

	// wg := sync.WaitGroup{}
	// wg.Add(len(resolutions))
	// for _, size := range resolutions {
	// 	if size < w {
	// 		go func(wg *sync.WaitGroup, size int) {

	// 			defer wg.Done()
	// 			_, err := image.Resize(size)
	// 			fmt.Println(err)

	// 		}(&wg, size)
	// 	}
	// }
	// wg.Wait()
	// fmt.Printf("Width: %d, Height: %d\n", w, h)
	// fmt.Println("done")
}
