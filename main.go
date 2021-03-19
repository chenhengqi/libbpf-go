package main

// #cgo CFLAGS: -I${SRCDIR}/.output
// #cgo LDFLAGS: -L${SRCDIR}/.output -lbpf -lelf -lz
// #include <strings.h>
// #include <bpf/libbpf.h>
//
// static bool str_has_suffix(const char *str, const char *suffix)
// {
//         size_t i, n1 = strlen(str), n2 = strlen(suffix);
//
//         if (n1 < n2)
//             return false;
//
//         for (i = 0; i < n2; i++) {
//             if (str[n1 - i - 1] != suffix[n2 - i - 1])
//                return false;
//         }
//
//         return true;
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
	bpfObj := C.bpf_object__open_mem(C.CBytes(data), C.ulong(bpfObjFile.Size()), &opts)

	// get maps
	var m *C.struct_bpf_map
	m = C.bpf_map__next(m, bpfObj)
	for m != nil {
		name := C.get_map_ident(m)
		fmt.Println(name)
		m = C.bpf_map__next(m, bpfObj)
	}
}
