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
	for _, includeFilter := range folder.IncludeFilter {
		folders, err := searchFolders(ctx, foldersClient, includeFilter)
		if err != nil {
			return fmt.Errorf("failed to get organizations with filter: %w", err)
		}
		c.includeFolders = append(c.includeFolders, folders...)
	}
	for _, excludeFilter := range folder.ExcludeFilter {
		folders, err := searchFolders(ctx, foldersClient, excludeFilter)
		if err != nil {
			return fmt.Errorf("failed to get organizations with filter: %w", err)
		}
		c.excludeFolders = append(c.excludeFolders, folders...)
	}
	// Resolve folder from gcpSpec.Projects.Folders.id_include_list and add to c.include_folders
	for _, folderId := range folder.IncludeListId {
		folder, err := getFolderFromId(ctx, foldersClient, folderId)
		if err != nil {
			return fmt.Errorf("failed to get folder with id %s: %w", folderId, err)
		}
		c.includeFolders = append(c.includeFolders, folder)
	}

	allFolders, err := searchFolders(ctx, foldersClient, "name:*")
	if err != nil {
		return fmt.Errorf("failed to get all folders: %w", err)
	}
	for {

		for _, folder := range allFolders {
			existingNode := findNodeByID(c.graph, folder.Parent)
			if existingNode != nil {
				newNode := node{
					parentID: &folder.Parent,
					included: true,
					folder:   folder,
				}
				existingNode.relations = append(existingNode.relations, &newNode)
			}
		}
		if len(allFolders) == 0 {
			break
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
