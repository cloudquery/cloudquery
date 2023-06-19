package client

import (
	"context"
	"fmt"

	resourcemanager "cloud.google.com/go/resourcemanager/apiv3"
	"cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	"google.golang.org/api/iterator"
)

func (c *Client) resolveFolders(ctx context.Context, folder ResourceDiscovery) error {
	var err error
	foldersClient, err := resourcemanager.NewFoldersClient(ctx, c.ClientOptions...)
	if err != nil {
		return fmt.Errorf("failed to create folders client: %w", err)
	}

	// Cannot directly add all excluded/included folders to the graph because the graph is not yet fully populated
	// So first must grab all folders and then add them to the graph
	allFolders, err := searchFolders(ctx, foldersClient, "name:*")
	if err != nil {
		return fmt.Errorf("failed to get all folders: %w", err)
	}
	added := 0
	for {
		for _, folder := range allFolders {
			if addFolder(c.graph, folder, nil) {
				added++
			}
		}
		if len(allFolders) == added {
			break
		}
	}

	// Resolve folder from gcpSpec.Projects.Folders.id_include_filter and add to graph
	for _, includeFilter := range folder.IncludeFilter {
		folders, err := searchFolders(ctx, foldersClient, includeFilter)
		if err != nil {
			return fmt.Errorf("failed to get organizations with filter: %w", err)
		}
		for _, folder := range folders {
			if !updateFolder(c.graph, folder, &boolTrue) {
				c.logger.Warn().Msgf("folder %s is included but could not be added to the dependency graph", folder.Name)
			}
		}
	}
	// Resolve folder from gcpSpec.Projects.Folders.id_include_list and add to graph
	for _, folderId := range folder.IncludeListId {
		folder, err := getFolderFromId(ctx, foldersClient, folderId)
		if err != nil {
			return fmt.Errorf("failed to get folder with id %s: %w", folderId, err)
		}
		if !updateFolder(c.graph, folder, &boolTrue) {
			c.logger.Warn().Msgf("folder %s is included but could not be added to the dependency graph", folder.Name)
		}
	}

	// Resolve folder from gcpSpec.Projects.Folders.id_exclude_list and add to graph
	for _, folderId := range folder.ExcludeListId {
		folder, err := getFolderFromId(ctx, foldersClient, folderId)
		if err != nil {
			return fmt.Errorf("failed to get folder with id %s: %w", folderId, err)
		}
		if !updateFolder(c.graph, folder, &boolFalse) {
			c.logger.Warn().Msgf("folder %s is included but could not be added to the dependency graph", folder.Name)
		}
	}

	return err
}

func getFolderFromId(ctx context.Context, foldersClient *resourcemanager.FoldersClient, id string) (*resourcemanagerpb.Folder, error) {
	return foldersClient.GetFolder(ctx, &resourcemanagerpb.GetFolderRequest{Name: "folders/" + id})
}

// searchFolders finds all folders that match the filter.
func searchFolders(ctx context.Context, folderClient *resourcemanager.FoldersClient, filter string) ([]*resourcemanagerpb.Folder, error) {
	folders := make([]*resourcemanagerpb.Folder, 0)

	it := folderClient.SearchFolders(ctx, &resourcemanagerpb.SearchFoldersRequest{
		Query: filter,
	})

	for {
		folder, err := it.Next()

		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		folders = append(folders, folder)
	}

	return folders, nil
}
