package repos

import (
	"github.com/GCapeJasmine/ronin-follower/config"
	"github.com/GCapeJasmine/ronin-follower/internal/domains/repos"
)

type RepoInMem struct {
	cfg *config.InMemConfig
}

func NewRepoInMem(cfg *config.InMemConfig) IRepo {
	return &RepoInMem{
		cfg: cfg,
	}
}

func (r *RepoInMem) Blocks() repos.IBlockRepo {
	return NewBlockRepo(r.cfg)
}
