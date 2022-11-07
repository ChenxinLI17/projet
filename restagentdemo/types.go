package restagentdemo

type Profile [][]int
type Count map[int]int

type BallotRequest struct {
	Rule      string `json:"rule"`
	Deadline  string `json:"deadline"`
	Voter_ids string `json:"voter-ids"`
	Alts      int    `json:"#alts"`
}

type Ballot struct {
	Ballot_id string
	Rule      string
	Deadline  string
	Voter_ids []string
	Alts      int
	Prof      Profile
	Options   [][]int
}

type BallotResponse struct {
	Ballot_id string `json:"ballot-id"`
}

type VoteRequest struct {
	Agent_id string `json:"agent-id"`
	Vote_id  string `json:"vote-id"`
	Prefs    []int  `json:"prefs"`
	Options  []int  `json:"options"`
}

type ResultRequest struct {
	Ballot_id string `json:"ballot-id"`
}

type ResultResponse struct {
	Winner  int   `json:"winner"`
	Ranking []int `json:"ranking"`
}
