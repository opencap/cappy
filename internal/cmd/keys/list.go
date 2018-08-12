package keys

import (
	"fmt"
	"github.com/opencap/cappy/internal/pkg/context"
)

func List() error {
	keys, err := context.Instance().KeyManager().List()
	if err != nil {
		return err
	}

	fmt.Printf("%d key(s) found\n", len(keys))
	for _, key := range keys {
		fmt.Printf(" - %s\n", key)
	}

	return nil
}
