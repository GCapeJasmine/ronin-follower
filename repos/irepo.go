package repos

import (
	"github.com/GCapeJasmine/ronin-follower/internal/domains/repos"
)

type IRepo interface {
	Blocks() repos.IBlockRepo
}
