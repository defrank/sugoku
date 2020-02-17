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

	"github.com/defrank/sugoku/sudoku"
	"github.com/gdamore/tcell"
	"github.com/urfave/cli/v2" // imports as package "cli"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

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
		return nil
	}
	app.Commands = []*cli.Command{
		&cli.Command{
			Name:    "play",
			Aliases: []string{"p"},
			Usage:   "Play Sudoku of default size 9",
			Action:  play,
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

func play(c *cli.Context) error {
	s := initScreen()
	// _, screenHeight := s.Size()

	st := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	s.SetStyle(st)

	g := sudoku.NewGrid(9)
	for box := range g.Iter() {
		if v := rand.Int() % 10; v != 0 {
			g.Set(box.X, box.Y, v)
		}
	}

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
		makegrid(s, g)
		cnt++
		dur += time.Now().Sub(start)
	}

	s.Fini()
	return nil
}

func makegrid(s tcell.Screen, g *sudoku.Grid) {
	size := g.Size()
	w, h := s.Size()

	if w < size || h < size {
		return
	}

	st := tcell.StyleDefault.
		Foreground(tcell.ColorGreen).
		Background(tcell.ColorBlack)

	glyphs := []rune{' ', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

	for box := range g.Iter() {
		s.SetCell(box.Y, box.X, st, glyphs[box.V])
	}

	s.Show()
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
