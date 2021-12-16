package main

import (
	"fmt"
	"gopl.io/ch8/thumbnail"
	"os"
	"sync"
)

func main() {
	filenames := make(chan string, 3)
	filenames <- "pic.jpg"
	filenames <- "pic12.jpg"
	filenames <- "pic.jpg"
	close(filenames)
	size := makeThumbnails6(filenames)

	fmt.Println(size)
}

func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)

	var wg sync.WaitGroup

	for filename := range filenames {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()

			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				fmt.Println(err)
				return
			}
			info, _ := os.Stat(thumb)
			sizes <- info.Size()
		}(filename)
	}

	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}

	return total

}

func makeThumbnails2(filenames []string) {
	for _, f := range filenames {
		go thumbnail.ImageFile(f)
	}
}

func makeThumbnails3(filenames []string) {
	ch := make(chan struct{})
	fmt.Println(cap(ch))
	for _, filename := range filenames {
		go func(f string) {
			thumbnail.ImageFile(f)
			ch <- struct{}{}
		}(filename)
	}
	for range filenames {
		<-ch
	}
}

func makeThumbnails4(filenames []string) error {
	errors := make(chan error)

	for _, f := range filenames {
		go func(f string) {
			_, err := thumbnail.ImageFile(f)
			errors <- err
		}(f)
	}

	for range filenames {
		if err := <-errors; err != nil {
			return err
		}
	}

	return nil
}

func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err       error
	}

	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbfile, it.err = thumbnail.ImageFile(f)
			ch <- it
		}(f)
	}

	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}

	return thumbfiles, nil
}
