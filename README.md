# isikukood
Package for learning GO and for validating Estonian Personal Identification Number

```go
package your_package

import (
	"fmt"
	isikukood "github.com/texmex5/isikukood"
)

func main() {
	i := isikukood.Isikukood{Code: "37605030299"}
	fmt.Printf("Is valid? %d", i.Validate())
	fmt.Printf("Birthday: %v", i.Birthday())
	fmt.Printf("Sex: %v", i.Sex())
}


```
