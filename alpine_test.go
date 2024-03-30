package alpine_test

import (
	"testing"

	a "github.com/jlucasnsilva/gomponents-alpine"
	g "github.com/maragudk/gomponents"
	"github.com/stretchr/testify/assert"
)

func TestCases(t *testing.T) {
	cases := map[string]struct {
		input    g.Node
		expected g.Node
	}{
		"Data": {
			input:    a.Data("{ open: false }"),
			expected: g.Attr("x-data", "{ open: false }"),
		},
		"Bind": {
			input:    a.Bind("class", "! open ? 'hidden' : ''"),
			expected: g.Attr("x-bind:class", "! open ? 'hidden' : ''"),
		},
		"On": {
			input:    a.On("click", "open = ! open"),
			expected: g.Attr("x-on:click", "open = ! open"),
		},
		"Text": {
			input:    a.Text("new Date().getFullYear()"),
			expected: g.Attr("x-text", "new Date().getFullYear()"),
		},
		"HTML": {
			input:    a.HTML("(await axios.get('/some/html/partial')).data"),
			expected: g.Attr("x-html", "(await axios.get('/some/html/partial')).data"),
		},
		"Model": {
			input:    a.Model("search"),
			expected: g.Attr("x-model", "search"),
		},
		"Show": {
			input:    a.Show("open"),
			expected: g.Attr("x-show", "open"),
		},
		"Transition": {
			input:    a.Transition(),
			expected: g.Attr("x-transition"),
		},
		"For": {
			input:    a.For("post in posts"),
			expected: g.Attr("x-for", "post in posts"),
		},
		"If": {
			input:    a.If("open"),
			expected: g.Attr("x-if", "open"),
		},
		"Init": {
			input:    a.Init("date = new Date()"),
			expected: g.Attr("x-init", "date = new Date()"),
		},
		"Effect": {
			input:    a.Effect("console.log('Count is '+count)"),
			expected: g.Attr("x-effect", "console.log('Count is '+count)"),
		},
		"Ref": {
			input:    a.Ref("content"),
			expected: g.Attr("x-ref", "content"),
		},
		"Cloak": {
			input:    a.Cloak(),
			expected: g.Attr("x-cloak"),
		},
		"Ignore": {
			input:    a.Ignore(),
			expected: g.Attr("x-ignore"),
		},
	}

	for label, c := range cases {
		t.Run(label, func(t *testing.T) {
			assert.Equal(t, c.expected, c.input)
		})
	}
}
