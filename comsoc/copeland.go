package comsoc

import(
	. "projet/restagentdemo"
)

func CopelandSWF(p Profile) (count Count, err error) {
	err = checkProfile(p)
	count = make(Count)
	for candidat := 1; candidat <= len(p[0]); candidat++ {
		sum := 0
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
			if candidatScore > candidatcompScore {
				sum++
			} else {
				sum--
			}
		}
		count[candidat] = sum
	}
	return
}
func CopelandSCF(p Profile) (bestAlts []int, err error) {
	err = checkProfile(p)
	count, err := CopelandSWF(p)
	if err != nil {
		return nil, err
	}
	bestAlts = maxCount(count)
	return
}
