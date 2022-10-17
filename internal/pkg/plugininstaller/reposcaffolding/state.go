package reposcaffolding

import (
	"github.com/devstream-io/devstream/internal/pkg/plugininstaller"
	"github.com/devstream-io/devstream/internal/pkg/statemanager"
	"github.com/devstream-io/devstream/pkg/util/log"
	"github.com/devstream-io/devstream/pkg/util/scm/github"
)

func GetDynamicState(options plugininstaller.RawOptions) (statemanager.ResourceStatus, error) {
	opts, err := NewOptions(options)
	if err != nil {
		return nil, err
	}

	scmClient, err := github.NewClient(opts.DestinationRepo.BuildRepoInfo())
	if err != nil {
		log.Debugf("reposcaffolding status init repo failed: %+v", err)
		return nil, err
	}

	repoInfo, err := scmClient.DescribeRepo()
	if err != nil {
		log.Debugf("reposcaffolding status describe repo failed: %+v", err)
		return nil, err
	}

	resState := statemanager.ResourceStatus{
		"repo":     repoInfo.Repo,
		"owner":    repoInfo.Owner,
		"org":      repoInfo.Org,
		"repoURL":  repoInfo.CloneURL,
		"repoType": repoInfo.Type,
		"source":   opts.SourceRepo.BuildScmURL(),
	}
	return resState, nil
}
