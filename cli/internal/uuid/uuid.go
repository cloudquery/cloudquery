package uuid

import (
	guuid "github.com/google/uuid"
)

// https://github.com/spf13/pflag/issues/236#issuecomment-931600452

type UUID struct {
	guuid.UUID
}

func (u *UUID) Set(str string) error {
	var err error
	u.UUID, err = guuid.Parse(str)
	return err
}

func (*UUID) Type() string {
	return "uuid"
}
