package comsoc

import (
	"fmt"
	. "projet/restagentdemo"
)

type Pair struct {
	e1 int
	e2 int
}

func (p *Pair) Equal(p2 Pair) bool {
	if p.e1 == p2.e1 && p.e2 == p2.e2 {
		return true
	}
	return false
}

func CalculDistEdition(pref1 []int, pref2 []int) float64 {
	pairs1 := make([]Pair, 0)
	pairs2 := make([]Pair, 0)
	for i := 0; i < len(pref1); i++ {
		for j := i + 1; j <= len(pref1)-1; j++ {
			pairs1 = append(pairs1, Pair{int(pref1[i]), int(pref1[j])})
		}
	}
	for i := 0; i < len(pref1); i++ {
		for j := i + 1; j <= len(pref1)-1; j++ {
			pairs2 = append(pairs2, Pair{int(pref2[i]), int(pref2[j])})
		}
	}

	count := 0
	for _, pair1 := range pairs1 {
		for _, pair2 := range pairs2 {
			if pair1.Equal(pair2) {
				count++
				break
			}
		}
	}
	disEdition := len(pairs1) - count
	taux := float64((count - disEdition)) / float64(len(pairs1))
	return taux
}

func CalculDistEdiRangeProfile(pref []int, profil Profile) float64 {
	sum := 0.0
	for _, prefprofil := range profil {
		sum += CalculDistEdition(pref, prefprofil)
	}
	return sum
}

func KemenySWF(p Profile) (alts []int) {
	minIndex := 0
	mindistance := CalculDistEdiRangeProfile(p[0], p)
	for i, pref := range p {
		fmt.Printf("rangement %d : %v\n", i, pref)
		distance := CalculDistEdiRangeProfile(pref, p)
		if mindistance > distance {
			mindistance = distance
			minIndex = i
		}
	}
	return p[minIndex]
}

func KemenySCF(p Profile) (alt int) {
	alts := KemenySWF(p)
	return alts[0]
}
