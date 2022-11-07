package comsoc

import(
	. "projet/restagentdemo"
)

func CondorcetWinner(p Profile) (bestAlts []int, err error) {
	err = checkProfile(p)
	bestAlts = make([]int, 0)
	for candidat := 1; candidat <= len(p[0]); candidat++ {
		isMajorPref := true
		for candidatcomp := 1; candidatcomp <= len(p[0]); candidatcomp++ {
			if candidat == candidatcomp {
				continue
			}
			candidatScore := 0
			candidatcompScore := 0
			for _, prefs := range p {
				if isPref(candidat, candidatcomp, prefs) {
					candidatScore++
				} else {
					candidatcompScore++
				}
			}
			if candidatcompScore > candidatScore {
				isMajorPref = false
				break
			}
		}
		if isMajorPref {
			bestAlts = append(bestAlts, candidat)
			break
		}
	}
	return bestAlts, err
}
