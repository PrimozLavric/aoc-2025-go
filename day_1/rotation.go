package day_1

type Dir int

const (
	DirLeft Dir = iota
	DirRight
)

type Rotation struct {
	Dir       Dir
	NumClicks int
}

func wraparoundMod(a, b int) int {
	m := a % b
	if m < 0 {
		m += b
	}
	return m
}

func (r Rotation) apply(dialPosition int) (int, int) {
	if r.Dir == DirLeft {
		if dialPosition == 0 {
			// [Special case]: Starting at 0
			return wraparoundMod(-r.NumClicks, 100), r.NumClicks / 100
		}

		if r.NumClicks < dialPosition {
			// [Special case]: Never wraparound
			return dialPosition - r.NumClicks, 0
		}

		clicksAfterFirstZero := r.NumClicks - dialPosition
		zeroHits := 1 + (clicksAfterFirstZero / 100)

		return wraparoundMod(dialPosition-r.NumClicks, 100), zeroHits
	}

	newPos := dialPosition + r.NumClicks
	return newPos % 100, newPos / 100
}
