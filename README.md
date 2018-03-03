# gtranslate
Google translate simple api connector with posibility to translate ukrainian.

Example:
--------
    
    package main
    import (
        "fmt"

        "github.com/test/gtranslate/gtranslate"
    )
    func main() {
        fmt.Println(gtranslate.Translate("Hello", "ru", "ua")) //Привет
    }
