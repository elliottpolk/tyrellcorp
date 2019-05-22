package tyrellcorp

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func GenerateReplicant(spec *Spec) error {
	log.Debugf("generating project for %s replicant", spec.Name)
	if err := CreateProject(spec); err != nil {
		return errors.Wrapf(err, "unable to generate the project directory for %s", spec.Name)
	}

	log.Debugf("generating project assets for %s replicant", spec.Name)
	if err := GenerateAssets(spec); err != nil {
		return errors.Wrapf(err, "unable to generate repository assets for %s", spec.Name)
	}

	return nil
}
