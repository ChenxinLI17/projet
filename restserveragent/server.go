package restserveragent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	. "projet/comsoc"
	. "projet/restagentdemo"
	"strconv"
	"strings"
	"sync"
	"time"
)

type RestServerAgent struct {
	sync.Mutex
	id        string
	addr      string
	ballotmap map[string]*Ballot
}

func InStringSlice(haystack []string, needle string) bool {
	for _, e := range haystack {
		if e == needle {
			return true
		}
	}
	return false
}

func DeleteAgentId(haystack []string, needle string) {
	for i, e := range haystack {
		if e == needle {
			haystack[i] = ""
		}
	}
}

func TranslateDate(ddl string) string {
	var date string
	stringarray := strings.Split(ddl, " ")
	switch stringarray[1] {
	case "Jan":
		date = stringarray[5] + "-01-" + stringarray[2] + " " + stringarray[3]
	case "Feb":
		date = stringarray[5] + "-02-" + stringarray[2] + " " + stringarray[3]
	case "Mar":
		date = stringarray[5] + "-03-" + stringarray[2] + " " + stringarray[3]
	case "Apr":
		date = stringarray[5] + "-04-" + stringarray[2] + " " + stringarray[3]
	case "May":
		date = stringarray[5] + "-05-" + stringarray[2] + " " + stringarray[3]
	case "Jun":
		date = stringarray[5] + "-06-" + stringarray[2] + " " + stringarray[3]
	case "Jul":
		date = stringarray[5] + "-07-" + stringarray[2] + " " + stringarray[3]
	case "Aug":
		date = stringarray[5] + "-08-" + stringarray[2] + " " + stringarray[3]
	case "Sept":
		date = stringarray[5] + "-09-" + stringarray[2] + " " + stringarray[3]
	case "Oct":
		date = stringarray[5] + "-10-" + stringarray[2] + " " + stringarray[3]
	case "Nov":
		date = stringarray[5] + "-11-" + stringarray[2] + " " + stringarray[3]
	case "Dec":
		date = stringarray[5] + "-12-" + stringarray[2] + " " + stringarray[3]
	}
	return date
}

func NewRestServerAgent(addr string) *RestServerAgent {
	var rsa = new(RestServerAgent)
	rsa.id = addr
	rsa.addr = addr
	rsa.ballotmap = make(map[string]*Ballot)
	return rsa
}

func (rsa *RestServerAgent) checkMethod(method string, w http.ResponseWriter, r *http.Request) bool {
	if r.Method != method {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "method %q not allowed", r.Method)
		return false
	}
	return true
}

func (*RestServerAgent) decodeBallotRequest(r *http.Request) (req BallotRequest, err error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	err = json.Unmarshal(buf.Bytes(), &req)
	return
}

func (*RestServerAgent) decodeVoteRequest(r *http.Request) (req VoteRequest, err error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	err = json.Unmarshal(buf.Bytes(), &req)
	return
}

func (*RestServerAgent) decodeResultRequest(r *http.Request) (req ResultRequest, err error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	err = json.Unmarshal(buf.Bytes(), &req)
	return
}

func (rsa *RestServerAgent) donewBallot(w http.ResponseWriter, r *http.Request) {
	if !rsa.checkMethod("POST", w, r) {
		w.WriteHeader(501)
		fmt.Fprintf(w, "not implemented")
		return
	}
	req, err := rsa.decodeBallotRequest(r)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "bad request")
		return
	}
	var newballot Ballot
	newballot.Ballot_id = "vote" + strconv.Itoa(req.Alts)
	newballot.Deadline = TranslateDate(req.Deadline)
	newballot.Rule = req.Rule
	fmt.Println(newballot.Rule)
	fmt.Println(newballot.Deadline)
	newballot.Voter_ids = strings.Split(req.Voter_ids, ",")
	fmt.Println(newballot.Voter_ids)
	rsa.ballotmap[newballot.Ballot_id] = &newballot

	var resp BallotResponse
	resp.Ballot_id = "vote" + strconv.Itoa(req.Alts)
	w.WriteHeader(http.StatusOK)
	serial, _ := json.Marshal(resp)
	w.WriteHeader(201)
	fmt.Fprintf(w, "vote créé")
	w.Write(serial)
}

func (rsa *RestServerAgent) doVote(w http.ResponseWriter, r *http.Request) {
	if !rsa.checkMethod("POST", w, r) {
		w.WriteHeader(501)
		fmt.Fprintf(w, "not implemented")
		return
	}
	req, err := rsa.decodeVoteRequest(r)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "bad request")
		return
	}

	ballot, ok := rsa.ballotmap[req.Vote_id]
	if ok != true {
		w.WriteHeader(400)
		fmt.Fprintf(w, "bad request")
		return
	}
	rsa.Lock()
	defer rsa.Unlock()

	deadline, err := time.ParseInLocation("2006-01-02 15:04:05", ballot.Deadline, time.Local)
	if time.Now().After(deadline) {
		w.WriteHeader(503)
		fmt.Fprintf(w, "la deadline est depasse")
		return
	}

	if InStringSlice((*ballot).Voter_ids, req.Agent_id) {
		(*ballot).Prof = append((*ballot).Prof, req.Prefs)
		(*ballot).Options = append((*ballot).Options, req.Options)
		//fmt.Println((*ballot).Prof)
		//fmt.Println((*ballot).Options)
		DeleteAgentId((*ballot).Voter_ids, req.Agent_id)
		w.WriteHeader(200)
		fmt.Fprintf(w, "vote pris en compte")
	} else {
		w.WriteHeader(403)
		fmt.Fprintf(w, "vote déja éffectué")
	}
}

func (rsa *RestServerAgent) returnResult(w http.ResponseWriter, r *http.Request) {
	if !rsa.checkMethod("POST", w, r) {
		w.WriteHeader(501)
		fmt.Fprintf(w, "not implemented")
		return
	}

	rsa.Lock()
	defer rsa.Unlock()
	req, _ := rsa.decodeResultRequest(r)

	ballot, ok := rsa.ballotmap[req.Ballot_id]
	if ok != true {
		w.WriteHeader(404)
		fmt.Fprintf(w, "not found")
		return
	}

	//fmt.Println(ballot.Prof)

	deadline, _ := time.ParseInLocation("2006-01-02 15:04:05", ballot.Deadline, time.Local)
	if !time.Now().After(deadline) {
		w.WriteHeader(425)
		fmt.Fprintf(w, "Too early")
		return
	}

	var resp ResultResponse
	method := ballot.Rule

	switch method {
	case "majority":
		count, _ := MajoritySWF(ballot.Prof)
		fmt.Println(count)
		resp.Ranking = RankCount(count)
		scf, _ := MajoritySCF(ballot.Prof)
		resp.Winner, _ = TieBreak(scf)

	case "borda":
		count, _ := BordaSWF(ballot.Prof)
		resp.Ranking = RankCount(count)
		scf, _ := BordaSCF(ballot.Prof)
		resp.Winner, _ = TieBreak(scf)

	case "approval":
		thresholds := make([]int, len(ballot.Options))
		for i := 0; i < len(ballot.Options); i++ {
			thresholds[i] = ballot.Options[i][0]
		}
		count, _ := ApprovalSWF(ballot.Prof, thresholds)
		fmt.Println(count)
		resp.Ranking = RankCount(count)
		scf, _ := ApprovalSCF(ballot.Prof, thresholds)
		resp.Winner, _ = TieBreak(scf)

	case "stv":
		count, _ := STV_SWF(ballot.Prof)
		fmt.Println(count)
		resp.Ranking = RankCount(count)
		scf, _ := STV_SCF(ballot.Prof)
		resp.Winner, _ = TieBreak(scf)

	case "kemeny":
		resp.Ranking = KemenySWF(ballot.Prof)
		resp.Winner = KemenySCF(ballot.Prof)

	case "copeland":
		count, _ := CopelandSWF(ballot.Prof)
		resp.Ranking = RankCount(count)
		scf, _ := CopelandSCF(ballot.Prof)
		resp.Winner, _ = TieBreak(scf)
	}

	w.WriteHeader(http.StatusOK)
	serial, _ := json.Marshal(resp)
	w.Write(serial)

}

func (rsa *RestServerAgent) Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/new_ballot", rsa.donewBallot)
	mux.HandleFunc("/vote", rsa.doVote)
	mux.HandleFunc("/result", rsa.returnResult)

	s := &http.Server{
		Addr:           rsa.addr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20}

	log.Println("Listening on", rsa.addr)
	go log.Fatal(s.ListenAndServe())
}
