package playback

import (
	"fmt"

	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/effects"
)

func NewVolume(ctrl *beep.Ctrl, base, vol float64) *effects.Volume {
	return &effects.Volume{
		Streamer: ctrl,
		Base:     base,
		Volume:   vol,
		Silent:   false,
	}
}

type VolumeBoost struct {
	boosted bool
	step    float64
}

func NewVolumeBoost() *VolumeBoost {
	return &VolumeBoost{step: 1.0}
}

func (b *VolumeBoost) Toggle(v *effects.Volume) error {
	if v == nil {
		return fmt.Errorf("input cannot be nil")
	}

	if b.boosted {
		v.Volume -= b.step
	} else {
		v.Volume += b.step
	}
	b.boosted = !b.boosted
	return nil
}

func toggleMute(v *effects.Volume) error {
	if v == nil {
		return fmt.Errorf("input cannot be nil")
	}
	v.Silent = !v.Silent
	return nil
}
