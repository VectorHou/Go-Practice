package main

import (
	"os"
	"github.com/olekukonko/tablewriter"
)

func main(){
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"A", "B", "C", "D", "E"})
	table.Render()
}