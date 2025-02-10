package util

import (
	"fmt"
	"os"
	"bufio"
)

func ReadChar() (rune, error) {
	var char rune

	_, err := fmt.Scanf("%c", &char)
	if err != nil {
		return '0', err
	}

	return char, nil
}

func ReadLine() (string, error) {
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Scan()
  err := scanner.Err()
  if err != nil {
    return "", err
  }
  
  return scanner.Text(), nil
}