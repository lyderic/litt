package main

import "fmt"

func list() {
	// what is the longest name?
	ln := 0
	for _, montage := range configuration.Montages {
		i := len(montage.Name)
		if i > ln {
			ln = i
		}
	}
	format := fmt.Sprintf("  [%%02d] %%-%d.%ds [%%s]", ln, ln)
	for _, montage := range configuration.Montages {
		fmt.Printf(format, montage.Id, montage.Name, montage.Path)
		if montage.Id == 1 {
			fmt.Println(" (default)")
		} else {
			fmt.Println()
		}
	}
}
