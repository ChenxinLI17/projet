package comsoc

import(
	. "projet/restagentdemo"
)

func MajoritySWF(p Profile) (count Count, err error) {
	err = checkProfile(p)
	if count == nil {
		count = make(map[int]int)
		for _, pref := range p {
			count[pref[0]]++
		}
	}
	return
}

func MajoritySCF(p Profile) (bestAlts []int, err error) {
	err = checkProfile(p)
	count, err := MajoritySWF(p)
	bestAlts = maxCount(count)
	return
}
