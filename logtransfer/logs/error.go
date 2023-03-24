package logs

import (
	"context"
	"os"
)

const errorFilePath = "error.log"

func Error(ctx context.Context, errc chan error) error {
	f, err := os.Create(errorFilePath)
	if err != nil {
		return err
	}
	defer f.Close()

	for {
		select {
		case <-ctx.Done():
			return nil
		case err := <-errc:
			_, _ = f.WriteString(err.Error())
		}
	}
}
