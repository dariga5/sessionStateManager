package main

import (
	"SSM/main/session_s"
)

func main() {

	var s = session_s.CreateSessionObj()

	var err = s.GetCache().GetData(3).Err

	s.GetCache().PullRequest("1")
	s.GetCache().PullRequest("2")
	s.GetCache().PullRequest("3")
	s.GetCache().PullRequest("4")
	s.GetCache().PullRequest("5")
	s.GetCache().PullRequest("6")
	s.GetCache().PullRequest("7")
	s.GetCache().PullRequest("8")
	s.GetCache().PullRequest("9")
	var err1 = s.GetCache().PullRequest("10")

	if err1 != nil || err != nil {

		print(err, err1)
	}
}
