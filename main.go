package main

import (
	"github.com/nsf/termbox-go"
	"time"
	"math/rand"
)

type Point struct{ X, Y int }

var (
	width, height int
	snake         []Point
	dir           = Point{1, 0}
	food          Point
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	width, height = termbox.Size()
	snake = append(snake, Point{width / 2, height / 2})
	rand.Seed(time.Now().UnixNano())
	food = Point{rand.Intn(width), rand.Intn(height)}

	go func() {
		for {
			switch ev := termbox.PollEvent(); ev.Type {
			case termbox.EventKey:
				switch ev.Key {
				case termbox.KeyArrowUp:
					dir = Point{0, -1}
				case termbox.KeyArrowDown:
					dir = Point{0, 1}
				case termbox.KeyArrowLeft:
					dir = Point{-1, 0}
				case termbox.KeyArrowRight:
					dir = Point{1, 0}
				}
			}
		}
	}()

	for {
		draw()
		time.Sleep(200 * time.Millisecond)
		head := snake[0]
		next := Point{head.X + dir.X, head.Y + dir.Y}
		if next.X < 0 || next.Y < 0 || next.X >= width || next.Y >= height {
			return
		}
		for _, v := range snake {
			if v == next {
				return
			}
		}
		snake = append([]Point{next}, snake...)
		if next == food {
			food = Point{rand.Intn(width), rand.Intn(height)}
		} else {
			snake = snake[:len(snake)-1]
		}
	}
}

func draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	for _, v := range snake {
		termbox.SetCell(v.X, v.Y, 'o', termbox.ColorGreen, termbox.ColorBlack)
	}
	termbox.SetCell(food.X, food.Y, 'x', termbox.ColorRed, termbox.ColorBlack)
	termbox.Flush()
}
