package alpine

import g "github.com/maragudk/gomponents"

// Data render the x-data attribute.
func Data(value string) g.Node {
	return g.Attr("x-data", value)
}

// Bind render the x-bind attribute.
func Bind(attr, value string) g.Node {
	return g.Attr("x-bind:"+attr, value)
}

// Bind render the x-on attribute.
func On(attr, value string) g.Node {
	return g.Attr("x-on:"+attr, value)
}

// Bind render the x-text attribute.
func Text(value string) g.Node {
	return g.Attr("x-text", value)
}

// Bind render the x-html attribute.
func HTML(value string) g.Node {
	return g.Attr("x-html", value)
}

// Bind render the x-model attribute.
func Model(value string) g.Node {
	return g.Attr("x-model", value)
}

// Bind render the x-show attribute.
func Show(value string) g.Node {
	return g.Attr("x-show", value)
}

// Bind render the x-transition attribute.
func Transition() g.Node {
	return g.Attr("x-transition")
}

// Bind render the x-for attribute.
func For(value string) g.Node {
	return g.Attr("x-for", value)
}

// Bind render the x-if attribute.
func If(value string) g.Node {
	return g.Attr("x-if", value)
}

// Bind render the x-init attribute.
func Init(value string) g.Node {
	return g.Attr("x-init", value)
}

// Bind render the x-effect attribute.
func Effect(value string) g.Node {
	return g.Attr("x-effect", value)
}

// Bind render the x-ref attribute.
func Ref(value string) g.Node {
	return g.Attr("x-ref", value)
}

// Bind render the x-cloak attribute.
func Cloak() g.Node {
	return g.Attr("x-cloak")
}

// Bind render the x-ignore attribute.
func Ignore() g.Node {
	return g.Attr("x-ignore")
}
