package IsDocker

import (
	"fmt"
	"testing"
)

func TestIsDocker(t *testing.T) {
	if IsDocker() {
		fmt.Println("Info: Running inside a docker container")
	} else {
		fmt.Println("Info: Running outside a docker container")
	}
}
