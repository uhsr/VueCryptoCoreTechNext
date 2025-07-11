// cmd/vuecryptocoretechnext/main.go
package main

import (
"flag"
"log"
"os"

"vuecryptocoretechnext/internal/vuecryptocoretechnext"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := vuecryptocoretechnext.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
