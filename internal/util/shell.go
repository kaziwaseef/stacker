package util

import (
	"context"
	"os/exec"
)

func GitExits(ctx context.Context) bool {
	cmd := exec.Command("git", "--version")
	stdout, err := cmd.Output()
	Logger(ctx).Log(string(stdout))
	return err == nil
}

func GithubExits(ctx context.Context) bool {
	cmd := exec.Command("gh", "--version")
	stdout, err := cmd.Output()
	Logger(ctx).Log(string(stdout))
	return err == nil
}

func GitlabExits(ctx context.Context) bool {
	cmd := exec.Command("glab", "--version")
	stdout, err := cmd.Output()
	Logger(ctx).Log(string(stdout))
	return err == nil
}
