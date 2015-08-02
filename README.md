# isikukood
Package for learning GO and for validating Estonian Personal Identification Number

```go
package your_package

import (
	"fmt"
	"github.com/texmex5/isikukood"
)

func main() {
	isikukood := Isikukood{code: "37605030299"}
	fmt.Printf("Is valid? %d", isikukood.Validate())
	fmt.Printf("Birthday: %v", isikukood.Birthday())
	fmt.Printf("Sex: %v", isikukood.Sex())
}

```
