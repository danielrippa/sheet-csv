package main

import (
  "os"
  "fmt"
  "encoding/csv"
  "github.com/xuri/excelize/v2"
)

func main() {
  args := os.Args[1:]

  xlf, err := excelize.OpenFile(args[0])
  if err != nil {
    fmt.Println(err)
    return
  }

  rows, err := xlf.GetRows(args[1])
  if err != nil {
    fmt.Println(err)
    return
  }

  csvf, err := os.Create(args[2])
  defer csvf.Close()
  if err != nil {
    fmt.Println(err)
    return
  }

  w := csv.NewWriter(csvf)
  defer w.Flush()

  err = w.WriteAll(rows)
  if err != nil {
    fmt.Println(err)
  }

}