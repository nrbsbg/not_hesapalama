package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*

	ORTALAMA  VİZE NOTUNUN %40 ile FİNAL NOTUNUN %60 ile bulunur.

*/

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Vize Notunu Giriniz: ")
	scanner.Scan()

	vizenot, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Print("Final Notunu Giriniz: ")
	scanner.Scan()

	finalnot, _ := strconv.ParseFloat(scanner.Text(), 64)

	ortalama := (vizenot * 0.4) + (finalnot * 0.6)

	fmt.Println("Ortalama: ", ortalama)

	if ortalama >= 85 && ortalama <= 100 {
		fmt.Println("AA")
	} else if ortalama >= 70 && ortalama < 85 {
		fmt.Println("BB")
	} else if ortalama >= 60 && ortalama < 70 {
		fmt.Println("CB")
	} else if ortalama >= 50 && ortalama < 60 {
		fmt.Println("CC")
	} else if ortalama < 50 && ortalama >= 0 {
		fmt.Println("FF")
	} else {
		fmt.Println("0-100 arasında bir değer olmalıdır.")
	}

}
