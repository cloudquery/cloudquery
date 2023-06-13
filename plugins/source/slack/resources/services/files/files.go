package files

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/slack/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/cloudquery/plugin-sdk/v3/types"
	"github.com/slack-go/slack"
)

func Files() *schema.Table {
	return &schema.Table{
		Name:        "slack_files",
		Description: `https://api.slack.com/methods/files.list`,
		Resolver:    fetchFiles,
		Multiplex:   client.TeamMultiplex,
		Transform:   transformers.TransformWithStruct(&slack.File{}),
		Columns: []schema.Column{
			{
				Name:       "team_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveTeamID,
				PrimaryKey: true,
			},
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
			{
				Name:     "created",
				Type:     arrow.FixedWidthTypes.Timestamp_us,
				Resolver: client.JSONTimeResolver("Created"),
			},
			{
				Name:     "timestamp",
				Type:     arrow.FixedWidthTypes.Timestamp_us,
				Resolver: client.JSONTimeResolver("Timestamp"),
			},
			{
				Name:     "name",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "title",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("Title"),
			},
			{
				Name:     "mimetype",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("Mimetype"),
			},
			{
				Name:     "image_exif_rotation",
				Type:     arrow.PrimitiveTypes.Int64,
				Resolver: schema.PathResolver("ImageExifRotation"),
			},
			{
				Name:     "filetype",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("Filetype"),
			},
			{
				Name:     "pretty_type",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("PrettyType"),
			},
			{
				Name:     "user",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("User"),
			},
			{
				Name:     "mode",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("Mode"),
			},
			{
				Name:     "editable",
				Type:     arrow.FixedWidthTypes.Boolean,
				Resolver: schema.PathResolver("Editable"),
			},
			{
				Name:     "is_external",
				Type:     arrow.FixedWidthTypes.Boolean,
				Resolver: schema.PathResolver("IsExternal"),
			},
			{
				Name:     "external_type",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("ExternalType"),
			},
			{
				Name:     "size",
				Type:     arrow.PrimitiveTypes.Int64,
				Resolver: schema.PathResolver("Size"),
			},
			{
				Name:     "url",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("URL"),
			},
			{
				Name:     "url_download",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("URLDownload"),
			},
			{
				Name:     "url_private",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("URLPrivate"),
			},
			{
				Name:     "url_private_download",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("URLPrivateDownload"),
			},
			{
				Name:     "original_h",
				Type:     arrow.PrimitiveTypes.Int64,
				Resolver: schema.PathResolver("OriginalH"),
			},
			{
				Name:     "original_w",
				Type:     arrow.PrimitiveTypes.Int64,
				Resolver: schema.PathResolver("OriginalW"),
			},
			{
				Name:     "thumb_64",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("Thumb64"),
			},
			{
				Name:     "thumb_80",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("Thumb80"),
			},
			{
				Name:     "thumb_160",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("Thumb160"),
			},
			{
				Name:     "thumb_360",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("Thumb360"),
			},
			{
				Name:     "thumb_360_gif",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("Thumb360Gif"),
			},
			{
				Name:     "thumb_360_w",
				Type:     arrow.PrimitiveTypes.Int64,
				Resolver: schema.PathResolver("Thumb360W"),
			},
			{
				Name:     "thumb_360_h",
				Type:     arrow.PrimitiveTypes.Int64,
				Resolver: schema.PathResolver("Thumb360H"),
			},
			{
				Name:     "thumb_480",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("Thumb480"),
			},
			{
				Name:     "thumb_480_w",
				Type:     arrow.PrimitiveTypes.Int64,
				Resolver: schema.PathResolver("Thumb480W"),
			},
			{
				Name:     "thumb_480_h",
				Type:     arrow.PrimitiveTypes.Int64,
				Resolver: schema.PathResolver("Thumb480H"),
			},
			{
				Name:     "thumb_720",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("Thumb720"),
			},
			{
				Name:     "thumb_720_w",
				Type:     arrow.PrimitiveTypes.Int64,
				Resolver: schema.PathResolver("Thumb720W"),
			},
			{
				Name:     "thumb_720_h",
				Type:     arrow.PrimitiveTypes.Int64,
				Resolver: schema.PathResolver("Thumb720H"),
			},
			{
				Name:     "thumb_960",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("Thumb960"),
			},
			{
				Name:     "thumb_960_w",
				Type:     arrow.PrimitiveTypes.Int64,
				Resolver: schema.PathResolver("Thumb960W"),
			},
			{
				Name:     "thumb_960_h",
				Type:     arrow.PrimitiveTypes.Int64,
				Resolver: schema.PathResolver("Thumb960H"),
			},
			{
				Name:     "thumb_1024",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("Thumb1024"),
			},
			{
				Name:     "thumb_1024_w",
				Type:     arrow.PrimitiveTypes.Int64,
				Resolver: schema.PathResolver("Thumb1024W"),
			},
			{
				Name:     "thumb_1024_h",
				Type:     arrow.PrimitiveTypes.Int64,
				Resolver: schema.PathResolver("Thumb1024H"),
			},
			{
				Name:     "permalink",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("Permalink"),
			},
			{
				Name:     "permalink_public",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("PermalinkPublic"),
			},
			{
				Name:     "edit_link",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("EditLink"),
			},
			{
				Name:     "preview",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("Preview"),
			},
			{
				Name:     "preview_highlight",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("PreviewHighlight"),
			},
			{
				Name:     "lines",
				Type:     arrow.PrimitiveTypes.Int64,
				Resolver: schema.PathResolver("Lines"),
			},
			{
				Name:     "lines_more",
				Type:     arrow.PrimitiveTypes.Int64,
				Resolver: schema.PathResolver("LinesMore"),
			},
			{
				Name:     "is_public",
				Type:     arrow.FixedWidthTypes.Boolean,
				Resolver: schema.PathResolver("IsPublic"),
			},
			{
				Name:     "public_url_shared",
				Type:     arrow.FixedWidthTypes.Boolean,
				Resolver: schema.PathResolver("PublicURLShared"),
			},
			{
				Name:     "channels",
				Type:     arrow.ListOf(arrow.BinaryTypes.String),
				Resolver: schema.PathResolver("Channels"),
			},
			{
				Name:     "groups",
				Type:     arrow.ListOf(arrow.BinaryTypes.String),
				Resolver: schema.PathResolver("Groups"),
			},
			{
				Name:     "ims",
				Type:     arrow.ListOf(arrow.BinaryTypes.String),
				Resolver: schema.PathResolver("IMs"),
			},
			{
				Name:     "initial_comment",
				Type:     types.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("InitialComment"),
			},
			{
				Name:     "comments_count",
				Type:     arrow.PrimitiveTypes.Int64,
				Resolver: schema.PathResolver("CommentsCount"),
			},
			{
				Name:     "num_stars",
				Type:     arrow.PrimitiveTypes.Int64,
				Resolver: schema.PathResolver("NumStars"),
			},
			{
				Name:     "is_starred",
				Type:     arrow.FixedWidthTypes.Boolean,
				Resolver: schema.PathResolver("IsStarred"),
			},
			{
				Name:     "shares",
				Type:     types.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("Shares"),
			},
		},
	}
}
