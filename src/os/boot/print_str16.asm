; bits 16

; print null-terminated string stored in si
print_str:
  pushf
  pusha

  cld  ; make sure we're incrementing on lodsb

  mov bx, 0

.iter:
  lodsb
  or al, al  ; this sets zf to 1 if the result is zero
  jz .done ; if zf is set

  mov ah, 0x0e
  int 0x10  ; http://www.ctyme.com/intr/rb-0106.htm

  jmp .iter

.done:
  popa
  popf
  ret
