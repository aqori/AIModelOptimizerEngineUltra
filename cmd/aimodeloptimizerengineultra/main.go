// cmd/aimodeloptimizerengineultra/main.go
package main

import (
"flag"
"log"
"os"

"aimodeloptimizerengineultra/internal/aimodeloptimizerengineultra"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := aimodeloptimizerengineultra.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
