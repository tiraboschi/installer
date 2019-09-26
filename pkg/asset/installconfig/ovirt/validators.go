package ovirt

import (
	"github.com/openshift/installer/pkg/types/ovirt"
	ovirtsdk "github.com/ovirt/go-ovirt"
	"gopkg.in/AlecAivazis/survey.v1"
)

func Authenticated(p ovirt.Platform) survey.Validator {
	connection, err := ovirtsdk.NewConnectionBuilder().
		URL(p.Url).
		Username(p.Username).
		Password(p.Password).
		CAFile(p.Cafile).
		Insecure(p.Insecure).
		Build()

	return func(val interface{}) error {
		defer connection.Close()
		if err != nil {
			return err
		}
		err = connection.Test()
		if err != nil {
			return err
		}

		return nil
	}

}
