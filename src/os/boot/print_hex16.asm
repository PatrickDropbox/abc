; bits 16

; print dl's value via BIOS (without 0x prefix)
print_hex_byte:
  pushf
  pusha

  mov ah, 0x0e
  mov bx, 0

  mov al, dl
  shr al, 4
  call .format_hex
  int 0x10

  mov al, dl
  and al, 0x0f
  call .format_hex
  int 0x10

  popa
  popf
  ret

.format_hex:
  cmp al, 10
  jge .else

  add al, '0'
  ret
.else:
  add al, ('A'-10)
  ret

; print dl's value via BIOS (with 0x prefix)
print_hex8:
  pushf
  pusha

  mov si, _0x
  call print_str

  call print_hex_byte

  popa
  popf
  ret

; print dx's value via BIOS
print_hex16:
  pushf
  pusha
  push dx

  mov ah, 0x0e
  mov bx, 0

  mov al, '0'
  int 0x10

  mov al, 'x'
  int 0x10

  mov dl, dh
  call print_hex_byte

  pop dx
  call print_hex_byte

  popa
  popf
  ret

; print si's content via BIOS (assuming little endian), cx specifies number of
; bytes to print
print_hex_number:
  pushf
  pusha

  push si

  mov si, _0x
  call print_str

  pop si

  mov ax, 0
  add si, cx
  sub si, 1

.iter:
  cmp ax, cx
  jge .done

  mov dl, [si]
  call print_hex_byte

  sub si, 1
  add ax, 1
  jmp .iter

.done:
  popa
  popf
  ret

; print si's content via BIOS, cx specifies number of bytes to print
print_hex_bytes:
  pushf
  pusha

  push si

  mov si, _0x
  call print_str

  pop si
  mov ax, 0

.iter:
  cmp ax, cx
  jge .done

  mov dl, [si]
  call print_hex_byte

  add si, 1
  add ax, 1
  jmp .iter

.done:
  popa
  popf
  ret
