; bits 16

; print dx's value via BIOS
print_hex16:
  pushf
  push ax
  push bx
  push dx

  mov ah, 0x0e
  mov bx, 0

  mov al, '0'
  int 0x10

  mov al, 'x'
  int 0x10

  mov dl, dh
  call .print_char

  pop dx
  call .print_char

  pop bx
  pop ax
  popf
  ret

.print_char:
  mov al, dl
  shr al, 4
  call .format_hex
  int 0x10

  mov al, dl
  and al, 0x0f
  call .format_hex
  int 0x10

  ret

.format_hex:
  cmp al, 10
  jge .else

  add al, '0'
  ret
.else:
  add al, ('A'-10)
  ret
