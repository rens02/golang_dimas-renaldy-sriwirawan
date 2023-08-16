package main
import (
	"fmt"
)

func main() {

	// deklarasi variabel dan type data nya
	var top, bottom, height float32
	
	fmt.Print("Panjang bagian atas trapesium : ")
	fmt.Scanln(&top) // Tanda & digunakan untuk mengambil alamat variabel teratas

	fmt.Print("Panjang bagian bawah trapesium : ")
	fmt.Scanln(&bottom)

	fmt.Print("tinggi trapesium : ")
	fmt.Scanln(&height)

	
	wide := 0.5 * (top + bottom)* height

	fmt.Printf("Luas trapesium adalah : %.2f\n",wide)
}