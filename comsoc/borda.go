package comsoc

import(
	. "projet/restagentdemo"
)

func BordaSWF(p Profile) (Count, error) {
	err := checkProfile(p)
	count := make(map[int]int)
	for _, prefsIndivid := range p {
		highScore := len(p[0]) - 1
		for _, pref := range prefsIndivid {
			count[pref] += highScore
			highScore--
		}
	}
	return count,err
}

func BordaSCF(p Profile) (bestAlts []int, err error) {
	count, err := BordaSWF(p)
	bestAlts = maxCount(count)
	return
}
