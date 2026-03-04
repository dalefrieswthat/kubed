package layout

import (
	"fmt"
	"os"
)

// RunShow reads layout.json and prints it to stdout.
// If the file is missing, prints "run kubed layout capture" to stderr and returns an error.
func RunShow() error {
	layoutPath, err := LayoutPath()
	if err != nil {
		return err
	}
	data, err := os.ReadFile(layoutPath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "run kubed layout capture\n")
			return err
		}
		return fmt.Errorf("read %s: %w", layoutPath, err)
	}
	fmt.Print(string(data))
	return nil
}
