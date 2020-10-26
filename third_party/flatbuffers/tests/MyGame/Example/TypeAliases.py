# automatically generated by the FlatBuffers compiler, do not modify

# namespace: Example

import flatbuffers

class TypeAliases(object):
    __slots__ = ['_tab']

    @classmethod
    def GetRootAsTypeAliases(cls, buf, offset):
        n = flatbuffers.encode.Get(flatbuffers.packer.uoffset, buf, offset)
        x = TypeAliases()
        x.Init(buf, n + offset)
        return x

    @classmethod
    def TypeAliasesBufferHasIdentifier(cls, buf, offset, size_prefixed=False):
        return flatbuffers.util.BufferHasIdentifier(buf, offset, b"\x4D\x4F\x4E\x53", size_prefixed=size_prefixed)

    # TypeAliases
    def Init(self, buf, pos):
        self._tab = flatbuffers.table.Table(buf, pos)

    # TypeAliases
    def I8(self):
        o = flatbuffers.number_types.UOffsetTFlags.py_type(self._tab.Offset(4))
        if o != 0:
            return self._tab.Get(flatbuffers.number_types.Int8Flags, o + self._tab.Pos)
        return 0

    # TypeAliases
    def U8(self):
        o = flatbuffers.number_types.UOffsetTFlags.py_type(self._tab.Offset(6))
        if o != 0:
            return self._tab.Get(flatbuffers.number_types.Uint8Flags, o + self._tab.Pos)
        return 0

    # TypeAliases
    def I16(self):
        o = flatbuffers.number_types.UOffsetTFlags.py_type(self._tab.Offset(8))
        if o != 0:
            return self._tab.Get(flatbuffers.number_types.Int16Flags, o + self._tab.Pos)
        return 0

    # TypeAliases
    def U16(self):
        o = flatbuffers.number_types.UOffsetTFlags.py_type(self._tab.Offset(10))
        if o != 0:
            return self._tab.Get(flatbuffers.number_types.Uint16Flags, o + self._tab.Pos)
        return 0

    # TypeAliases
    def I32(self):
        o = flatbuffers.number_types.UOffsetTFlags.py_type(self._tab.Offset(12))
        if o != 0:
            return self._tab.Get(flatbuffers.number_types.Int32Flags, o + self._tab.Pos)
        return 0

    # TypeAliases
    def U32(self):
        o = flatbuffers.number_types.UOffsetTFlags.py_type(self._tab.Offset(14))
        if o != 0:
            return self._tab.Get(flatbuffers.number_types.Uint32Flags, o + self._tab.Pos)
        return 0

    # TypeAliases
    def I64(self):
        o = flatbuffers.number_types.UOffsetTFlags.py_type(self._tab.Offset(16))
        if o != 0:
            return self._tab.Get(flatbuffers.number_types.Int64Flags, o + self._tab.Pos)
        return 0

    # TypeAliases
    def U64(self):
        o = flatbuffers.number_types.UOffsetTFlags.py_type(self._tab.Offset(18))
        if o != 0:
            return self._tab.Get(flatbuffers.number_types.Uint64Flags, o + self._tab.Pos)
        return 0

    # TypeAliases
    def F32(self):
        o = flatbuffers.number_types.UOffsetTFlags.py_type(self._tab.Offset(20))
        if o != 0:
            return self._tab.Get(flatbuffers.number_types.Float32Flags, o + self._tab.Pos)
        return 0.0

    # TypeAliases
    def F64(self):
        o = flatbuffers.number_types.UOffsetTFlags.py_type(self._tab.Offset(22))
        if o != 0:
            return self._tab.Get(flatbuffers.number_types.Float64Flags, o + self._tab.Pos)
        return 0.0

    # TypeAliases
    def V8(self, j):
        o = flatbuffers.number_types.UOffsetTFlags.py_type(self._tab.Offset(24))
        if o != 0:
            a = self._tab.Vector(o)
            return self._tab.Get(flatbuffers.number_types.Int8Flags, a + flatbuffers.number_types.UOffsetTFlags.py_type(j * 1))
        return 0

    # TypeAliases
    def V8AsNumpy(self):
        o = flatbuffers.number_types.UOffsetTFlags.py_type(self._tab.Offset(24))
        if o != 0:
            return self._tab.GetVectorAsNumpy(flatbuffers.number_types.Int8Flags, o)
        return 0

    # TypeAliases
    def V8Length(self):
        o = flatbuffers.number_types.UOffsetTFlags.py_type(self._tab.Offset(24))
        if o != 0:
            return self._tab.VectorLen(o)
        return 0

    # TypeAliases
    def Vf64(self, j):
        o = flatbuffers.number_types.UOffsetTFlags.py_type(self._tab.Offset(26))
        if o != 0:
            a = self._tab.Vector(o)
            return self._tab.Get(flatbuffers.number_types.Float64Flags, a + flatbuffers.number_types.UOffsetTFlags.py_type(j * 8))
        return 0

    # TypeAliases
    def Vf64AsNumpy(self):
        o = flatbuffers.number_types.UOffsetTFlags.py_type(self._tab.Offset(26))
        if o != 0:
            return self._tab.GetVectorAsNumpy(flatbuffers.number_types.Float64Flags, o)
        return 0

    # TypeAliases
    def Vf64Length(self):
        o = flatbuffers.number_types.UOffsetTFlags.py_type(self._tab.Offset(26))
        if o != 0:
            return self._tab.VectorLen(o)
        return 0

def TypeAliasesStart(builder): builder.StartObject(12)
def TypeAliasesAddI8(builder, i8): builder.PrependInt8Slot(0, i8, 0)
def TypeAliasesAddU8(builder, u8): builder.PrependUint8Slot(1, u8, 0)
def TypeAliasesAddI16(builder, i16): builder.PrependInt16Slot(2, i16, 0)
def TypeAliasesAddU16(builder, u16): builder.PrependUint16Slot(3, u16, 0)
def TypeAliasesAddI32(builder, i32): builder.PrependInt32Slot(4, i32, 0)
def TypeAliasesAddU32(builder, u32): builder.PrependUint32Slot(5, u32, 0)
def TypeAliasesAddI64(builder, i64): builder.PrependInt64Slot(6, i64, 0)
def TypeAliasesAddU64(builder, u64): builder.PrependUint64Slot(7, u64, 0)
def TypeAliasesAddF32(builder, f32): builder.PrependFloat32Slot(8, f32, 0.0)
def TypeAliasesAddF64(builder, f64): builder.PrependFloat64Slot(9, f64, 0.0)
def TypeAliasesAddV8(builder, v8): builder.PrependUOffsetTRelativeSlot(10, flatbuffers.number_types.UOffsetTFlags.py_type(v8), 0)
def TypeAliasesStartV8Vector(builder, numElems): return builder.StartVector(1, numElems, 1)
def TypeAliasesAddVf64(builder, vf64): builder.PrependUOffsetTRelativeSlot(11, flatbuffers.number_types.UOffsetTFlags.py_type(vf64), 0)
def TypeAliasesStartVf64Vector(builder, numElems): return builder.StartVector(8, numElems, 8)
def TypeAliasesEnd(builder): return builder.EndObject()
