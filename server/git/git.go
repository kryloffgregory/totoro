package git

import (
	"context"
	"fmt"
	"log"
	"time"


	"github.com/google/go-github/github"
	"github.com/kryloffgregory/totoro/server/execute"
	"golang.org/x/oauth2"
)

const owner = "kryloffgregory"
const token = "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"
const token2 = "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"
const token3 = "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"
const repo = "apt-server"

var client *github.Client

func init() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token3},
	)
	tc := oauth2.NewClient(ctx, ts)

	client = github.NewClient(tc)
}

func CreatePR(user string, command string, reviewers []string) string {
	ctx := context.Background()
	branch:="branch"+fmt.Sprint(time.Now().UnixNano())

	ref, err := getRef(ctx, client, branch)
	if err != nil {
		log.Fatalf("Unable to get/create the commit reference: %s\n", err)
	}
	if ref == nil {
		log.Fatalf("No error where returned but the reference is nil")
	}

	tree, err := getTree(ctx, client, ref, command)
	if err != nil {
		log.Fatalf("Unable to create the tree based on the provided files: %s\n", err)
	}

	if err := pushCommit(ctx, client, ref, tree, user); err != nil {
		log.Fatalf("Unable to create the commit: %s\n", err)
	}

	url, err := createPR1(ctx, client, user, command, reviewers, branch)
	if err!=nil {
		log.Fatalf("Error while creating the pull request: %s", err)
	}

	return url

}

func getRef(ctx context.Context, client *github.Client, branch string) (ref *github.Reference, err error) {
	if ref, _, err = client.Git.GetRef(ctx, owner, repo, "refs/heads/"+branch); err == nil {
		return ref, nil
	}

	var baseRef *github.Reference
	if baseRef, _, err = client.Git.GetRef(ctx, owner, repo, "refs/heads/main"); err != nil {
		return nil, err
	}
	newRef := &github.Reference{Ref: github.String("refs/heads/"+branch), Object: &github.GitObject{SHA: baseRef.Object.SHA}}
	ref, _, err = client.Git.CreateRef(ctx, owner, repo, newRef)
	return ref, err
}

func getTree(ctx context.Context, client *github.Client, ref *github.Reference, command string) (tree *github.Tree, err error) {
	//github.RepositoriesService.GetContents()
	// Create a tree with what to commit.
	entries := []*github.TreeEntry{
		{
			Path: github.String("oplog"),
			Type: github.String("blob"),
			Content: github.String(command),
			Mode: github.String("100644"),
		},
	}

	tree, _, err = client.Git.CreateTree(ctx, owner, repo, *ref.Object.SHA, entries)
	return tree, err
}

func pushCommit(ctx context.Context, client *github.Client, ref *github.Reference, tree *github.Tree, user string) (err error) {
	// Get the parent commit to attach the commit to.
	parent, _, err := client.Repositories.GetCommit(ctx, owner, repo, *ref.Object.SHA)
	if err != nil {
		return err
	}
	// This is not always populated, but is needed.
	parent.Commit.SHA = parent.SHA

	// Create the commit using the tree.
	date := time.Now()
	author := &github.CommitAuthor{Date: &date, Name: github.String(user), Email:github.String("kryloffgv@yahoo.com")}
	commit := &github.Commit{Author: author, Message:github.String("New command to execute"), Tree: tree, Parents: []*github.Commit{parent.Commit}}
	newCommit, _, err := client.Git.CreateCommit(ctx, owner, repo, commit)
	if err != nil {
		return err
	}

	// Attach the commit to the master branch.
	ref.Object.SHA = newCommit.SHA
	_, _, err = client.Git.UpdateRef(ctx, owner, repo, ref, false)
	return err
}

func createPR1(ctx context.Context, client *github.Client, user string, command string, reviewers []string, branch string) (url string, err error) {
	newPR := &github.NewPullRequest{
		Title:               github.String("Request from " + user),
		Head:                github.String(owner+":"+branch),
		Base:                github.String("main"),
		Body: github.String(command),
	}

	pr, _, err := client.PullRequests.Create(ctx, owner, repo, newPR)
	if err != nil {
		return "", err
	}

	fmt.Printf("PR created: %s\n", pr.GetHTMLURL())

	revReq:=github.ReviewersRequest{
		NodeID:        pr.NodeID,
		Reviewers:     reviewers,
	}
	pr, _, err = client.PullRequests.RequestReviewers(ctx, owner, repo, *pr.Number,revReq)
	if err!=nil {
		return "", err
	}
	return pr.GetHTMLURL(), err
}

func ProcessPRs(shutdown chan bool, done chan bool) error {
	ctx := context.Background()
	
	opts:=&github.PullRequestListOptions{
		State:       "open",
		Base:        "main",
	}
	prs, _, err:=client.PullRequests.List(ctx, owner, repo, opts)
	if err!=nil {
		return err
	}
	for _, pr:=range prs {
		select {
		case <-shutdown:
			done <-true
		default:
			err:=ProcessPR(ctx, pr)
			if err!=nil {
				log.Println(fmt.Sprintf("Error occured while processing pr %v: %v", *pr.Body, err))
			}
		}
	}

	return nil
}

func ProcessPR(ctx context.Context, pr *github.PullRequest) error{
	log.Println(fmt.Sprintf("Processing pr %v", pr.GetHTMLURL()))
	reviews, _,  err:=client.PullRequests.ListReviews(ctx, owner, repo, *pr.Number, nil)
	if err!=nil {
		return err
	}

	reviewsLeft:=len(pr.RequestedReviewers)
	fmt.Println(reviewsLeft)
	for _, review:=range reviews {
		if *review.State == "APPROVED" {
			reviewsLeft--
		}
	}

	if reviewsLeft > 0 {
		log.Println(fmt.Sprintf("Skipping pr %v", pr.GetHTMLURL()))
		return nil
	}

	result, err:=execute.ExecuteString(*pr.Body)
	if err!=nil {
		return err
	}
	fmt.Println(result)


	comment:=&github.IssueComment{
		Body:              github.String(result),
	}

	_, _, err=client.Issues.CreateComment(ctx, owner, repo, *pr.Number, comment)
	if err!=nil {
		return err
	}

	_, _, err = client.PullRequests.Merge(ctx, owner, repo, *pr.Number, "merging", nil)
	if err!=nil {
		return err
	}

	log.Println(fmt.Sprintf("Merged pr %v", pr.GetHTMLURL()))

	return nil
}