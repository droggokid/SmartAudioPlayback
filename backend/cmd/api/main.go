package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func main() {
	f, err := os.Open("backend/fma_test_file.mp3")
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	sr := format.SampleRate * 2
	speaker.Init(sr, sr.N(time.Second/10))

	resampled := beep.Resample(4, format.SampleRate, sr, streamer)

	//loop := beep.Loop(3, streamer)
	//fast := beep.ResampleRatio(4, 5, loop)
	ctrl := beep.Ctrl{Streamer: resampled, Paused: false}

	done := make(chan bool)
	speaker.Play(beep.Seq(&ctrl, beep.Callback(func() {
		done <- true
	})))

	input := make(chan string)
	go func() {
		r := bufio.NewReader(os.Stdin)
		for {
			line, err := r.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			input <- strings.ToLower(strings.TrimSpace(line))
		}
	}()

	fmt.Println("Press [ENTER] to pause/resume. ")
	var playbackPosition time.Duration
	for {
		select {
		case <-done:
			return
		case <-time.After(time.Second):
			speaker.Lock()
			timeCheck := format.SampleRate.D(streamer.Position()).Round(time.Second)
			if timeCheck > playbackPosition {
				playbackPosition = timeCheck
				fmt.Println(playbackPosition)
			}
			speaker.Unlock()
		case s := <-input:
			switch s {
			case "":
				speaker.Lock()
				ctrl.Paused = !ctrl.Paused
				speaker.Unlock()
			case "q":
				return
			}
		}
	}
}
