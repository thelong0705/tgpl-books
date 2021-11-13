package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}

type customSort struct {
	tracks []*Track
	less   func(i, j *Track) bool
}

func (s *customSort) Len() int {
	return len(s.tracks)
}

func (s *customSort) Less(i, j int) bool {
	return s.less(s.tracks[i], s.tracks[j])
}

func (s *customSort) Swap(i, j int) {
	s.tracks[i], s.tracks[j] = s.tracks[j], s.tracks[i]
}

func main() {
	sort.Sort(&customSort{
		tracks: tracks,
		less: func(i, j *Track) bool {
			if i.Title != j.Title {
				return i.Title < j.Title
			}
			if i.Year != j.Year {
				return i.Year < j.Year
			}
			if i.Length != j.Length {
				return i.Length < j.Length
			}
			return false
		},
	})

	printTracks(tracks)
}
