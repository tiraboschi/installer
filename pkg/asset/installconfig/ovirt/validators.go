package ovirt

import (
	"github.com/openshift/installer/pkg/types/ovirt"
	"gopkg.in/AlecAivazis/survey.v1"
)
func Authenticated(p ovirt.Platform) survey.Validator {
	c, err := NewConnectionBuilder().
		URL(p.Url).
		Username(p.Username).
		Password(p.Password).
		CAFile(p.Cafile).
		Insecure(p.Insecure).
		Build()
	return func(val interface{}) error {
		defer c.Close()
		if err != nil {
			return err
		}
		err = c.Test()
		if err != nil {
			return err
		}

		return nil
	}

}
