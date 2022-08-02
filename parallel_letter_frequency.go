package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	totalFreq := FreqMap{}
	ch := make(chan FreqMap)

	for _, s := range l {
		go func(s1 string) {
			m := Frequency(s1)
			ch <- m
		}(s)
	}

	for range l {
		for k, v := range <-ch {
			totalFreq[k] += v
		}
	}

	return totalFreq
}
