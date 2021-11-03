// Copyright 2021 Matrix Origin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package min

import (
	"github.com/matrixorigin/matrixone/pkg/container/types"
	"github.com/matrixorigin/matrixone/pkg/container/vector"
	"github.com/matrixorigin/matrixone/pkg/encoding"
	"github.com/matrixorigin/matrixone/pkg/sql/colexec/aggregation"
	"github.com/matrixorigin/matrixone/pkg/vectorize/min"
	"github.com/matrixorigin/matrixone/pkg/vm/process"
)

func NewInt8(typ types.Type) *int8Min {
	return &int8Min{typ: typ}
}

func (a *int8Min) Reset() {
	a.v = 0
	a.cnt = 0
}

func (a *int8Min) Type() types.Type {
	return a.typ
}

func (a *int8Min) Dup() aggregation.Aggregation {
	return &int8Min{typ: a.typ}
}

func (a *int8Min) Fill(sels []int64, vec *vector.Vector) error {
	if n := len(sels); n > 0 {
		v := min.Int8MinSels(vec.Col.([]int8), sels)
		if a.cnt == 0 || v < a.v {
			a.v = v
		}
		a.cnt += int64(n - vec.Nsp.FilterCount(sels))
	} else {
		v := min.Int8Min(vec.Col.([]int8))
		if a.cnt == 0 || v < a.v {
			a.v = v
		}
		a.cnt += int64(vec.Length() - vec.Nsp.Length())
	}
	return nil
}

func (a *int8Min) Eval() interface{} {
	if a.cnt == 0 {
		return nil
	}
	return a.v
}

func (a *int8Min) EvalCopy(proc *process.Process) (*vector.Vector, error) {
	data, err := proc.Alloc(1)
	if err != nil {
		return nil, err
	}
	vec := vector.New(a.typ)
	vs := encoding.DecodeInt8Slice(data[:1])
	vs[0] = a.v
	if a.cnt == 0 {
		vec.Nsp.Add(0)
	}
	vec.Col = vs
	vec.Data = data
	return vec, nil
}