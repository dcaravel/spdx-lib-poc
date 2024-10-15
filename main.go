package main

import (
	"fmt"
	"os"
	"time"

	spdxjson "github.com/spdx/tools-golang/json"
	"github.com/spdx/tools-golang/spdx/v2/common"
	spdx "github.com/spdx/tools-golang/spdx/v2/v2_3"
	spdxtagvalue "github.com/spdx/tools-golang/tagvalue"
	spdxyaml "github.com/spdx/tools-golang/yaml"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	doc := gimmieDocV23()

	// Create JSON
	path := "poc.spdx.json"
	writeJson(doc, path)
	fmt.Printf("Wrote: %s\n", path)

	// Create Yaml
	path = "poc.spdx.yaml"
	writeYaml(doc, path)
	fmt.Printf("Wrote: %s\n", path)

	// Create TagValue
	path = "poc.spdx.tv"
	writeTagValue(doc, path)
	fmt.Printf("Wrote: %s\n", path)

}

func writeJson(doc *spdx.Document, path string) {
	f, err := os.Create(path)
	check(err)

	err = spdxjson.Write(doc, f, spdxjson.Indent("  "))
	check(err)
}

func writeYaml(doc *spdx.Document, path string) {
	f, err := os.Create(path)
	check(err)

	err = spdxyaml.Write(doc, f)
	check(err)
}

func writeTagValue(doc *spdx.Document, path string) {
	f, err := os.Create(path)
	check(err)

	err = spdxtagvalue.Write(doc, f)
	check(err)
}

func gimmieDocV23() *spdx.Document {
	return &spdx.Document{
		DataLicense:       spdx.DataLicense,
		SPDXVersion:       spdx.Version,
		SPDXIdentifier:    "DOCUMENT",
		DocumentName:      "doc name hello!",
		DocumentNamespace: "what should namespace be for ACS produced sboms?",
		CreationInfo: &spdx.CreationInfo{
			Created: time.Now().UTC().Format(time.RFC3339),
			Creators: []common.Creator{
				{Creator: "StackRox-1.2.3", CreatorType: "Tool"},
			},
		},
		Packages: []*spdx.Package{
			{
				PackageName:           "glibc",
				PackageSPDXIdentifier: "Package",
				PackageVersion:        "2.11.1",
				PackageFileName:       "glibc-2.11.1.tar.gz",
				PackageSupplier: &common.Supplier{
					Supplier:     "Jane Doe (jane.doe@example.com)",
					SupplierType: "Person",
				},
				PackageOriginator: &common.Originator{
					Originator:     "ExampleCodeInspect (contact@example.com)",
					OriginatorType: "Organization",
				},
				PackageDownloadLocation: "http://ftp.gnu.org/gnu/glibc/glibc-ports-2.15.tar.gz",
				PackageHomePage:         "http://ftp.gnu.org/gnu/glibc",
				PackageSourceInfo:       "uses glibc-2_11-branch from git://sourceware.org/git/glibc.git.",
				PackageCopyrightText:    "Copyright 2008-2010 John Smith",
				PackageSummary:          "GNU C library.",
				PackageDescription:      "The GNU C Library defines functions that are specified by the ISO C standard, as well as additional features specific to POSIX and other derivatives of the Unix operating system, and extensions specific to GNU systems.",
				PackageComment:          "",
				PackageExternalReferences: []*spdx.PackageExternalReference{
					{
						Category: "SECURITY",
						RefType:  "cpe23Type",
						Locator:  "cpe:2.3:a:pivotal_software:spring_framework:4.1.0:*:*:*:*:*:*:*",
					},
				},
				PackageAttributionTexts: []string{
					"The GNU C Library is free software.  See the file COPYING.LIB for copying conditions, and LICENSES for notices about a few contributions that require these additional notices to be distributed.  License copyright years may be listed using range notation, e.g., 1996-2015, indicating that every year in the range, inclusive, is a copyrightable year that would otherwise be listed individually.",
				},
			},
		},
		Relationships: []*spdx.Relationship{
			{
				RefA:                common.MakeDocElementID("", "DOCUMENT"),
				RefB:                common.MakeDocElementID("", "Package"),
				Relationship:        "CONTAINS",
				RelationshipComment: "A relationship comment",
			},
		},
	}
}
