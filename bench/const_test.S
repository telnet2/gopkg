<unlinkable>.GetHeader STEXT size=100 args=0x10 locals=0x28 funcid=0x0 align=0x0
	0x0000 00000 (./const_test.go:19)	TEXT	<unlinkable>.GetHeader(SB), ABIInternal, $40-16
	0x0000 00000 (./const_test.go:19)	CMPQ	SP, 16(R14)
	0x0004 00004 (./const_test.go:19)	PCDATA	$0, $-2
	0x0004 00004 (./const_test.go:19)	JLS	73
	0x0006 00006 (./const_test.go:19)	PCDATA	$0, $-1
	0x0006 00006 (./const_test.go:19)	SUBQ	$40, SP
	0x000a 00010 (./const_test.go:19)	MOVQ	BP, 32(SP)
	0x000f 00015 (./const_test.go:19)	LEAQ	32(SP), BP
	0x0014 00020 (./const_test.go:19)	MOVQ	AX, <unlinkable>.key+48(FP)
	0x0019 00025 (./const_test.go:19)	FUNCDATA	$0, gclocals·wgcWObbY2HYnK2SU/U22lA==(SB)
	0x0019 00025 (./const_test.go:19)	FUNCDATA	$1, gclocals·J5F+7Qw7O7ve2QcWC7DpeQ==(SB)
	0x0019 00025 (./const_test.go:19)	FUNCDATA	$5, <unlinkable>.GetHeader.arginfo1(SB)
	0x0019 00025 (./const_test.go:19)	FUNCDATA	$6, <unlinkable>.GetHeader.argliveinfo(SB)
	0x0019 00025 (./const_test.go:19)	PCDATA	$3, $1
	0x0019 00025 (./const_test.go:20)	MOVQ	<unlinkable>.headers(SB), DX
	0x0020 00032 (./const_test.go:20)	MOVQ	AX, CX
	0x0023 00035 (./const_test.go:20)	MOVQ	BX, DI
	0x0026 00038 (./const_test.go:20)	LEAQ	type.map[string]string(SB), AX
	0x002d 00045 (./const_test.go:20)	MOVQ	DX, BX
	0x0030 00048 (./const_test.go:20)	PCDATA	$1, $1
	0x0030 00048 (./const_test.go:20)	CALL	runtime.mapaccess1_faststr(SB)
	0x0035 00053 (./const_test.go:20)	MOVQ	(AX), DX
	0x0038 00056 (./const_test.go:20)	MOVQ	8(AX), BX
	0x003c 00060 (./const_test.go:20)	MOVQ	DX, AX
	0x003f 00063 (./const_test.go:20)	MOVQ	32(SP), BP
	0x0044 00068 (./const_test.go:20)	ADDQ	$40, SP
	0x0048 00072 (./const_test.go:20)	RET
	0x0049 00073 (./const_test.go:20)	NOP
	0x0049 00073 (./const_test.go:19)	PCDATA	$1, $-1
	0x0049 00073 (./const_test.go:19)	PCDATA	$0, $-2
	0x0049 00073 (./const_test.go:19)	MOVQ	AX, 8(SP)
	0x004e 00078 (./const_test.go:19)	MOVQ	BX, 16(SP)
	0x0053 00083 (./const_test.go:19)	CALL	runtime.morestack_noctxt(SB)
	0x0058 00088 (./const_test.go:19)	MOVQ	8(SP), AX
	0x005d 00093 (./const_test.go:19)	MOVQ	16(SP), BX
	0x0062 00098 (./const_test.go:19)	PCDATA	$0, $-1
	0x0062 00098 (./const_test.go:19)	JMP	0
	0x0000 49 3b 66 10 76 43 48 83 ec 28 48 89 6c 24 20 48  I;f.vCH..(H.l$ H
	0x0010 8d 6c 24 20 48 89 44 24 30 48 8b 15 00 00 00 00  .l$ H.D$0H......
	0x0020 48 89 c1 48 89 df 48 8d 05 00 00 00 00 48 89 d3  H..H..H......H..
	0x0030 e8 00 00 00 00 48 8b 10 48 8b 58 08 48 89 d0 48  .....H..H.X.H..H
	0x0040 8b 6c 24 20 48 83 c4 28 c3 48 89 44 24 08 48 89  .l$ H..(.H.D$.H.
	0x0050 5c 24 10 e8 00 00 00 00 48 8b 44 24 08 48 8b 5c  \$......H.D$.H.\
	0x0060 24 10 eb 9c                                      $...
	rel 28+4 t=14 <unlinkable>.headers+0
	rel 41+4 t=14 type.map[string]string+0
	rel 49+4 t=7 runtime.mapaccess1_faststr+0
	rel 84+4 t=7 runtime.morestack_noctxt+0
<unlinkable>.BenchmarkRefConst STEXT size=168 args=0x8 locals=0x30 funcid=0x0 align=0x0
	0x0000 00000 (./const_test.go:23)	TEXT	<unlinkable>.BenchmarkRefConst(SB), ABIInternal, $48-8
	0x0000 00000 (./const_test.go:23)	CMPQ	SP, 16(R14)
	0x0004 00004 (./const_test.go:23)	PCDATA	$0, $-2
	0x0004 00004 (./const_test.go:23)	JLS	148
	0x000a 00010 (./const_test.go:23)	PCDATA	$0, $-1
	0x000a 00010 (./const_test.go:23)	SUBQ	$48, SP
	0x000e 00014 (./const_test.go:23)	MOVQ	BP, 40(SP)
	0x0013 00019 (./const_test.go:23)	LEAQ	40(SP), BP
	0x0018 00024 (./const_test.go:23)	FUNCDATA	$0, gclocals·wgcWObbY2HYnK2SU/U22lA==(SB)
	0x0018 00024 (./const_test.go:23)	FUNCDATA	$1, gclocals·J5F+7Qw7O7ve2QcWC7DpeQ==(SB)
	0x0018 00024 (./const_test.go:23)	FUNCDATA	$5, <unlinkable>.BenchmarkRefConst.arginfo1(SB)
	0x0018 00024 (./const_test.go:23)	FUNCDATA	$6, <unlinkable>.BenchmarkRefConst.argliveinfo(SB)
	0x0018 00024 (./const_test.go:23)	PCDATA	$3, $1
	0x0018 00024 (./const_test.go:23)	MOVQ	AX, <unlinkable>.b+56(SP)
	0x001d 00029 (./const_test.go:23)	PCDATA	$3, $-1
	0x001d 00029 (./const_test.go:23)	XORL	CX, CX
	0x001f 00031 (./const_test.go:23)	NOP
	0x0020 00032 (./const_test.go:24)	JMP	47
	0x0022 00034 (./const_test.go:24)	MOVQ	<unlinkable>.i+32(SP), CX
	0x0027 00039 (./const_test.go:24)	INCQ	CX
	0x002a 00042 (./const_test.go:24)	MOVQ	<unlinkable>.b+56(SP), AX
	0x002f 00047 (./const_test.go:24)	CMPQ	400(AX), CX
	0x0036 00054 (./const_test.go:24)	JLE	138
	0x0038 00056 (./const_test.go:24)	MOVQ	CX, <unlinkable>.i+32(SP)
	0x003d 00061 (<unknown line number>)	NOP
	0x003d 00061 (./const_test.go:20)	MOVQ	<unlinkable>.headers(SB), BX
	0x0044 00068 (./const_test.go:20)	LEAQ	type.map[string]string(SB), AX
	0x004b 00075 (./const_test.go:20)	LEAQ	go.string."X-Header-Key"(SB), CX
	0x0052 00082 (./const_test.go:20)	MOVL	$12, DI
	0x0057 00087 (./const_test.go:20)	PCDATA	$1, $0
	0x0057 00087 (./const_test.go:20)	CALL	runtime.mapaccess1_faststr(SB)
	0x005c 00092 (./const_test.go:20)	MOVQ	(AX), DX
	0x005f 00095 (./const_test.go:20)	MOVQ	8(AX), SI
	0x0063 00099 (./const_test.go:25)	MOVQ	SI, <unlinkable>.headerKeyValue+8(SB)
	0x006a 00106 (./const_test.go:25)	PCDATA	$0, $-2
	0x006a 00106 (./const_test.go:25)	CMPL	runtime.writeBarrier(SB), $0
	0x0071 00113 (./const_test.go:25)	JNE	124
	0x0073 00115 (./const_test.go:25)	MOVQ	DX, <unlinkable>.headerKeyValue(SB)
	0x007a 00122 (./const_test.go:25)	JMP	34
	0x007c 00124 (./const_test.go:25)	LEAQ	<unlinkable>.headerKeyValue(SB), DI
	0x0083 00131 (./const_test.go:25)	CALL	runtime.gcWriteBarrierDX(SB)
	0x0088 00136 (./const_test.go:25)	JMP	34
	0x008a 00138 (./const_test.go:27)	PCDATA	$0, $-1
	0x008a 00138 (./const_test.go:27)	PCDATA	$1, $-1
	0x008a 00138 (./const_test.go:27)	MOVQ	40(SP), BP
	0x008f 00143 (./const_test.go:27)	ADDQ	$48, SP
	0x0093 00147 (./const_test.go:27)	RET
	0x0094 00148 (./const_test.go:27)	NOP
	0x0094 00148 (./const_test.go:23)	PCDATA	$1, $-1
	0x0094 00148 (./const_test.go:23)	PCDATA	$0, $-2
	0x0094 00148 (./const_test.go:23)	MOVQ	AX, 8(SP)
	0x0099 00153 (./const_test.go:23)	CALL	runtime.morestack_noctxt(SB)
	0x009e 00158 (./const_test.go:23)	MOVQ	8(SP), AX
	0x00a3 00163 (./const_test.go:23)	PCDATA	$0, $-1
	0x00a3 00163 (./const_test.go:23)	JMP	0
	0x0000 49 3b 66 10 0f 86 8a 00 00 00 48 83 ec 30 48 89  I;f.......H..0H.
	0x0010 6c 24 28 48 8d 6c 24 28 48 89 44 24 38 31 c9 90  l$(H.l$(H.D$81..
	0x0020 eb 0d 48 8b 4c 24 20 48 ff c1 48 8b 44 24 38 48  ..H.L$ H..H.D$8H
	0x0030 39 88 90 01 00 00 7e 52 48 89 4c 24 20 48 8b 1d  9.....~RH.L$ H..
	0x0040 00 00 00 00 48 8d 05 00 00 00 00 48 8d 0d 00 00  ....H......H....
	0x0050 00 00 bf 0c 00 00 00 e8 00 00 00 00 48 8b 10 48  ............H..H
	0x0060 8b 70 08 48 89 35 00 00 00 00 83 3d 00 00 00 00  .p.H.5.....=....
	0x0070 00 75 09 48 89 15 00 00 00 00 eb a6 48 8d 3d 00  .u.H........H.=.
	0x0080 00 00 00 e8 00 00 00 00 eb 98 48 8b 6c 24 28 48  ..........H.l$(H
	0x0090 83 c4 30 c3 48 89 44 24 08 e8 00 00 00 00 48 8b  ..0.H.D$......H.
	0x00a0 44 24 08 e9 58 ff ff ff                          D$..X...
	rel 64+4 t=14 <unlinkable>.headers+0
	rel 71+4 t=14 type.map[string]string+0
	rel 78+4 t=14 go.string."X-Header-Key"+0
	rel 88+4 t=7 runtime.mapaccess1_faststr+0
	rel 102+4 t=14 <unlinkable>.headerKeyValue+8
	rel 108+4 t=14 runtime.writeBarrier+-1
	rel 118+4 t=14 <unlinkable>.headerKeyValue+0
	rel 127+4 t=14 <unlinkable>.headerKeyValue+0
	rel 132+4 t=7 runtime.gcWriteBarrierDX+0
	rel 154+4 t=7 runtime.morestack_noctxt+0
<unlinkable>.BenchmarkRefStruct STEXT size=175 args=0x8 locals=0x30 funcid=0x0 align=0x0
	0x0000 00000 (./const_test.go:29)	TEXT	<unlinkable>.BenchmarkRefStruct(SB), ABIInternal, $48-8
	0x0000 00000 (./const_test.go:29)	CMPQ	SP, 16(R14)
	0x0004 00004 (./const_test.go:29)	PCDATA	$0, $-2
	0x0004 00004 (./const_test.go:29)	JLS	150
	0x000a 00010 (./const_test.go:29)	PCDATA	$0, $-1
	0x000a 00010 (./const_test.go:29)	SUBQ	$48, SP
	0x000e 00014 (./const_test.go:29)	MOVQ	BP, 40(SP)
	0x0013 00019 (./const_test.go:29)	LEAQ	40(SP), BP
	0x0018 00024 (./const_test.go:29)	FUNCDATA	$0, gclocals·wgcWObbY2HYnK2SU/U22lA==(SB)
	0x0018 00024 (./const_test.go:29)	FUNCDATA	$1, gclocals·J5F+7Qw7O7ve2QcWC7DpeQ==(SB)
	0x0018 00024 (./const_test.go:29)	FUNCDATA	$5, <unlinkable>.BenchmarkRefStruct.arginfo1(SB)
	0x0018 00024 (./const_test.go:29)	FUNCDATA	$6, <unlinkable>.BenchmarkRefStruct.argliveinfo(SB)
	0x0018 00024 (./const_test.go:29)	PCDATA	$3, $1
	0x0018 00024 (./const_test.go:29)	MOVQ	AX, <unlinkable>.b+56(SP)
	0x001d 00029 (./const_test.go:29)	PCDATA	$3, $-1
	0x001d 00029 (./const_test.go:29)	XORL	CX, CX
	0x001f 00031 (./const_test.go:29)	NOP
	0x0020 00032 (./const_test.go:30)	JMP	47
	0x0022 00034 (./const_test.go:30)	MOVQ	<unlinkable>.i+32(SP), CX
	0x0027 00039 (./const_test.go:30)	INCQ	CX
	0x002a 00042 (./const_test.go:30)	MOVQ	<unlinkable>.b+56(SP), AX
	0x002f 00047 (./const_test.go:30)	CMPQ	400(AX), CX
	0x0036 00054 (./const_test.go:30)	JLE	140
	0x0038 00056 (./const_test.go:30)	MOVQ	CX, <unlinkable>.i+32(SP)
	0x003d 00061 (<unknown line number>)	NOP
	0x003d 00061 (./const_test.go:31)	MOVQ	<unlinkable>.Headers(SB), CX
	0x0044 00068 (./const_test.go:31)	MOVQ	<unlinkable>.Headers+8(SB), DI
	0x004b 00075 (./const_test.go:20)	MOVQ	<unlinkable>.headers(SB), BX
	0x0052 00082 (./const_test.go:20)	LEAQ	type.map[string]string(SB), AX
	0x0059 00089 (./const_test.go:20)	PCDATA	$1, $0
	0x0059 00089 (./const_test.go:20)	CALL	runtime.mapaccess1_faststr(SB)
	0x005e 00094 (./const_test.go:20)	MOVQ	(AX), DX
	0x0061 00097 (./const_test.go:20)	MOVQ	8(AX), SI
	0x0065 00101 (./const_test.go:31)	MOVQ	SI, <unlinkable>.headerKeyValue+8(SB)
	0x006c 00108 (./const_test.go:31)	PCDATA	$0, $-2
	0x006c 00108 (./const_test.go:31)	CMPL	runtime.writeBarrier(SB), $0
	0x0073 00115 (./const_test.go:31)	JNE	126
	0x0075 00117 (./const_test.go:31)	MOVQ	DX, <unlinkable>.headerKeyValue(SB)
	0x007c 00124 (./const_test.go:31)	JMP	34
	0x007e 00126 (./const_test.go:31)	LEAQ	<unlinkable>.headerKeyValue(SB), DI
	0x0085 00133 (./const_test.go:31)	CALL	runtime.gcWriteBarrierDX(SB)
	0x008a 00138 (./const_test.go:31)	JMP	34
	0x008c 00140 (./const_test.go:33)	PCDATA	$0, $-1
	0x008c 00140 (./const_test.go:33)	PCDATA	$1, $-1
	0x008c 00140 (./const_test.go:33)	MOVQ	40(SP), BP
	0x0091 00145 (./const_test.go:33)	ADDQ	$48, SP
	0x0095 00149 (./const_test.go:33)	RET
	0x0096 00150 (./const_test.go:33)	NOP
	0x0096 00150 (./const_test.go:29)	PCDATA	$1, $-1
	0x0096 00150 (./const_test.go:29)	PCDATA	$0, $-2
	0x0096 00150 (./const_test.go:29)	MOVQ	AX, 8(SP)
	0x009b 00155 (./const_test.go:29)	NOP
	0x00a0 00160 (./const_test.go:29)	CALL	runtime.morestack_noctxt(SB)
	0x00a5 00165 (./const_test.go:29)	MOVQ	8(SP), AX
	0x00aa 00170 (./const_test.go:29)	PCDATA	$0, $-1
	0x00aa 00170 (./const_test.go:29)	JMP	0
	0x0000 49 3b 66 10 0f 86 8c 00 00 00 48 83 ec 30 48 89  I;f.......H..0H.
	0x0010 6c 24 28 48 8d 6c 24 28 48 89 44 24 38 31 c9 90  l$(H.l$(H.D$81..
	0x0020 eb 0d 48 8b 4c 24 20 48 ff c1 48 8b 44 24 38 48  ..H.L$ H..H.D$8H
	0x0030 39 88 90 01 00 00 7e 54 48 89 4c 24 20 48 8b 0d  9.....~TH.L$ H..
	0x0040 00 00 00 00 48 8b 3d 00 00 00 00 48 8b 1d 00 00  ....H.=....H....
	0x0050 00 00 48 8d 05 00 00 00 00 e8 00 00 00 00 48 8b  ..H...........H.
	0x0060 10 48 8b 70 08 48 89 35 00 00 00 00 83 3d 00 00  .H.p.H.5.....=..
	0x0070 00 00 00 75 09 48 89 15 00 00 00 00 eb a4 48 8d  ...u.H........H.
	0x0080 3d 00 00 00 00 e8 00 00 00 00 eb 96 48 8b 6c 24  =...........H.l$
	0x0090 28 48 83 c4 30 c3 48 89 44 24 08 0f 1f 44 00 00  (H..0.H.D$...D..
	0x00a0 e8 00 00 00 00 48 8b 44 24 08 e9 51 ff ff ff     .....H.D$..Q...
	rel 64+4 t=14 <unlinkable>.Headers+0
	rel 71+4 t=14 <unlinkable>.Headers+8
	rel 78+4 t=14 <unlinkable>.headers+0
	rel 85+4 t=14 type.map[string]string+0
	rel 90+4 t=7 runtime.mapaccess1_faststr+0
	rel 104+4 t=14 <unlinkable>.headerKeyValue+8
	rel 110+4 t=14 runtime.writeBarrier+-1
	rel 120+4 t=14 <unlinkable>.headerKeyValue+0
	rel 129+4 t=14 <unlinkable>.headerKeyValue+0
	rel 134+4 t=7 runtime.gcWriteBarrierDX+0
	rel 161+4 t=7 runtime.morestack_noctxt+0
<unlinkable>.TestRef STEXT size=202 args=0x8 locals=0x28 funcid=0x0 align=0x0
	0x0000 00000 (./const_test.go:35)	TEXT	<unlinkable>.TestRef(SB), ABIInternal, $40-8
	0x0000 00000 (./const_test.go:35)	CMPQ	SP, 16(R14)
	0x0004 00004 (./const_test.go:35)	PCDATA	$0, $-2
	0x0004 00004 (./const_test.go:35)	JLS	188
	0x000a 00010 (./const_test.go:35)	PCDATA	$0, $-1
	0x000a 00010 (./const_test.go:35)	SUBQ	$40, SP
	0x000e 00014 (./const_test.go:35)	MOVQ	BP, 32(SP)
	0x0013 00019 (./const_test.go:35)	LEAQ	32(SP), BP
	0x0018 00024 (./const_test.go:35)	FUNCDATA	$0, gclocals·Plqv2ff52JtlYaDd2Rwxbg==(SB)
	0x0018 00024 (./const_test.go:35)	FUNCDATA	$1, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
	0x0018 00024 (./const_test.go:35)	FUNCDATA	$5, <unlinkable>.TestRef.arginfo1(SB)
	0x0018 00024 (./const_test.go:35)	FUNCDATA	$6, <unlinkable>.TestRef.argliveinfo(SB)
	0x0018 00024 (./const_test.go:35)	PCDATA	$3, $1
	0x0018 00024 (<unknown line number>)	NOP
	0x0018 00024 (./const_test.go:36)	MOVQ	<unlinkable>.headers(SB), BX
	0x001f 00031 (./const_test.go:20)	LEAQ	type.map[string]string(SB), AX
	0x0026 00038 (./const_test.go:20)	LEAQ	go.string."X-Header-Key"(SB), CX
	0x002d 00045 (./const_test.go:20)	MOVL	$12, DI
	0x0032 00050 (./const_test.go:20)	PCDATA	$1, $0
	0x0032 00050 (./const_test.go:20)	CALL	runtime.mapaccess1_faststr(SB)
	0x0037 00055 (./const_test.go:20)	MOVQ	(AX), DX
	0x003a 00058 (./const_test.go:20)	MOVQ	8(AX), SI
	0x003e 00062 (./const_test.go:36)	MOVQ	SI, <unlinkable>.headerKeyValue+8(SB)
	0x0045 00069 (./const_test.go:36)	PCDATA	$0, $-2
	0x0045 00069 (./const_test.go:36)	CMPL	runtime.writeBarrier(SB), $0
	0x004c 00076 (./const_test.go:36)	JNE	87
	0x004e 00078 (./const_test.go:36)	MOVQ	DX, <unlinkable>.headerKeyValue(SB)
	0x0055 00085 (./const_test.go:36)	JMP	101
	0x0057 00087 (./const_test.go:36)	LEAQ	<unlinkable>.headerKeyValue(SB), DI
	0x005e 00094 (./const_test.go:36)	NOP
	0x0060 00096 (./const_test.go:36)	CALL	runtime.gcWriteBarrierDX(SB)
	0x0065 00101 (./const_test.go:37)	PCDATA	$0, $-1
	0x0065 00101 (./const_test.go:37)	PCDATA	$1, $-1
	0x0065 00101 (<unknown line number>)	NOP
	0x0065 00101 (./const_test.go:37)	MOVQ	<unlinkable>.Headers(SB), CX
	0x006c 00108 (./const_test.go:37)	MOVQ	<unlinkable>.Headers+8(SB), DI
	0x0073 00115 (./const_test.go:20)	MOVQ	<unlinkable>.headers(SB), BX
	0x007a 00122 (./const_test.go:20)	LEAQ	type.map[string]string(SB), AX
	0x0081 00129 (./const_test.go:20)	PCDATA	$1, $0
	0x0081 00129 (./const_test.go:20)	CALL	runtime.mapaccess1_faststr(SB)
	0x0086 00134 (./const_test.go:20)	MOVQ	(AX), DX
	0x0089 00137 (./const_test.go:20)	MOVQ	8(AX), SI
	0x008d 00141 (./const_test.go:37)	MOVQ	SI, <unlinkable>.headerKeyValue+8(SB)
	0x0094 00148 (./const_test.go:37)	PCDATA	$0, $-2
	0x0094 00148 (./const_test.go:37)	CMPL	runtime.writeBarrier(SB), $0
	0x009b 00155 (./const_test.go:37)	JNE	166
	0x009d 00157 (./const_test.go:37)	MOVQ	DX, <unlinkable>.headerKeyValue(SB)
	0x00a4 00164 (./const_test.go:37)	JMP	178
	0x00a6 00166 (./const_test.go:37)	LEAQ	<unlinkable>.headerKeyValue(SB), DI
	0x00ad 00173 (./const_test.go:37)	CALL	runtime.gcWriteBarrierDX(SB)
	0x00b2 00178 (./const_test.go:38)	PCDATA	$0, $-1
	0x00b2 00178 (./const_test.go:38)	PCDATA	$1, $-1
	0x00b2 00178 (./const_test.go:38)	MOVQ	32(SP), BP
	0x00b7 00183 (./const_test.go:38)	ADDQ	$40, SP
	0x00bb 00187 (./const_test.go:38)	RET
	0x00bc 00188 (./const_test.go:38)	NOP
	0x00bc 00188 (./const_test.go:35)	PCDATA	$1, $-1
	0x00bc 00188 (./const_test.go:35)	PCDATA	$0, $-2
	0x00bc 00188 (./const_test.go:35)	NOP
	0x00c0 00192 (./const_test.go:35)	CALL	runtime.morestack_noctxt(SB)
	0x00c5 00197 (./const_test.go:35)	PCDATA	$0, $-1
	0x00c5 00197 (./const_test.go:35)	JMP	0
	0x0000 49 3b 66 10 0f 86 b2 00 00 00 48 83 ec 28 48 89  I;f.......H..(H.
	0x0010 6c 24 20 48 8d 6c 24 20 48 8b 1d 00 00 00 00 48  l$ H.l$ H......H
	0x0020 8d 05 00 00 00 00 48 8d 0d 00 00 00 00 bf 0c 00  ......H.........
	0x0030 00 00 e8 00 00 00 00 48 8b 10 48 8b 70 08 48 89  .......H..H.p.H.
	0x0040 35 00 00 00 00 83 3d 00 00 00 00 00 75 09 48 89  5.....=.....u.H.
	0x0050 15 00 00 00 00 eb 0e 48 8d 3d 00 00 00 00 66 90  .......H.=....f.
	0x0060 e8 00 00 00 00 48 8b 0d 00 00 00 00 48 8b 3d 00  .....H......H.=.
	0x0070 00 00 00 48 8b 1d 00 00 00 00 48 8d 05 00 00 00  ...H......H.....
	0x0080 00 e8 00 00 00 00 48 8b 10 48 8b 70 08 48 89 35  ......H..H.p.H.5
	0x0090 00 00 00 00 83 3d 00 00 00 00 00 75 09 48 89 15  .....=.....u.H..
	0x00a0 00 00 00 00 eb 0c 48 8d 3d 00 00 00 00 e8 00 00  ......H.=.......
	0x00b0 00 00 48 8b 6c 24 20 48 83 c4 28 c3 0f 1f 40 00  ..H.l$ H..(...@.
	0x00c0 e8 00 00 00 00 e9 36 ff ff ff                    ......6...
	rel 27+4 t=14 <unlinkable>.headers+0
	rel 34+4 t=14 type.map[string]string+0
	rel 41+4 t=14 go.string."X-Header-Key"+0
	rel 51+4 t=7 runtime.mapaccess1_faststr+0
	rel 65+4 t=14 <unlinkable>.headerKeyValue+8
	rel 71+4 t=14 runtime.writeBarrier+-1
	rel 81+4 t=14 <unlinkable>.headerKeyValue+0
	rel 90+4 t=14 <unlinkable>.headerKeyValue+0
	rel 97+4 t=7 runtime.gcWriteBarrierDX+0
	rel 104+4 t=14 <unlinkable>.Headers+0
	rel 111+4 t=14 <unlinkable>.Headers+8
	rel 118+4 t=14 <unlinkable>.headers+0
	rel 125+4 t=14 type.map[string]string+0
	rel 130+4 t=7 runtime.mapaccess1_faststr+0
	rel 144+4 t=14 <unlinkable>.headerKeyValue+8
	rel 150+4 t=14 runtime.writeBarrier+-1
	rel 160+4 t=14 <unlinkable>.headerKeyValue+0
	rel 169+4 t=14 <unlinkable>.headerKeyValue+0
	rel 174+4 t=7 runtime.gcWriteBarrierDX+0
	rel 193+4 t=7 runtime.morestack_noctxt+0
<unlinkable>.init STEXT size=154 args=0x0 locals=0x30 funcid=0x0 align=0x0
	0x0000 00000 (./const_test.go:8)	TEXT	<unlinkable>.init(SB), ABIInternal, $48-0
	0x0000 00000 (./const_test.go:8)	CMPQ	SP, 16(R14)
	0x0004 00004 (./const_test.go:8)	PCDATA	$0, $-2
	0x0004 00004 (./const_test.go:8)	JLS	144
	0x000a 00010 (./const_test.go:8)	PCDATA	$0, $-1
	0x000a 00010 (./const_test.go:8)	SUBQ	$48, SP
	0x000e 00014 (./const_test.go:8)	MOVQ	BP, 40(SP)
	0x0013 00019 (./const_test.go:8)	LEAQ	40(SP), BP
	0x0018 00024 (./const_test.go:8)	FUNCDATA	$0, gclocals·J5F+7Qw7O7ve2QcWC7DpeQ==(SB)
	0x0018 00024 (./const_test.go:8)	FUNCDATA	$1, gclocals·CnDyI2HjYXFz19SsOj98tw==(SB)
	0x0018 00024 (./const_test.go:8)	PCDATA	$1, $0
	0x0018 00024 (./const_test.go:8)	CALL	runtime.makemap_small(SB)
	0x001d 00029 (./const_test.go:8)	MOVQ	AX, <unlinkable>..autotmp_4+32(SP)
	0x0022 00034 (./const_test.go:8)	MOVQ	AX, BX
	0x0025 00037 (./const_test.go:8)	LEAQ	go.string."X-Header-Key"(SB), CX
	0x002c 00044 (./const_test.go:8)	MOVL	$12, DI
	0x0031 00049 (./const_test.go:8)	LEAQ	type.map[string]string(SB), AX
	0x0038 00056 (./const_test.go:8)	PCDATA	$1, $1
	0x0038 00056 (./const_test.go:8)	CALL	runtime.mapassign_faststr(SB)
	0x003d 00061 (./const_test.go:8)	MOVQ	$5, 8(AX)
	0x0045 00069 (./const_test.go:8)	PCDATA	$0, $-2
	0x0045 00069 (./const_test.go:8)	CMPL	runtime.writeBarrier(SB), $0
	0x004c 00076 (./const_test.go:8)	JNE	102
	0x004e 00078 (./const_test.go:8)	LEAQ	go.string."value"(SB), CX
	0x0055 00085 (./const_test.go:8)	MOVQ	CX, (AX)
	0x0058 00088 (./const_test.go:8)	MOVQ	<unlinkable>..autotmp_4+32(SP), AX
	0x005d 00093 (./const_test.go:8)	MOVQ	AX, <unlinkable>.headers(SB)
	0x0064 00100 (./const_test.go:8)	JMP	134
	0x0066 00102 (./const_test.go:8)	MOVQ	AX, DI
	0x0069 00105 (./const_test.go:8)	LEAQ	go.string."value"(SB), AX
	0x0070 00112 (./const_test.go:8)	CALL	runtime.gcWriteBarrier(SB)
	0x0075 00117 (./const_test.go:8)	LEAQ	<unlinkable>.headers(SB), DI
	0x007c 00124 (./const_test.go:8)	MOVQ	<unlinkable>..autotmp_4+32(SP), AX
	0x0081 00129 (./const_test.go:8)	CALL	runtime.gcWriteBarrier(SB)
	0x0086 00134 (./const_test.go:8)	PCDATA	$0, $-1
	0x0086 00134 (./const_test.go:8)	PCDATA	$1, $-1
	0x0086 00134 (./const_test.go:8)	MOVQ	40(SP), BP
	0x008b 00139 (./const_test.go:8)	ADDQ	$48, SP
	0x008f 00143 (./const_test.go:8)	RET
	0x0090 00144 (./const_test.go:8)	NOP
	0x0090 00144 (./const_test.go:8)	PCDATA	$1, $-1
	0x0090 00144 (./const_test.go:8)	PCDATA	$0, $-2
	0x0090 00144 (./const_test.go:8)	CALL	runtime.morestack_noctxt(SB)
	0x0095 00149 (./const_test.go:8)	PCDATA	$0, $-1
	0x0095 00149 (./const_test.go:8)	JMP	0
	0x0000 49 3b 66 10 0f 86 86 00 00 00 48 83 ec 30 48 89  I;f.......H..0H.
	0x0010 6c 24 28 48 8d 6c 24 28 e8 00 00 00 00 48 89 44  l$(H.l$(.....H.D
	0x0020 24 20 48 89 c3 48 8d 0d 00 00 00 00 bf 0c 00 00  $ H..H..........
	0x0030 00 48 8d 05 00 00 00 00 e8 00 00 00 00 48 c7 40  .H...........H.@
	0x0040 08 05 00 00 00 83 3d 00 00 00 00 00 75 18 48 8d  ......=.....u.H.
	0x0050 0d 00 00 00 00 48 89 08 48 8b 44 24 20 48 89 05  .....H..H.D$ H..
	0x0060 00 00 00 00 eb 20 48 89 c7 48 8d 05 00 00 00 00  ..... H..H......
	0x0070 e8 00 00 00 00 48 8d 3d 00 00 00 00 48 8b 44 24  .....H.=....H.D$
	0x0080 20 e8 00 00 00 00 48 8b 6c 24 28 48 83 c4 30 c3   .....H.l$(H..0.
	0x0090 e8 00 00 00 00 e9 66 ff ff ff                    ......f...
	rel 25+4 t=7 runtime.makemap_small+0
	rel 40+4 t=14 go.string."X-Header-Key"+0
	rel 52+4 t=14 type.map[string]string+0
	rel 57+4 t=7 runtime.mapassign_faststr+0
	rel 71+4 t=14 runtime.writeBarrier+-1
	rel 81+4 t=14 go.string."value"+0
	rel 96+4 t=14 <unlinkable>.headers+0
	rel 108+4 t=14 go.string."value"+0
	rel 113+4 t=7 runtime.gcWriteBarrier+0
	rel 120+4 t=14 <unlinkable>.headers+0
	rel 130+4 t=7 runtime.gcWriteBarrier+0
	rel 145+4 t=7 runtime.morestack_noctxt+0
go.cuinfo.producer.<unlinkable> SDWARFCUINFO dupok size=0
	0x0000 72 65 67 61 62 69                                regabi
go.cuinfo.packagename.<unlinkable> SDWARFCUINFO dupok size=0
	0x0000 62 65 6e 63 68                                   bench
go.info.<unlinkable>.GetHeader$abstract SDWARFABSFCN dupok size=37
	0x0000 05 3c 75 6e 6c 69 6e 6b 61 62 6c 65 3e 2e 47 65  .<unlinkable>.Ge
	0x0010 74 48 65 61 64 65 72 00 01 01 13 6b 65 79 00 00  tHeader....key..
	0x0020 00 00 00 00 00                                   .....
	rel 32+4 t=31 go.info.string+0
go.string."X-Header-Key" SRODATA dupok size=12
	0x0000 58 2d 48 65 61 64 65 72 2d 4b 65 79              X-Header-Key
<unlinkable>..inittask SNOPTRDATA size=40
	0x0000 00 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00  ................
	0x0010 01 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 testing..inittask+0
	rel 32+8 t=1 <unlinkable>.init+0
go.string."value" SRODATA dupok size=5
	0x0000 76 61 6c 75 65                                   value
<unlinkable>.headers SBSS size=8
go.info.<unlinkable>.headers SDWARFVAR dupok size=37
	0x0000 0a 3c 75 6e 6c 69 6e 6b 61 62 6c 65 3e 2e 68 65  .<unlinkable>.he
	0x0010 61 64 65 72 73 00 09 03 00 00 00 00 00 00 00 00  aders...........
	0x0020 00 00 00 00 01                                   .....
	rel 24+8 t=1 <unlinkable>.headers+0
	rel 32+4 t=31 go.info.map[string]string+0
<unlinkable>.headerKeyValue SBSS size=16
go.info.<unlinkable>.headerKeyValue SDWARFVAR dupok size=44
	0x0000 0a 3c 75 6e 6c 69 6e 6b 61 62 6c 65 3e 2e 68 65  .<unlinkable>.he
	0x0010 61 64 65 72 4b 65 79 56 61 6c 75 65 00 09 03 00  aderKeyValue....
	0x0020 00 00 00 00 00 00 00 00 00 00 00 01              ............
	rel 31+8 t=1 <unlinkable>.headerKeyValue+0
	rel 39+4 t=31 go.info.string+0
<unlinkable>.Headers SDATA size=16
	0x0000 00 00 00 00 00 00 00 00 0c 00 00 00 00 00 00 00  ................
	rel 0+8 t=1 go.string."X-Header-Key"+0
go.info.<unlinkable>.Headers SDWARFVAR dupok size=37
	0x0000 0a 3c 75 6e 6c 69 6e 6b 61 62 6c 65 3e 2e 48 65  .<unlinkable>.He
	0x0010 61 64 65 72 73 00 09 03 00 00 00 00 00 00 00 00  aders...........
	0x0020 00 00 00 00 01                                   .....
	rel 24+8 t=1 <unlinkable>.Headers+0
	rel 32+4 t=31 go.info.struct { Key string }+0
runtime.memequal64·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.memequal64+0
runtime.gcbits.01 SRODATA dupok size=1
	0x0000 01                                               .
type..namedata.*[]uint8- SRODATA dupok size=10
	0x0000 00 08 2a 5b 5d 75 69 6e 74 38                    ..*[]uint8
type.*[]uint8 SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 1f bb cf 43 08 08 08 36 00 00 00 00 00 00 00 00  ...C...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]uint8-+0
	rel 48+8 t=1 type.[]uint8+0
type.[]uint8 SRODATA dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 65 40 04 81 02 08 08 17 00 00 00 00 00 00 00 00  e@..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]uint8-+0
	rel 44+4 t=-32763 type.*[]uint8+0
	rel 48+8 t=1 type.uint8+0
type..namedata.*[8]uint8- SRODATA dupok size=11
	0x0000 00 09 2a 5b 38 5d 75 69 6e 74 38                 ..*[8]uint8
type.*[8]uint8 SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 f8 9a 95 1a 08 08 08 36 00 00 00 00 00 00 00 00  .......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[8]uint8-+0
	rel 48+8 t=1 type.[8]uint8+0
runtime.gcbits. SRODATA dupok size=0
type.[8]uint8 SRODATA dupok size=72
	0x0000 08 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 89 18 9c b4 0a 01 01 11 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 08 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.+0
	rel 40+4 t=5 type..namedata.*[8]uint8-+0
	rel 44+4 t=-32763 type.*[8]uint8+0
	rel 48+8 t=1 type.uint8+0
	rel 56+8 t=1 type.[]uint8+0
type..namedata.*[]string- SRODATA dupok size=11
	0x0000 00 09 2a 5b 5d 73 74 72 69 6e 67                 ..*[]string
type.*[]string SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 76 be 68 d5 08 08 08 36 00 00 00 00 00 00 00 00  v.h....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]string-+0
	rel 48+8 t=1 type.[]string+0
type.[]string SRODATA dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 b5 0e fe e8 02 08 08 17 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]string-+0
	rel 44+4 t=-32763 type.*[]string+0
	rel 48+8 t=1 type.string+0
type..namedata.*[8]string- SRODATA dupok size=12
	0x0000 00 0a 2a 5b 38 5d 73 74 72 69 6e 67              ..*[8]string
type.*[8]string SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 e3 bf d7 63 08 08 08 36 00 00 00 00 00 00 00 00  ...c...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[8]string-+0
	rel 48+8 t=1 type.noalg.[8]string+0
runtime.gcbits.5555 SRODATA dupok size=2
	0x0000 55 55                                            UU
type.noalg.[8]string SRODATA dupok size=72
	0x0000 80 00 00 00 00 00 00 00 78 00 00 00 00 00 00 00  ........x.......
	0x0010 0c 1c ff 04 02 08 08 11 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 08 00 00 00 00 00 00 00                          ........
	rel 32+8 t=1 runtime.gcbits.5555+0
	rel 40+4 t=5 type..namedata.*[8]string-+0
	rel 44+4 t=-32763 type.*[8]string+0
	rel 48+8 t=1 type.string+0
	rel 56+8 t=1 type.[]string+0
type..namedata.*map.bucket[string]string- SRODATA dupok size=27
	0x0000 00 19 2a 6d 61 70 2e 62 75 63 6b 65 74 5b 73 74  ..*map.bucket[st
	0x0010 72 69 6e 67 5d 73 74 72 69 6e 67                 ring]string
type.*map.bucket[string]string SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 40 92 79 ff 08 08 08 36 00 00 00 00 00 00 00 00  @.y....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*map.bucket[string]string-+0
	rel 48+8 t=1 type.noalg.map.bucket[string]string+0
runtime.gcbits.aaaaaaaa02 SRODATA dupok size=5
	0x0000 aa aa aa aa 02                                   .....
type..importpath.. SRODATA dupok size=2
	0x0000 00 00                                            ..
type..namedata.topbits- SRODATA dupok size=9
	0x0000 00 07 74 6f 70 62 69 74 73                       ..topbits
type..namedata.keys- SRODATA dupok size=6
	0x0000 00 04 6b 65 79 73                                ..keys
type..namedata.elems- SRODATA dupok size=7
	0x0000 00 05 65 6c 65 6d 73                             ..elems
type..namedata.overflow- SRODATA dupok size=10
	0x0000 00 08 6f 76 65 72 66 6c 6f 77                    ..overflow
type.noalg.map.bucket[string]string SRODATA dupok size=176
	0x0000 10 01 00 00 00 00 00 00 10 01 00 00 00 00 00 00  ................
	0x0010 4d c0 63 4d 02 08 08 19 00 00 00 00 00 00 00 00  M.cM............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 04 00 00 00 00 00 00 00 04 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0070 00 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0080 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0090 88 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x00a0 00 00 00 00 00 00 00 00 08 01 00 00 00 00 00 00  ................
	rel 32+8 t=1 runtime.gcbits.aaaaaaaa02+0
	rel 40+4 t=5 type..namedata.*map.bucket[string]string-+0
	rel 44+4 t=-32763 type.*map.bucket[string]string+0
	rel 48+8 t=1 type..importpath..+0
	rel 56+8 t=1 type.noalg.map.bucket[string]string+80
	rel 80+8 t=1 type..namedata.topbits-+0
	rel 88+8 t=1 type.[8]uint8+0
	rel 104+8 t=1 type..namedata.keys-+0
	rel 112+8 t=1 type.noalg.[8]string+0
	rel 128+8 t=1 type..namedata.elems-+0
	rel 136+8 t=1 type.noalg.[8]string+0
	rel 152+8 t=1 type..namedata.overflow-+0
	rel 160+8 t=1 type.unsafe.Pointer+0
runtime.strhash·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.strhash+0
type..namedata.*map[string]string- SRODATA dupok size=20
	0x0000 00 12 2a 6d 61 70 5b 73 74 72 69 6e 67 5d 73 74  ..*map[string]st
	0x0010 72 69 6e 67                                      ring
type.*map[string]string SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 d8 6c ad 45 08 08 08 36 00 00 00 00 00 00 00 00  .l.E...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*map[string]string-+0
	rel 48+8 t=1 type.map[string]string+0
type.map[string]string SRODATA dupok size=88
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 de 62 2b 92 02 08 08 35 00 00 00 00 00 00 00 00  .b+....5........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0050 10 10 10 01 0c 00 00 00                          ........
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*map[string]string-+0
	rel 44+4 t=-32763 type.*map[string]string+0
	rel 48+8 t=1 type.string+0
	rel 56+8 t=1 type.string+0
	rel 64+8 t=1 type.noalg.map.bucket[string]string+0
	rel 72+8 t=1 runtime.strhash·f+0
runtime.strequal·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.strequal+0
type..namedata.*struct { Key string }- SRODATA dupok size=24
	0x0000 00 16 2a 73 74 72 75 63 74 20 7b 20 4b 65 79 20  ..*struct { Key 
	0x0010 73 74 72 69 6e 67 20 7d                          string }
type.*struct { Key string } SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 3a 46 49 4c 08 08 08 36 00 00 00 00 00 00 00 00  :FIL...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*struct { Key string }-+0
	rel 48+8 t=1 type.struct { Key string }+0
type..namedata.Key. SRODATA dupok size=5
	0x0000 01 03 4b 65 79                                   ..Key
type.struct { Key string } SRODATA dupok size=104
	0x0000 10 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 34 bb e2 39 02 08 08 19 00 00 00 00 00 00 00 00  4..9............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 01 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0060 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.strequal·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*struct { Key string }-+0
	rel 44+4 t=-32763 type.*struct { Key string }+0
	rel 56+8 t=1 type.struct { Key string }+80
	rel 80+8 t=1 type..namedata.Key.+0
	rel 88+8 t=1 type.string+0
type..importpath.testing. SRODATA dupok size=9
	0x0000 00 07 74 65 73 74 69 6e 67                       ..testing
gclocals·wgcWObbY2HYnK2SU/U22lA== SRODATA dupok size=10
	0x0000 02 00 00 00 01 00 00 00 01 00                    ..........
gclocals·J5F+7Qw7O7ve2QcWC7DpeQ== SRODATA dupok size=8
	0x0000 02 00 00 00 00 00 00 00                          ........
<unlinkable>.GetHeader.arginfo1 SRODATA static dupok size=7
	0x0000 fe 00 08 08 08 fd ff                             .......
<unlinkable>.GetHeader.argliveinfo SRODATA static dupok size=2
	0x0000 00 00                                            ..
<unlinkable>.BenchmarkRefConst.arginfo1 SRODATA static dupok size=3
	0x0000 00 08 ff                                         ...
<unlinkable>.BenchmarkRefConst.argliveinfo SRODATA static dupok size=2
	0x0000 00 00                                            ..
<unlinkable>.BenchmarkRefStruct.arginfo1 SRODATA static dupok size=3
	0x0000 00 08 ff                                         ...
<unlinkable>.BenchmarkRefStruct.argliveinfo SRODATA static dupok size=2
	0x0000 00 00                                            ..
gclocals·Plqv2ff52JtlYaDd2Rwxbg== SRODATA dupok size=9
	0x0000 01 00 00 00 01 00 00 00 00                       .........
gclocals·g2BeySu+wFnoycgXfElmcg== SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
<unlinkable>.TestRef.arginfo1 SRODATA static dupok size=3
	0x0000 00 08 ff                                         ...
<unlinkable>.TestRef.argliveinfo SRODATA static dupok size=2
	0x0000 00 00                                            ..
gclocals·CnDyI2HjYXFz19SsOj98tw== SRODATA dupok size=10
	0x0000 02 00 00 00 01 00 00 00 00 01                    ..........
