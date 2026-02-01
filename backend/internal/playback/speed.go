package playback

import (
	"fmt"

	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/effects"
)

type Speed struct {
	resampler    *beep.Resampler
	ratio        float64
	ratioChanged chan float64
}

// NewSpeed Ratio must be at least one
func NewSpeed(input *effects.Volume, ratio float64) (*Speed, error) {
	if input == nil {
		return nil, fmt.Errorf("input cannot be nil")
	}

	if ratio <= 0 {
		return nil, fmt.Errorf("invalid ratio: %f", ratio)
	}

	r := beep.ResampleRatio(4, ratio, input)
	return &Speed{resampler: r, ratio: ratio, ratioChanged: make(chan float64, 1)}, nil
}

// ChangeRatio between 0.5 and 4, should be called under speaker.Lock()
func (s *Speed) ChangeRatio(r float64) error {
	if r < 0.5 || r > 4 {
		return fmt.Errorf("invalid ratio: %f", r)
	}

	s.ratio = r
	s.resampler.SetRatio(r)

	select {
	case s.ratioChanged <- r:
	default:
	}
	return nil
}
