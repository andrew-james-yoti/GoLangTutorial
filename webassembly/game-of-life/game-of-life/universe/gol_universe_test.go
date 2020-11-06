package universe

import "testing"

func TestCreateUniverse(t *testing.T) {
	testUniverse := createUniverse()
	testCell := &testUniverse.Cells[0][0]
	testCell.Alive = true

	if testCell.Alive != true {
		t.Errorf("testUniverse.Cells[0][0] should be true %t", testUniverse.Cells[0][0].Alive)
	}
}

func TestGetRowStartIndex(t *testing.T) {
	if getRowStartIndex(0) != 0 {
		t.Errorf("row at index 0 should not be less than 0 %d", getRowStartIndex(0))
	}

	if getRowStartIndex(10) != 9 {
		t.Errorf("row at index 10 should be -1 from specified index %d", getRowStartIndex(10))
	}
}

func TestGetRowEndIndex(t *testing.T) {
	testUniverse := createUniverse()

	if (getRowEndIndex(&testUniverse.Cells, 0)) != 1 {
		t.Errorf("row at index 0 should be +1 from the specified index %d", getRowEndIndex(&testUniverse.Cells, 0))
	}

	if (getRowEndIndex(&testUniverse.Cells, 10)) != 10 {
		t.Errorf("row at index 10 should not be greater than the length of the array %d", getRowEndIndex(&testUniverse.Cells, 10))
	}
}

func TestGetColStartIndex(t *testing.T) {
	if getColStartIndex(0) != 0 {
		t.Errorf("column at index 0 should not be less than 0 %d", getColStartIndex(0))
	}

	if getColStartIndex(10) != 9 {
		t.Errorf("column at index 10 should be -1 from specified index %d", getColStartIndex(10))
	}
}

func TestGetColEndIndex(t *testing.T) {
	testUniverse := createUniverse()

	if getColEndIndex(&testUniverse.Cells[0], 0) != 1 {
		t.Errorf("column at index 0 should be +1 from the specified index %d", getColEndIndex(&testUniverse.Cells[0], 0))
	}

	if getColEndIndex(&testUniverse.Cells[0], 10) != 10 {
		t.Errorf("column at index 10 should not be greater than the length of the array %d", getColEndIndex(&testUniverse.Cells[0], 10))
	}
}

func TestCountAdjacentLiveCells(t *testing.T) {
	// initialise
	universe := Universe{}

	for i := 0; i < len(universe.Cells); i++ {
		for j := 0; j < len(universe.Cells[i]); j++ {
			universe.Cells[i][j].Alive = false
		}
	}

	// cell 1 - 1 is alive
	universe.Cells[1][1].Alive = true

	// test for no alive adjacent cells
	test1 := countAdjacentLiveCells(&universe.Cells, 1, 1)

	if test1 != 0 {
		t.Errorf("test1 should be 0 - actual value %d", test1)
	}

	// test for 3 alive adjacent cells
	universe.Cells[0][0].Alive = true
	universe.Cells[0][1].Alive = true
	universe.Cells[2][2].Alive = true

	test2 := countAdjacentLiveCells(&universe.Cells, 1, 1)

	if test2 != 3 {
		t.Errorf("test2 should be 3 - actual value %d", test2)
	}
}

func TestGolRule1(t *testing.T) {
	// initialise
	universe := Universe{}

	for i := 0; i < len(universe.Cells); i++ {
		for j := 0; j < len(universe.Cells[i]); j++ {
			universe.Cells[i][j].Alive = false
		}
	}

	// cell 1 - 1 is alive
	universe.Cells[1][1].Alive = true

	// Test 1 - cell should die
	golRule1(&universe.Cells, 1, 1)

	if universe.Cells[1][1].Alive {
		t.Errorf("Cell should be dead")
	}

	// Test 2 - cell should live
	// cell 1 - 1 is alive
	universe.Cells[1][1].Alive = true
	// adjacent cells
	universe.Cells[0][0].Alive = true
	universe.Cells[0][1].Alive = true

	golRule1(&universe.Cells, 1, 1)

	if !universe.Cells[1][1].Alive {
		t.Errorf("Cell should be alive")
	}
}
