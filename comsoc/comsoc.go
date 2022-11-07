package comsoc

import (
	"errors"
	. "projet/restagentdemo"
	"sort"
)

// renvoie l'indice ou se trouve alt dans prefs
func rank(alt int, prefs []int) int {
	for i, pref := range prefs {
		if pref == alt {
			return i
		}
	}
	return -1
}

// renvoie vrai ssi alt1 est préférée à alt2
func isPref(alt1, alt2 int, prefs []int) bool {
	pref1 := rank(alt1, prefs)
	pref2 := rank(alt2, prefs)
	if pref1 < pref2 {
		return true
	} else {
		return false
	}
}

// renvoie les meilleures alternatives pour un décomtpe donné
func maxCount(count Count) (bestAlts []int) {
	maxc := 0
	for _, v := range count {
		if v > maxc {
			maxc = v
		}
	}
	for k, v := range count {
		if maxc == v {
			bestAlts = append(bestAlts, k)
		}
	}
	return
}

// vérifie le profil donné, par ex. qu'ils sont tous complets et que chaque alternative n'apparaît qu'une seule fois par préférences
func checkProfile(prefs Profile) error {
	for _, prefsIndivid := range prefs {
		for _, pref := range prefsIndivid {
			if pref == 0 {
				return errors.New("profil non complet")
			}
		}
		prefsIndividCopy := make([]int, len(prefsIndivid))
		copy(prefsIndividCopy, prefsIndivid)
		sort.Slice(prefsIndividCopy, func(m, n int) bool { return prefsIndividCopy[m] < prefsIndividCopy[n] })
		InferPref := -1
		for _, pref := range prefsIndividCopy {
			if InferPref == pref {
				return errors.New("doublon preference")
			}
			InferPref = pref
		}
	}
	return nil
}

// vérifie le profil donné, par ex. qu'ils sont tous complets et que chaque alternative de alts apparaît exactement une fois par préférences
func checkProfileAlternative(prefs Profile, alts []int) error {
	checkProfile(prefs)
	altsCopy := make([]int, len(alts))
	copy(altsCopy, alts)
	sort.Slice(altsCopy, func(m, n int) bool {
		return altsCopy[m] < altsCopy[n]
	})
	InferAlt := -1
	for _, alt := range altsCopy {
		if InferAlt == alt {
			return errors.New("doublon alts")
		}
		InferAlt = alt
	}
	return nil
}

func RankCount(count Count) []int {
	type temp struct {
		key int
		val int
	}
	var temps []temp
	for k, v := range count {
		temps = append(temps, temp{k, v})
	}
	sort.Slice(temps, func(i, j int) bool { return temps[i].val > temps[j].val })

	rankresult := make([]int, len(temps))
	for i, ele := range temps {
		rankresult[i] = ele.key
	}
	return rankresult
}
