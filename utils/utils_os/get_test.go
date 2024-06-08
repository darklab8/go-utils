package utils_os

import (
	"fmt"
	"testing"
)

func TestFolderName(t *testing.T) {
	fmt.Println(GetCurrentFolder().Base())
}
