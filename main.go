/*
 * @author Marko Skorupan
 * @version 1.0.0
 * @date 2025-11-12
 * @fileoverview "Program to round several numbers to specific decimal places and display them inside fields of a given size."
 */

package main

import "fmt"

func main() {
	fmt.Printf("%10.3f\n", 8.5467)
	fmt.Printf("%8.5f\n", 9.6382)
	fmt.Printf("%6.1f\n", 18.5146)
	fmt.Printf("%3.1f\n", 125.496)

	fmt.Printf("\nDone.")
} 