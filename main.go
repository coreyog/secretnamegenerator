//go:generate go-bindata -pkg embedded -prefix assets/ -o ./embedded/embedded.go -pkg embedded assets
package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"math/rand/v2"
	"os"
	"runtime/debug"
	"strings"

	"github.com/jessevdk/go-flags"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	AdjectiveCount uint32 = 1305
	NounCount      uint32 = 1520
)

type Arguments struct {
	LowerCase bool `short:"l" long:"lower" description:"Lowercase output"`
	TitleCase bool `short:"t" long:"title" description:"Titlecase output"`
	OneWord   bool `short:"1" long:"one" description:"Output as one word"`
	Dash      bool `short:"d" long:"dash" description:"Output with a dash"`
}

func (a *Arguments) Validate() error {
	if a.LowerCase && a.TitleCase {
		return fmt.Errorf("cannot specify both lower and title case")
	}

	if a.OneWord && a.Dash {
		return fmt.Errorf("cannot specify both one and dash")
	}

	return nil
}

var args Arguments

//go:embed assets/adjectives.txt
var rawAdj []byte

//go:embed assets/nouns.txt
var rawNoun []byte

func main() {
	_, err := flags.Parse(&args)
	if err != nil {
		if flags.WroteHelp(err) {
			info, ok := debug.ReadBuildInfo()
			if ok {
				fmt.Println(info.Main.Version)
			}

			return
		}

		os.Exit(1)
	}

	err = args.Validate()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	adj := wordFromList(rawAdj, AdjectiveCount)
	noun := wordFromList(rawNoun, NounCount)

	if args.TitleCase {
		titler := cases.Title(language.English)
		adj = titler.String(adj)
		noun = titler.String(noun)
	} else if args.LowerCase {
		adj = strings.ToLower(adj)
		noun = strings.ToLower(noun)
	}

	sep := " "
	if args.OneWord {
		sep = ""
	} else if args.Dash {
		sep = "-"
	}

	fmt.Printf("%s%s%s\n", adj, sep, noun)
}

func wordFromList(list []byte, count uint32) string {
	scanner := bufio.NewScanner(bytes.NewBuffer(list))
	scanner.Scan() // must scan at least once to get anything from the scanner

	num := rand.Uint32() % count

	for range num {
		scanner.Scan()
	}

	return scanner.Text()
}
