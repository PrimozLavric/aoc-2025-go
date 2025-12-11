package day_1

type SafeDial struct {
	dialValue         int
	zeroHits          int
	zeroHitsAndPasses int
}

func NewSafeDial() *SafeDial {
	return &SafeDial{dialValue: 50, zeroHits: 0}
}

func (sd *SafeDial) GetZeroHits() int {
	return sd.zeroHits
}

func (sd *SafeDial) GetZeroHitsAndPasses() int {
	return sd.zeroHitsAndPasses
}

func (sd *SafeDial) ExecuteRotationSequence(rotationSeq []Rotation) (int, int) {

	for _, rot := range rotationSeq {
		newDialValue, zeroPointPasses := rot.apply(sd.dialValue)
		sd.dialValue = newDialValue
		sd.zeroHitsAndPasses += zeroPointPasses

		if sd.dialValue == 0 {
			sd.zeroHits += 1
		}
	}

	return sd.zeroHits, sd.zeroHitsAndPasses
}
