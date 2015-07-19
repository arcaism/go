// autogenerated: do not edit!
// generated from gen/*Ops.go
package ssa

import "cmd/internal/obj/x86"

const (
	blockInvalid BlockKind = iota

	BlockAMD64EQ
	BlockAMD64NE
	BlockAMD64LT
	BlockAMD64LE
	BlockAMD64GT
	BlockAMD64GE
	BlockAMD64ULT
	BlockAMD64ULE
	BlockAMD64UGT
	BlockAMD64UGE

	BlockExit
	BlockDead
	BlockPlain
	BlockIf
	BlockCall
)

var blockString = [...]string{
	blockInvalid: "BlockInvalid",

	BlockAMD64EQ:  "EQ",
	BlockAMD64NE:  "NE",
	BlockAMD64LT:  "LT",
	BlockAMD64LE:  "LE",
	BlockAMD64GT:  "GT",
	BlockAMD64GE:  "GE",
	BlockAMD64ULT: "ULT",
	BlockAMD64ULE: "ULE",
	BlockAMD64UGT: "UGT",
	BlockAMD64UGE: "UGE",

	BlockExit:  "Exit",
	BlockDead:  "Dead",
	BlockPlain: "Plain",
	BlockIf:    "If",
	BlockCall:  "Call",
}

func (k BlockKind) String() string { return blockString[k] }

const (
	OpInvalid Op = iota

	OpAMD64ADDQ
	OpAMD64ADDQconst
	OpAMD64SUBQ
	OpAMD64SUBQconst
	OpAMD64MULQ
	OpAMD64MULQconst
	OpAMD64ANDQ
	OpAMD64ANDQconst
	OpAMD64SHLQ
	OpAMD64SHLQconst
	OpAMD64SHRQ
	OpAMD64SHRQconst
	OpAMD64SARQ
	OpAMD64SARQconst
	OpAMD64NEGQ
	OpAMD64XORQconst
	OpAMD64CMPQ
	OpAMD64CMPQconst
	OpAMD64TESTQ
	OpAMD64TESTB
	OpAMD64SBBQcarrymask
	OpAMD64SETEQ
	OpAMD64SETNE
	OpAMD64SETL
	OpAMD64SETLE
	OpAMD64SETG
	OpAMD64SETGE
	OpAMD64SETB
	OpAMD64CMOVQCC
	OpAMD64MOVLQSX
	OpAMD64MOVWQSX
	OpAMD64MOVBQSX
	OpAMD64MOVQconst
	OpAMD64LEAQ
	OpAMD64LEAQ1
	OpAMD64LEAQ2
	OpAMD64LEAQ4
	OpAMD64LEAQ8
	OpAMD64MOVBload
	OpAMD64MOVBQZXload
	OpAMD64MOVBQSXload
	OpAMD64MOVWload
	OpAMD64MOVLload
	OpAMD64MOVQload
	OpAMD64MOVQloadidx8
	OpAMD64MOVBstore
	OpAMD64MOVWstore
	OpAMD64MOVLstore
	OpAMD64MOVQstore
	OpAMD64MOVQstoreidx8
	OpAMD64MOVXzero
	OpAMD64REPSTOSQ
	OpAMD64MOVQloadglobal
	OpAMD64MOVQstoreglobal
	OpAMD64CALLstatic
	OpAMD64CALLclosure
	OpAMD64REPMOVSB
	OpAMD64ADDL
	OpAMD64ADDW
	OpAMD64ADDB
	OpAMD64SUBL
	OpAMD64SUBW
	OpAMD64SUBB
	OpAMD64InvertFlags

	OpAdd8
	OpAdd16
	OpAdd32
	OpAdd64
	OpAdd8U
	OpAdd16U
	OpAdd32U
	OpAdd64U
	OpAddPtr
	OpSub8
	OpSub16
	OpSub32
	OpSub64
	OpSub8U
	OpSub16U
	OpSub32U
	OpSub64U
	OpMul
	OpLsh8
	OpLsh16
	OpLsh32
	OpLsh64
	OpRsh8
	OpRsh8U
	OpRsh16
	OpRsh16U
	OpRsh32
	OpRsh32U
	OpRsh64
	OpRsh64U
	OpEq8
	OpEq16
	OpEq32
	OpEq64
	OpNeq8
	OpNeq16
	OpNeq32
	OpNeq64
	OpLess8
	OpLess8U
	OpLess16
	OpLess16U
	OpLess32
	OpLess32U
	OpLess64
	OpLess64U
	OpLeq8
	OpLeq8U
	OpLeq16
	OpLeq16U
	OpLeq32
	OpLeq32U
	OpLeq64
	OpLeq64U
	OpGreater8
	OpGreater8U
	OpGreater16
	OpGreater16U
	OpGreater32
	OpGreater32U
	OpGreater64
	OpGreater64U
	OpGeq8
	OpGeq8U
	OpGeq16
	OpGeq16U
	OpGeq32
	OpGeq32U
	OpGeq64
	OpGeq64U
	OpNot
	OpPhi
	OpCopy
	OpConst
	OpArg
	OpAddr
	OpSP
	OpSB
	OpFunc
	OpLoad
	OpStore
	OpMove
	OpZero
	OpClosureCall
	OpStaticCall
	OpConvert
	OpConvNop
	OpIsNonNil
	OpIsInBounds
	OpArrayIndex
	OpPtrIndex
	OpOffPtr
	OpStructSelect
	OpSliceMake
	OpSlicePtr
	OpSliceLen
	OpSliceCap
	OpStringMake
	OpStringPtr
	OpStringLen
	OpStoreReg8
	OpLoadReg8
	OpFwdRef
)

var opcodeTable = [...]opInfo{
	{name: "OpInvalid"},

	{
		name: "ADDQ",
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			clobbers: 0,
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "ADDQconst",
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			clobbers: 0,
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "SUBQ",
		asm:  x86.ASUBQ,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			clobbers: 0,
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "SUBQconst",
		asm:  x86.ASUBQ,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			clobbers: 0,
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "MULQ",
		asm:  x86.AIMULQ,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			clobbers: 0,
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "MULQconst",
		asm:  x86.AIMULQ,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			clobbers: 0,
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "ANDQ",
		asm:  x86.AANDQ,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			clobbers: 0,
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "ANDQconst",
		asm:  x86.AANDQ,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			clobbers: 0,
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "SHLQ",
		asm:  x86.ASHLQ,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				2,     // .CX
			},
			clobbers: 0,
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "SHLQconst",
		asm:  x86.ASHLQ,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			clobbers: 0,
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "SHRQ",
		asm:  x86.ASHRQ,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				2,     // .CX
			},
			clobbers: 0,
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "SHRQconst",
		asm:  x86.ASHRQ,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			clobbers: 0,
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "SARQ",
		asm:  x86.ASARQ,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				2,     // .CX
			},
			clobbers: 0,
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "SARQconst",
		asm:  x86.ASARQ,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			clobbers: 0,
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "NEGQ",
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			clobbers: 0,
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "XORQconst",
		asm:  x86.AXORQ,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			clobbers: 0,
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "CMPQ",
		asm:  x86.ACMPQ,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			clobbers: 0,
			outputs: []regMask{
				8589934592, // .FLAGS
			},
		},
	},
	{
		name: "CMPQconst",
		asm:  x86.ACMPQ,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			clobbers: 0,
			outputs: []regMask{
				8589934592, // .FLAGS
			},
		},
	},
	{
		name: "TESTQ",
		asm:  x86.ATESTQ,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			clobbers: 0,
			outputs: []regMask{
				8589934592, // .FLAGS
			},
		},
	},
	{
		name: "TESTB",
		asm:  x86.ATESTB,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			clobbers: 0,
			outputs: []regMask{
				8589934592, // .FLAGS
			},
		},
	},
	{
		name: "SBBQcarrymask",
		asm:  x86.ASBBQ,
		reg: regInfo{
			inputs: []regMask{
				8589934592, // .FLAGS
			},
			clobbers: 0,
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "SETEQ",
		reg: regInfo{
			inputs: []regMask{
				8589934592, // .FLAGS
			},
			clobbers: 0,
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "SETNE",
		reg: regInfo{
			inputs: []regMask{
				8589934592, // .FLAGS
			},
			clobbers: 0,
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "SETL",
		reg: regInfo{
			inputs: []regMask{
				8589934592, // .FLAGS
			},
			clobbers: 0,
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "SETLE",
		reg: regInfo{
			inputs: []regMask{
				8589934592, // .FLAGS
			},
			clobbers: 0,
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "SETG",
		reg: regInfo{
			inputs: []regMask{
				8589934592, // .FLAGS
			},
			clobbers: 0,
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "SETGE",
		reg: regInfo{
			inputs: []regMask{
				8589934592, // .FLAGS
			},
			clobbers: 0,
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "SETB",
		reg: regInfo{
			inputs: []regMask{
				8589934592, // .FLAGS
			},
			clobbers: 0,
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "CMOVQCC",
		reg: regInfo{
			inputs: []regMask{
				8589934592, // .FLAGS
				65519,      // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				65519,      // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			clobbers: 0,
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "MOVLQSX",
		asm:  x86.AMOVLQSX,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			clobbers: 0,
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "MOVWQSX",
		asm:  x86.AMOVWQSX,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			clobbers: 0,
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "MOVBQSX",
		asm:  x86.AMOVBQSX,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			clobbers: 0,
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "MOVQconst",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "LEAQ",
		reg: regInfo{
			inputs: []regMask{
				4295032831, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15 .SB
			},
			clobbers: 0,
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "LEAQ1",
		reg: regInfo{
			inputs: []regMask{
				4295032831, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15 .SB
				65535,      // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			clobbers: 0,
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "LEAQ2",
		reg: regInfo{
			inputs: []regMask{
				4295032831, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15 .SB
				65535,      // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			clobbers: 0,
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "LEAQ4",
		reg: regInfo{
			inputs: []regMask{
				4295032831, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15 .SB
				65535,      // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			clobbers: 0,
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "LEAQ8",
		reg: regInfo{
			inputs: []regMask{
				4295032831, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15 .SB
				65535,      // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			clobbers: 0,
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "MOVBload",
		asm:  x86.AMOVB,
		reg: regInfo{
			inputs: []regMask{
				4295032831, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15 .SB
				0,
			},
			clobbers: 0,
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "MOVBQZXload",
		reg: regInfo{
			inputs: []regMask{
				4295032831, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15 .SB
				0,
			},
			clobbers: 0,
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "MOVBQSXload",
		reg: regInfo{
			inputs: []regMask{
				4295032831, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15 .SB
				0,
			},
			clobbers: 0,
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "MOVWload",
		asm:  x86.AMOVW,
		reg: regInfo{
			inputs: []regMask{
				4295032831, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15 .SB
				0,
			},
			clobbers: 0,
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "MOVLload",
		asm:  x86.AMOVL,
		reg: regInfo{
			inputs: []regMask{
				4295032831, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15 .SB
				0,
			},
			clobbers: 0,
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "MOVQload",
		asm:  x86.AMOVQ,
		reg: regInfo{
			inputs: []regMask{
				4295032831, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15 .SB
				0,
			},
			clobbers: 0,
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "MOVQloadidx8",
		asm:  x86.AMOVQ,
		reg: regInfo{
			inputs: []regMask{
				4295032831, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15 .SB
				65535,      // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				0,
			},
			clobbers: 0,
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "MOVBstore",
		asm:  x86.AMOVB,
		reg: regInfo{
			inputs: []regMask{
				4295032831, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15 .SB
				65535,      // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				0,
			},
			clobbers: 0,
			outputs:  []regMask{},
		},
	},
	{
		name: "MOVWstore",
		asm:  x86.AMOVW,
		reg: regInfo{
			inputs: []regMask{
				4295032831, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15 .SB
				65535,      // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				0,
			},
			clobbers: 0,
			outputs:  []regMask{},
		},
	},
	{
		name: "MOVLstore",
		asm:  x86.AMOVL,
		reg: regInfo{
			inputs: []regMask{
				4295032831, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15 .SB
				65535,      // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				0,
			},
			clobbers: 0,
			outputs:  []regMask{},
		},
	},
	{
		name: "MOVQstore",
		asm:  x86.AMOVQ,
		reg: regInfo{
			inputs: []regMask{
				4295032831, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15 .SB
				65535,      // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				0,
			},
			clobbers: 0,
			outputs:  []regMask{},
		},
	},
	{
		name: "MOVQstoreidx8",
		reg: regInfo{
			inputs: []regMask{
				4295032831, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15 .SB
				65535,      // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				65535,      // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				0,
			},
			clobbers: 0,
			outputs:  []regMask{},
		},
	},
	{
		name: "MOVXzero",
		reg: regInfo{
			inputs: []regMask{
				4295032831, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15 .SB
				0,
			},
			clobbers: 0,
			outputs:  []regMask{},
		},
	},
	{
		name: "REPSTOSQ",
		reg: regInfo{
			inputs: []regMask{
				128, // .DI
				2,   // .CX
			},
			clobbers: 131, // .AX .CX .DI
			outputs:  []regMask{},
		},
	},
	{
		name: "MOVQloadglobal",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
	},
	{
		name: "MOVQstoreglobal",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
	},
	{
		name: "CALLstatic",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
	},
	{
		name: "CALLclosure",
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				4,     // .DX
				0,
			},
			clobbers: 0,
			outputs:  []regMask{},
		},
	},
	{
		name: "REPMOVSB",
		reg: regInfo{
			inputs: []regMask{
				128, // .DI
				64,  // .SI
				2,   // .CX
			},
			clobbers: 194, // .CX .SI .DI
			outputs:  []regMask{},
		},
	},
	{
		name: "ADDL",
		asm:  x86.AADDL,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			clobbers: 0,
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "ADDW",
		asm:  x86.AADDW,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			clobbers: 0,
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "ADDB",
		asm:  x86.AADDB,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			clobbers: 0,
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "SUBL",
		asm:  x86.ASUBL,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			clobbers: 0,
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "SUBW",
		asm:  x86.ASUBW,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			clobbers: 0,
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "SUBB",
		asm:  x86.ASUBB,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			clobbers: 0,
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "InvertFlags",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
	},

	{
		name: "Add8",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Add16",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Add32",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Add64",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Add8U",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Add16U",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Add32U",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Add64U",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "AddPtr",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Sub8",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Sub16",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Sub32",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Sub64",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Sub8U",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Sub16U",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Sub32U",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Sub64U",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Mul",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Lsh8",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Lsh16",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Lsh32",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Lsh64",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Rsh8",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Rsh8U",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Rsh16",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Rsh16U",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Rsh32",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Rsh32U",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Rsh64",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Rsh64U",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Eq8",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Eq16",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Eq32",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Eq64",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Neq8",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Neq16",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Neq32",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Neq64",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Less8",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Less8U",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Less16",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Less16U",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Less32",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Less32U",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Less64",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Less64U",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Leq8",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Leq8U",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Leq16",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Leq16U",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Leq32",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Leq32U",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Leq64",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Leq64U",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Greater8",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Greater8U",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Greater16",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Greater16U",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Greater32",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Greater32U",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Greater64",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Greater64U",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Geq8",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Geq8U",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Geq16",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Geq16U",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Geq32",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Geq32U",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Geq64",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Geq64U",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Not",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Phi",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Copy",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Const",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Arg",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Addr",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "SP",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "SB",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Func",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Load",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Store",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Move",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Zero",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "ClosureCall",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "StaticCall",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "Convert",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "ConvNop",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "IsNonNil",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "IsInBounds",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "ArrayIndex",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "PtrIndex",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "OffPtr",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "StructSelect",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "SliceMake",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "SlicePtr",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "SliceLen",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "SliceCap",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "StringMake",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "StringPtr",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "StringLen",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "StoreReg8",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "LoadReg8",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
	{
		name: "FwdRef",
		reg: regInfo{
			inputs:   []regMask{},
			clobbers: 0,
			outputs:  []regMask{},
		},
		generic: true,
	},
}

func (o Op) Asm() int       { return opcodeTable[o].asm }
func (o Op) String() string { return opcodeTable[o].name }