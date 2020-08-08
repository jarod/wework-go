package robot

import "testing"

func TestSend(t *testing.T) {
	c := New("test api key")
	r, err := c.Send(Text("test"))
	t.Log(r, err)
	println(r, err)
}
