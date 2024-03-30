package alpine

import (
	"strings"
	"testing"

	"github.com/maragudk/gomponents"
	"github.com/stretchr/testify/assert"
)

type (
	testCase struct {
		label    string
		expected string
		node     gomponents.Node
	}
)

func TestAlpine(t *testing.T) {
	tests := []testCase{
		{
			label:    "x-data",
			expected: ` x-data="{ open: false }"`,
			node:     Data("{ open: false }"),
		},
		{
			label:    "x-bind",
			expected: ` x-bind:class="! open ? 'hidden' : ''"`,
			node:     Bind("class", "! open ? 'hidden' : ''"),
		},
		{
			label:    "x-on",
			expected: ` x-on:click="open = ! open"`,
			node:     On("click", "open = ! open"),
		},
		{
			label:    "x-text",
			expected: ` x-text="new Date().getFullYear()"`,
			node:     Text("new Date().getFullYear()"),
		},
		{
			label:    "x-html",
			expected: ` x-html="(await axios.get('/some/html/partial')).data"`,
			node:     HTML("(await axios.get('/some/html/partial')).data"),
		},
		{
			label:    "x-model",
			expected: ` x-model="search"`,
			node:     Model("search"),
		},
		{
			label:    "x-show",
			expected: ` x-show="open"`,
			node:     Show("open"),
		},
		{
			label:    "x-transition",
			expected: ` x-transition`,
			node:     Transition(),
		},
		{
			label:    "x-for",
			expected: ` x-for="post in posts"`,
			node:     For("post in posts"),
		},
		{
			label:    "x-if",
			expected: ` x-if="open"`,
			node:     If("open"),
		},
		{
			label:    "x-init",
			expected: ` x-init="date = new Date()"`,
			node:     Init("date = new Date()"),
		},
		{
			label:    "x-effect",
			expected: ` x-effect="console.log('Count is '+count)"`,
			node:     Effect("console.log('Count is '+count)"),
		},
		{
			label:    "x-ref",
			expected: ` x-ref="content"`,
			node:     Ref("content"),
		},
		{
			label:    "x-cloak",
			expected: ` x-cloak`,
			node:     Cloak(),
		},
		{
			label:    "x-ignore",
			expected: ` x-ignore`,
			node:     Ignore(),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.label, func(t *testing.T) {
			res, err := render(test.node)
			assert.Nil(t, err)
			assert.Equal(t, test.expected, res)
		})
	}

	assert.Panics(t, func() {
		attribute("foo", "bar", "baz")
	})
}

func render(n gomponents.Node) (string, error) {
	b := strings.Builder{}
	if err := n.Render(&b); err != nil {
		return "", err
	}
	return b.String(), nil
}
