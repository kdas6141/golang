package main
import (
	"fmt"
	"os"
	"io/ioutil"
	"path/filepath"
	"flag"
	"time"
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

var verbose = flag.Bool("v", false, "show verbose progress message")

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
	// print the results periodically
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-filesizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsages(nfiles, nbytes)		
		}
	}
	printDiskUsages(nfiles, nbytes)
}

func printDiskUsages(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
}