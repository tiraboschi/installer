package ovirt

import (
	"fmt"
	"github.com/openshift/installer/pkg/types/ovirt"
	ovirtsdk "github.com/ovirt/go-ovirt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/AlecAivazis/survey.v1"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func Authenticated(p *ovirt.Platform) survey.Validator {
	return func(val interface{}) error {
		if p.Cafile == "" {
			certUrl := url.URL{
				Scheme:   "http",
				Path:     "ovirt-engine/services/pki-resource",
				RawQuery: url.PathEscape("resource=ca-certificate&format=X509-PEM-CA"),
			}
			surveyUrl, err := url.Parse(p.Url)
			if err != nil {
				return errors.Errorf("failed parse ovirt-engine survey Url %s", err)
			}

			if strings.Contains(surveyUrl.Host, ":443") {
				certUrl.Host = strings.Replace(surveyUrl.Host, ":443", ":80", 1)
			}
			if strings.Contains(surveyUrl.Host, ":8443") {
				certUrl.Host = strings.Replace(surveyUrl.Host, ":8443", ":8080", 1)
			}
			if !strings.Contains(surveyUrl.Host, ":") && surveyUrl.Scheme == "https" {
				certUrl.Host = surveyUrl.Host
			}

			logrus.Infof("ovirt cert url %s", certUrl.String())

			resp, err := http.Get(certUrl.String())
			if err != nil || resp.StatusCode != http.StatusOK {
				return fmt.Errorf("error downloading ovirt-engine certificate %s with status %s", err, resp.Status)
			}

			file, err := os.Create("/tmp/ovirt-engine.ca")
			if err != nil {
				return fmt.Errorf("failed writing ovirt-engine certificate %s", err)
			}
			io.Copy(file, resp.Body)
			logrus.Infof("downloaded ovirt-engine certificate to %s", file.Name())
			p.Cafile = file.Name()
		}

		connection, err := ovirtsdk.NewConnectionBuilder().
			URL(p.Url).
			Username(p.Username).
			Password(fmt.Sprint(val)).
			CAFile(p.Cafile).
			Insecure(false).
			Build()

		if err != nil {
			return errors.Errorf("failed to construct connection to oVirt platform %s", err)
		}

		defer connection.Close()

		err = connection.Test()
		if err != nil {
			return errors.Errorf("connection to oVirt platform test failed %s", err)
		}

		if err != nil {
			return err
		}

		return nil
	}

}
