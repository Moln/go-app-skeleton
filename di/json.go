package di

import "github.com/json-iterator/go/extra"

func init() {
	extra.SetNamingStrategy(extra.LowerCaseWithUnderscores)
}
