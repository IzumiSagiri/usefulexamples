package main

import (
	"fmt"
	"github.com/IzumiSagiri/record"
	"github.com/mjibson/go-dsp/dsputils"
	"github.com/mjibson/go-dsp/fft"
	"log"
	"math"
	"math/cmplx"
	"os"
	"time"
)

func main() {
	data := record.GetRecord()
	var datainfloat []float64
	for _, v := range data {
		datainfloat = append(datainfloat, float64(v)-128.0)
	}
	log.Print(datainfloat)
	X := fft.FFTReal(datainfloat)
	presentDate := time.Now()
	filePTR, err := os.Create(fmt.Sprint(presentDate.Year(), presentDate.Month(), presentDate.Day(), presentDate.Hour(), presentDate.Minute(), presentDate.Second(), ".csv"))
	if err != nil {
		log.Fatal(err)
	}
	// Print the magnitude and phase at each frequency.
	for i := 0; i < 44100; i++ {
		r, θ := cmplx.Polar(X[i])
		θ *= 360.0 / (2 * math.Pi)
		if dsputils.Float64Equal(r, 0) {
			θ = 0 // (When the magnitude is close to 0, the angle is meaningless)
		}
		filePTR.WriteString(fmt.Sprintf("X(%d), %.1f, %.1f\n", i, r, θ))
	}
}
