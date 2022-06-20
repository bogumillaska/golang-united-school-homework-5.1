package square

import "testing"

func TestArea(t *testing.T) {
	got := Square{start: Point{x: 1, y: 2}, a: 10}
	if got.Area() != 100 {
		t.Errorf("Incorrect area value, got :%v want :%v", got.Area(), 100)
	}
}

func TestPeremiter(t *testing.T) {
	got := Square{start: Point{x: 1, y: 2}, a: 10}
	if got.Perimeter() != 40 {
		t.Errorf("Incorrect peremiter value, got :%v want :%v", got.Perimeter(), 40)
	}
}

func TestEnd(t *testing.T) {
	got := Square{start: Point{x: 1, y: 2}, a: 4}

	if got.End().x != 5 {
		t.Errorf("Incorrect x value, got :%v want :%v", got.End().x, 5)
	}

	if got.End().y != 6 {
		t.Errorf("Incorrect y value, got :%v want :%v", got.End().y, 6)
	}
}
