package playback

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/effects"
	"github.com/gopxl/beep/v2/mp3"
	"github.com/gopxl/beep/v2/speaker"
)

func RunPlayer() {
	f, err := os.Open("backend/fma_test_file.mp3")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	sr := format.SampleRate * 2
	speaker.Init(sr, sr.N(time.Second/10))

	resampled := beep.Resample(4, format.SampleRate, sr, streamer)

	//loop := beep.Loop(3, streamer)
	ctrl := &beep.Ctrl{Streamer: resampled, Paused: false}
	volume := &effects.Volume{
		Streamer: ctrl,
		Base:     2,
		Volume:   0,
		Silent:   false,
	}

	speedy, err := SetSpeed(volume, 2)
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan bool)
	speaker.Play(beep.Seq(speedy.streamer, beep.Callback(func() {
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

	interval := time.Duration(float64(time.Second) / speedy.ratio)
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	fmt.Println("Press [ENTER] to pause/resume. ")
	var playbackPosition time.Duration
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
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
