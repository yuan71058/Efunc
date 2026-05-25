// x86 汇编指令集模块
// 提供常用 x86 汇编指令的机器码生成函数。
// 每调用一个函数即向内部机器码缓冲区追加对应指令的字节码。
// 配合 ASM_Execute / ASM_ExecuteRemote 使用。
package utils

// ===================== Stack Operations =====================

func ASM_LEAVE() {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("C9")
}

func ASM_PUSHAD() {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("60")
}

func ASM_POPAD() {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("61")
}

func ASM_NOP() {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("90")
}

func ASM_RET() {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("C3")
}

func ASM_IN_AL_DX() {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("EC")
}

func ASM_TEST_EAX_EAX() {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("85C0")
}

// ===================== MOV reg ← imm =====================

func ASM_MOV_EAX(v int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("B8")
	asmAppendInt32(v)
}

func ASM_MOV_EBX(v int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("BB")
	asmAppendInt32(v)
}

func ASM_MOV_ECX(v int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("B9")
	asmAppendInt32(v)
}

func ASM_MOV_EDX(v int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("BA")
	asmAppendInt32(v)
}

func ASM_MOV_ESI(v int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("BE")
	asmAppendInt32(v)
}

func ASM_MOV_ESP(v int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("BC")
	asmAppendInt32(v)
}

func ASM_MOV_EBP(v int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("BD")
	asmAppendInt32(v)
}

func ASM_MOV_EDI(v int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("BF")
	asmAppendInt32(v)
}

// ===================== MOV reg ← [dword ptr addr] =====================

func ASM_MOV_EAX_DWORD_PTR(addr int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("A1")
	asmAppendInt32(addr)
}

func ASM_MOV_EBX_DWORD_PTR(addr int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("8B1D")
	asmAppendInt32(addr)
}

func ASM_MOV_EBP_DWORD_PTR(addr int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("8B2D")
	asmAppendInt32(addr)
}

func ASM_MOV_ECX_DWORD_PTR(addr int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("8B0D")
	asmAppendInt32(addr)
}

func ASM_MOV_EDX_DWORD_PTR(addr int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("8B15")
	asmAppendInt32(addr)
}

func ASM_MOV_ESI_DWORD_PTR(addr int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("8B35")
	asmAppendInt32(addr)
}

func ASM_MOV_EDI_DWORD_PTR(addr int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("8B3D")
	asmAppendInt32(addr)
}

func ASM_MOV_ESP_DWORD_PTR(addr int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("8B25")
	asmAppendInt32(addr)
}

// ===================== MOV [dword ptr addr], EAX =====================

func ASM_MOV_DWORD_PTR_EAX(addr int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("A3")
	asmAppendInt32(addr)
}

// ===================== MOV reg ← [reg] =====================

func ASM_MOV_EAX_DWORD_PTR_EAX() {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("8B00")
}

func ASM_MOV_EAX_DWORD_PTR_EBX() {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("8B03")
}

func ASM_MOV_EAX_DWORD_PTR_ECX() {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("8B01")
}

func ASM_MOV_EAX_DWORD_PTR_EDX() {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("8B02")
}

func ASM_MOV_EAX_DWORD_PTR_EBP() {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("8B4500")
}

func ASM_MOV_EAX_DWORD_PTR_ESI() {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("8B06")
}

func ASM_MOV_EAX_DWORD_PTR_EDI() {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("8B07")
}

func ASM_MOV_EAX_DWORD_PTR_ESP() {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("8B0424")
}

// ===================== MOV reg ← [reg + disp] =====================

func asmMovRegDisp(shortHex, longHex string, disp int32) {
	if disp >= -128 && disp <= 127 {
		asmAppendHex(shortHex)
		asmAppendInt8(int8(disp))
	} else {
		asmAppendHex(longHex)
		asmAppendInt32(disp)
	}
}

func ASM_MOV_EAX_DWORD_PTR_EAX_ADD(disp int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmMovRegDisp("8B40", "8B80", disp)
}

func ASM_MOV_EAX_DWORD_PTR_EBX_ADD(disp int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmMovRegDisp("8B43", "8B83", disp)
}

func ASM_MOV_EAX_DWORD_PTR_ECX_ADD(disp int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmMovRegDisp("8B41", "8B81", disp)
}

func ASM_MOV_EAX_DWORD_PTR_EDX_ADD(disp int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmMovRegDisp("8B42", "8B82", disp)
}

func ASM_MOV_EAX_DWORD_PTR_EBP_ADD(disp int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmMovRegDisp("8B45", "8B85", disp)
}

func ASM_MOV_EAX_DWORD_PTR_ESI_ADD(disp int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmMovRegDisp("8B46", "8B86", disp)
}

func ASM_MOV_EAX_DWORD_PTR_EDI_ADD(disp int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmMovRegDisp("8B47", "8B87", disp)
}

func ASM_MOV_EAX_DWORD_PTR_ESP_ADD(disp int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	if disp >= -128 && disp <= 127 {
		asmAppendHex("8B4424")
		asmAppendInt8(int8(disp))
	} else {
		asmAppendHex("8B8424")
		asmAppendInt32(disp)
	}
}

func ASM_MOV_EBX_DWORD_PTR_EAX_ADD(disp int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmMovRegDisp("8B58", "8B98", disp)
}

func ASM_MOV_EBX_DWORD_PTR_EBX_ADD(disp int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmMovRegDisp("8B5B", "8B9B", disp)
}

func ASM_MOV_EBX_DWORD_PTR_ECX_ADD(disp int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmMovRegDisp("8B59", "8B99", disp)
}

func ASM_MOV_EBX_DWORD_PTR_EDX_ADD(disp int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmMovRegDisp("8B5A", "8B9A", disp)
}

func ASM_MOV_EBX_DWORD_PTR_EBP_ADD(disp int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmMovRegDisp("8B5D", "8B9D", disp)
}

func ASM_MOV_EBX_DWORD_PTR_ESI_ADD(disp int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmMovRegDisp("8B5E", "8B9E", disp)
}

func ASM_MOV_EBX_DWORD_PTR_EDI_ADD(disp int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmMovRegDisp("8B5F", "8B9F", disp)
}

func ASM_MOV_EBX_DWORD_PTR_ESP_ADD(disp int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	if disp >= -128 && disp <= 127 {
		asmAppendHex("8B5C24")
		asmAppendInt8(int8(disp))
	} else {
		asmAppendHex("8B9C24")
		asmAppendInt32(disp)
	}
}

func ASM_MOV_ECX_DWORD_PTR_EAX_ADD(disp int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmMovRegDisp("8B48", "8B88", disp)
}

func ASM_MOV_ECX_DWORD_PTR_EBX_ADD(disp int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmMovRegDisp("8B4B", "8B8B", disp)
}

func ASM_MOV_ECX_DWORD_PTR_ECX_ADD(disp int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmMovRegDisp("8B49", "8B89", disp)
}

func ASM_MOV_ECX_DWORD_PTR_EDX_ADD(disp int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmMovRegDisp("8B4A", "8B8A", disp)
}

func ASM_MOV_ECX_DWORD_PTR_EBP_ADD(disp int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmMovRegDisp("8B4D", "8B8D", disp)
}

func ASM_MOV_ECX_DWORD_PTR_ESI_ADD(disp int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmMovRegDisp("8B4E", "8B8E", disp)
}

func ASM_MOV_ECX_DWORD_PTR_EDI_ADD(disp int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmMovRegDisp("8B4F", "8B8F", disp)
}

func ASM_MOV_ECX_DWORD_PTR_ESP_ADD(disp int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	if disp >= -128 && disp <= 127 {
		asmAppendHex("8B4C24")
		asmAppendInt8(int8(disp))
	} else {
		asmAppendHex("8B8C24")
		asmAppendInt32(disp)
	}
}

func ASM_MOV_EDX_DWORD_PTR_EAX_ADD(disp int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmMovRegDisp("8B50", "8B90", disp)
}

func ASM_MOV_EDX_DWORD_PTR_EBX_ADD(disp int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmMovRegDisp("8B53", "8B93", disp)
}

func ASM_MOV_EDX_DWORD_PTR_ECX_ADD(disp int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmMovRegDisp("8B51", "8B91", disp)
}

func ASM_MOV_EDX_DWORD_PTR_EDX_ADD(disp int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmMovRegDisp("8B52", "8B92", disp)
}

func ASM_MOV_EDX_DWORD_PTR_EBP_ADD(disp int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmMovRegDisp("8B55", "8B95", disp)
}

func ASM_MOV_EDX_DWORD_PTR_ESI_ADD(disp int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmMovRegDisp("8B56", "8B96", disp)
}

func ASM_MOV_EDX_DWORD_PTR_EDI_ADD(disp int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmMovRegDisp("8B57", "8B97", disp)
}

func ASM_MOV_ESI_DWORD_PTR_EAX_ADD(disp int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmMovRegDisp("8B70", "8BB0", disp)
}

// ===================== ADD reg ← imm =====================

func ASM_ADD_EAX(v int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	if v >= -128 && v <= 127 {
		asmAppendHex("83C0")
		asmAppendInt8(int8(v))
	} else {
		asmAppendHex("05")
		asmAppendInt32(v)
	}
}

func ASM_ADD_EBX(v int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	if v >= -128 && v <= 127 {
		asmAppendHex("83C3")
		asmAppendInt8(int8(v))
	} else {
		asmAppendHex("81C3")
		asmAppendInt32(v)
	}
}

func ASM_ADD_ECX(v int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	if v >= -128 && v <= 127 {
		asmAppendHex("83C1")
		asmAppendInt8(int8(v))
	} else {
		asmAppendHex("81C1")
		asmAppendInt32(v)
	}
}

func ASM_ADD_EDX(v int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	if v >= -128 && v <= 127 {
		asmAppendHex("83C2")
		asmAppendInt8(int8(v))
	} else {
		asmAppendHex("81C2")
		asmAppendInt32(v)
	}
}

func ASM_ADD_ESI(v int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	if v >= -128 && v <= 127 {
		asmAppendHex("83C6")
		asmAppendInt8(int8(v))
	} else {
		asmAppendHex("81C6")
		asmAppendInt32(v)
	}
}

func ASM_ADD_ESP(v int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	if v >= -128 && v <= 127 {
		asmAppendHex("83C4")
		asmAppendInt8(int8(v))
	} else {
		asmAppendHex("81C4")
		asmAppendInt32(v)
	}
}

// ===================== ADD reg ← reg =====================

func ASM_ADD_EAX_EDX() {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("03C2")
}

func ASM_ADD_EBX_EAX() {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("03D8")
}

// ===================== ADD reg ← [dword ptr addr] =====================

func ASM_ADD_EAX_DWORD_PTR(addr int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("0305")
	asmAppendInt32(addr)
}

func ASM_ADD_EBX_DWORD_PTR(addr int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("031D")
	asmAppendInt32(addr)
}

func ASM_ADD_EBP_DWORD_PTR(addr int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("032D")
	asmAppendInt32(addr)
}

// ===================== CALL instructions =====================

func ASM_CALL(addr int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("E8")
	asmAppendInt32(addr)
}

func ASM_CALL_EAX() {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("FFD0")
}

func ASM_CALL_EBX() {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("FFD3")
}

func ASM_CALL_ECX() {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("FFD1")
}

func ASM_CALL_EDX() {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("FFD2")
}

func ASM_CALL_EBP() {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("FFD5")
}

func ASM_CALL_ESI() {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("FFD6")
}

func ASM_CALL_EDI() {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("FFD7")
}

func ASM_CALL_ESP() {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("FFD4")
}

func ASM_CALL_DWORD_PTR(addr int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("FF15")
	asmAppendInt32(addr)
}

func ASM_CALL_DWORD_PTR_EAX() {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("FF10")
}

func ASM_CALL_DWORD_PTR_EBX() {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("FF13")
}

// ===================== JMP instructions =====================

func ASM_JMP(addr int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("E9")
	asmAppendInt32(addr)
}

func ASM_JMP_EAX() {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("FFE0")
}

// ===================== CMP instructions =====================

func ASM_CMP_EAX(v int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	if v >= -128 && v <= 127 {
		asmAppendHex("83F8")
		asmAppendInt8(int8(v))
	} else {
		asmAppendHex("3D")
		asmAppendInt32(v)
	}
}

func ASM_CMP_EAX_EDX() {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("3BC2")
}

func ASM_CMP_EAX_DWORD_PTR(addr int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("3B05")
	asmAppendInt32(addr)
}

func ASM_CMP_DWORD_PTR_EAX(addr int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("3905")
	asmAppendInt32(addr)
}

// ===================== INC/DEC instructions =====================

func ASM_INC_EAX() {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("40")
}

func ASM_INC_EBX() {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("43")
}

func ASM_INC_ECX() {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("41")
}

func ASM_INC_EDX() {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("42")
}

func ASM_INC_ESI() {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("46")
}

func ASM_INC_EDI() {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("47")
}

func ASM_INC_DWORD_PTR_EAX() {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("FF00")
}

func ASM_INC_DWORD_PTR_EBX() {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("FF03")
}

func ASM_INC_DWORD_PTR_ECX() {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("FF01")
}

func ASM_INC_DWORD_PTR_EDX() {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("FF02")
}

func ASM_DEC_EAX() {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("48")
}

func ASM_DEC_EBX() {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("4B")
}

func ASM_DEC_ECX() {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("49")
}

func ASM_DEC_EDX() {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("4A")
}

// ===================== IDIV/IMUL instructions =====================

func ASM_IDIV_EAX() {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("F7F8")
}

func ASM_IDIV_EBX() {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("F7FB")
}

func ASM_IDIV_ECX() {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("F7F9")
}

func ASM_IDIV_EDX() {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("F7FA")
}

func ASM_IMUL_EAX_EDX() {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("0FAFC2")
}

func ASM_IMUL_EAX(v int8) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("6BC0")
	asmAppendInt8(v)
}

func ASM_IMULB_EAX(v int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("69C0")
	asmAppendInt32(v)
}

func ASM_ADDR(v int32) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	asmAppendHex("BD")
	asmAppendInt32(v)
}