package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/google/go-github/github"
	"github.com/sethvargo/go-githubactions"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

func main() {
	//gh_token := os.Getenv("GITHUB_TOKEN")
	releaseNamePrefix := githubactions.GetInput("release_name_Prefix") + "_"
	tagName := githubactions.GetInput("tag_name")
	githubToken := githubactions.GetInput("github_token")
	protectBranch := (githubactions.GetInput("protect_branch") == "")
	fmt.Println(tagName)
	if releaseNamePrefix == "" {
		githubactions.Fatalf("missing input 'release_name'")
	}
	if tagName == "" {
		githubactions.Fatalf("missing input 'tag_name'")
	}
	if githubToken == "" {
		githubactions.Fatalf("missing input 'github_token'")
	}
	githubactions.AddMask(githubToken)
	githubReference := os.Getenv("GITHUB_REF")
	repository := os.Getenv("GITHUB_REPOSITORY")
	regexp := regexp.MustCompile(`refs/tags/v.*`)
	if regexp.Match([]byte(githubReference)) {
		//releaseBranch := strings.Replace(githubReference, "refs/tags/", releaseNamePrefix, 1)
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: githubToken})
		tc := oauth2.NewClient(ctx, ts)
		toolkit := github.NewClient(tc)
		repo := strings.Split(repository, "/")
		branches, _, err := toolkit.Repositories.ListBranches(ctx, repo[0], repo[1], nil)
		if err != nil {
			githubactions.Fatalf("Fetching Branches resultet in %v", err)
		}
		branch, _, err := toolkit.Repositories.GetBranch(ctx, repo[0], repo[1], tagName)
		if err != nil {
			fmt.Println("test")
			githubactions.Fatalf(" Failed to get branch %v", err)
		}
		fmt.Println(branch)
		fmt.Println(branches)
		fmt.Println(toolkit.BaseURL)
		fmt.Println(repository)
		fmt.Println(protectBranch)
		fmt.Println("-------------")
	}

}
