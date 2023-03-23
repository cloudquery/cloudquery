package client

import (
	"errors"
	"io"

	"github.com/shenzhencenter/google-ads-pb/services"
)

type (
	Receiver[G GoogleAdsRowGetter] func() (G, error)
	GoogleAdsRowGetter             interface {
		GetResults() []*services.GoogleAdsRow
	}
	RowExtractor[A any] func(*services.GoogleAdsRow) A
)

func ReceiveStream[A any, G GoogleAdsRowGetter](receive Receiver[G], extract RowExtractor[A], res chan<- any) error {
	for {
		chunk, err := receive()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}
		for _, row := range chunk.GetResults() {
			res <- extract(row)
		}
	}
}
