package main

// #cgo CFLAGS: -I${SRCDIR}/.output
// #cgo LDFLAGS: -L${SRCDIR}/.output -lbpf -lelf -lz
// #include <strings.h>
// #include <bpf/bpf.h>
// #include <bpf/btf.h>
// #include <bpf/libbpf.h>
//
// static bool str_has_suffix(const char *str, const char *suffix)
// {
//     size_t i, n1 = strlen(str), n2 = strlen(suffix);
//
//     if (n1 < n2)
//         return false;
//
//     for (i = 0; i < n2; i++) {
//         if (str[n1 - i - 1] != suffix[n2 - i - 1])
//            return false;
//     }
//
//     return true;
// }
//
// static const char *get_map_ident(const struct bpf_map *map)
// {
//     const char *name = bpf_map__name(map);
//
//     if (!bpf_map__is_internal(map))
//         return name;
//
//     if (str_has_suffix(name, ".data"))
//         return "data";
//     else if (str_has_suffix(name, ".rodata"))
//         return "rodata";
//     else if (str_has_suffix(name, ".bss"))
//         return "bss";
//     else if (str_has_suffix(name, ".kconfig"))
//         return "kconfig";
//     else
//         return NULL;
// }
//
// void codegen_btf_dump_printf(void *ctx, const char *fmt, va_list args)
// {
//     vprintf(fmt, args);
// }
import "C"
import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

func main() {
	bpfObjFileName := os.Args[1]
	bpfObjFile, err := os.Stat(bpfObjFileName)
	if err != nil {
		panic(err)
	}

	size := roundUp(int(bpfObjFile.Size()), os.Getpagesize())
	file, err := os.Open(bpfObjFileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := syscall.Mmap(int(file.Fd()), 0, size, syscall.PROT_READ, syscall.MAP_PRIVATE)
	if err != nil {
		panic(err)
	}
	defer syscall.Munmap(data)

	opts := C.struct_bpf_object_open_opts{}
	C.bzero(unsafe.Pointer(&opts), C.sizeof_struct_bpf_object_open_opts)
	opts.object_name = C.CString(objectName(bpfObjFileName))
	opts.sz = C.sizeof_struct_bpf_object_open_opts
	bpfObj := C.bpf_object__open_mem(C.CBytes(data), C.ulong(bpfObjFile.Size()), &opts)

	// get maps
	var m *C.struct_bpf_map
	m = C.bpf_map__next(m, bpfObj)
	for m != nil {
		name := C.get_map_ident(m)
		fmt.Println(C.GoString(name))
		m = C.bpf_map__next(m, bpfObj)
	}

	// get progs
	var p *C.struct_bpf_program
	p = C.bpf_program__next(p, bpfObj)
	for p != nil {
		name := C.bpf_program__name(p)
		fmt.Println(C.GoString(name))
		p = C.bpf_program__next(p, bpfObj)
	}

	btfObj := C.bpf_object__btf(bpfObj)
	var btfExt *C.struct_btf_ext
	var btfOpts *C.struct_btf_dump_opts
	btfDump := C.btf_dump__new(btfObj, btfExt, btfOpts, (*[0]byte)(C.codegen_btf_dump_printf))

	n := int(C.btf__get_nr_types(btfObj))
	for i := 1; i <= n; i++ {
		section := C.btf__type_by_id(btfObj, C.uint(i))
		if !C.btf_is_datasec(section) {
			continue
		}

		var stripMods bool

		sectionName := C.btf__name_by_offset(btfObj, section.name_off)
		switch C.GoString(sectionName) {
		case ".data":
		case ".bss":
		case ".rodata":
			stripMods = true
		case ".kconfig":
		default:
			continue
		}
	}
}
