package bitmap


// 暂时只支持 1 << 32 位（可以扩展到 1 << 64)
// The Max Size is 0x01 << 32 at present(can expand to 0x01 << 64)
const BitmapSize = 0x01 << 32

// Bitmap 数据结构定义
type Bitmap struct {
	// 保存实际的 bit 数据
	Data []byte
	// 指示该 Bitmap 的 bit 容量
	Bitsize uint64
	// 该 Bitmap 被设置为 1 的最大位置（方便遍历）
	Maxpos uint64
}

// NewBitmap 使用默认容量实例化一个 Bitmap
func NewBitmap() *Bitmap {
	return NewBitmapSize(BitmapSize)
}

// NewBitmapSize 根据指定的 size 实例化一个 Bitmap
func NewBitmapSize(size uint64) *Bitmap {
	if size == 0 || size > BitmapSize {
		size = BitmapSize
	} else if remainder := size % 8; remainder != 0 {
		size += 8 - remainder
	}

	return &Bitmap{Data: make([]byte, size>>3), Bitsize: uint64(size - 1)}
}

// SetBit 将 offset 位置的 bit 置为 value (0/1)
func (this *Bitmap) SetBit(offset uint64, value uint8) bool {
	index, pos := offset/8, offset%8

	if this.Bitsize < offset {
		return false
	}

	if value == 0 {
		// &^ 清位
		this.Data[index] &^= 0x01 << pos
	} else {
		this.Data[index] |= 0x01 << pos

		// 记录曾经设置为 1 的最大位置
		if this.Maxpos < offset {
			this.Maxpos = offset
		}
	}

	return true
}

// GetBit 获得 offset 位置处的 value
func (this *Bitmap) GetBit(offset uint64) uint8 {
	index, pos := offset/8, offset%8

	if this.Bitsize < offset {
		return 0
	}

	return (this.Data[index] >> pos) & 0x01
}

// Maxpos 获的置为 1 的最大位置
func (this *Bitmap) Max() uint64 {
	return this.Maxpos
}
