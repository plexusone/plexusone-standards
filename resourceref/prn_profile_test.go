// Package prn validates PlexusOne's own "prn" scheme instance
// (prn-profile.json) against the resourceref engine. This is not a
// library other repos import — it exists so CI can prove the profile
// document is both schema-valid and functionally correct against the
// real reference examples used across the ecosystem's design docs, and
// so a broken edit to prn-profile.json fails a test, not a review.
//
// Named prn rather than matching this directory's own name (resourceref)
// to avoid a file where "resourceref" ambiguously means both the local
// package and the imported github.com/plexusone/resourceref module.
package prn

import (
	_ "embed"
	"errors"
	"testing"

	"github.com/plexusone/resourceref"
	"github.com/plexusone/resourceref/scheme"
)

//go:embed prn-profile.json
var prnProfileJSON []byte

func loadPRNProfile(t *testing.T) *scheme.Profile {
	t.Helper()
	p, err := scheme.Load(prnProfileJSON)
	if err != nil {
		t.Fatalf("scheme.Load(prn-profile.json): %v", err)
	}
	return p
}

func TestPRNProfileLoads(t *testing.T) {
	p := loadPRNProfile(t)
	if p.Scheme != "prn" {
		t.Fatalf("Scheme = %q, want %q", p.Scheme, "prn")
	}
	if len(p.Authorities) != 3 {
		t.Fatalf("len(Authorities) = %d, want 3", len(p.Authorities))
	}
	if len(p.Shapes) != 2 {
		t.Fatalf("len(Shapes) = %d, want 2", len(p.Shapes))
	}
}

// TestPRNProfileGoldenReferences exercises the real prn:// examples used
// throughout the ecosystem's design docs (crossforge's TRD, the
// resourceref README, the plexusone-rest API style profile) against the
// actual instance profile, not a test fixture — proving those examples
// are not just illustrative but genuinely valid PRNs under the profile
// that governs them.
func TestPRNProfileGoldenReferences(t *testing.T) {
	p := loadPRNProfile(t)

	cases := []struct {
		name string
		ref  string
		want scheme.Parts
	}{
		{
			name: "dashforge notebook pinned to an alias",
			ref:  "prn://plexusone/dashforge/org_7f3/notebook/nb_churn@prod",
			want: scheme.Parts{
				Authority: "plexusone", App: "dashforge", Org: "org_7f3",
				ResourceType: "notebook", ID: "nb_churn", Version: "prod",
			},
		},
		{
			name: "actionforge workflow pinned to a numeric version",
			ref:  "prn://plexusone/actionforge/org_1/workflow/nightly-etl@14",
			want: scheme.Parts{
				Authority: "plexusone", App: "actionforge", Org: "org_1",
				ResourceType: "workflow", ID: "nightly-etl", Version: "14",
			},
		},
		{
			name: "actionforge workflow-run, a run resource",
			ref:  "prn://plexusone/actionforge/org_1/workflow-run/run_abc",
			want: scheme.Parts{
				Authority: "plexusone", App: "actionforge", Org: "org_1",
				ResourceType: "workflow-run", ID: "run_abc",
			},
		},
		{
			name: "productbuildershq roadmap item, an org-less ecosystem artifact",
			ref:  "prn://productbuildershq/roadmap/RMI-142",
			want: scheme.Parts{
				Authority: "productbuildershq", ResourceType: "roadmap", ID: "RMI-142",
			},
		},
		{
			name: "app with no resourceType registry entry is unrestricted",
			ref:  "prn://plexusone/agentforge/org_1/agent/data-analyst",
			want: scheme.Parts{
				Authority: "plexusone", App: "agentforge", Org: "org_1",
				ResourceType: "agent", ID: "data-analyst",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ref := resourceref.MustParse(tc.ref)
			got, err := p.Split(ref)
			if err != nil {
				t.Fatalf("Split(%q): unexpected error: %v", tc.ref, err)
			}
			if got != tc.want {
				t.Fatalf("Split(%q) = %+v, want %+v", tc.ref, got, tc.want)
			}
		})
	}
}

func TestPRNProfileRejectsInvalid(t *testing.T) {
	p := loadPRNProfile(t)

	cases := []struct {
		name string
		ref  string
	}{
		{name: "unregistered authority", ref: "prn://unknown-org/dashforge/org_1/notebook/nb_1"},
		{name: "dashforge notebook with unregistered resourceType", ref: "prn://plexusone/dashforge/org_1/bogus-type/x"},
		{name: "wrong scheme", ref: "https://plexusone/dashforge/org_1/notebook/nb_1"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if err := p.Validate(resourceref.MustParse(tc.ref)); !errors.Is(err, scheme.ErrNoMatch) {
				t.Fatalf("Validate(%q): err = %v, want wrapping scheme.ErrNoMatch", tc.ref, err)
			}
		})
	}
}
