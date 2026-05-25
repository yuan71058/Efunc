package utils

// ===================== 栈操作指令 =====================

func H汇编_LEAVE() {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("C9")
}

func H汇编_PUSHAD() {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("60")
}

func H汇编_POPAD() {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("61")
}

func H汇编_NOP() {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("90")
}

func H汇编_RET() {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("C3")
}

func H汇编_IN_AL_DX() {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("EC")
}

func H汇编_TEST_EAX_EAX() {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("85C0")
}

// ===================== MOV 寄存器 ← 立即数 =====================

func H汇编_MOV_EAX(v int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("B8")
	h汇编追加Int32(v)
}

func H汇编_MOV_EBX(v int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("BB")
	h汇编追加Int32(v)
}

func H汇编_MOV_ECX(v int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("B9")
	h汇编追加Int32(v)
}

func H汇编_MOV_EDX(v int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("BA")
	h汇编追加Int32(v)
}

func H汇编_MOV_ESI(v int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("BE")
	h汇编追加Int32(v)
}

func H汇编_MOV_ESP(v int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("BC")
	h汇编追加Int32(v)
}

func H汇编_MOV_EBP(v int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("BD")
	h汇编追加Int32(v)
}

func H汇编_MOV_EDI(v int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("BF")
	h汇编追加Int32(v)
}

// ===================== MOV 寄存器 ← [dword ptr addr] =====================

func H汇编_MOV_EAX_DWORD_PTR(addr int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("A1")
	h汇编追加Int32(addr)
}

func H汇编_MOV_EBX_DWORD_PTR(addr int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("8B1D")
	h汇编追加Int32(addr)
}

func H汇编_MOV_EBP_DWORD_PTR(addr int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("8B2D")
	h汇编追加Int32(addr)
}

func H汇编_MOV_ECX_DWORD_PTR(addr int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("8B0D")
	h汇编追加Int32(addr)
}

func H汇编_MOV_EDX_DWORD_PTR(addr int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("8B15")
	h汇编追加Int32(addr)
}

func H汇编_MOV_ESI_DWORD_PTR(addr int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("8B35")
	h汇编追加Int32(addr)
}

func H汇编_MOV_EDI_DWORD_PTR(addr int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("8B3D")
	h汇编追加Int32(addr)
}

func H汇编_MOV_ESP_DWORD_PTR(addr int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("8B25")
	h汇编追加Int32(addr)
}

// ===================== MOV [dword ptr addr], EAX =====================

func H汇编_MOV_DWORD_PTR_EAX(addr int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("A3")
	h汇编追加Int32(addr)
}

// ===================== MOV 寄存器 ← [寄存器] =====================

func H汇编_MOV_EAX_DWORD_PTR_EAX() {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("8B00")
}

func H汇编_MOV_EAX_DWORD_PTR_EBX() {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("8B03")
}

func H汇编_MOV_EAX_DWORD_PTR_ECX() {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("8B01")
}

func H汇编_MOV_EAX_DWORD_PTR_EDX() {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("8B02")
}

func H汇编_MOV_EAX_DWORD_PTR_EBP() {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("8B4500")
}

func H汇编_MOV_EAX_DWORD_PTR_ESI() {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("8B06")
}

func H汇编_MOV_EAX_DWORD_PTR_EDI() {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("8B07")
}

func H汇编_MOV_EAX_DWORD_PTR_ESP() {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("8B0424")
}

// ===================== MOV 寄存器 ← [寄存器 + 偏移] =====================

func h汇编MovRegDisp(shortHex, longHex string, disp int32) {
	if disp >= -128 && disp <= 127 {
		h汇编追加密文(shortHex)
		h汇编追加Int8(int8(disp))
	} else {
		h汇编追加密文(longHex)
		h汇编追加Int32(disp)
	}
}

func H汇编_MOV_EAX_DWORD_PTR_EAX_ADD(disp int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编MovRegDisp("8B40", "8B80", disp)
}

func H汇编_MOV_EAX_DWORD_PTR_EBX_ADD(disp int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编MovRegDisp("8B43", "8B83", disp)
}

func H汇编_MOV_EAX_DWORD_PTR_ECX_ADD(disp int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编MovRegDisp("8B41", "8B81", disp)
}

func H汇编_MOV_EAX_DWORD_PTR_EDX_ADD(disp int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编MovRegDisp("8B42", "8B82", disp)
}

func H汇编_MOV_EAX_DWORD_PTR_EBP_ADD(disp int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编MovRegDisp("8B45", "8B85", disp)
}

func H汇编_MOV_EAX_DWORD_PTR_ESI_ADD(disp int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编MovRegDisp("8B46", "8B86", disp)
}

func H汇编_MOV_EAX_DWORD_PTR_EDI_ADD(disp int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编MovRegDisp("8B47", "8B87", disp)
}

func H汇编_MOV_EAX_DWORD_PTR_ESP_ADD(disp int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	if disp >= -128 && disp <= 127 {
		h汇编追加密文("8B4424")
		h汇编追加Int8(int8(disp))
	} else {
		h汇编追加密文("8B8424")
		h汇编追加Int32(disp)
	}
}

func H汇编_MOV_EBX_DWORD_PTR_EAX_ADD(disp int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编MovRegDisp("8B58", "8B98", disp)
}

func H汇编_MOV_EBX_DWORD_PTR_EBX_ADD(disp int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编MovRegDisp("8B5B", "8B9B", disp)
}

func H汇编_MOV_EBX_DWORD_PTR_ECX_ADD(disp int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编MovRegDisp("8B59", "8B99", disp)
}

func H汇编_MOV_EBX_DWORD_PTR_EDX_ADD(disp int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编MovRegDisp("8B5A", "8B9A", disp)
}

func H汇编_MOV_EBX_DWORD_PTR_EBP_ADD(disp int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编MovRegDisp("8B5D", "8B9D", disp)
}

func H汇编_MOV_EBX_DWORD_PTR_ESI_ADD(disp int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编MovRegDisp("8B5E", "8B9E", disp)
}

func H汇编_MOV_EBX_DWORD_PTR_EDI_ADD(disp int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编MovRegDisp("8B5F", "8B9F", disp)
}

func H汇编_MOV_EBX_DWORD_PTR_ESP_ADD(disp int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	if disp >= -128 && disp <= 127 {
		h汇编追加密文("8B5C24")
		h汇编追加Int8(int8(disp))
	} else {
		h汇编追加密文("8B9C24")
		h汇编追加Int32(disp)
	}
}

func H汇编_MOV_ECX_DWORD_PTR_EAX_ADD(disp int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编MovRegDisp("8B48", "8B88", disp)
}

func H汇编_MOV_ECX_DWORD_PTR_EBX_ADD(disp int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编MovRegDisp("8B4B", "8B8B", disp)
}

func H汇编_MOV_ECX_DWORD_PTR_ECX_ADD(disp int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编MovRegDisp("8B49", "8B89", disp)
}

func H汇编_MOV_ECX_DWORD_PTR_EDX_ADD(disp int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编MovRegDisp("8B4A", "8B8A", disp)
}

func H汇编_MOV_ECX_DWORD_PTR_EBP_ADD(disp int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编MovRegDisp("8B4D", "8B8D", disp)
}

func H汇编_MOV_ECX_DWORD_PTR_ESI_ADD(disp int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编MovRegDisp("8B4E", "8B8E", disp)
}

func H汇编_MOV_ECX_DWORD_PTR_EDI_ADD(disp int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编MovRegDisp("8B4F", "8B8F", disp)
}

func H汇编_MOV_ECX_DWORD_PTR_ESP_ADD(disp int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	if disp >= -128 && disp <= 127 {
		h汇编追加密文("8B4C24")
		h汇编追加Int8(int8(disp))
	} else {
		h汇编追加密文("8B8C24")
		h汇编追加Int32(disp)
	}
}

func H汇编_MOV_EDX_DWORD_PTR_EAX_ADD(disp int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编MovRegDisp("8B50", "8B90", disp)
}

func H汇编_MOV_EDX_DWORD_PTR_EBX_ADD(disp int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编MovRegDisp("8B53", "8B93", disp)
}

func H汇编_MOV_EDX_DWORD_PTR_ECX_ADD(disp int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编MovRegDisp("8B51", "8B91", disp)
}

func H汇编_MOV_EDX_DWORD_PTR_EDX_ADD(disp int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编MovRegDisp("8B52", "8B92", disp)
}

func H汇编_MOV_EDX_DWORD_PTR_EBP_ADD(disp int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编MovRegDisp("8B55", "8B95", disp)
}

func H汇编_MOV_EDX_DWORD_PTR_ESI_ADD(disp int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编MovRegDisp("8B56", "8B96", disp)
}

func H汇编_MOV_EDX_DWORD_PTR_EDI_ADD(disp int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编MovRegDisp("8B57", "8B97", disp)
}

func H汇编_MOV_ESI_DWORD_PTR_EAX_ADD(disp int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编MovRegDisp("8B70", "8BB0", disp)
}

// ===================== ADD 寄存器 ← 立即数 =====================

func h汇编AddImmReg(shortImm, shortOp, longImm, longOp string, v int32) {
	if v >= -128 && v <= 127 {
		h汇编追加密文(shortImm)
		h汇编追加Int8(int8(v))
	} else {
		h汇编追加密文(longImm)
		h汇编追加Int32(v)
	}
}

func H汇编_ADD_EAX(v int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	if v >= -128 && v <= 127 {
		h汇编追加密文("83C0")
		h汇编追加Int8(int8(v))
	} else {
		h汇编追加密文("05")
		h汇编追加Int32(v)
	}
}

func H汇编_ADD_EBX(v int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	if v >= -128 && v <= 127 {
		h汇编追加密文("83C3")
		h汇编追加Int8(int8(v))
	} else {
		h汇编追加密文("81C3")
		h汇编追加Int32(v)
	}
}

func H汇编_ADD_ECX(v int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	if v >= -128 && v <= 127 {
		h汇编追加密文("83C1")
		h汇编追加Int8(int8(v))
	} else {
		h汇编追加密文("81C1")
		h汇编追加Int32(v)
	}
}

func H汇编_ADD_EDX(v int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	if v >= -128 && v <= 127 {
		h汇编追加密文("83C2")
		h汇编追加Int8(int8(v))
	} else {
		h汇编追加密文("81C2")
		h汇编追加Int32(v)
	}
}

func H汇编_ADD_ESI(v int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	if v >= -128 && v <= 127 {
		h汇编追加密文("83C6")
		h汇编追加Int8(int8(v))
	} else {
		h汇编追加密文("81C6")
		h汇编追加Int32(v)
	}
}

func H汇编_ADD_ESP(v int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	if v >= -128 && v <= 127 {
		h汇编追加密文("83C4")
		h汇编追加Int8(int8(v))
	} else {
		h汇编追加密文("81C4")
		h汇编追加Int32(v)
	}
}

// ===================== ADD 寄存器 ← 寄存器 =====================

func H汇编_ADD_EAX_EDX() {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("03C2")
}

func H汇编_ADD_EBX_EAX() {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("03D8")
}

// ===================== ADD 寄存器 ← [dword ptr addr] =====================

func H汇编_ADD_EAX_DWORD_PTR(addr int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("0305")
	h汇编追加Int32(addr)
}

func H汇编_ADD_EBX_DWORD_PTR(addr int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("031D")
	h汇编追加Int32(addr)
}

func H汇编_ADD_EBP_DWORD_PTR(addr int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("032D")
	h汇编追加Int32(addr)
}

// ===================== CALL 指令 =====================

func H汇编_CALL(addr int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("E8")
	h汇编追加Int32(addr)
}

func H汇编_CALL_EAX() {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("FFD0")
}

func H汇编_CALL_EBX() {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("FFD3")
}

func H汇编_CALL_ECX() {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("FFD1")
}

func H汇编_CALL_EDX() {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("FFD2")
}

func H汇编_CALL_EBP() {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("FFD5")
}

func H汇编_CALL_ESI() {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("FFD6")
}

func H汇编_CALL_EDI() {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("FFD7")
}

func H汇编_CALL_ESP() {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("FFD4")
}

func H汇编_CALL_DWORD_PTR(addr int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("FF15")
	h汇编追加Int32(addr)
}

func H汇编_CALL_DWORD_PTR_EAX() {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("FF10")
}

func H汇编_CALL_DWORD_PTR_EBX() {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("FF13")
}

// ===================== JMP 指令 =====================

func H汇编_JMP(addr int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("E9")
	h汇编追加Int32(addr)
}

func H汇编_JMP_EAX() {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("FFE0")
}

// ===================== CMP 指令 =====================

func H汇编_CMP_EAX(v int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	if v >= -128 && v <= 127 {
		h汇编追加密文("83F8")
		h汇编追加Int8(int8(v))
	} else {
		h汇编追加密文("3D")
		h汇编追加Int32(v)
	}
}

func H汇编_CMP_EAX_EDX() {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("3BC2")
}

func H汇编_CMP_EAX_DWORD_PTR(addr int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("3B05")
	h汇编追加Int32(addr)
}

func H汇编_CMP_DWORD_PTR_EAX(addr int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("3905")
	h汇编追加Int32(addr)
}

// ===================== INC/DEC 指令 =====================

func H汇编_INC_EAX() {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("40")
}

func H汇编_INC_EBX() {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("43")
}

func H汇编_INC_ECX() {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("41")
}

func H汇编_INC_EDX() {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("42")
}

func H汇编_INC_ESI() {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("46")
}

func H汇编_INC_EDI() {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("47")
}

func H汇编_INC_DWORD_PTR_EAX() {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("FF00")
}

func H汇编_INC_DWORD_PTR_EBX() {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("FF03")
}

func H汇编_INC_DWORD_PTR_ECX() {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("FF01")
}

func H汇编_INC_DWORD_PTR_EDX() {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("FF02")
}

func H汇编_DEC_EAX() {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("48")
}

func H汇编_DEC_EBX() {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("4B")
}

func H汇编_DEC_ECX() {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("49")
}

func H汇编_DEC_EDX() {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("4A")
}

// ===================== IDIV/IMUL 指令 =====================

func H汇编_IDIV_EAX() {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("F7F8")
}

func H汇编_IDIV_EBX() {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("F7FB")
}

func H汇编_IDIV_ECX() {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("F7F9")
}

func H汇编_IDIV_EDX() {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("F7FA")
}

func H汇编_IMUL_EAX_EDX() {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("0FAFC2")
}

func H汇编_IMUL_EAX(v int8) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("6BC0")
	h汇编追加Int8(v)
}

func H汇编_IMULB_EAX(v int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("69C0")
	h汇编追加Int32(v)
}

func H汇编_ADDR(v int32) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	h汇编追加密文("BD")
	h汇编追加Int32(v)
}