package terminal

import (
	"context"
	"fmt"
	"github.com/cloudquery/cloudquery/client"
	"github.com/cloudquery/cloudquery/config"
	"github.com/cloudquery/cloudquery/hub"
	"github.com/fatih/color"
	"github.com/vbauerster/mpb/v6"
	"github.com/vbauerster/mpb/v6/decor"
	"io"
	"time"
)

func Fetch(ctx context.Context, configPath string) error {
	cfg, err := config.Parse(configPath)
	if err != nil {
		return err
	}

	// initialize progress container, with custom width
	p := ProgressUpdater{mpb.NewWithContext(ctx, mpb.WithWidth(64), mpb.WithRefreshRate(180*time.Millisecond)), make(map[string]string), make(map[string]*mpb.Bar), ""}

	c, err := client.New(func(client *client.Client) {
		client.Hub = hub.NewRegistryHub(hub.CloudQueryRegistryURl, func(h *hub.Hub) {
			h.ProgressUpdater = p
		})
	})

	fmt.Println("")

	if err != nil {
		color.Red("❌ Failed to initialize client.\n\n")
		return err
	}
	color.White("Initializing Cloudquery Providers...\n\n")

	err = c.Initialize(ctx, cfg.CloudQuery.Providers)
	if err != nil {
		time.Sleep(100 * time.Millisecond)
		color.Red("❌ Failed to initialize provider: %s.\n\n", err.Error())
		return err
	}
	// sleep some extra 100 milliseconds for refresh
	time.Sleep(100 * time.Millisecond)
	p.progress.Wait()

	fmt.Println("")
	color.Cyan("Finished provider intialization...\n")

	//if err := c.Fetch(ctx, client.FetchRequest{Providers: cfg.Providers}); err != nil {
	//	return err
	//}

	return err
}

type ProgressUpdater struct {
	progress  *mpb.Progress
	barStatus map[string]string
	bars      map[string]*mpb.Bar
	status    string
}

func (p ProgressUpdater) OnDownload(providerName string, version string, size int64, data io.Reader) io.Reader {
	total := size
	name := fmt.Sprintf("cq-provider-%s@%s", providerName, version)
	// adding a single bar, which will inherit container's width
	bar := p.progress.Add(total+2, // total of file + 2 verify results
		// progress bar filler with customized style
		mpb.NewBarFiller("[=>-|"),
		mpb.BarFillerOnComplete(p.status),
		mpb.PrependDecorators(
			decor.Any(func(statistics decor.Statistics) string {
				status := p.barStatus[providerName]
				if status == string(hub.Verified) {
					return color.GreenString("✓")
				}
				if status == "downloading" || status == "verifying"{
					return color.WhiteString("⌛")
				}
				return color.RedString("❌")
			}, decor.WC{W: 2, C: 1}),
			// display our name with one space on the right
			decor.Name(name, decor.WC{W: len(name) + 1, C: decor.DidentRight}),
			decor.Any(func(statistics decor.Statistics) string {
				if p.barStatus[providerName] == string(hub.Verified) {
					return string(hub.Verified)
				}
				return p.barStatus[providerName] + "..."
			},
			decor.WC{W: len(name) + 1, C: decor.DidentRight}),
			// replace ETA decorator with "done" message, OnComplete event
			decor.OnComplete(
				decor.AverageETA(decor.ET_STYLE_GO, decor.WC{W: 4}), "",
			),
		),
		mpb.AppendDecorators(
			decor.Percentage(),
		),
	)
	p.barStatus[providerName] = "downloading"
	p.bars[providerName] = bar
	return bar.ProxyReader(io.LimitReader(data, size))
}

func (p ProgressUpdater) OnVerify(providerName string, status hub.VerifyStatus) {
	p.barStatus[providerName] = string(status)
	p.bars[providerName].Increment()
}
