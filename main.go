package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func isBitString(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, c := range s {
		if c != '0' && c != '1' {
			return false
		}
	}
	return true
}

func convertStringToBin(bitString, outputPath string) error {
	// Create output file
	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output: %w", err)
	}
	defer outFile.Close()

	var currentByte byte
	bitCount := 0
	totalBytesWritten := 0

	for _, r := range bitString {
		// Only process '0' and '1'
		if r == '0' || r == '1' {
			currentByte <<= 1
			if r == '1' {
				currentByte |= 1
			}
			bitCount++

			// Every 8 bits, write exactly one byte
			if bitCount == 8 {
				_, err := outFile.Write([]byte{currentByte})
				if err != nil {
					return err
				}
				bitCount = 0
				currentByte = 0
				totalBytesWritten++
			}
		}
	}

	// PADDING: If bitCount > 0, pad with zeros
	if bitCount > 0 {
		paddingNeeded := 8 - bitCount
		currentByte <<= paddingNeeded
		_, err := outFile.Write([]byte{currentByte})
		if err != nil {
			return err
		}
		totalBytesWritten++
		fmt.Printf("Note: Padded last byte with %d zeros to complete the byte.\n", paddingNeeded)
	}

	fmt.Printf("Successfully wrote %d bytes to %s\n", totalBytesWritten, outputPath)
	return nil
}

func main() {
	// 1. Check for command line arguments
	if len(os.Args) < 3 {
		fmt.Println("Usage: bit2bin <input.txt> <output.bin>   (file input)")
		fmt.Println("   or: bit2bin \"10101010\" output.bin    (string input)")
		os.Exit(1)
	}

	inputArg := os.Args[1]
	outputPath := os.Args[2]

	var err error

	// 2. Auto-detect input mode
	if isBitString(inputArg) {
		err = convertStringToBin(inputArg, outputPath)
	} else {
		err = convertTextToBin(inputArg, outputPath)
	}

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Conversion complete.")
}

func convertTextToBin(inputPath, outputPath string) error {
	// 2. Open input file
	inFile, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("failed to open input: %w", err)
	}
	defer inFile.Close()

	// 3. Create output file
	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output: %w", err)
	}
	defer outFile.Close()

	reader := bufio.NewReader(inFile)
	var currentByte byte
	bitCount := 0
	totalBytesWritten := 0

	for {
		r, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		// Only process '0' and '1'
		if r == '0' || r == '1' {
			currentByte <<= 1
			if r == '1' {
				currentByte |= 1
			}
			bitCount++

			// Every 8 bits, write exactly one byte
			if bitCount == 8 {
				_, err := outFile.Write([]byte{currentByte})
				if err != nil {
					return err
				}
				bitCount = 0
				currentByte = 0
				totalBytesWritten++
			}
		}
	}

	// PADDING LOGIC:
	// If bitCount > 0, we have between 1 and 7 bits leftover.
	// We shift them to the left to fill the byte and pad the rest with zeros.
	if bitCount > 0 {
		paddingNeeded := 8 - bitCount
		currentByte <<= paddingNeeded

		_, err := outFile.Write([]byte{currentByte})
		if err != nil {
			return err
		}
		totalBytesWritten++
		fmt.Printf("Note: Padded last byte with %d zeros to complete the byte.\n", paddingNeeded)
	}
	fmt.Printf("Successfully wrote %d bytes to %s\n", totalBytesWritten, outputPath)
	return nil
}
