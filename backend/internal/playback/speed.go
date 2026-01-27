package playback

import (
	"fmt"

	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/effects"
)

type Speed struct {
	streamer *beep.Resampler
	ratio    float64
}

// SetSpeed Ratio must be at least one
func SetSpeed(input *effects.Volume, ratio float64) (*Speed, error) {
	if input == nil {
		return nil, fmt.Errorf("input cannot be nil")
	}

	if ratio <= 0 {
		return nil, fmt.Errorf("invalid ratio: %f", ratio)
	}

	r := beep.ResampleRatio(4, ratio, input)
	return &Speed{streamer: r, ratio: ratio}, nil
}
