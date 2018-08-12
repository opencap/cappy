package keys

import (
	"fmt"
	"github.com/AlecAivazis/survey"
	"github.com/opencap/cappy/internal/pkg/context"
)

func Remove(idPtr *string) error {
	var id string

	if idPtr == nil {
		ids, err := context.Instance().KeyManager().List()
		if err != nil {
			return fmt.Errorf("failed to list key ids: %v", err)
		}

		err = survey.AskOne(&survey.Input{
			Message: "\xF0\x9F\x94\x91  Enter Key ID: ",
		}, &id, func(i interface{}) (err error) {
			for _, id := range ids {
				if id == i.(string) {
					return nil
				}
			}
			return fmt.Errorf("id not found")
		})
		if err != nil {
			return fmt.Errorf("key selection failed: %v", err)
		}
	} else {
		id = *idPtr
	}

	if err := context.Instance().KeyManager().Delete(id); err != nil {
		return fmt.Errorf("failed to delete key: %v", err)
	}

	return nil
}
