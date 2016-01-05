package bulk_test

import (
	"errors"
	"reflect"
	"strings"
	"testing"

	"github.com/kusubooru/local-tagger/bulk"
)

var errInvalidFormat = errors.New("invalid csv file format")
var loadCSVTests = []struct {
	in   string
	out  []bulk.Image
	oerr error
}{
	{"", []bulk.Image{}, nil},
	{",,,,", []bulk.Image{}, nil},
	{",,,,,,,,,,,,,,,,,,,", []bulk.Image(nil), errInvalidFormat},
	{"invalid format", []bulk.Image(nil), errInvalidFormat},
	{",,,,\n,,,,\n,,,,", []bulk.Image{}, nil},
	{
		"/server/path/img1,tag1 tag2,source,s,",
		[]bulk.Image{
			{Name: "img1", Tags: (bulk.Tags)([]string{"tag1", "tag2"}), Source: "source", Rating: "s"},
		},
		nil,
	},
	{
		"/server/path/img1,tag1 tag2,source1,s,\n" +
			",,,,\n" +
			"/server/path/img2,tag1 tag2,source2,q,\n",
		[]bulk.Image{
			{Name: "img1", Tags: (bulk.Tags)([]string{"tag1", "tag2"}), Source: "source1", Rating: "s"},
			{Name: "img2", Tags: (bulk.Tags)([]string{"tag1", "tag2"}), Source: "source2", Rating: "q"},
		},
		nil,
	},
}

func TestLoadCSV(t *testing.T) {

	for _, tt := range loadCSVTests {
		got, err := bulk.LoadCSV(strings.NewReader(tt.in))
		if want := tt.oerr; !reflect.DeepEqual(err, want) {
			t.Errorf("LoadCSV(%q) returned err %q, want %q", tt.in, err, want)
		}
		if want := tt.out; !reflect.DeepEqual(got, want) {
			t.Errorf("LoadCSV(%q) => %q, want %q", tt.in, got, want)
		}
	}
}

var currentPrefixTests = []struct {
	indir  string
	infile string
	out    string
	oerr   error
}{
	{"/local/path/dir", ",,,,", "/", nil},
	{"/local/path/dir", "/server/path/dir,,,,", "/server/path", nil},
	{"/local/path/dir", "", "", nil},
}

func TestCurrentPrefix(t *testing.T) {

	for _, tt := range currentPrefixTests {
		got, err := bulk.CurrentPrefix(tt.indir, strings.NewReader(tt.infile))
		if want := tt.oerr; !reflect.DeepEqual(err, want) {
			t.Errorf("CurrentPrefix(%q, %q) returned err %q, want %q", tt.indir, tt.infile, err, want)
		}
		if want := tt.out; !reflect.DeepEqual(got, want) {
			t.Errorf("CurrentPrefix(%q, %q) => %q, want %q", tt.indir, tt.infile, got, want)
		}
	}
}

var findByIDTests = []struct {
	images []bulk.Image
	id     int
	out    *bulk.Image
}{
	{nil, 0, nil},
	{[]bulk.Image{}, 0, nil},
	{[]bulk.Image{{ID: 1}}, 0, nil},
	{[]bulk.Image{{ID: 1}}, 1, &bulk.Image{ID: 1}},
	{[]bulk.Image{{ID: 2}, {ID: 1}}, 1, &bulk.Image{ID: 1}},
	{[]bulk.Image{{ID: 2}, {ID: 3}, {ID: 1}}, 3, &bulk.Image{ID: 3}},
	{[]bulk.Image{{ID: 2}, {ID: 3}, {ID: 1}}, 5, nil},
}

func TestFindByID(t *testing.T) {

	for _, tt := range findByIDTests {
		got := bulk.FindByID(tt.images, tt.id)
		if want := tt.out; !reflect.DeepEqual(got, want) {
			t.Errorf("FindByID(%q, %q) => %q, want %q", tt.images, tt.id, got, want)
		}
	}
}