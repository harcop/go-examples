package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

// compressFile compresses the source file and writes the compressed data to the destination file.
func compressFile(sourceFile, destFile string) error {
	// Open the source file for reading.
	src, err := os.Open(sourceFile)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer src.Close()

	// Create the destination file for writing compressed data.
	dst, err := os.Create(destFile)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	defer dst.Close()

	// Create a gzip writer to compress the file.
	gzipWriter := gzip.NewWriter(dst)
	defer gzipWriter.Close()

	// Copy the data from the source file to the gzip writer (which compresses it).
	_, err = io.Copy(gzipWriter, src)
	if err != nil {
		return fmt.Errorf("failed to compress file: %w", err)
	}

	return nil
}

func main() {
	// Input and output files.
	sourceFile := "input.txt"      // Replace with the path to your input file.
	destFile := "output.txt.gz"    // Replace with the desired output file name.

	// Compress the file.
	err := compressFile(sourceFile, destFile)
	if err != nil {
		fmt.Printf("Error compressing file: %v\n", err)
	} else {
		fmt.Printf("File successfully compressed to %s\n", destFile)
	}
}
