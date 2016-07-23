; BIOS loads the boot sector at 0x7c00.
; See http://wiki.osdev.org/Memory_Map_(x86)
org 0x7c00
bits 16  ; x86 always boots in real mode

; Boot loader's entry point is just the first byte of the sector.
boot_sector_entry_point:

%include "src/os/boot/init_registers.asm"

;
; Now that the stack have been setup, we can make calls
;

; This should be first call in case dl gets clobbered.
call save_boot_drive_id

; Make sure we have space to load additional boot data.
call real_mode_memory_check

; Read 1 sectors from boot disk, and load that into 0x7e00.
; [0x7e00, 0x9fc00) is free for use. See http://wiki.osdev.org/Memory_Map_(x86)
mov dh, 1
mov bx, 0x7e00
call load_boot_data

;
; At this point, we can only fit another 30 bytes or so into the boot sector.
; Time to move off the boot sector.
;
; Since the boot sector is in [0x7c00, 0x7e00), and the newly loaded sectors
; are in [0x7e00, ...), the two memory ranges are adjacent.  Hence, the label
; address offset are identical, and the boot sector can referr to the second
; sector via label (and vice versa).
;
jmp second_stage_entry_point  ; can also use "0x0000:0x7e00"

;
; Global variables
;

_crlf:
  db 13, 10, 0  ; BIOS printing requires explicit \r

%include "src/os/boot/boot_disk.asm"
%include "src/os/boot/halt.asm"
%include "src/os/boot/print_hex16.asm"
%include "src/os/boot/print_str16.asm"
%include "src/os/boot/real_mode_memory_check.asm"

;
; End of boot sector.  The last two bytes on the boot sector must end with
; 0xaa55.
;

times 510-($-$$) db 0
dw 0xaa55

; -----------------------------------------------------------------------------

; This is loaded into 0x7e00
second_stage_entry_point:
  mov si, _second_stage_msg
  call print_str16

;
; To be continue ...
;

jmp halt

_second_stage_msg:
  db "Entered second stage booting", 13, 10, 0

; VBoxManage requires the generate bin file to be multiples of 512 bytes ...
times 1024-($-$$) db 0

