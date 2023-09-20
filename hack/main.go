package main

import (
	"fmt"
)

func main() {
	consecutiveDifference := []int32{-2, -1, -2, 5}
	lowerBound := int32(3)
	upperBound := int32(10)

	result := countAnalogousArrays(consecutiveDifference, lowerBound, upperBound)
	fmt.Println("Result:", result) // Hasil yang diharapkan: 3
}
func countAnalogousArrays(consecutiveDifference []int32, lowerBound int32, upperBound int32) int32 {
    // Menghitung jumlah cara untuk mengisi elemen pertama
    firstDiff := consecutiveDifference[0]
    firstCount := ((upperBound - lowerBound) + firstDiff) / firstDiff

    return firstCount
}
func getNetProfit(events []string) []int64 {
	// buat map
    porto := make(map[string]int64)
	//deklarasi variable
    var profit int64
    var results []int64

    for _, eventStr := range events {
        eventParts := strings.Split(eventStr, " ")
        category := eventParts[0]

        switch category {
        case "BUY", "SELL", "CHANGE":
            value1 := eventParts[len(eventParts)-1]
            value2 := eventParts[1]

            switch category {
			//case BUY	
            case "BUY":
                quantity, _ := strconv.ParseInt(value1, 10, 64) 
                porto[value2] += quantity
			// case SELL	
            case "SELL":
                quantity, _ := strconv.ParseInt(value1, 10, 64) 
                porto[value2] -= quantity
			// case CHANGE	
            case "CHANGE":
                priceChange, _ := strconv.ParseInt(value1, 10, 64) 
                tempValue, exists := porto[value2]
                if exists {
                    profit += tempValue * priceChange
                }
            }
        case "QUERY":
            results = append(results, profit)
        }
    }

    return results
}

func getNetProfit(events []string) []int64 {
    // Write your code here
    porto := make(map[string]int64) // Change the value type to int64
    var keuntungan int64
    var hasil []int64

    for _, eventString := range events {
        pecahanEvent := strings.Split(eventString, " ")
        kategori := pecahanEvent[0]

        switch kategori {
        case "BUY", "SELL", "CHANGE":
            value1 := pecahanEvent[len(pecahanEvent)-1]
            aset := pecahanEvent[1]

            switch kategori {
            case "BUY":
                jumlah, _ := strconv.ParseInt(value1, 10, 64) 
                porto[aset] += jumlah
            case "SELL":
                jumlah, _ := strconv.ParseInt(value1, 10, 64) 
                porto[aset] -= jumlah
            case "CHANGE":
                gantiHarga, _ := strconv.ParseInt(value1, 10, 64) 
                valTemp, exists := porto[aset]
                if exists {
                    keuntungan += valTemp * gantiHarga
                }
            }
        case "QUERY":
            hasil = append(hasil, keuntungan)
        }
    }

    return hasil
    

}
