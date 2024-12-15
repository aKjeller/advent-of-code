package main

type dir [2]int

var (
	UP    dir = [2]int{-1, 0}
	DOWN  dir = [2]int{1, 0}
	LEFT  dir = [2]int{0, -1}
	RIGHT dir = [2]int{0, 1}
)

type mover interface {
	move(m [][]mover, d dir)
	canMove(m [][]mover, d dir) bool
	getType() string
}

type object struct {
	x       int
	y       int
	movable bool
	t       string
}

type wideObject struct {
	object
}

func (o *object) getType() string {
	return o.t
}

func (o *object) canMove(m [][]mover, d dir) bool {
	if !o.movable {
		return false
	}
	other := m[o.x+d[0]][o.y+d[1]]
	if other != nil {
		return other.canMove(m, d)
	}
	return true
}

func (o *object) move(m [][]mover, d dir) {
	other := m[o.x+d[0]][o.y+d[1]]
	if other != nil {
		other.move(m, d)
	}

	m[o.x][o.y] = nil
	o.x += d[0]
	o.y += d[1]
	m[o.x][o.y] = o
}

func (o *wideObject) getType() string {
	return o.t
}

func (o *wideObject) canMove(m [][]mover, d dir) bool {
	if d == LEFT || d == RIGHT {
		return o.object.canMove(m, d)
	}

	var other *wideObject
	if o.getType() == "[" {
		other, _ = m[o.x][o.y+1].(*wideObject)
	} else {
		other, _ = m[o.x][o.y-1].(*wideObject)
	}

	return o.object.canMove(m, d) && other.object.canMove(m, d)
}

func (o *wideObject) move(m [][]mover, d dir) {
	if d == UP || d == DOWN {
		other := m[o.x+d[0]][o.y+d[1]]
		if other != nil {
			other.move(m, d)
		}

		if o.getType() == "[" {
			other = m[o.x][o.y+1]
		} else {
			other = m[o.x][o.y-1]
		}

		if other != nil {
			m[o.x][o.y] = nil
			o.x += d[0]
			o.y += d[1]
			m[o.x][o.y] = o
			other.move(m, d)
		} else {
			m[o.x][o.y] = nil
			o.x += d[0]
			o.y += d[1]
			m[o.x][o.y] = o
		}
	} else {
		other := m[o.x+d[0]][o.y+d[1]]
		if other != nil {
			other.move(m, d)
		}

		m[o.x][o.y] = nil
		o.x += d[0]
		o.y += d[1]
		m[o.x][o.y] = o
	}
}
