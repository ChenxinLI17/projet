package comsoc

import (
	. "projet/restagentdemo"
	"sort"
)

func minCount(count Count) int {
	type temp struct {
		key int
		val int
	}
	var temps []temp
	for k, v := range count {
		temps = append(temps, temp{k, v})
	}
	sort.Slice(temps, func(i, j int) bool { return temps[i].val < temps[j].val })

	rankresult := make([]int, len(temps))
	for i, ele := range temps {
		rankresult[i] = ele.key
	}
	return rankresult[0]
}

func STV_SWF(p Profile) (count Count, err error) {
	err = checkProfile(p)
	count = make(Count)
	pcopy := make(Profile, len(p))
	for i, _ := range pcopy {
		pcopy[i] = make([]int, len(p[i]))
		copy(pcopy[i], p[i])
	}
	n := len(p[0])
	for t := 1; t < n; t++ {
		counttemp := make(Count)
		for _, pref := range pcopy[0] {
			if pref != 0 {
				counttemp[pref] = 0
			}
		}
		for _, prefs := range pcopy {
			i := 0
			for prefs[i] == 0 {
				i++
			}
			counttemp[prefs[i]]++
		}
		var elim int
		elim = minCount(counttemp)
		count[elim] = t
		for i, prefs := range pcopy {
			for j, _ := range prefs {
				if pcopy[i][j] == elim {
					pcopy[i][j] = 0
				}
			}
		}
	}
	count[p[0][0]] = n
	return
}

func STV_SCF(p Profile) (bestAlts []int, err error) {
	err = checkProfile(p)
	count, err := STV_SWF(p)
	if err != nil {
		return nil, err
	}
	bestAlts = maxCount(count)
	return
}
