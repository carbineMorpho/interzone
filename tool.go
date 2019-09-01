package main

type toolf func(w *world, p pos)

type tool struct {
	name string
	i interface{}
	f toolf
}

func (t *tool) Demon(w *world, p pos){
	w.Build(rorschach(p), colorRandom())
}

func (t *tool) Base(w *world, p pos){
	w.Build(house(p, pos{p.X+24, p.Y+12}), colorRandom())
}

func (t *tool) PosSet(i interface{}) {
	t.i = i
}
