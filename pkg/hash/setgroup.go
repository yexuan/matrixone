package hash

import (
	"bytes"
	"matrixbase/pkg/container/batch"
	"matrixbase/pkg/container/types"
	"matrixbase/pkg/container/vector"
	"matrixbase/pkg/vm/process"
)

func NewSetGroup(idx, sel int64) *SetGroup {
	return &SetGroup{
		Idx: idx,
		Sel: sel,
	}
}

func (g *SetGroup) Free(_ *process.Process) {
}

func (g *SetGroup) Probe(sels, matched []int64, vecs []*vector.Vector,
	bats []*batch.Batch, diffs []bool, proc *process.Process) (int64, []int64, error) {
	for i, vec := range vecs {
		switch vec.Typ.Oid {
		case types.T_int8:
			gvec := bats[g.Idx].Vecs[i]
			lnull := vec.Nsp.Any()
			rnull := gvec.Nsp.Contains(uint64(g.Sel))
			switch {
			case lnull && rnull:
				for i, sel := range sels {
					if !vec.Nsp.Contains(uint64(sel)) { // only null eq null
						diffs[i] = true
					}
				}
			case lnull && !rnull: // null is not value
				vs := vec.Col.([]int8)
				gv := gvec.Col.([]int8)[g.Sel]
				for i, sel := range sels {
					if vec.Nsp.Contains(uint64(sel)) {
						diffs[i] = true
					} else {
						diffs[i] = diffs[i] || (gv != vs[sel])
					}
				}
			case !lnull && rnull: // null is not value
				for i := range sels {
					diffs[i] = true
				}
			default:
				vs := vec.Col.([]int8)
				gv := gvec.Col.([]int8)[g.Sel]
				for i, sel := range sels {
					diffs[i] = diffs[i] || (gv != vs[sel])
				}
			}
		case types.T_int16:
			gvec := bats[g.Idx].Vecs[i]
			lnull := vec.Nsp.Any()
			rnull := gvec.Nsp.Contains(uint64(g.Sel))
			switch {
			case lnull && rnull:
				for i, sel := range sels {
					if !vec.Nsp.Contains(uint64(sel)) { // only null eq null
						diffs[i] = true
					}
				}
			case lnull && !rnull: // null is not value
				vs := vec.Col.([]int16)
				gv := gvec.Col.([]int16)[g.Sel]
				for i, sel := range sels {
					if vec.Nsp.Contains(uint64(sel)) {
						diffs[i] = true
					} else {
						diffs[i] = diffs[i] || (gv != vs[sel])
					}
				}
			case !lnull && rnull: // null is not value
				for i := range sels {
					diffs[i] = true
				}
			default:
				vs := vec.Col.([]int16)
				gv := gvec.Col.([]int16)[g.Sel]
				for i, sel := range sels {
					diffs[i] = diffs[i] || (gv != vs[sel])
				}
			}
		case types.T_int32:
			gvec := bats[g.Idx].Vecs[i]
			lnull := vec.Nsp.Any()
			rnull := gvec.Nsp.Contains(uint64(g.Sel))
			switch {
			case lnull && rnull:
				for i, sel := range sels {
					if !vec.Nsp.Contains(uint64(sel)) { // only null eq null
						diffs[i] = true
					}
				}
			case lnull && !rnull: // null is not value
				vs := vec.Col.([]int32)
				gv := gvec.Col.([]int32)[g.Sel]
				for i, sel := range sels {
					if vec.Nsp.Contains(uint64(sel)) {
						diffs[i] = true
					} else {
						diffs[i] = diffs[i] || (gv != vs[sel])
					}
				}
			case !lnull && rnull: // null is not value
				for i := range sels {
					diffs[i] = true
				}
			default:
				vs := vec.Col.([]int32)
				gv := gvec.Col.([]int32)[g.Sel]
				for i, sel := range sels {
					diffs[i] = diffs[i] || (gv != vs[sel])
				}
			}
		case types.T_int64:
			gvec := bats[g.Idx].Vecs[i]
			lnull := vec.Nsp.Any()
			rnull := gvec.Nsp.Contains(uint64(g.Sel))
			switch {
			case lnull && rnull:
				for i, sel := range sels {
					if !vec.Nsp.Contains(uint64(sel)) { // only null eq null
						diffs[i] = true
					}
				}
			case lnull && !rnull: // null is not value
				vs := vec.Col.([]int64)
				gv := gvec.Col.([]int64)[g.Sel]
				for i, sel := range sels {
					if vec.Nsp.Contains(uint64(sel)) {
						diffs[i] = true
					} else {
						diffs[i] = diffs[i] || (gv != vs[sel])
					}
				}
			case !lnull && rnull: // null is not value
				for i := range sels {
					diffs[i] = true
				}
			default:
				vs := vec.Col.([]int64)
				gv := gvec.Col.([]int64)[g.Sel]
				for i, sel := range sels {
					diffs[i] = diffs[i] || (gv != vs[sel])
				}
			}
		case types.T_uint8:
			gvec := bats[g.Idx].Vecs[i]
			lnull := vec.Nsp.Any()
			rnull := gvec.Nsp.Contains(uint64(g.Sel))
			switch {
			case lnull && rnull:
				for i, sel := range sels {
					if !vec.Nsp.Contains(uint64(sel)) { // only null eq null
						diffs[i] = true
					}
				}
			case lnull && !rnull: // null is not value
				vs := vec.Col.([]uint8)
				gv := gvec.Col.([]uint8)[g.Sel]
				for i, sel := range sels {
					if vec.Nsp.Contains(uint64(sel)) {
						diffs[i] = true
					} else {
						diffs[i] = diffs[i] || (gv != vs[sel])
					}
				}
			case !lnull && rnull: // null is not value
				for i := range sels {
					diffs[i] = true
				}
			default:
				vs := vec.Col.([]uint8)
				gv := gvec.Col.([]uint8)[g.Sel]
				for i, sel := range sels {
					diffs[i] = diffs[i] || (gv != vs[sel])
				}
			}
		case types.T_uint16:
			gvec := bats[g.Idx].Vecs[i]
			lnull := vec.Nsp.Any()
			rnull := gvec.Nsp.Contains(uint64(g.Sel))
			switch {
			case lnull && rnull:
				for i, sel := range sels {
					if !vec.Nsp.Contains(uint64(sel)) { // only null eq null
						diffs[i] = true
					}
				}
			case lnull && !rnull: // null is not value
				vs := vec.Col.([]uint16)
				gv := gvec.Col.([]uint16)[g.Sel]
				for i, sel := range sels {
					if vec.Nsp.Contains(uint64(sel)) {
						diffs[i] = true
					} else {
						diffs[i] = diffs[i] || (gv != vs[sel])
					}
				}
			case !lnull && rnull: // null is not value
				for i := range sels {
					diffs[i] = true
				}
			default:
				vs := vec.Col.([]uint16)
				gv := gvec.Col.([]uint16)[g.Sel]
				for i, sel := range sels {
					diffs[i] = diffs[i] || (gv != vs[sel])
				}
			}
		case types.T_uint32:
			gvec := bats[g.Idx].Vecs[i]
			lnull := vec.Nsp.Any()
			rnull := gvec.Nsp.Contains(uint64(g.Sel))
			switch {
			case lnull && rnull:
				for i, sel := range sels {
					if !vec.Nsp.Contains(uint64(sel)) { // only null eq null
						diffs[i] = true
					}
				}
			case lnull && !rnull: // null is not value
				vs := vec.Col.([]uint32)
				gv := gvec.Col.([]uint32)[g.Sel]
				for i, sel := range sels {
					if vec.Nsp.Contains(uint64(sel)) {
						diffs[i] = true
					} else {
						diffs[i] = diffs[i] || (gv != vs[sel])
					}
				}
			case !lnull && rnull: // null is not value
				for i := range sels {
					diffs[i] = true
				}
			default:
				vs := vec.Col.([]uint32)
				gv := gvec.Col.([]uint32)[g.Sel]
				for i, sel := range sels {
					diffs[i] = diffs[i] || (gv != vs[sel])
				}
			}
		case types.T_uint64:
			gvec := bats[g.Idx].Vecs[i]
			lnull := vec.Nsp.Any()
			rnull := gvec.Nsp.Contains(uint64(g.Sel))
			switch {
			case lnull && rnull:
				for i, sel := range sels {
					if !vec.Nsp.Contains(uint64(sel)) { // only null eq null
						diffs[i] = true
					}
				}
			case lnull && !rnull: // null is not value
				vs := vec.Col.([]uint64)
				gv := gvec.Col.([]uint64)[g.Sel]
				for i, sel := range sels {
					if vec.Nsp.Contains(uint64(sel)) {
						diffs[i] = true
					} else {
						diffs[i] = diffs[i] || (gv != vs[sel])
					}
				}
			case !lnull && rnull: // null is not value
				for i := range sels {
					diffs[i] = true
				}
			default:
				vs := vec.Col.([]uint64)
				gv := gvec.Col.([]uint64)[g.Sel]
				for i, sel := range sels {
					diffs[i] = diffs[i] || (gv != vs[sel])
				}
			}
		case types.T_float32:
			gvec := bats[g.Idx].Vecs[i]
			lnull := vec.Nsp.Any()
			rnull := gvec.Nsp.Contains(uint64(g.Sel))
			switch {
			case lnull && rnull:
				for i, sel := range sels {
					if !vec.Nsp.Contains(uint64(sel)) { // only null eq null
						diffs[i] = true
					}
				}
			case lnull && !rnull: // null is not value
				vs := vec.Col.([]float32)
				gv := gvec.Col.([]float32)[g.Sel]
				for i, sel := range sels {
					if vec.Nsp.Contains(uint64(sel)) {
						diffs[i] = true
					} else {
						diffs[i] = diffs[i] || (gv != vs[sel])
					}
				}
			case !lnull && rnull: // null is not value
				for i := range sels {
					diffs[i] = true
				}
			default:
				vs := vec.Col.([]float32)
				gv := gvec.Col.([]float32)[g.Sel]
				for i, sel := range sels {
					diffs[i] = diffs[i] || (gv != vs[sel])
				}
			}
		case types.T_float64:
			gvec := bats[g.Idx].Vecs[i]
			lnull := vec.Nsp.Any()
			rnull := gvec.Nsp.Contains(uint64(g.Sel))
			switch {
			case lnull && rnull:
				for i, sel := range sels {
					if !vec.Nsp.Contains(uint64(sel)) { // only null eq null
						diffs[i] = true
					}
				}
			case lnull && !rnull: // null is not value
				vs := vec.Col.([]float64)
				gv := gvec.Col.([]float64)[g.Sel]
				for i, sel := range sels {
					if vec.Nsp.Contains(uint64(sel)) {
						diffs[i] = true
					} else {
						diffs[i] = diffs[i] || (gv != vs[sel])
					}
				}
			case !lnull && rnull: // null is not value
				for i := range sels {
					diffs[i] = true
				}
			default:
				vs := vec.Col.([]float64)
				gv := gvec.Col.([]float64)[g.Sel]
				for i, sel := range sels {
					diffs[i] = diffs[i] || (gv != vs[sel])
				}
			}
		case types.T_decimal:
		case types.T_date:
		case types.T_datetime:
		case types.T_char:
		case types.T_varchar:
			gvec := bats[g.Idx].Vecs[i]
			lnull := vec.Nsp.Any()
			rnull := gvec.Nsp.Contains(uint64(g.Sel))
			switch {
			case lnull && rnull:
				for i, sel := range sels {
					if !vec.Nsp.Contains(uint64(sel)) { // only null eq null
						diffs[i] = true
					}
				}
			case lnull && !rnull: // null is not value
				vs := vec.Col.(*types.Bytes)
				gvs := gvec.Col.(*types.Bytes)
				gv := gvs.Get(int(g.Sel))
				for i, sel := range sels {
					if vec.Nsp.Contains(uint64(sel)) {
						diffs[i] = true
					} else {
						diffs[i] = diffs[i] || (bytes.Compare(gv, vs.Get(int(sel))) != 0)
					}
				}
			case !lnull && rnull: // null is not value
				for i := range sels {
					diffs[i] = true
				}
			default:
				vs := vec.Col.(*types.Bytes)
				gvs := gvec.Col.(*types.Bytes)
				gv := gvs.Get(int(g.Sel))
				for i, sel := range sels {
					diffs[i] = diffs[i] || (bytes.Compare(gv, vs.Get(int(sel))) != 0)
				}
			}
		case types.T_json:
			gvec := bats[g.Idx].Vecs[i]
			lnull := vec.Nsp.Any()
			rnull := gvec.Nsp.Contains(uint64(g.Sel))
			switch {
			case lnull && rnull:
				for i, sel := range sels {
					if !vec.Nsp.Contains(uint64(sel)) { // only null eq null
						diffs[i] = true
					}
				}
			case lnull && !rnull: // null is not value
				vs := vec.Col.(*types.Bytes)
				gvs := gvec.Col.(*types.Bytes)
				gv := gvs.Get(int(g.Sel))
				for i, sel := range sels {
					if vec.Nsp.Contains(uint64(sel)) {
						diffs[i] = true
					} else {
						diffs[i] = diffs[i] || (bytes.Compare(gv, vs.Get(int(sel))) != 0)
					}
				}
			case !lnull && rnull: // null is not value
				for i := range sels {
					diffs[i] = true
				}
			default:
				vs := vec.Col.(*types.Bytes)
				gvs := gvec.Col.(*types.Bytes)
				gv := gvs.Get(int(g.Sel))
				for i, sel := range sels {
					diffs[i] = diffs[i] || (bytes.Compare(gv, vs.Get(int(sel))) != 0)
				}
			}
		}
	}
	n := len(sels)
	matched = matched[:0]
	remaining := sels[:0]
	for i := 0; i < n; i++ {
		if diffs[i] {
			remaining = append(remaining, sels[i])
		} else if len(matched) == 0 {
			matched = append(matched, sels[i])
		}
	}
	if len(matched) == 0 {
		return -1, remaining, nil
	}
	return matched[0], remaining, nil
}

func (g *SetGroup) Fill(sels, matched []int64, vecs []*vector.Vector,
	bats []*batch.Batch, diffs []bool, proc *process.Process) ([]int64, error) {
	for i, vec := range vecs {
		switch vec.Typ.Oid {
		case types.T_int8:
			gvec := bats[g.Idx].Vecs[i]
			lnull := vec.Nsp.Any()
			rnull := gvec.Nsp.Contains(uint64(g.Sel))
			switch {
			case lnull && rnull:
				for i, sel := range sels {
					if !vec.Nsp.Contains(uint64(sel)) { // only null eq null
						diffs[i] = true
					}
				}
			case lnull && !rnull: // null is not value
				vs := vec.Col.([]int8)
				gv := gvec.Col.([]int8)[g.Sel]
				for i, sel := range sels {
					if vec.Nsp.Contains(uint64(sel)) {
						diffs[i] = true
					} else {
						diffs[i] = diffs[i] || (gv != vs[sel])
					}
				}
			case !lnull && rnull: // null is not value
				for i := range sels {
					diffs[i] = true
				}
			default:
				vs := vec.Col.([]int8)
				gv := gvec.Col.([]int8)[g.Sel]
				for i, sel := range sels {
					diffs[i] = diffs[i] || (gv != vs[sel])
				}
			}
		case types.T_int16:
			gvec := bats[g.Idx].Vecs[i]
			lnull := vec.Nsp.Any()
			rnull := gvec.Nsp.Contains(uint64(g.Sel))
			switch {
			case lnull && rnull:
				for i, sel := range sels {
					if !vec.Nsp.Contains(uint64(sel)) { // only null eq null
						diffs[i] = true
					}
				}
			case lnull && !rnull: // null is not value
				vs := vec.Col.([]int16)
				gv := gvec.Col.([]int16)[g.Sel]
				for i, sel := range sels {
					if vec.Nsp.Contains(uint64(sel)) {
						diffs[i] = true
					} else {
						diffs[i] = diffs[i] || (gv != vs[sel])
					}
				}
			case !lnull && rnull: // null is not value
				for i := range sels {
					diffs[i] = true
				}
			default:
				vs := vec.Col.([]int16)
				gv := gvec.Col.([]int16)[g.Sel]
				for i, sel := range sels {
					diffs[i] = diffs[i] || (gv != vs[sel])
				}
			}
		case types.T_int32:
			gvec := bats[g.Idx].Vecs[i]
			lnull := vec.Nsp.Any()
			rnull := gvec.Nsp.Contains(uint64(g.Sel))
			switch {
			case lnull && rnull:
				for i, sel := range sels {
					if !vec.Nsp.Contains(uint64(sel)) { // only null eq null
						diffs[i] = true
					}
				}
			case lnull && !rnull: // null is not value
				vs := vec.Col.([]int32)
				gv := gvec.Col.([]int32)[g.Sel]
				for i, sel := range sels {
					if vec.Nsp.Contains(uint64(sel)) {
						diffs[i] = true
					} else {
						diffs[i] = diffs[i] || (gv != vs[sel])
					}
				}
			case !lnull && rnull: // null is not value
				for i := range sels {
					diffs[i] = true
				}
			default:
				vs := vec.Col.([]int32)
				gv := gvec.Col.([]int32)[g.Sel]
				for i, sel := range sels {
					diffs[i] = diffs[i] || (gv != vs[sel])
				}
			}
		case types.T_int64:
			gvec := bats[g.Idx].Vecs[i]
			lnull := vec.Nsp.Any()
			rnull := gvec.Nsp.Contains(uint64(g.Sel))
			switch {
			case lnull && rnull:
				for i, sel := range sels {
					if !vec.Nsp.Contains(uint64(sel)) { // only null eq null
						diffs[i] = true
					}
				}
			case lnull && !rnull: // null is not value
				vs := vec.Col.([]int64)
				gv := gvec.Col.([]int64)[g.Sel]
				for i, sel := range sels {
					if vec.Nsp.Contains(uint64(sel)) {
						diffs[i] = true
					} else {
						diffs[i] = diffs[i] || (gv != vs[sel])
					}
				}
			case !lnull && rnull: // null is not value
				for i := range sels {
					diffs[i] = true
				}
			default:
				vs := vec.Col.([]int64)
				gv := gvec.Col.([]int64)[g.Sel]
				for i, sel := range sels {
					diffs[i] = diffs[i] || (gv != vs[sel])
				}
			}
		case types.T_uint8:
			gvec := bats[g.Idx].Vecs[i]
			lnull := vec.Nsp.Any()
			rnull := gvec.Nsp.Contains(uint64(g.Sel))
			switch {
			case lnull && rnull:
				for i, sel := range sels {
					if !vec.Nsp.Contains(uint64(sel)) { // only null eq null
						diffs[i] = true
					}
				}
			case lnull && !rnull: // null is not value
				vs := vec.Col.([]uint8)
				gv := gvec.Col.([]uint8)[g.Sel]
				for i, sel := range sels {
					if vec.Nsp.Contains(uint64(sel)) {
						diffs[i] = true
					} else {
						diffs[i] = diffs[i] || (gv != vs[sel])
					}
				}
			case !lnull && rnull: // null is not value
				for i := range sels {
					diffs[i] = true
				}
			default:
				vs := vec.Col.([]uint8)
				gv := gvec.Col.([]uint8)[g.Sel]
				for i, sel := range sels {
					diffs[i] = diffs[i] || (gv != vs[sel])
				}
			}
		case types.T_uint16:
			gvec := bats[g.Idx].Vecs[i]
			lnull := vec.Nsp.Any()
			rnull := gvec.Nsp.Contains(uint64(g.Sel))
			switch {
			case lnull && rnull:
				for i, sel := range sels {
					if !vec.Nsp.Contains(uint64(sel)) { // only null eq null
						diffs[i] = true
					}
				}
			case lnull && !rnull: // null is not value
				vs := vec.Col.([]uint16)
				gv := gvec.Col.([]uint16)[g.Sel]
				for i, sel := range sels {
					if vec.Nsp.Contains(uint64(sel)) {
						diffs[i] = true
					} else {
						diffs[i] = diffs[i] || (gv != vs[sel])
					}
				}
			case !lnull && rnull: // null is not value
				for i := range sels {
					diffs[i] = true
				}
			default:
				vs := vec.Col.([]uint16)
				gv := gvec.Col.([]uint16)[g.Sel]
				for i, sel := range sels {
					diffs[i] = diffs[i] || (gv != vs[sel])
				}
			}
		case types.T_uint32:
			gvec := bats[g.Idx].Vecs[i]
			lnull := vec.Nsp.Any()
			rnull := gvec.Nsp.Contains(uint64(g.Sel))
			switch {
			case lnull && rnull:
				for i, sel := range sels {
					if !vec.Nsp.Contains(uint64(sel)) { // only null eq null
						diffs[i] = true
					}
				}
			case lnull && !rnull: // null is not value
				vs := vec.Col.([]uint32)
				gv := gvec.Col.([]uint32)[g.Sel]
				for i, sel := range sels {
					if vec.Nsp.Contains(uint64(sel)) {
						diffs[i] = true
					} else {
						diffs[i] = diffs[i] || (gv != vs[sel])
					}
				}
			case !lnull && rnull: // null is not value
				for i := range sels {
					diffs[i] = true
				}
			default:
				vs := vec.Col.([]uint32)
				gv := gvec.Col.([]uint32)[g.Sel]
				for i, sel := range sels {
					diffs[i] = diffs[i] || (gv != vs[sel])
				}
			}
		case types.T_uint64:
			gvec := bats[g.Idx].Vecs[i]
			lnull := vec.Nsp.Any()
			rnull := gvec.Nsp.Contains(uint64(g.Sel))
			switch {
			case lnull && rnull:
				for i, sel := range sels {
					if !vec.Nsp.Contains(uint64(sel)) { // only null eq null
						diffs[i] = true
					}
				}
			case lnull && !rnull: // null is not value
				vs := vec.Col.([]uint64)
				gv := gvec.Col.([]uint64)[g.Sel]
				for i, sel := range sels {
					if vec.Nsp.Contains(uint64(sel)) {
						diffs[i] = true
					} else {
						diffs[i] = diffs[i] || (gv != vs[sel])
					}
				}
			case !lnull && rnull: // null is not value
				for i := range sels {
					diffs[i] = true
				}
			default:
				vs := vec.Col.([]uint64)
				gv := gvec.Col.([]uint64)[g.Sel]
				for i, sel := range sels {
					diffs[i] = diffs[i] || (gv != vs[sel])
				}
			}
		case types.T_float32:
			gvec := bats[g.Idx].Vecs[i]
			lnull := vec.Nsp.Any()
			rnull := gvec.Nsp.Contains(uint64(g.Sel))
			switch {
			case lnull && rnull:
				for i, sel := range sels {
					if !vec.Nsp.Contains(uint64(sel)) { // only null eq null
						diffs[i] = true
					}
				}
			case lnull && !rnull: // null is not value
				vs := vec.Col.([]float32)
				gv := gvec.Col.([]float32)[g.Sel]
				for i, sel := range sels {
					if vec.Nsp.Contains(uint64(sel)) {
						diffs[i] = true
					} else {
						diffs[i] = diffs[i] || (gv != vs[sel])
					}
				}
			case !lnull && rnull: // null is not value
				for i := range sels {
					diffs[i] = true
				}
			default:
				vs := vec.Col.([]float32)
				gv := gvec.Col.([]float32)[g.Sel]
				for i, sel := range sels {
					diffs[i] = diffs[i] || (gv != vs[sel])
				}
			}
		case types.T_float64:
			gvec := bats[g.Idx].Vecs[i]
			lnull := vec.Nsp.Any()
			rnull := gvec.Nsp.Contains(uint64(g.Sel))
			switch {
			case lnull && rnull:
				for i, sel := range sels {
					if !vec.Nsp.Contains(uint64(sel)) { // only null eq null
						diffs[i] = true
					}
				}
			case lnull && !rnull: // null is not value
				vs := vec.Col.([]float64)
				gv := gvec.Col.([]float64)[g.Sel]
				for i, sel := range sels {
					if vec.Nsp.Contains(uint64(sel)) {
						diffs[i] = true
					} else {
						diffs[i] = diffs[i] || (gv != vs[sel])
					}
				}
			case !lnull && rnull: // null is not value
				for i := range sels {
					diffs[i] = true
				}
			default:
				vs := vec.Col.([]float64)
				gv := gvec.Col.([]float64)[g.Sel]
				for i, sel := range sels {
					diffs[i] = diffs[i] || (gv != vs[sel])
				}
			}
		case types.T_decimal:
		case types.T_date:
		case types.T_datetime:
		case types.T_char:
		case types.T_varchar:
			gvec := bats[g.Idx].Vecs[i]
			lnull := vec.Nsp.Any()
			rnull := gvec.Nsp.Contains(uint64(g.Sel))
			switch {
			case lnull && rnull:
				for i, sel := range sels {
					if !vec.Nsp.Contains(uint64(sel)) { // only null eq null
						diffs[i] = true
					}
				}
			case lnull && !rnull: // null is not value
				vs := vec.Col.(*types.Bytes)
				gvs := gvec.Col.(*types.Bytes)
				gv := gvs.Get(int(g.Sel))
				for i, sel := range sels {
					if vec.Nsp.Contains(uint64(sel)) {
						diffs[i] = true
					} else {
						diffs[i] = diffs[i] || (bytes.Compare(gv, vs.Get(int(sel))) != 0)
					}
				}
			case !lnull && rnull: // null is not value
				for i := range sels {
					diffs[i] = true
				}
			default:
				vs := vec.Col.(*types.Bytes)
				gvs := gvec.Col.(*types.Bytes)
				gv := gvs.Get(int(g.Sel))
				for i, sel := range sels {
					diffs[i] = diffs[i] || (bytes.Compare(gv, vs.Get(int(sel))) != 0)
				}
			}
		case types.T_json:
			gvec := bats[g.Idx].Vecs[i]
			lnull := vec.Nsp.Any()
			rnull := gvec.Nsp.Contains(uint64(g.Sel))
			switch {
			case lnull && rnull:
				for i, sel := range sels {
					if !vec.Nsp.Contains(uint64(sel)) { // only null eq null
						diffs[i] = true
					}
				}
			case lnull && !rnull: // null is not value
				vs := vec.Col.(*types.Bytes)
				gvs := gvec.Col.(*types.Bytes)
				gv := gvs.Get(int(g.Sel))
				for i, sel := range sels {
					if vec.Nsp.Contains(uint64(sel)) {
						diffs[i] = true
					} else {
						diffs[i] = diffs[i] || (bytes.Compare(gv, vs.Get(int(sel))) != 0)
					}
				}
			case !lnull && rnull: // null is not value
				for i := range sels {
					diffs[i] = true
				}
			default:
				vs := vec.Col.(*types.Bytes)
				gvs := gvec.Col.(*types.Bytes)
				gv := gvs.Get(int(g.Sel))
				for i, sel := range sels {
					diffs[i] = diffs[i] || (bytes.Compare(gv, vs.Get(int(sel))) != 0)
				}
			}
		}
	}
	n := len(sels)
	remaining := sels[:0]
	for i := 0; i < n; i++ {
		if diffs[i] {
			remaining = append(remaining, sels[i])
		}
	}
	return remaining, nil
}