package log

import "math"

func Log2(number float64) float64 {
  return math.Log(number) / math.Log(2)
}

func Logb(number float64, base float64) float64 {
  return math.Log(number) / math.Log(base)
}
