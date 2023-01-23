package files

import (
	"github.com/cloudquery/cloudquery/plugins/source/slack/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
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
				Name:     "team_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveTeamID,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "created",
				Type:     schema.TypeTimestamp,
				Resolver: client.JSONTimeResolver("Created"),
			},
			{
				Name:     "timestamp",
				Type:     schema.TypeTimestamp,
				Resolver: client.JSONTimeResolver("Timestamp"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "title",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Title"),
			},
			{
				Name:     "mimetype",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Mimetype"),
			},
			{
				Name:     "image_exif_rotation",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ImageExifRotation"),
			},
			{
				Name:     "filetype",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Filetype"),
			},
			{
				Name:     "pretty_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PrettyType"),
			},
			{
				Name:     "user",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("User"),
			},
			{
				Name:     "mode",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Mode"),
			},
			{
				Name:     "editable",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Editable"),
			},
			{
				Name:     "is_external",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsExternal"),
			},
			{
				Name:     "external_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ExternalType"),
			},
			{
				Name:     "size",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Size"),
			},
			{
				Name:     "url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("URL"),
			},
			{
				Name:     "url_download",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("URLDownload"),
			},
			{
				Name:     "url_private",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("URLPrivate"),
			},
			{
				Name:     "url_private_download",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("URLPrivateDownload"),
			},
			{
				Name:     "original_h",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("OriginalH"),
			},
			{
				Name:     "original_w",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("OriginalW"),
			},
			{
				Name:     "thumb_64",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Thumb64"),
			},
			{
				Name:     "thumb_80",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Thumb80"),
			},
			{
				Name:     "thumb_160",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Thumb160"),
			},
			{
				Name:     "thumb_360",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Thumb360"),
			},
			{
				Name:     "thumb_360_gif",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Thumb360Gif"),
			},
			{
				Name:     "thumb_360_w",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Thumb360W"),
			},
			{
				Name:     "thumb_360_h",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Thumb360H"),
			},
			{
				Name:     "thumb_480",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Thumb480"),
			},
			{
				Name:     "thumb_480_w",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Thumb480W"),
			},
			{
				Name:     "thumb_480_h",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Thumb480H"),
			},
			{
				Name:     "thumb_720",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Thumb720"),
			},
			{
				Name:     "thumb_720_w",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Thumb720W"),
			},
			{
				Name:     "thumb_720_h",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Thumb720H"),
			},
			{
				Name:     "thumb_960",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Thumb960"),
			},
			{
				Name:     "thumb_960_w",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Thumb960W"),
			},
			{
				Name:     "thumb_960_h",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Thumb960H"),
			},
			{
				Name:     "thumb_1024",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Thumb1024"),
			},
			{
				Name:     "thumb_1024_w",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Thumb1024W"),
			},
			{
				Name:     "thumb_1024_h",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Thumb1024H"),
			},
			{
				Name:     "permalink",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Permalink"),
			},
			{
				Name:     "permalink_public",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PermalinkPublic"),
			},
			{
				Name:     "edit_link",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EditLink"),
			},
			{
				Name:     "preview",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Preview"),
			},
			{
				Name:     "preview_highlight",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PreviewHighlight"),
			},
			{
				Name:     "lines",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Lines"),
			},
			{
				Name:     "lines_more",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("LinesMore"),
			},
			{
				Name:     "is_public",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsPublic"),
			},
			{
				Name:     "public_url_shared",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("PublicURLShared"),
			},
			{
				Name:     "channels",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Channels"),
			},
			{
				Name:     "groups",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Groups"),
			},
			{
				Name:     "ims",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("IMs"),
			},
			{
				Name:     "initial_comment",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("InitialComment"),
			},
			{
				Name:     "comments_count",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("CommentsCount"),
			},
			{
				Name:     "num_stars",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("NumStars"),
			},
			{
				Name:     "is_starred",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsStarred"),
			},
			{
				Name:     "shares",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Shares"),
			},
		},
	}
}
