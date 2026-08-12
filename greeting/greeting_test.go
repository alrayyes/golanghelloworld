package greeting_test

import (
	"fmt"
	"testing"

	"github.com/alrayyes/golanghelloworld/greeting"
	"github.com/stretchr/testify/assert"
)

func TestGreet(t *testing.T) {
	assert.Equal(t, "Hello World", greeting.Greet())
}

func ExampleGreet() {
	fmt.Println(greeting.Greet())
	// Output: Hello World
}
