package main

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	. "github.com/onsi/gomega"
)

func TestParseJSON(t *testing.T) {
	g := NewGomegaWithT(t)

	wd, _ := os.Getwd()
	content, _ := ioutil.ReadFile(
		path.Join(wd, "example/response.json"),
	)

	parsed, err := parseJSON(content)
	g.Expect(err).To(Not(HaveOccurred()))

	converted := convert(parsed)
	expected := []string{
		"bastion",
		"t85_fls1_axsprt1",
		"100.00GB",
		"5.65GB",
		"94.35GB (94.35%)",
		"2020-06-04T04:09:00Z",
	}

	g.Expect(converted[0].toStringSlice()).To(Equal(expected))
}
