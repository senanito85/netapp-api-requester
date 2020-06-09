package main

import (
	"fmt"

	"github.com/inhies/go-bytesize"
)

type output struct {
	Name      string
	SvmName   string
	Size      string
	Used      string
	Available string
	Timestamp string
}

func convert(resp response) []output {
	tbl := []output{}
	for _, r := range resp.Records {
		totalSize := bytesize.New(r.Space.Size)
		usedSize := bytesize.New(r.Space.Size - r.Space.Available)
		availableSize := bytesize.New(r.Space.Available)
		availablePercent := r.Space.Available * 100 / r.Space.Size

		oRow := output{
			Name:      r.Name,
			SvmName:   r.Svm.Name,
			Size:      totalSize.String(),
			Used:      usedSize.String(),
			Available: fmt.Sprintf("%s (%.2f%%)", availableSize, availablePercent),
			Timestamp: r.Metric.Timestamp,
		}

		tbl = append(tbl, oRow)
	}

	return tbl
}

func (o *output) toStringSlice() []string {
	return []string{o.Name, o.SvmName, o.Size, o.Used, o.Available, o.Timestamp}
}
