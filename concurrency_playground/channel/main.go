package main

import "fmt"

// >> channel Direction
// hanya bisa mengirim
// func sendOnly(ch chan<- int){
// 	for i := 1; i <= 5; i++{
// 		ch <- i
// 	}
// 	close(ch)
// }

// // hanya bisa mengirim
// func receiveOnly(ch <-chan int){
// 	for v := range ch{
// 		fmt.Println("Received: ", v)
// 	}
// }
	
func main() {
	// ch := make(chan int)

	// go sendOnly(ch)
	// receiveOnly(ch)


	// >> Contoh Range Channel
	// ch := make(chan int, 3)

	// ch <- 1
	// ch <- 2
	// ch <- 3

	/* >> close menandakan tidak ada lagi data yang akan dikrim melalui channel
	>> hanya sender yang bisa menutup channel
	>> tidak boleh mengirim data ke channel yang sudah ditutup
	>> receiver masih bisa membaca sampai data habis (nilai terakhir 0 karena channel sudah kosong dan tertutup)
	*/

	// close(ch)

	// for v := range ch {
	// 	fmt.Println(v)
	// }

	// >> contoh 1
	// ch := make(chan string)

	// go func() {
	// 	ch <- "Hello from goroutine"
	// }()

	// message := <- ch
	// fmt.Println(message)

	// >> contoh 2
	// ch := make(chan string)

	// go func(){
	// 	fmt.Println("Sending...")
	// 	ch <- "hello"
	// 	fmt.Println("Sent!")
	// }()

	// message := <- ch
	// fmt.Println("Received", message)

	// >>contoh dengan membuat channel menjadi 2 slot (buffered Channel)
	// ch := make(chan string, 2)

	// go func(){
	// fmt.Println("Start...")
	// ch <- "satu"
	// ch <- "dua"
	// fmt.Println("end")
	// }()
	// message := <- ch
	// message2 := <- ch
	// fmt.Println(message)
	// fmt.Println(message2)

	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3

	close(ch)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println("Done")
}