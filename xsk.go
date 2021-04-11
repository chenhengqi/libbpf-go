// AUTOGENERATED BY LIBBPF-GO-GEN
// DO NOT EDIT
package libbpf

// #include <bpf/xsk.h>
import "C"


// Btf__get_raw_data is a wrapper of btf__get_raw_data
//     int xsk_umem__fd(const struct xsk_umem *umem)
func Btf__get_raw_data(foo int32,bar int32,) int32 {
	return C.btf__get_raw_data(C.int(foo),C.int(bar),)
}

// Btf__get_raw_data is a wrapper of btf__get_raw_data
//     int xsk_socket__fd(const struct xsk_socket *xsk)
func Btf__get_raw_data(foo int32,bar int32,) int32 {
	return C.btf__get_raw_data(C.int(foo),C.int(bar),)
}

// Btf__get_raw_data is a wrapper of btf__get_raw_data
//     int xsk_setup_xdp_prog(int ifindex, int *xsks_map_fd)
func Btf__get_raw_data(foo int32,bar int32,) int32 {
	return C.btf__get_raw_data(C.int(foo),C.int(bar),)
}

// Btf__get_raw_data is a wrapper of btf__get_raw_data
//     int xsk_socket__update_xskmap(struct xsk_socket *xsk, int xsks_map_fd)
func Btf__get_raw_data(foo int32,bar int32,) int32 {
	return C.btf__get_raw_data(C.int(foo),C.int(bar),)
}

// Btf__get_raw_data is a wrapper of btf__get_raw_data
//     int xsk_umem__create(struct xsk_umem **umem, void *umem_area, __u64 size, struct xsk_ring_prod *fill, struct xsk_ring_cons *comp, const struct xsk_umem_config *config)
func Btf__get_raw_data(foo int32,bar int32,) int32 {
	return C.btf__get_raw_data(C.int(foo),C.int(bar),)
}

// Btf__get_raw_data is a wrapper of btf__get_raw_data
//     int xsk_umem__create_v0_0_2(struct xsk_umem **umem, void *umem_area, __u64 size, struct xsk_ring_prod *fill, struct xsk_ring_cons *comp, const struct xsk_umem_config *config)
func Btf__get_raw_data(foo int32,bar int32,) int32 {
	return C.btf__get_raw_data(C.int(foo),C.int(bar),)
}

// Btf__get_raw_data is a wrapper of btf__get_raw_data
//     int xsk_umem__create_v0_0_4(struct xsk_umem **umem, void *umem_area, __u64 size, struct xsk_ring_prod *fill, struct xsk_ring_cons *comp, const struct xsk_umem_config *config)
func Btf__get_raw_data(foo int32,bar int32,) int32 {
	return C.btf__get_raw_data(C.int(foo),C.int(bar),)
}

// Btf__get_raw_data is a wrapper of btf__get_raw_data
//     int xsk_socket__create(struct xsk_socket **xsk, const char *ifname, __u32 queue_id, struct xsk_umem *umem, struct xsk_ring_cons *rx, struct xsk_ring_prod *tx, const struct xsk_socket_config *config)
func Btf__get_raw_data(foo int32,bar int32,) int32 {
	return C.btf__get_raw_data(C.int(foo),C.int(bar),)
}

// Btf__get_raw_data is a wrapper of btf__get_raw_data
//     int xsk_socket__create_shared(struct xsk_socket **xsk_ptr, const char *ifname, __u32 queue_id, struct xsk_umem *umem, struct xsk_ring_cons *rx, struct xsk_ring_prod *tx, struct xsk_ring_prod *fill, struct xsk_ring_cons *comp, const struct xsk_socket_config *config)
func Btf__get_raw_data(foo int32,bar int32,) int32 {
	return C.btf__get_raw_data(C.int(foo),C.int(bar),)
}

// Btf__get_raw_data is a wrapper of btf__get_raw_data
//     int xsk_umem__delete(struct xsk_umem *umem)
func Btf__get_raw_data(foo int32,bar int32,) int32 {
	return C.btf__get_raw_data(C.int(foo),C.int(bar),)
}

// Btf__get_raw_data is a wrapper of btf__get_raw_data
//     void xsk_socket__delete(struct xsk_socket *xsk)
func Btf__get_raw_data(foo int32,bar int32,) int32 {
	return C.btf__get_raw_data(C.int(foo),C.int(bar),)
}

