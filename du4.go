package main
import (
	"fmt"
	"os"
	"io/ioutil"
	"path/filepath"
	"flag"
	"time"
	"sync"
)

func walkDir(dir string, n *sync.WaitGroup, filesizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, filesizes)
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

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <- done:
		return true
	default:
		return false
	}
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
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, filesizes)
	}
	go func() {
		n.Wait()	
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