package ds

type DsuItem struct {
	Rank   int
	Parent int
}

type Dsu []DsuItem

func (u *Dsu) Add() int {
	index := len(*u)
	item := DsuItem{
		Rank:   0,
		Parent: index,
	}
	*u = append(*u, item)
	return index
}

func (u Dsu) Find(x int) int {
	if u[x].Parent != x {
		u[x].Parent = u.Find(u[x].Parent)
		return u[x].Parent
	}
	return x
}

func (u Dsu) Union(x, y int) bool {
	x = u.Find(x)
	y = u.Find(y)

	if x == y {
		return false
	}

	if u[x].Rank < u[y].Rank {
		x, y = y, x
	}

	u[y].Parent = x

	if u[x].Rank == u[y].Rank {
		u[y].Rank += 1
	}
	return true
}
