package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// An artificial input source.
	const input = "It is not the critic who counts; not the man who points out how the strong man stumbles, or where the doer of deeds could have done them better. The credit belongs to the man who is actually in the arena, whose face is marred by dust and sweat and blood; who strives valiantly; who errs, who comes short again and again, because there is no effort without error and shortcoming; but who does actually strive to do the deeds; who knows great enthusiasms, the great devotions; who spends himself in a worthy cause; who at the best knows in the end the triumph of high achievement, and who at the worst, if he fails, at least fails while daring greatly, so that his place shall never be with those cold and timid souls who neither know victory nor defeat. "
	scanner := bufio.NewScanner(strings.NewReader(input))

	fmt.Print("Scaner is", scanner, " and the Type of scanner is")
	fmt.Printf("Scaner is %T\n", scanner)
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	// Count the words.
	iteration := 0
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		fmt.Println("Iteration: ", iteration)
		iteration++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}
