package col

import (
	"errors"
	"fmt"
	"io"
	bmgrif "matrixone/pkg/vm/engine/aoe/storage/buffer/manager/iface"
	nif "matrixone/pkg/vm/engine/aoe/storage/buffer/node/iface"
	"matrixone/pkg/vm/engine/aoe/storage/common"
	ldio "matrixone/pkg/vm/engine/aoe/storage/layout/dataio"
	"sync"
	// log "github.com/sirupsen/logrus"
)

func initUnsortedBlkNode(part *ColumnPart, fsMgr ldio.IManager) {
	sf := fsMgr.GetUnsortedFile(part.ID.AsSegmentID())
	sf.RefBlock(part.ID.AsBlockID())
	csf := sf.MakeColSegmentFile(part.ColIdx)
	part.BufNode = part.BufMgr.RegisterNode(part.Capacity, part.NodeID, csf)
	part.SegFile = sf
}

func initSortedBlkNode(part *ColumnPart, fsMgr ldio.IManager) {
	sf := fsMgr.GetSortedFile(part.ID.AsSegmentID())
	sf.RefBlock(part.ID.AsBlockID())
	csf := sf.MakeColSegmentFile(part.ColIdx)
	part.BufNode = part.BufMgr.RegisterNode(part.Capacity, part.NodeID, csf)
	part.SegFile = sf
}

type IColumnPart interface {
	io.Closer
	GetNext() IColumnPart
	SetNext(IColumnPart)
	InitScanCursor(cursor *ScanCursor) error
	GetID() common.ID
	// GetBlock() IColumnBlock
	GetBuf() []byte
	GetColIdx() int
	CloneWithUpgrade(IColumnBlock, bmgrif.IBufferManager, ldio.IManager) IColumnPart
	GetNodeID() common.ID
}

type ColumnPart struct {
	sync.RWMutex
	ID          common.ID
	Next        IColumnPart
	BufMgr      bmgrif.IBufferManager
	BufNode     nif.INodeHandle
	TypeSize    uint64
	MaxRowCount uint64
	RowCount    uint64
	Size        uint64
	Capacity    uint64
	NodeID      common.ID
	ColIdx      int
	SegFile     ldio.ISegmentFile
}

func NewColumnPart(fsMgr ldio.IManager, bmgr bmgrif.IBufferManager, blk IColumnBlock, id common.ID,
	rowCount uint64, typeSize uint64) IColumnPart {
	defer blk.UnRef()
	part := &ColumnPart{
		BufMgr:      bmgr,
		ID:          id,
		TypeSize:    typeSize,
		MaxRowCount: rowCount,
		NodeID:      id,
		ColIdx:      blk.GetColIdx(),
		Capacity:    typeSize * rowCount,
	}
	part.NodeID.Idx = uint16(blk.GetColIdx())

	switch blk.GetBlockType() {
	case TRANSIENT_BLK:
		bNode := bmgr.RegisterSpillableNode(typeSize*rowCount, part.NodeID)
		if bNode == nil {
			return nil
		}
		part.BufNode = bNode
	case PERSISTENT_BLK:
		initUnsortedBlkNode(part, fsMgr)
	case PERSISTENT_SORTED_BLK:
		initSortedBlkNode(part, fsMgr)
	default:
		panic("not support")
	}

	blk.Append(part)
	return part
}

func (part *ColumnPart) GetNodeID() common.ID {
	return part.NodeID
}

func (part *ColumnPart) CloneWithUpgrade(blk IColumnBlock, sstBufMgr bmgrif.IBufferManager, fsMgr ldio.IManager) IColumnPart {
	cloned := &ColumnPart{
		ID:          part.ID,
		BufMgr:      sstBufMgr,
		TypeSize:    part.TypeSize,
		MaxRowCount: part.MaxRowCount,
		RowCount:    part.RowCount,
		Size:        part.Size,
		Capacity:    part.Capacity,
		NodeID:      part.NodeID.NextIter(),
		ColIdx:      part.GetColIdx(),
	}
	switch blk.GetBlockType() {
	case TRANSIENT_BLK:
		panic("logic error")
	case PERSISTENT_BLK:
		initUnsortedBlkNode(cloned, fsMgr)
	case PERSISTENT_SORTED_BLK:
		initSortedBlkNode(cloned, fsMgr)
	default:
		panic("not supported")
	}

	blk.UnRef()
	return cloned
}

func (part *ColumnPart) GetColIdx() int {
	return part.ColIdx
}

func (part *ColumnPart) GetBuf() []byte {
	return part.BufNode.GetBuffer().GetDataNode().Data
}

func (part *ColumnPart) SetRowCount(cnt uint64) {
	if cnt > part.MaxRowCount {
		panic("logic error")
	}
	part.Lock()
	defer part.Unlock()
	part.RowCount = cnt
}

func (part *ColumnPart) SetSize(size uint64) {
	if size > part.Capacity {
		panic("logic error")
	}
	part.Lock()
	defer part.Unlock()
	part.Size = size
}

func (part *ColumnPart) GetID() common.ID {
	return part.ID
}

func (part *ColumnPart) SetNext(next IColumnPart) {
	part.Lock()
	defer part.Unlock()
	part.Next = next
}

func (part *ColumnPart) GetNext() IColumnPart {
	part.RLock()
	defer part.RUnlock()
	return part.Next
}

func (part *ColumnPart) Close() error {
	if part.BufNode != nil {
		err := part.BufNode.Close()
		if err != nil {
			panic("logic error")
		}
		part.BufNode = nil
	}
	if part.SegFile != nil {
		part.SegFile.UnrefBlock(part.ID.AsSegmentID())
	}
	return nil
}

func (part *ColumnPart) InitScanCursor(cursor *ScanCursor) error {
	cursor.Handle = part.BufMgr.Pin(part.BufNode)
	if cursor.Handle == nil {
		return errors.New(fmt.Sprintf("Cannot pin part %v", part.ID))
	}
	// cursor.Inited = true
	return nil
}