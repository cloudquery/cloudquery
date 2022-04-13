package plugin

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cloudquery/cloudquery/pkg/plugin/registry"
	"github.com/golang/mock/gomock"
)

func TestManager_DownloadProvider(t *testing.T) {
	ctrl := gomock.NewController(t)
	r := registry.NewMockRegistry(ctrl)
	manager, err := NewManager(r)
	assert.Nil(t, err)

	r.EXPECT().Download(gomock.Any(), registry.Provider{Name: "test", Version: "latest"}, false).Return(
		registry.ProviderBinary{
			Provider: registry.Provider{
				Name:    "test",
				Version: "v0.0.3",
				Source:  "cloudquery",
			},
			FilePath: "some/file/path",
		}, nil).Times(2)
	assert.Nil(t, manager.DownloadProviders(context.TODO(), []registry.Provider{{
		Name:    "test",
		Version: "latest",
	}}, false))

	assert.Nil(t, manager.DownloadProviders(context.TODO(), []registry.Provider{{
		Name:    "test",
		Version: "latest",
	}}, false))

	r.EXPECT().Download(gomock.Any(), registry.Provider{Name: "test", Version: "latest"}, false).Return(registry.ProviderBinary{}, errors.New("failed to download")).Times(1)
	assert.Error(t, manager.DownloadProviders(context.TODO(), []registry.Provider{{
		Name:    "test",
		Version: "latest",
	}}, false))
}

func TestManager_DownloadProviderWithReattach(t *testing.T) {
	ctrl := gomock.NewController(t)
	r := registry.NewMockRegistry(ctrl)
	manager, err := NewManager(r)
	assert.Nil(t, err)
	r.EXPECT().Download(gomock.Any(), registry.Provider{Name: "test", Version: "latest"}, false).Return(
		registry.ProviderBinary{
			Provider: registry.Provider{
				Name:    "test",
				Version: "v0.0.3",
				Source:  "cloudquery",
			},
			FilePath: "some/file/path",
		}, nil).Times(2)
	assert.Nil(t, manager.DownloadProviders(context.TODO(), []registry.Provider{{
		Name:    "test",
		Version: "latest",
	}}, false))

	assert.Nil(t, manager.DownloadProviders(context.TODO(), []registry.Provider{{
		Name:    "test",
		Version: "latest",
	}}, false))

	r.EXPECT().Download(gomock.Any(), registry.Provider{Name: "test", Version: "latest"}, false).Return(registry.ProviderBinary{}, errors.New("failed to download")).Times(1)
	assert.Error(t, manager.DownloadProviders(context.TODO(), []registry.Provider{{
		Name:    "test",
		Version: "latest",
	}}, false))
}
