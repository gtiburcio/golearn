package matematica

import (
	"fmt"
	"strconv"
)

func Media(numeros ...float64) float64 {
	total := 0.0
	for _, value := range numeros {
		total += value
	}
	media := total / float64(len(numeros))
	mediaArrendondada, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", media), 64)
	return mediaArrendondada
}
