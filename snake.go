package main

type Point struct {
	x, y int
}

type Snake struct {
	Dots []Point
	Dir  Point
}

func NewSnake() *Snake {
	return &Snake{
		Dots: []Point{{1, 1}, {2, 1}, {3, 1}},
		Dir:  Point{1, 0},
	}
}

func (s *Snake) Eat(food Point) {
	newHead := food

	s.Dots = append([]Point{newHead}, s.Dots...)
}

func (s *Snake) Move() {
	head := s.Dots[0]
	newHead := Point{
		head.x + s.Dir.x,
		head.y + s.Dir.y,
	}

	if newHead.x < 0 {
		newHead.x = 19
	}
	if newHead.x > 19 {
		newHead.x = 0
	}
	if newHead.y < 0 {
		newHead.y = 19
	}
	if newHead.y > 19 {
		newHead.y = 0
	}

	s.Dots = append([]Point{newHead}, s.Dots...)

	s.Dots = s.Dots[:len(s.Dots)-1]
}

func (s *Snake) ChangeDirection(dir Point) {
	s.Dir = dir
}

func (s *Snake) ReachFood(food Point) bool {
	head := s.Dots[0]
	if head.x == food.x && head.y == food.y {
		return true
	}
	return false
}
