package domains

import (
	"fmt"
	"github.com/AlecAivazis/survey"
	"github.com/opencap/cappy/pkg/api"
	"github.com/opencap/opencap/pkg/resolver"
)

func Notify(domainPtr *string) (err error) {
	var domain string
	if domainPtr == nil {
		err = survey.AskOne(&survey.Input{
			Message: "\xF0\x9F\x8C\x90  Domain name:",
		}, &domain, survey.Required)
		if err != nil {
			return
		}
	} else {
		domain = *domainPtr
	}

	res, err := resolver.Resolve(domain)
	if err != nil {
		return fmt.Errorf("dns request failed: %v", err)
	}

	_ = res

	var client *api.Client
	client, err = api.NewClient( /*res.GetServer()*/ "localhost", 41145)
	if err != nil {
		return
	}

	err = client.AssociateDomain(domain)
	if err != nil {
		return
	}

	return
}
