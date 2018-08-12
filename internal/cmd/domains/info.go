package domains

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/opencap/cappy/internal/pkg/context"
	"github.com/opencap/cappy/internal/pkg/key"
	"github.com/opencap/opencap/pkg/resolver"
)

func Info(domainPtr *string) error {
	var domain string
	if domainPtr == nil {

	} else {
		domain = *domainPtr
	}

	res, err := resolver.Resolve(domain)
	if err != nil {
		return fmt.Errorf("resolving domain name failed: %v", err)
	}

	color.New(color.Underline).Println("Information")

	color.New(color.Bold).Print("Domain: ")
	fmt.Println(domain)

	if res.PublicKey != nil {
		id := key.Identifier(res.PublicKey)
		color.New(color.Bold).Print("Key ID: ")
		fmt.Print(id)

		_, err := context.Instance().KeyManager().Read(id)
		if err != nil {
			fmt.Printf(" (%s)\n", color.RedString("%s", "Not loaded"))
		} else {
			fmt.Printf(" (%s)\n", color.GreenString("%s", "Loaded"))
		}

		color.New(color.Bold).Print("DNSSig: ")
		if res.DNSSig {
			fmt.Println("Enabled")
		} else {
			fmt.Println("Disabled")
		}
	}

	color.New(color.Bold).Print("Server: ")

	for i, entry := range res.Servers {
		if i > 0 {
			fmt.Println("        ")
		}

		fmt.Printf("%s:%d\n", entry.Host, entry.Port)
	}

	return nil
}
