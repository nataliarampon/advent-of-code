package main

type SeedRange struct {
	start      int
	rangeValue int
	end        int
}

func newSeedRange(start int, rangeValue int) *SeedRange {
	seedRange := SeedRange{start: start, rangeValue: rangeValue}
	seedRange.end = seedRange.start + seedRange.rangeValue - 1
	return &seedRange
}

type SeedMapping struct {
	start            int
	rangeValue       int
	end              int
	destinationStart int
}

func newSeedMapping(start int, rangeValue int, destination int) *SeedMapping {
	seedMapping := SeedMapping{start: start, rangeValue: rangeValue, destinationStart: destination}
	seedMapping.end = seedMapping.start + seedMapping.rangeValue - 1
	return &seedMapping
}