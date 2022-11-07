Le serveur gère les requêtes POST reçues par les trois routes.

/new_ballot : Responsable du traitement des requêtes des clients pour la création de nouveaux bulletins de vote. Le serveur dispose d'une map <ballotmap> qui stocke les bulletins, en utilisant le ballot_id comme key et la structure du bulletin comme value.

type Ballot struct {
	Ballot_id string
	Rule      string
	Deadline  string
	Voter_ids []string
	Alts      int
	Prof      Profile
	Options   [][]int
}

Le contenu de la requête est reçu dans le format de structure <BallotRequest>. Ensuite, le serveur génère le ballot_id et crée une nouvelle instance de <Ballot> structure, convertit les informations reçues au format json dans le format correspondant et les stocke dans l'instance. L'information de deadline, "Tue Nov 10 23:00:00 UTC 2009" est converti en "2009-11-10 23:00:00" et puis stocké dans le valeur de <Ballot.Deadline>. <BallotRequest.Voters_id> obtenu en tant que type de chaîne "ag_id1, ag_id2, ag_id3" est converti en telles chaînes "ag_id1","ag_id2" ,"ag_id3"et stockés dans le tableau de chaîne <Ballot.Voters_id>. Après avoir terminé l'initialisation du bulletin de vote, le serveur renvoie un code d'état 200 et <ballot_id> au client.

/vote : Responsable de l'enregistrement du vote du client

Le serveur trouve le bulletin de vote correspondant dans le map <ballotmap> en fonction du <vote_id> du client. Le serveur vérifie si ce <agent_id> qui a envoyé la requête est dans le <Ballot.Voter_ids> et si la deadline n'est pas dépassée. Si oui, il ajoute la préférence du client <VoteRequest.Prefs> à le profile de vote <Ballot.Prof>, ajoute l'option de vote <VoteRequest.Options> à <Ballot.Options>. Après, le serveur supprime <Agent_id> de ce client de <Ballot.Voter-ids> pour éviter la double enregistrement de vote.

/result : Responsable d'obtention et renvoi de résultats du vote

Le serveur vérifie si ce <ballot_id> qui a envoyé la requête est dans le <ballotmap> et si la deadline de vote est dépassée. Si oui, il compter les résultats du vote selon ce <Ballot.Rule>. <Ballot.Rule> inclue "majority", "borda", "approval", "stv", "kemeny", "copeland".