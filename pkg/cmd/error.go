package cmd

import "fmt"

type ArgumentAndFlagCantBeUsedTogetherError struct {
	Arg  string
	Flag string
}

func (e *ArgumentAndFlagCantBeUsedTogetherError) Error() string {
	return fmt.Sprintf("%q argument and %q flag can not be used together", e.Arg, e.Flag)
}

type MustIncludeEitherOrFlagError struct {
	Flag1 string
	Flag2 string
}

func (e *MustIncludeEitherOrFlagError) Error() string {
	return fmt.Sprintf("must include either %q flag or %q flag", e.Flag1, e.Flag2)
}
