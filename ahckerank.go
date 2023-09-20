func getQueryAnswers(cache_entries [][]string, queries [][]string) []int32 {
    // Write your code here
    cacheResult := make(map[string]string)
    
    for _, entry := range cache_entries{
        key := entry[1] + "-" + entry[0]
        value := entry[2]
        cacheResult[key] = value
    }
    
    results := make([]int32, len(queries))
    for i, query := range queries{
        key := query[0]
        timestamp := query[1]
        cacheKey := key + "-" + timestamp
        
        if value, exists := cacheResult[cacheKey]; exists{
            dataValue, _ := strconv.Atoi(value)
            results[i] = int32(dataValue)
        }else{
            results[i] = -1
        }
    }
    return results
}

func isWithinBounds(current, lowerBound, upperBound int32) bool {
    return current >= lowerBound && current <= upperBound
}

func isWithinBounds(current, lowerBound, upperBound int32) bool {
    return current >= lowerBound && current <= upperBound
}

func getNetProfit(events []string) []int64 {
    portfolio := make(map[string]int64) // Change the value type to int64
    var profit int64
    var results []int64

    for _, eventStr := range events {
        eventParts := strings.Split(eventStr, " ")
        category := eventParts[0]

        switch category {
        case "BUY", "SELL", "CHANGE":
            value := eventParts[len(eventParts)-1]
            value1 := eventParts[1]

            switch category {
            case "BUY":
                quantity, _ := strconv.ParseInt(value, 10, 64) 
                portfolio[value1] += quantity
            case "SELL":
                quantity, _ := strconv.ParseInt(value, 10, 64) 
                portfolio[value1] -= quantity
            case "CHANGE":
                priceChange, _ := strconv.ParseInt(value, 10, 64) 
                tempValue, exists := portfolio[value1]
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



func countAnalogousArrays(consecutiveDifference []int32, lowerBound int32, upperBound int32) int32 {
    count := int32(0)
    var i, current int32
    var valid bool
    
    for i = lowerBound; i <= upperBound; i++{
        current = i
        valid = true
        
        for _, diff := range consecutiveDifference{
            current -= diff
            
            if current < lowerBound || current > upperBound{
                valid = false
                break
            }
        }
        
        if valid {
            count++
        }
    }
    
    return count
    
}


func countAnalogousArrays(consecutiveDifference []int32, lowerBound int32, upperBound int32) int32 {
	// Write your code here
	
	count := int32(0)
	seen := make(map[int32]bool)
	
	for i := lowerBound; i <= upperBound; i++ {
		current := i
	
		// Check if the current value is in the seen map
		if seen[current] {
			continue
		}
	
		// Check if the current value is within the range [lowerBound, upperBound]
		if !binarySearch(consecutiveDifference, lowerBound, upperBound, current) {
			continue
		}
	
		// The current value is valid
		count++
		seen[current] = true
	}
	
	return count
	}
	
func binarySearch(consecutiveDifference []int32, lowerBound int32, upperBound int32, target int32) bool {
	mid := (lowerBound + upperBound) / 2
	
	if consecutiveDifference[mid] == target {
		return true
	} else if consecutiveDifference[mid] < target {
		return binarySearch(consecutiveDifference, mid + 1, upperBound, target)
	} else {
		return binarySearch(consecutiveDifference, lowerBound, mid - 1, target)
	}
}

func binarySearch(consecutiveDifference []int32, lowerBound int32, upperBound int32, target int32) bool {
// Check if the slice is empty
if len(consecutiveDifference) == 0 {
return false
}

mid := (lowerBound + upperBound) / 2

if consecutiveDifference[mid] == target {
    return true
} else if consecutiveDifference[mid] < target {
    return binarySearch(consecutiveDifference, mid + 1, upperBound, target)
} else {
    return binarySearch(consecutiveDifference, lowerBound, mid - 1, target)
}
}
