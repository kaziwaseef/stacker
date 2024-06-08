package util

import (
	"context"
	"os/exec"
)

func GitExits(ctx context.Context) bool {
	Logger(ctx).Log("Running \"git --version\"")
	cmd := exec.Command("git", "--version")
	stdout, err := cmd.Output()
	Logger(ctx).Log(string(stdout))
	return err == nil
}

func GithubExits(ctx context.Context) bool {
	Logger(ctx).Log("Running \"gh --version\"")
	cmd := exec.Command("gh", "--version")
	stdout, err := cmd.Output()
	Logger(ctx).Log(string(stdout))
	return err == nil
}

func GitlabExits(ctx context.Context) bool {
	Logger(ctx).Log("Running \"glab --version\"")
	cmd := exec.Command("glab", "--version")
	stdout, err := cmd.Output()
	Logger(ctx).Log(string(stdout))
	return err == nil
}
