package golang_united_school_homework

import "errors"

var (
	errorOutOfRangeCapacity = errors.New("box already contain maximum quantity of shapes")
	errorIndexOutOfRange    = errors.New("requested shape index is out of range (no shape found or shapes capacity exceded)")
	errorCirclesNotFound    = errors.New("there are no circles in list of shapes)")
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) < b.shapesCapacity {
		b.shapes = append(b.shapes, shape)
		return nil
	}
	return errorOutOfRangeCapacity
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i >= 0 && i < len(b.shapes) {
		return b.shapes[i], nil
	}
	return nil, errorIndexOutOfRange
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	s, e := b.GetByIndex(i)
	if e != nil {
		return nil, e
	}
	b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
	return s, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	s, e := b.GetByIndex(i)
	if e != nil {
		return nil, e
	}
	b.shapes[i] = shape
	return s, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var result float64
	for _, s := range b.shapes {
		result += s.CalcPerimeter()
	}
	return result
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var result float64
	for _, s := range b.shapes {
		result += s.CalcArea()
	}
	return result
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	newShapes := make([]Shape, 0)
	for _, s := range b.shapes {
		switch s.(type) {
		default:
			newShapes = append(newShapes, s)
		case *Circle, Circle:
			// skip appending to new shapes slice
		}
	}
	if len(newShapes) == len(b.shapes) {
		return errorCirclesNotFound
	}
	b.shapes = newShapes
	return nil
}
