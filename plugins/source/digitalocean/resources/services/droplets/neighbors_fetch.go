package droplets

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/digitalocean/godo"
)

type NeighborWrapper struct {
	DropletId  int
	NeighborId int
}

func fetchDropletsNeighbors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client)
	droplet := parent.Item.(godo.Droplet)

	neighbors, _, err := svc.Services.Droplets.Neighbors(ctx, droplet.ID)
	if err != nil {
		return err
	}
	if neighbors == nil {
		return nil
	}
	nn := make([]NeighborWrapper, len(neighbors))
	for i, n := range neighbors {
		nn[i] = NeighborWrapper{
			DropletId:  droplet.ID,
			NeighborId: n.ID,
		}
	}
	res <- nn
	return nil
}
