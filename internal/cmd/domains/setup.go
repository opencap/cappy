package domains

import (
	"fmt"
	"github.com/AlecAivazis/survey"
	"github.com/fatih/color"
	"github.com/opencap/cappy/internal/cmd/keys"
	"github.com/opencap/cappy/internal/pkg/context"
	"github.com/opencap/cappy/internal/pkg/key"
	"strconv"
)

func Setup(domainPtr, hostPtr *string, portPtr *uint16, newKeyPtr, dnssigPtr *bool) (err error) {
	var (
		domain string
		host   string
		port   uint16
		newKey bool
		dnssig bool
	)

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

	if hostPtr == nil {
		err = survey.AskOne(&survey.Input{
			Message: "\xF0\x9F\x8F\xA0  Server hostname or ip:",
			Default: domain,
		}, &host, survey.Required)
		if err != nil {
			return
		}
	} else {
		host = *hostPtr
	}

	if portPtr == nil {
		var p string

		err = survey.AskOne(&survey.Input{
			Message: "\xF0\x9F\x94\x8C  Server port:",
			Default: "41145",
		}, &p, func(i interface{}) error {
			_, err := strconv.ParseUint(i.(string), 10, 16)
			if err != nil {
				return fmt.Errorf("invalid port")
			}
			return nil
		})
		if err != nil {
			return
		}

		var p64 uint64
		p64, err = strconv.ParseUint(p, 10, 16)
		if err != nil {
			return
		}

		port = uint16(p64)
	} else {
		port = *portPtr
	}

	if newKeyPtr == nil {
		err = survey.AskOne(&survey.Confirm{
			Message: "\xF0\x9F\x8E\xB2  Generate a new key?",
			Default: true,
		}, &newKey, nil)
		if err != nil {
			return
		}
	} else {
		newKey = *newKeyPtr
	}

	var k key.Key
	if newKey {
		fmt.Println()
		k, err = keys.Add()
		if err != nil {
			return
		}
		fmt.Println()
	} else {
		var unused string
		err = survey.AskOne(&survey.Input{
			Message: "\xF0\x9F\x94\x91  Enter Key ID: ",
		}, &unused, func(i interface{}) (err error) {
			k, err = context.Instance().KeyManager().Read(i.(string))
			return
		})
		if err != nil {
			return
		}
	}

	if dnssigPtr == nil {
		err = survey.AskOne(&survey.Confirm{
			Message: "Would you like to use DNS signatures?",
			Default: true,
		}, &dnssig, nil)
		if err != nil {
			return
		}
	} else {
		dnssig = *dnssigPtr
	}

	fmt.Println()

	color.New(color.Bold).Printf("Please open the DNS settings for %s and create the following entries:\n", domain)
	fmt.Printf(" (1) _opencap._tcp.%s.  IN  SRV  %d  %s.\n", domain, port, host)
	if dnssig {
		fmt.Printf(" (2)               %s.  IN  TXT  \"opencap_key=%s opencap_dnssig=1\"\n", domain, k.Public())
	} else {
		fmt.Printf(" (2)               %s.  IN  TXT  \"opencap_key=%s\"\n", domain, k.Public())
	}

	fmt.Println()

	fmt.Printf("Please run `cappy domains notify %s` when you're done.\n", domain)

	return
}
