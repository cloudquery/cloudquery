package client

import (
	"crypto/sha256"
	"time"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
	analyticsdata "google.golang.org/api/analyticsdata/v1beta"
)

type row struct {
	Date          time.Time
	Dimensions    map[string]string
	DimensionHash []byte
	Metrics       map[string]string
	DataLoss      bool
}

func convertRows(resp *analyticsdata.RunReportResponse, date time.Time) []*row {
	res := make([]*row, len(resp.Rows))

	for idx, r := range resp.Rows {
		mx := make(map[string]string, len(resp.MetricHeaders))
		for i, mHeader := range resp.MetricHeaders {
			mx[mHeader.Name] = r.MetricValues[i].Value
		}

		dim := make(map[string]string, len(resp.DimensionHeaders))
		for i, dHeader := range resp.DimensionHeaders {
			mx[dHeader.Name] = r.DimensionValues[i].Value
		}

		res[idx] = &row{
			Date:          date,
			Dimensions:    dim,
			DimensionHash: calcMapHash(dim),
			Metrics:       mx,
			DataLoss:      resp.Metadata.DataLossFromOtherRow,
		}
	}

	return res
}

func calcMapHash(m map[string]string) []byte {
	h := sha256.New()
	keys := maps.Keys(m)
	slices.Sort(keys)
	for _, k := range keys {
		h.Write([]byte(k))
		h.Write([]byte(m[k]))
	}

	return h.Sum(nil)
}
