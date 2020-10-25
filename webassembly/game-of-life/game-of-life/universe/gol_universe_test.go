package universe

import "testing"

func TestCreateUniverse(t *testing.T) {
	testUniverse := CreateUniverse()
	testCell := &testUniverse.Cells[0][0]
	testCell.Alive = true

	if testCell.Alive != true {
		t.Errorf("testUniverse.Cells[0][0] should be true %t", testUniverse.Cells[0][0].Alive)
	}
}

func TestRandCellValue(t *testing.T) {
	val := RandCellValue()
	if val != false {
		t.Errorf("actual value %t", val)
	}
}
