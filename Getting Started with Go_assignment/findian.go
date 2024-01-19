package main
 
import (
  "fmt"
  "regexp"
  "strings"
  "bufio"
  "os"
)

func main() {
  fmt.Print("Enter the string: ")   
  scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		s := strings.Replace(scanner.Text(), " ", "", -1)
    matched, err := regexp.MatchString("(?i)(([a-z0-9]*)i([a-z0-9]*)a([a-z0-9]*)n([a-z0-9]*))", s)
    if err != nil {
      fmt.Println(err)
    }
    if matched {
      fmt.Println("String contains ian")
    } else {
      fmt.Println("String does not contain ian")
    }
	}
}