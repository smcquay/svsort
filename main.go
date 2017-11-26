// command svsort sorts lines of text by semver at beginning of line
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/blang/semver"
)

type line struct {
	v    semver.Version
	line string
}

type lines []line

func (l lines) Len() int           { return len(l) }
func (l lines) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l lines) Less(i, j int) bool { return l[i].v.LT(l[j].v) }

func main() {
	s := bufio.NewScanner(os.Stdin)

	ls := lines{}
	for s.Scan() {
		f := strings.Fields(s.Text())
		if len(f) < 2 {
			continue
		}
		v, err := semver.ParseTolerant(f[0])
		if err != nil {
			log.Printf("%v: %q", err, s.Text())
			continue
		}
		l := line{
			v:    v,
			line: s.Text(),
		}
		ls = append(ls, l)
	}

	if err := s.Err(); err != nil {
		log.Fatalf("scan: %v")
	}

	sort.Sort(ls)
	for _, l := range ls {
		fmt.Printf("%v\n", l.line)
	}
}
