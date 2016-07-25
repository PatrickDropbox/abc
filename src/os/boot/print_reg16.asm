print_reg16:
  pushf
  pusha

  push dx

  mov dx, ax
  call print_hex16
  mov si, _space
  call print_str

  mov dx, bx
  call print_hex16
  mov si, _space
  call print_str

  mov dx, cx
  call print_hex16
  mov si, _space
  call print_str

  pop dx
  call print_hex16
  mov si, _crlf
  call print_str

  popa
  popf
  ret
