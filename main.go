package main

import (
    "fmt"
    "math"
)

// Function for calculating derivatives of x2+y2
func DerivateValueCalculator(x float64, y float64) (float64, float64, float64, float64) {
    
    // Value of x^2+y^2 
    firstDerivate := math.Pow(x, 2)+math.Pow(y, 2);

    // value of 2x+2yy'
    secondDerivate := 2*x+2*y*firstDerivate;
    // value of 2+2(yy'' + y'y')
    thirdDerivate := 2 + 2*(y*secondDerivate + firstDerivate * firstDerivate);
    // value of 2(yy''' + 3y''y')
    fourthDerivate := 2*(y * thirdDerivate+ 3 * secondDerivate * firstDerivate)

    return firstDerivate, secondDerivate, thirdDerivate, fourthDerivate; 
}

// Calculate the value of each term in expansion (upto 5th term)
func SeriesCalculator(x float64, x0 float64, y0 float64) (float64, float64, float64, float64) {

    // Gets the derivate value by calling DerivateValueCalculator function
    var firstDerivate, secondDerivate, thirdDerivate, fourthDerivate float64 = DerivateValueCalculator(x0, y0);

    // value of (x-x0)y'(x0)
    var secondTerm float64 = (x-x0)*firstDerivate;
    // value of (x-x0)^/2!*y''(x0)
    var thirdTerm float64 = (math.Pow((x-x0), 2)/2)*secondDerivate;
    // value of (x-x0)^3/3! * y'''(x0)
    var fourthTerm float64 = (math.Pow((x-x0), 3)/6)*thirdDerivate;
    // value of (x-x0)^4/4! * y''''(x0)
    var fifthTerm float64 = (math.Pow((x-x0), 4)/24)*fourthDerivate;

    return secondTerm, thirdTerm, fourthTerm, fifthTerm;
}

func main() {
  fmt.Println("Equation: dy/dx = x^2+y^2\n");
  
  var (
      x0 float64;
      y0 float64;
      x float64;
  )

  fmt.Println("Enter the initial value\n");
  fmt.Println("Enter the initial value of x0\t");
  fmt.Scanln(&x0);
  fmt.Println("\nEnter the initial value of y0\t");
  fmt.Scanln(&y0);
  fmt.Println("\nEnter the value where Taylor Series has to be estimated\t");
  fmt.Scanln(&x);

  fmt.Println("\nSeries Obtained\n");
  var firstTerm float64 = y0;
  var secondTerm, thirdTerm, fourthTerm, fifthTerm float64 = SeriesCalculator(x, x0, y0);
  fmt.Printf("%.4f + %.4f + %.4f + %.4f + %.4f", firstTerm, secondTerm, thirdTerm, fourthTerm, fifthTerm);

  fmt.Printf("\nTherefore, result\n=\t%.4f", firstTerm+secondTerm+thirdTerm+fourthTerm+fifthTerm);
  }
