package playback

import "time"

func makeNewTicker(r float64) *time.Ticker {
	interval := time.Duration(float64(time.Second) / r)
	ticker := time.NewTicker(interval)

	return ticker
}

func recreateTicker(ticker *time.Ticker, r float64) *time.Ticker {
	ticker.Stop()
	return makeNewTicker(r)
}
