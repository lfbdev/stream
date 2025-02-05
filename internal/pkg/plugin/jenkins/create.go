package jenkins

import (
	"fmt"

	"github.com/mitchellh/mapstructure"

	. "github.com/merico-dev/stream/internal/pkg/plugin/common/helm"
	"github.com/merico-dev/stream/pkg/util/log"
)

// Create creates jenkins with provided options.
func Create(options map[string]interface{}) (map[string]interface{}, error) {
	// 1. decode options
	var opts Options
	if err := mapstructure.Decode(options, &opts); err != nil {
		return nil, err
	}

	if errs := validate(&opts); len(errs) != 0 {
		for _, e := range errs {
			log.Errorf("Options error: %s.", e)
		}
		return nil, fmt.Errorf("opts are illegal")
	}

	// 2. deal with ns
	if err := DealWithNsWhenInstall(&opts); err != nil {
		return nil, err
	}

	var retErr error
	defer func() {
		if retErr == nil {
			return
		}
		if err := DealWithNsWhenInterruption(&opts); err != nil {
			log.Errorf("Failed to deal with namespace: %s.", err)
		}
		log.Debugf("Deal with namespace when interruption succeeded.")

		// Clear all the resources have been created if the creation process interruption.
		if err := postDelete(); err != nil {
			log.Errorf("Failed to clear the resources have been created: %s.", err)
		}
	}()

	// 3. pre-create
	if retErr = preCreate(); retErr != nil {
		log.Errorf("The pre-create logic failed: %s.", retErr)
		return nil, retErr
	}

	// 4. install or upgrade
	if retErr = InstallOrUpgradeChart(&opts); retErr != nil {
		return nil, retErr
	}

	// 5. fill the return map
	releaseName := opts.Chart.ReleaseName
	retMap := GetStaticState(releaseName).ToStringInterfaceMap()
	log.Debugf("Return map: %v.", retMap)

	return retMap, nil
}
