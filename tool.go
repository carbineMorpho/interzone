
package main

type toolf func()

type tool struct {
	name string
	i interface{}
	f toolf
}

func (w *world) Gun(){
	w.Monitor("Click Target")

	if p, prs := mouseGet(); prs{

		p.X += w.player.p.X
		p.Y += w.player.p.Y

		if e := w.CollideCheck(p); e != nil{
			log("bang!")
			e.hp -= 1
		}
	}	
}

func (w *world) Shoes(){
	w.player.p.Y += 1
}

func (t *tool) PosSet(i interface{}) {
	t.i = i
}
