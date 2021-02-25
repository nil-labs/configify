package main

import (
	"github.com/olekukonko/tablewriter"
	 "os"
)


func main(){
	Greet()
}


// Greet gives a nice greetings string 
func Greet() {
		data := [][]string{
		[]string{"A", "The Good", "500"},
		[]string{"B", "The Very very Bad Man", "288"},
		[]string{"C", "The Ugly", "120"},
		[]string{"D", "The Gopher", "800"},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Sign", "Rating"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output
}


