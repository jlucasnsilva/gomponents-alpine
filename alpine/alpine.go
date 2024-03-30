package alpine

import "github.com/maragudk/gomponents"

// Data render the x-data attribute.
func Data(value string) gomponents.Node {
	return attribute("x-data", value)
}

// Bind render the x-bind attribute.
func Bind(attr, value string) gomponents.Node {
	return attribute("x-bind:"+attr, value)
}

// Bind render the x-on attribute.
func On(attr, value string) gomponents.Node {
	return attribute("x-on:"+attr, value)
}

// Bind render the x-text attribute.
func Text(value string) gomponents.Node {
	return attribute("x-text", value)
}

// Bind render the x-html attribute.
func HTML(value string) gomponents.Node {
	return attribute("x-html", value)
}

// Bind render the x-model attribute.
func Model(value string) gomponents.Node {
	return attribute("x-model", value)
}

// Bind render the x-show attribute.
func Show(value string) gomponents.Node {
	return attribute("x-show", value)
}

// Bind render the x-transition attribute.
func Transition() gomponents.Node {
	return attribute("x-transition")
}

// Bind render the x-for attribute.
func For(value string) gomponents.Node {
	return attribute("x-for", value)
}

// Bind render the x-if attribute.
func If(value string) gomponents.Node {
	return attribute("x-if", value)
}

// Bind render the x-init attribute.
func Init(value string) gomponents.Node {
	return attribute("x-init", value)
}

// Bind render the x-effect attribute.
func Effect(value string) gomponents.Node {
	return attribute("x-effect", value)
}

// Bind render the x-ref attribute.
func Ref(value string) gomponents.Node {
	return attribute("x-ref", value)
}

// Bind render the x-cloak attribute.
func Cloak() gomponents.Node {
	return attribute("x-cloak")
}

// Bind render the x-ignore attribute.
func Ignore() gomponents.Node {
	return attribute("x-ignore")
}

func attribute(name string, value ...string) gomponents.Node {
	switch len(value) {
	case 0:
		return gomponents.Raw(" " + name)
	case 1:
		return gomponents.Rawf(` %v="%v"`, name, value[0])
	default:
		panic("attribute must be just name or name and value pair")
	}
}
