/**********************************************************************
 * References:
 *   https://github.com/urfave/cli/blob/master/docs/v2/manual.md
 *   https://appliedgo.net/tui/
 *********************************************************************/
package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/gdamore/tcell"
	"github.com/urfave/cli/v2" // imports as package "cli"
	// "github.com/defrank/sugoku/sudoku"
)

func main() {
	app := cli.NewApp()
	app.Name = "Sugoku CLI"
	app.Usage = "Sudoku CLI written in Go!"
	app.Authors = []*cli.Author{
		&cli.Author{
			Name:  "Derek F",
			Email: "derek@frank.sh",
		},
	}
	app.Version = "0.1.0"
	app.Action = func(c *cli.Context) error {
		fmt.Println(fmt.Sprintf("Welcome to %s!", app.Name))
		s := initScreen()
		// _, screenHeight := s.Size()

		quit := make(chan struct{})
		go func() {
			for {
				ev := s.PollEvent()
				switch ev := ev.(type) {
				case *tcell.EventKey:
					switch ev.Key() {
					case tcell.KeyEscape, tcell.KeyEnter:
						close(quit)
						return
					case tcell.KeyCtrlL:
						s.Sync()
					}
				case *tcell.EventResize:
					s.Sync()
				}
			}
		}()

		cnt := 0
		dur := time.Duration(0)
	loop:
		for {
			select {
			case <-quit:
				break loop
			case <-time.After(time.Millisecond * 50):
			}
			start := time.Now()
			makebox(s)
			cnt++
			dur += time.Now().Sub(start)
		}

		//s.Show()
		s.Fini()
		return nil
	}
	app.Commands = []*cli.Command{
		&cli.Command{
			Name:    "play",
			Aliases: []string{"p"},
			Usage:   "Play Sudoku of default size 9",
			Action: func(c *cli.Context) error {
				fmt.Println("Playing Sugoku!")
				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func initScreen() tcell.Screen {
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	s, err := tcell.NewScreen()
	if err != nil {
		log.Printf("%v\n", err)
		os.Exit(1)
	}
	if err = s.Init(); err != nil {
		log.Printf("%v\n", err)
		os.Exit(1)
	}
	s.SetStyle(tcell.StyleDefault.
		Foreground(tcell.ColorGreen).
		Background(tcell.ColorBlack))
	s.Clear()

	return s
}

func makebox(s tcell.Screen) {
	w, h := s.Size()

	if w == 0 || h == 0 {
		return
	}

	glyphs := []rune{'@', '#', '&', '*', '=', '%', 'Z', 'A'}

	lx := rand.Int() % w
	ly := rand.Int() % h
	lw := rand.Int() % (w - lx)
	lh := rand.Int() % (h - ly)
	st := tcell.StyleDefault
	gl := ' '
	if s.Colors() > 256 {
		rgb := tcell.NewHexColor(int32(rand.Int() & 0xffffff))
		st = st.Background(rgb)
	} else if s.Colors() > 1 {
		st = st.Background(tcell.Color(rand.Int() % s.Colors()))
	} else {
		st = st.Reverse(rand.Int()%2 == 0)
		gl = glyphs[rand.Int()%len(glyphs)]
	}

	for row := 0; row < lh; row++ {
		for col := 0; col < lw; col++ {
			s.SetCell(lx+col, ly+row, st, gl)
		}
	}
	s.Show()
}
