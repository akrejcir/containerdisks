package common

import (
	"strings"

	"github.com/sirupsen/logrus"
	"kubevirt.io/api/instancetype"

	"kubevirt.io/containerdisks/artifacts/centos"
	"kubevirt.io/containerdisks/artifacts/centosstream"
	"kubevirt.io/containerdisks/artifacts/fedora"
	"kubevirt.io/containerdisks/artifacts/generic"
	"kubevirt.io/containerdisks/artifacts/rhcos"
	"kubevirt.io/containerdisks/artifacts/rhcosprerelease"
	"kubevirt.io/containerdisks/artifacts/ubuntu"
	"kubevirt.io/containerdisks/pkg/api"
	"kubevirt.io/containerdisks/pkg/docs"
)

type Entry struct {
	Artifact           api.Artifact
	UseForDocs         bool
	UseForLatest       bool
	SkipWhenNotFocused bool
}

var staticRegistry = []Entry{
	{
		Artifact: rhcos.New(
			"4.9",
			true,
			defaultLabels("u1.small", "rhel.8"),
		),
		UseForDocs: false,
	},
	{
		Artifact: rhcos.New(
			"4.10",
			true,
			defaultLabels("u1.small", "rhel.8"),
		),
		UseForDocs: false,
	},
	{
		Artifact: rhcos.New(
			"4.11",
			true,
			defaultLabels("u1.small", "rhel.8"),
		),
		UseForDocs: false,
	},
	{
		Artifact: rhcos.New(
			"4.12",
			true,
			defaultLabels("u1.small", "rhel.8"),
		),
		UseForDocs: true,
	},
	{
		Artifact: rhcos.New(
			"latest",
			false,
			defaultLabels("u1.small", "rhel.9"),
		),
		UseForDocs: false,
	},
	{
		Artifact: rhcosprerelease.New(
			"latest-4.9",
			defaultLabels("u1.small", "rhel.8"),
		),
		UseForDocs: false,
	},
	{
		Artifact: rhcosprerelease.New(
			"latest-4.10",
			defaultLabels("u1.small", "rhel.8"),
		),
		UseForDocs: false,
	},
	{
		Artifact: rhcosprerelease.New(
			"latest-4.11",
			defaultLabels("u1.small", "rhel.8"),
		),
		UseForDocs: false,
	},
	{
		Artifact: rhcosprerelease.New(
			"latest-4.12",
			defaultLabels("u1.small", "rhel.8"),
		),
		UseForDocs: false,
	},
	{
		Artifact: rhcosprerelease.New(
			"latest-4.13",
			defaultLabels("u1.small", "rhel.9"),
		),
		UseForDocs: false,
	},
	{
		Artifact: rhcosprerelease.New(
			"latest",
			defaultLabels("u1.small", "rhel.9"),
		),
		UseForDocs: false,
	},
	{
		Artifact:   centos.New("8.4", nil),
		UseForDocs: false,
	},
	{
		Artifact: centos.New(
			"7-2009",
			defaultLabels("u1.small", "centos.7"),
		),
		UseForDocs: true,
	},
	{
		Artifact: centosstream.New(
			"9",
			&docs.UserData{
				Username: "cloud-user",
			},
			defaultLabels("u1.small", "centos.stream9"),
		),
		UseForDocs: true,
	},
	{
		Artifact: centosstream.New(
			"8",
			&docs.UserData{
				Username: "centos",
			},
			defaultLabels("u1.small", "centos.stream8"),
		),
		UseForDocs: false,
	},
	{
		Artifact: ubuntu.New(
			"22.04",
			defaultLabels("u1.small", "ubuntu"),
		),
		UseForDocs: true,
	},
	{
		Artifact: ubuntu.New(
			"20.04",
			defaultLabels("u1.small", "ubuntu"),
		),
		UseForDocs: false,
	},
	{
		Artifact: ubuntu.New(
			"18.04",
			defaultLabels("u1.small", "ubuntu"),
		),
		UseForDocs: false,
	},
	// for testing only
	{
		Artifact: generic.New(
			&api.ArtifactDetails{
				SHA256Sum:   "cc704ab14342c1c8a8d91b66a7fc611d921c8b8f1aaf4695f9d6463d913fa8d1",
				DownloadURL: "https://download.cirros-cloud.net/0.6.1/cirros-0.6.1-x86_64-disk.img",
			},
			&api.Metadata{
				Name:    "cirros",
				Version: "6.1",
			},
		),
		SkipWhenNotFocused: true,
		UseForDocs:         false,
	},
}

func gatherArtifacts(registry *[]Entry, gatherers []api.ArtifactsGatherer) {
	for _, gatherer := range gatherers {
		artifacts, err := gatherer.Gather()
		if err != nil {
			logrus.Warn("Failed to gather artifacts", err)
		} else {
			for i := range artifacts {
				*registry = append(*registry, Entry{
					Artifact:     artifacts[i],
					UseForDocs:   i == 0,
					UseForLatest: i == 0,
				})
			}
		}
	}
}

func defaultLabels(defaultInstancetype, defaultPreference string) map[string]string {
	return map[string]string{
		instancetype.DefaultInstancetypeLabel: defaultInstancetype,
		instancetype.DefaultPreferenceLabel:   defaultPreference,
	}
}

func NewRegistry() []Entry {
	registry := make([]Entry, len(staticRegistry))
	copy(registry, staticRegistry)

	gatherers := []api.ArtifactsGatherer{fedora.NewGatherer()}
	gatherArtifacts(&registry, gatherers)

	return registry
}

func ShouldSkip(focus string, entry *Entry) bool {
	if focus == "" {
		return entry.SkipWhenNotFocused
	}

	focusSplit := strings.Split(focus, ":")
	wildcardFocus := len(focusSplit) == 2 && focusSplit[1] == "*"

	if wildcardFocus {
		return focusSplit[0] != entry.Artifact.Metadata().Name
	}

	return focus != entry.Artifact.Metadata().Describe()
}
