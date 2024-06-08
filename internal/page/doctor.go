package page

import (
	"context"

	"github.com/kaziwaseef/stacker/internal/component"
	"github.com/kaziwaseef/stacker/internal/util"
)

func DoctorPage(ctx context.Context) {
	model := component.SpinnerComponent(ctx, checkDependencies)
	if model.Data == nil {
		return
	}

	title := ""

	if model.Data.isGitExits && model.Data.isGithubExits && model.Data.isGitlabExits {
		title = "✅ Everything is Perfect"
	} else if !model.Data.isGitExits {
		title = "❌ Required Dependencies are missing"
	} else if !model.Data.isGithubExits || !model.Data.isGitlabExits {
		title = "🆗 Optional Dependencies are missing"
	}

	statusList := []string{}

	if model.Data.isGitExits {
		statusList = append(statusList, "✅ Git is installed")
	} else {
		statusList = append(statusList, "❌ Git is not installed")
	}

	if model.Data.isGithubExits {
		statusList = append(statusList, "✅ Github CLI is installed")
	} else {
		statusList = append(statusList, "❌ Github CLI (optional) is not installed")
	}

	if model.Data.isGitlabExits {
		statusList = append(statusList, "✅ Gitlab CLI is installed")
	} else {
		statusList = append(statusList, "❌ Gitlab CLI (optional) is not installed")
	}

	component.StaticListComponent(title, statusList)
}

type DoctorPageModel struct {
	isGitExits    bool
	isGithubExits bool
	isGitlabExits bool
}

func checkDependencies(ctx context.Context) *DoctorPageModel {
	// time.Sleep(2 * time.Second)
	isGitExits := util.GitExits(ctx)
	isGithubExits := util.GithubExits(ctx)
	isGitlabExits := util.GitlabExits(ctx)
	return &DoctorPageModel{
		isGitExits:    isGitExits,
		isGithubExits: isGithubExits,
		isGitlabExits: isGitlabExits,
	}
}
