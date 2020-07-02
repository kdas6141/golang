package main
import (
	"fmt"
	"os"
	"io/ioutil"
	"path/filepath"
	"flag"
)

func walkDir(dir string, filesizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, filesizes)
		} else {
			filesizes <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

func main() {
	// detrmine the initial directories
	flag.Parse()
	roots := flag.Args()
	if (len(roots) == 0) {
		roots = []string{"."}
	}
	// traveres the file tree
	filesizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, filesizes)
		}
		close(filesizes)
	}()
	// print the results
	var nfiles, nbytes int64
	for size := range filesizes {
		nfiles++
		nbytes += size
	}
	printDiskUsages(nfiles, nbytes)
}

func printDiskUsages(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
}