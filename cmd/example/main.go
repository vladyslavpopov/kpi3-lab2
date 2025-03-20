package main

import (
  "flag"
  "fmt"
  "io"
  "os"
  "strings"

  lab2 "github.com/vladyslavpopov/kpi3-lab2"
)

var (
  flagExpr = flag.String("e", "", "Expression to compute")
  flagFile = flag.String("f", "", "File containing expression")
  flagOut  = flag.String("o", "", "Output file for result")
)

func main() {
  flag.Parse()

  if (*flagExpr == "" && *flagFile == "") || (*flagExpr != "" && *flagFile != "") {
    fmt.Fprintln(os.Stderr, "Specify either -e or -f, but not both.")
    os.Exit(1)
  }

  var inputReader io.Reader
  if *flagExpr != "" {
    inputReader = strings.NewReader(*flagExpr)
  } else {
    f, err := os.Open(*flagFile)
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error opening input file:", err)
      os.Exit(1)
    }
    defer f.Close()
    inputReader = f
  }

  var outputWriter io.Writer
  if *flagOut != "" {
    f, err := os.Create(*flagOut)
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error creating output file:", err)
      os.Exit(1)
    }
    defer f.Close()
    outputWriter = f
  } else {
    outputWriter = os.Stdout
  }

  handler := &lab2.ComputeHandler{
    Input:  inputReader,
    Output: outputWriter,
  }

  if err := handler.Compute(); err != nil {
    fmt.Fprintln(os.Stderr, "Error computing expression:", err)
    os.Exit(1)
  }
}
