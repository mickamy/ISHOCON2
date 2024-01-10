package main

import (
	"database/sql"
	"github.com/gin-contrib/multitemplate"
	"html/template"
	// "log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func createRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	pagesPath := "templates/includes"

	layout := "templates/layouts/layout.tmpl"
	files, err := os.ReadDir(pagesPath)
	if err != nil {
		panic(err)
	}
	funcs := template.FuncMap{"indexPlus1": func(i int) int { return i + 1 }}

	for _, file := range files {
		fileNameWithoutExt := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
		//fileName := file.Name()
		r.AddFromFilesFuncs(fileNameWithoutExt, funcs, layout, pagesPath+"/"+file.Name())
	}
	return r
}

func main() {
	// database setting
	user := getEnv("ISHOCON2_DB_USER", "ishocon")
	pass := getEnv("ISHOCON2_DB_PASSWORD", "ishocon")
	dbname := getEnv("ISHOCON2_DB_NAME", "ishocon2")
	db, _ = sql.Open("mysql", user+":"+pass+"@/"+dbname)
	db.SetMaxIdleConns(12)

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(static.Serve("/css", static.LocalFile("public/css", true)))

	r.HTMLRender = createRender()

	// session store
	store := sessions.NewCookieStore([]byte("mysession"))
	store.Options(sessions.Options{HttpOnly: true})
	r.Use(sessions.Sessions("showwin_happy", store))

	candidates := getAllCandidate()

	// GET /
	r.GET("/", func(c *gin.Context) {
		electionResults := getElectionResult()

		// 上位10人と最下位のみ表示
		tmp := make([]CandidateElectionResult, len(electionResults))
		copy(tmp, electionResults)
		candidates := tmp[:10]
		candidates = append(candidates, tmp[len(tmp)-1])

		partyNameMap := make(map[string]bool)
		partyNames := []string{}
		for _, can := range candidates {
			partyName := can.PoliticalParty
			if _, value := partyNameMap[partyName]; !value {
				partyNameMap[partyName] = true
				partyNames = append(partyNames, partyName)
			}
		}

		// Sort partyNames
		sort.Strings(partyNames)

		//partyNames := getAllPartyName()
		partyResultMap := map[string]int{}
		for _, name := range partyNames {
			partyResultMap[name] = 0
		}
		for _, r := range electionResults {
			partyResultMap[r.PoliticalParty] += r.VoteCount
		}
		partyResults := []PartyElectionResult{}
		for name, count := range partyResultMap {
			r := PartyElectionResult{}
			r.PoliticalParty = name
			r.VoteCount = count
			partyResults = append(partyResults, r)
		}
		// 投票数でソート
		sort.Slice(partyResults, func(i, j int) bool { return partyResults[i].VoteCount > partyResults[j].VoteCount })

		sexRatio := map[string]int{
			"men":   0,
			"women": 0,
		}
		for _, r := range electionResults {
			if r.Sex == "男" {
				sexRatio["men"] += r.VoteCount
			} else if r.Sex == "女" {
				sexRatio["women"] += r.VoteCount
			}
		}

		c.HTML(http.StatusOK, "index", gin.H{
			"candidates": candidates,
			"parties":    partyResults,
			"sexRatio":   sexRatio,
		})
	})

	// GET /candidates/:candidateID(int)
	r.GET("/candidates/:candidateID", func(c *gin.Context) {
		candidateID, _ := strconv.Atoi(c.Param("candidateID"))
		//var candidate *Candidate
		//for _, can := range candidates {
		//	if can.ID == candidateID {
		//		candidate = &can
		//	}
		//}
		//if candidate == nil {
		//	c.Redirect(http.StatusFound, "/")
		//}
		candidate, err := getCandidate(candidateID)
		if err != nil {
			c.Redirect(http.StatusFound, "/")
		}
		votes := getVoteCountByCandidateID(candidateID)
		candidateIDs := []int{candidateID}
		keywords := getVoiceOfSupporter(candidateIDs)

		c.HTML(http.StatusOK, "candidate", gin.H{
			"candidate": candidate,
			"votes":     votes,
			"keywords":  keywords,
		})
	})

	// GET /political_parties/:name(string)
	r.GET("/political_parties/:name", func(c *gin.Context) {
		partyName := c.Param("name")
		var votes int
		electionResults := getElectionResult()
		for _, r := range electionResults {
			if r.PoliticalParty == partyName {
				votes += r.VoteCount
			}
		}

		var candidatesByParty []Candidate
		for _, can := range candidates {
			if can.PoliticalParty == partyName {
				candidatesByParty = append(candidatesByParty, can)
			}
		}
		//candidatesByParty := getCandidatesByPoliticalParty(partyName)
		candidateIDs := []int{}
		for _, can := range candidatesByParty {
			candidateIDs = append(candidateIDs, can.ID)
		}
		keywords := getVoiceOfSupporter(candidateIDs)

		c.HTML(http.StatusOK, "political_party", gin.H{
			"politicalParty": partyName,
			"votes":          votes,
			"candidates":     candidatesByParty,
			"keywords":       keywords,
		})
	})

	// GET /vote
	r.GET("/vote", func(c *gin.Context) {
		c.HTML(http.StatusOK, "vote", gin.H{
			"candidates": candidates,
			"message":    "",
		})
	})

	// POST /vote
	r.POST("/vote", func(c *gin.Context) {
		user, userErr := getUser(c.PostForm("name"), c.PostForm("address"), c.PostForm("mynumber"))
		//var candidate *Candidate
		//for _, can := range candidates {
		//	if can.Name == c.PostForm("candidate") {
		//		candidate = &can
		//	}
		//}
		candidate, cndErr := getCandidateByName(c.PostForm("candidate"))

		votedCount := getUserVotedCount(user.ID)
		voteCount, _ := strconv.Atoi(c.PostForm("vote_count"))

		var message string
		if userErr != nil {
			message = "個人情報に誤りがあります"
		} else if user.Votes < voteCount+votedCount {
			message = "投票数が上限を超えています"
		} else if c.PostForm("candidate") == "" {
			message = "候補者を記入してください"
			//} else if candidate == nil {
		} else if cndErr != nil {
			message = "候補者を正しく記入してください"
		} else if c.PostForm("keyword") == "" {
			message = "投票理由を記入してください"
		} else {
			//for i := 1; i <= voteCount; i++ {
			//	createVote(user.ID, candidate.ID, c.PostForm("keyword"))
			//}
			err := createVote(c, user.ID, candidate.ID, c.PostForm("keyword"), voteCount)
			if err != nil {
				panic(err)
			}
			message = "投票に成功しました"
		}
		c.HTML(http.StatusOK, "vote", gin.H{
			"candidates": candidates,
			"message":    message,
		})
	})

	r.GET("/initialize", func(c *gin.Context) {
		db.Exec("DELETE FROM votes")

		c.String(http.StatusOK, "Finish")
	})

	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	r.Run(":8080")
}
