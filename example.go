package isikukood

import (
	"fmt"
)

func main() {
	isikukood := Isikukood{code: "37605030299"}
	fmt.Printf("Is valid? %d", isikukood.Validate())
	fmt.Printf("Birthday: %v", isikukood.Birthday())
	fmt.Printf("Sex: %v", isikukood.Sex())
}
