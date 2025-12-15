package pkg

import (
	"fmt"
	"testing"
)

func Test_Version(t *testing.T) {
	v := "v1.10.100"

	version := Version{}
	version.Parse(v)

	fmt.Println(version.String())

	if version.String() != "v1.10.100" {
		t.Fail()
	}

}

func Test_Version_without_prefix(t *testing.T) {
	v := "1.10.100"

	version := Version{}
	version.Parse(v)

	fmt.Println(version.String())

	if version.String() != "1.10.100" {
		t.Fail()
	}

}

func Test_Version_Revision(t *testing.T) {
	v := "v2.10.100+Build.01"

	version := Version{}
	version.Parse(v)
	version.ParseMetadata()

	fmt.Println(version.String())

	if version.String() != "v2.10.100+Build.01" {
		t.Fail()
	}
}

func Test_Version_Build(t *testing.T) {
	v := "v10.0.0-Alpha1"

	version := Version{}
	version.Parse(v)
	version.ParseMetadata()

	fmt.Println(version.String())

	if version.String() != "v10.0.0-Alpha1" {
		t.Fail()
	}
}

func Test_Version_Full(t *testing.T) {
	v := "v10.0.0-Alpha1+Build.01"

	version := Version{}
	version.Parse(v)
	version.ParseMetadata()

	fmt.Println(version.String())

	if version.String() != "v10.0.0-Alpha1+Build.01" {
		t.Fail()
	}
}

func Test_Version_Increment_Patch(t *testing.T) {
	v := "v10.0.0"

	version := Version{}
	version.Parse(v)
	version.ParseMetadata()
	version.UpPatch()

	fmt.Println(version.String())

	if version.String() != "v10.0.1" {
		t.Fail()
	}
}

func Test_Version_Increment_Patch_Revision(t *testing.T) {
	v := "v10.0.1"

	version := Version{}
	version.Parse(v)
	version.ParseMetadata()

	version.Revision = "08f5b8e8"

	version.UpPatch()

	fmt.Println(version.String())
	fmt.Println(version.GetTag())

	if version.String() != "v10.0.1+08f5b8e8" {
		t.Fail()
	}
}

func Test_Version_Increment_Minor(t *testing.T) {
	v := "v10.0.0"

	version := Version{}
	version.Parse(v)
	version.ParseMetadata()
	version.UpRelease()

	fmt.Println(version.String())

	if version.String() != "v10.1.0" {
		t.Fail()
	}
}

func Test_Version_Increment_Minor_Patch(t *testing.T) {
	v := "v10.0.0"

	version := Version{}
	version.Parse(v)
	version.ParseMetadata()
	version.UpPatch()
	version.UpRelease()

	fmt.Println(version.String())

	if version.String() != "v10.1.0" {
		t.Fail()
	}
}
