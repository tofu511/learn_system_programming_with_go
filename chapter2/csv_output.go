package main

import (
	"encoding/csv"
	"os"
)

func main()  {
	file, err := os.Create("test.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Write([]string{"csv", "test"})
	writer.Write([]string{"foo", "bar"})
	writer.Write([]string{"hoge", "fuga"})
	writer.Flush()
}
