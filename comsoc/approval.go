package comsoc

import(
	. "projet/restagentdemo"
)

func ApprovalSWF(p Profile, thresholds []int) (count Count, err error){
	err = checkProfile(p)
	if count == nil{
		count = make(map[int]int)
	}
	for i,pref := range p{
		for j:=0; j<thresholds[i]; j++{
			count[pref[j]] ++
		}
	}
	return
}

func ApprovalSCF(p Profile, thresholds []int) (bestAlts []int, err error){
	err = checkProfile(p)
	count, err := ApprovalSWF(p,thresholds)
	bestAlts = maxCount(count)
	return
}