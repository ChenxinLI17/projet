package comsoc

import (
	"errors"
	. "projet/restagentdemo"
)

// en cas d'égalité
// pouvoir utiliser différents critères :
// ex: choisir toujours le premier candidat
// ex: choisir le candidat qui est le plus proche de ...

func TieBreak(alts []int) (alt int, err error) {
	if len(alts) == 0 {
		return -1, errors.New("alternative est vide")
	}
	return alts[0], nil
}

func TieBreakFactory(altsFac []int) func([]int) (int, error) {
	AltPreOrder := make([]int, len(altsFac))
	copy(AltPreOrder, altsFac)
	tiebreak := func(alts []int) (int, error) {
		if len(alts) == 0 {
			return -1, errors.New("alternative est vide")
		}
		result := alts[0]
		for _, a := range alts {
			if isPref(a, result, AltPreOrder) {
				result = a
			}
		}
		return result, nil
	}
	return tiebreak
}

func SWFFactory(swf func(Profile) (Count, error), tiebreak func([]int) (int, error)) func(Profile) ([]int, error) {
	SWFProduct := func(p Profile) ([]int, error) {
		count, errSWF := swf(p)
		if errSWF != nil {
			return nil, errSWF
		}
		orderStrict := make([]int, len(count))
		bestAlts := maxCount(count)
		if len(bestAlts) == 0 {
			return nil, errors.New("count est vide")
		} else if len(bestAlts) == 1 {
			orderStrict = append(orderStrict, bestAlts[0])
		} else {
			// append to order list and remove from bestAlts one after another
			for len(bestAlts) != 0 {
				alt, err := tiebreak(bestAlts)
				if err != nil {
					return nil, err
				}
				orderStrict = append(orderStrict, alt)
				// remove all proceeded alts from count map
				delete(count, alt)
				bestAlts = maxCount(count)
			}
		}
		return orderStrict, nil
	}
	return SWFProduct
}

func SCFFactory(scf func(p Profile) (Count, error), tiebreak func([]int) (int, error)) func(Profile) (int, error) {
	SCFProduct := func(p Profile) (int, error) {
		count, errSCF := scf(p)
		bestAlts := maxCount(count)
		bestAlt, errTB := tiebreak(bestAlts)

		if errSCF != nil {
			return -1, errSCF
		}
		if errTB != nil {
			return -1, errTB
		}
		return bestAlt, nil
	}
	return SCFProduct
}
