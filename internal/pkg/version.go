package pkg

import (
	"bytes"
	"fmt"
	"hash/crc32"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Version struct {
	tmp      string
	Prefix   string
	Major    int
	Minor    int
	Patch    int
	Revision string
	Build    string
	Tag      string
}

func (ver *Version) Parse(v string) error {
	ver.prefix(v)

	for i := 0; i < len(ver.tmp); i++ {
		if ver.tmp[i] == 43 || ver.tmp[i] == 45 {
			version := strings.Split(ver.tmp[0:i], ".")
			ver.version(version)
			break
		}

		if i == len(ver.tmp)-1 {
			version := strings.Split(ver.tmp[0:i+1], ".")
			ver.version(version)
			break
		}
	}

	return nil
}

func (ver *Version) ParseMetadata() {
	if idx := strings.LastIndex(ver.tmp, "-"); idx != -1 {
		ver.Build, ver.Revision, _ = strings.Cut(ver.tmp[idx+1:], "+")
	}
	if idx := strings.LastIndex(ver.tmp, "+"); idx != -1 {
		ver.Revision = ver.tmp[idx+1:]

	}
}

func (ver *Version) prefix(v string) {

	if 48 < v[0] && v[0] <= 57 {
		ver.tmp = v
		return
	}

	ver.Prefix = "v"

	if v[0] > 127 {
		ver.tmp = v[2:]
		return
	}

	ver.tmp = v[1:]
}

func (ver *Version) version(v []string) {
	if len(v) > 3 {
		log.Fatal(fmt.Errorf("%s", "the version was not defined. version contains more than 3 block"))
	}

	block, err := ver.convertToInt(v[0])
	if err != nil {
		log.Fatal(err.Error())
	}
	ver.Major = block

	block, err = ver.convertToInt(v[1])
	if err != nil {
		log.Fatal(err.Error())
	}

	ver.Minor = block

	block, err = ver.convertToInt(v[2])
	if err != nil {
		log.Fatal(err.Error())
	}
	ver.Patch = block
}

func (ver *Version) convertToInt(v string) (int, error) {
	i, err := strconv.Atoi(v)
	if err != nil {
		return 0, fmt.Errorf("prefix from version does not contains digit | identifier: %s", v)
	}

	return i, nil
}

func (ver *Version) String() string {
	if ver.Build == "" && ver.Revision != "" {
		return fmt.Sprintf("%s%d.%d.%d+%s", ver.Prefix, ver.Major, ver.Minor, ver.Patch, ver.Revision)
	}

	if ver.Build != "" && ver.Revision == "" {
		return fmt.Sprintf("%s%d.%d.%d-%s", ver.Prefix, ver.Major, ver.Minor, ver.Patch, ver.Build)
	}
	if ver.Build != "" && ver.Revision != "" {
		return fmt.Sprintf("%s%d.%d.%d-%s+%s", ver.Prefix, ver.Major, ver.Minor, ver.Patch, ver.Build, ver.Revision)
	}

	return fmt.Sprintf("%s%d.%d.%d", ver.Prefix, ver.Major, ver.Minor, ver.Patch)
}

func (ver *Version) GetTag() string {
	if ver.Build == "" && ver.Revision != "" {
		return fmt.Sprintf("%s%d.%d.%d.%s", ver.Prefix, ver.Major, ver.Minor, ver.Patch, ver.Revision)
	}

	if ver.Build != "" && ver.Revision == "" {
		return fmt.Sprintf("%s%d.%d.%d.%s", ver.Prefix, ver.Major, ver.Minor, ver.Patch, ver.Build)
	}
	if ver.Build != "" && ver.Revision != "" {
		return fmt.Sprintf("%s%d.%d.%d.%s.%s", ver.Prefix, ver.Major, ver.Minor, ver.Patch, ver.Build, ver.Revision)
	}

	return fmt.Sprintf("%s%d.%d.%d", ver.Prefix, ver.Major, ver.Minor, ver.Patch)
}

func (ver *Version) UpPatch() {
	if ver.Patch >= 9999 {
		ver.Patch = 0
	}
	ver.Patch++
}

func (ver *Version) UpRelease() {
	ver.Patch = 0

	if ver.Minor >= 999 {
		ver.Minor = 0
		ver.Major++
	}
	ver.Minor++
}

func (ver *Version) InitRevision() {
	var xRevision = [10]byte{48, 49, 50, 51, 52, 53, 54, 55, 56, 57}

	nano := time.Now().UTC().Nanosecond()
	crcTable := crc32.MakeTable(crc32.IEEE)
	crc := crc32.New(crcTable)

	sum := crc.Sum([]byte(strconv.Itoa(nano)))

	if sum[len(sum)-1] == 0 {
		sum = bytes.Replace(sum, []byte{0, 0, 0, 0},
			[]byte{xRevision[rand.Intn(10)],
				xRevision[rand.Intn(10)],
				xRevision[rand.Intn(10)],
				xRevision[rand.Intn(10)]}, 3)
	}

	ver.Revision = string(sum)
}
