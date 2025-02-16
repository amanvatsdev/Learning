/*1.Define an inrerface 'Shape' with a method 'Area()Floar64'
2.Create a 'Rectangle' struct with 'Width' and 'Height' fields.
3.Create a 'Circle' struct with a 'Radius' field.
4.Implement the 'Area()' method for both 'Rectange' and 'Circle'.
5.Print the area of both shapes in thee 'main()' function.
*/

package main

import (
	"fmt"
	"math"
)

type Shape interface{
	Area()float64
}

type Rectangle struct {
	Width  float64
	Height float64
}
func (r Rectangle)Area()float64{
	return r.Height*r.Width
}


type Circle struct {
	Radius float64
}

func (c Circle)Area()float64{
	return math.Pi*c.Radius*c.Radius
}

func main() {
Rectangle1:=Rectangle{
	Width: 10,
	Height: 15,
}
Circle1:=Circle{
	Radius: 13,
}

fmt.Printf("Area of Rectangle:%.2f\n",Rectangle1.Area())
fmt.Printf("Area of Circle :%.2f\n",Circle1.Area())
}
