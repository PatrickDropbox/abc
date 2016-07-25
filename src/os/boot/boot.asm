; BIOS loads the boot sector at 0x7c00.
; See http://wiki.osdev.org/Memory_Map_(x86)
org 0x7c00
bits 16  ; x86 always boots in real mode

;
; Memory map:
;
; [0x0000, 0x0400)
;     real mode interrupt vector table (unusable)
; [0x0400, 0x0500)
;     BIOS data area (unusable)
; [0x0500, 0x????]
;     memory map (via detect_memory); should use boot sector's len/addr global
;     variables to refer to the map, in case it's moved.
; [0x????, 0x7c00)
;     stack (grows down)
; [0x7c00, 0x7e00)
;     boot sector
; [0x7e00, 0x7e00 + 512 * SECOND_STAGE_SECTORS)
;     second-stage boot sectors
;


%define SECOND_STAGE_SECTORS 4

; Boot loader's entry point is just the first byte of the sector.
boot_sector_entry_point:

%include "src/os/boot/init_registers16.asm"

;
; Now that the stack have been setup, we can make calls
;

; This should be first call in case dl gets clobbered.
call save_boot_drive_id

; Make sure we have space to load additional boot data.
call real_mode_memory_check

; Read SECOND_STAGE_SECTORS from boot disk, and load that into 0x7e00.
; [0x7e00, 0x9fc00) is free for use. See http://wiki.osdev.org/Memory_Map_(x86)
mov dh, SECOND_STAGE_SECTORS
mov bx, 0x7e00
call load_boot_data

;
; Time to move off the boot sector.
;
; Since the boot sector is in [0x7c00, 0x7e00), and the newly loaded sectors
; are in [0x7e00, ...), the two memory ranges are adjacent.  Hence, the label
; address offset are identical, and the boot sector can referr to the second
; sector via label (and vice versa).
;
jmp second_stage_entry_point  ; can also use "0x0000:0x7e00"

;
; Code
;

%include "src/os/boot/boot_disk16.asm"
%include "src/os/boot/halt16.asm"
%include "src/os/boot/print_hex16.asm"
%include "src/os/boot/print_str16.asm"
%include "src/os/boot/real_mode_memory_check16.asm"

;
; Shared global variables
;

; Which drive did the BIOS boot from.  Initialized by save_boot_drive_id call.
_boot_drive_id:
  db 0

%define MEMORY_MAP_ENTRY_SIZE 24

; Num entries (not byte count) in the memory map.  Initialized by detect_memory
; call.
_memory_map_len:
  dw 0  ; initialized by detect_memory call

; Where to save the memory map.  The map itself is initialized by call to
; detect_memory.  Each entry is a packed struct
;
; struct MemoryMapEntry {
;   uint64 base_addr;
;   uint64 len;  // in bytes
;   uint32 type;
;   unint32 flags;
; };
;
; See http://wiki.osdev.org/Detecting_Memory_(x86) and
; http://www.ctyme.com/intr/rb-1741.htm
_memory_map_addr:
  dw 0x0500


_space:
  db ' ', 0

_crlf:
  db 13, 10, 0  ; BIOS printing requires explicit \r

;
; End of boot sector.  The last two bytes on the boot sector must end with
; 0xaa55.
;

times 510-($-$$) db 0
dw 0xaa55

; -----------------------------------------------------------------------------

; This sector is loaded into 0x7e00.
second_sector:

;
; Code
;

%include "src/os/boot/a20_16.asm"
%include "src/os/boot/detect_memory16.asm"
%include "src/os/boot/print_reg16.asm"
%include "src/os/boot/sleep16.asm"

;
; Shared global variables
;

_second_stage_msg:
  db "Entered second stage booting", 13, 10, 0

; Unlike the boot sector entry point, the second stage entry point can be
; anywhere.
second_stage_entry_point:
  mov si, _second_stage_msg
  call print_str

; We want to use more than 1MB of memory ...
call enable_a20

; Find out how much memory do we have.
call detect_memory

;
; To be continue ...
;

jmp halt

; VBoxManage requires the generate bin file to be multiples of 512 bytes ...
times (512*(SECOND_STAGE_SECTORS+1))-($-$$) db 0

