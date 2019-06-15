package piscine

import "math"

func Sqrt(arg1 int) int{

	argFlottant := float64(arg1)
	var testeur float64 = float64(int(math.Sqrt(argFlottant)))

	if math.Sqrt(argFlottant) - testeur != 0 {
		return 0		
	}else{
		return int(math.Sqrt(argFlottant))
	}
}
