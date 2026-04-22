package v3

import (
	"fmt"
)

func HelloV3(name string) {
	message := fmt.Sprintf("Greetings from V3: %v", name);
	fmt.Println(message);
}