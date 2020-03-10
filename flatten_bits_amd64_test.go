package simdcsv

import (
	"testing"
	"reflect"
)

func TestFlattenBitsIncremental(t *testing.T) {

	testCases := []struct {
		masks    []uint64
		expected []uint32
	}{
		// Single mask
		{[]uint64{0x11}, []uint32{0x1, 0x4}},
		{[]uint64{0x100100100100}, []uint32{0x9, 0xc, 0xc, 0xc}},
		{[]uint64{0x100100100300}, []uint32{0x9, 0x1, 0xb, 0xc, 0xc}},
		{[]uint64{0x8101010101010101}, []uint32{0x1, 0x8, 0x8, 0x8, 0x8, 0x8, 0x8, 0x8, 0x7}},
		{[]uint64{0x4000000000000000}, []uint32{0x3f}},
		{[]uint64{0x8000000000000000}, []uint32{0x40}},
		{[]uint64{0xf000000000000000}, []uint32{0x3d, 0x1, 0x1, 0x1}},
		{[]uint64{0xffffffffffffffff}, []uint32{
			0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1,
			0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1,
			0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1,
			0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1,
		}},
		////
		//// Multiple masks
		{[]uint64{0x1, 0x1000}, []uint32{0x1, 0x4c}},
		{[]uint64{0x1, 0x4000000000000000}, []uint32{0x1, 0x7e}},
		{[]uint64{0x1, 0x8000000000000000}, []uint32{0x1, 0x7f}},
		{[]uint64{0x1, 0x0, 0x8000000000000000}, []uint32{0x1, 0xbf}},
		{[]uint64{0x1, 0x0, 0x0, 0x8000000000000000}, []uint32{0x1, 0xff}},
		{[]uint64{0x100100100100100, 0x100100100100100}, []uint32{0x9, 0xc, 0xc, 0xc, 0xc, 0x10, 0xc, 0xc, 0xc, 0xc}},
		{[]uint64{0xffffffffffffffff, 0xffffffffffffffff}, []uint32{
			0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1,
			0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1,
			0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1,
			0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1,
			0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1,
			0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1,
			0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1,
			0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1,
		}},
	}

	for i, tc := range testCases {

		indexes := &[INDEX_SIZE]uint32{}
		length := 0
		carried := 0
		position := ^uint64(0)

		for _, mask := range tc.masks {
			flatten_bits_incremental(indexes, &length, mask, &carried, &position)
		}

		if length != len(tc.expected) {
			t.Errorf("TestFlattenBitsIncremental(%d): got: %d want: %d", i, length, len(tc.expected))
		}

		compare := make([]uint32, 0, 1024)
		for idx := 0; idx < length; idx++ {
			compare = append(compare, indexes[idx])
		}

		if !reflect.DeepEqual(compare, tc.expected) {
			t.Errorf("TestFlattenBitsIncremental(%d): got: %v want: %v", i, compare, tc.expected)
		}
	}
}

