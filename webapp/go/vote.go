package main

import (
	"context"
	"strings"
)

// Vote Model
type Vote struct {
	ID          int
	UserID      int
	CandidateID int
	Keyword     string
}

func getVoteCountByCandidateID(candidateID int) (count int) {
	row := db.QueryRow("SELECT IFNULL(SUM(count), 0) AS count FROM votes WHERE candidate_id = ?", candidateID)
	row.Scan(&count)
	return
}

func getUserVotedCount(userID int) (count int) {
	row := db.QueryRow("SELECT count FROM votes WHERE user_id = ?", userID)
	row.Scan(&count)
	return
}

func createVote(c context.Context, userID int, candidateID int, keyword string, count int) error {
	tx, err := db.BeginTx(c, nil)
	if err != nil {
		return err
	}
	_, err = db.ExecContext(c, "INSERT INTO votes (user_id, candidate_id, keyword, count) VALUES (?, ?, ?, ?)",
		userID, candidateID, keyword, count)
	if err != nil {
		tx.Rollback()
		return err
	}
	rows, err := db.QueryContext(c,
		"SELECT id from candidate_keywords WHERE candidate_id = ? AND content = ? FOR UPDATE",
		candidateID,
		keyword,
	)
	if err != nil {
		tx.Rollback()
		return err
	}
	var can_key_id int64
	rows.Scan(&can_key_id)
	rows.Close()
	if &can_key_id == nil {
		_, err = db.ExecContext(c,
			"INSERT INTO candidate_keywords (candidate_id, content, count) VALUES (?, ?, ?)",
			candidateID,
			keyword,
			count,
		)
	} else {
		_, err = db.ExecContext(c,
			"UPDATE candidate_keywords SET count = count + ? WHERE id = ?",
			count,
			can_key_id,
		)
	}
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	return err
}

func getVoiceOfSupporter(candidateIDs []int) (voices []string) {
	args := []interface{}{}
	for _, candidateID := range candidateIDs {
		args = append(args, candidateID)
	}
	rows, err := db.Query(`
SELECT content as keyword
FROM candidate_keywords
WHERE candidate_id IN (`+strings.Join(strings.Split(strings.Repeat("?", len(candidateIDs)), ""), ",")+`)
group by content
ORDER BY IFNULL(SUM(count), 0) DESC
LIMIT 10;
`,
		args...,
	)
	if err != nil {
		return nil
	}

	defer rows.Close()
	for rows.Next() {
		var keyword string
		err = rows.Scan(&keyword)
		if err != nil {
			panic(err.Error())
		}
		voices = append(voices, keyword)
	}
	return
}
