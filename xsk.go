// AUTOGENERATED BY LIBBPF-GO-GEN
// DO NOT EDIT
package libbpf

// #include <bpf/xsk.h>
import "C"
import "unsafe"

var _ unsafe.Pointer

// Xsk_umem__fd is a wrapper of xsk_umem__fd
//     int xsk_umem__fd(const struct xsk_umem *umem)
func Xsk_umem__fd(umem *Struct_xsk_umem) int32 {
	return C.xsk_umem__fd()
}

// Xsk_socket__fd is a wrapper of xsk_socket__fd
//     int xsk_socket__fd(const struct xsk_socket *xsk)
func Xsk_socket__fd(xsk *Struct_xsk_socket) int32 {
	return C.xsk_socket__fd()
}

// Xsk_setup_xdp_prog is a wrapper of xsk_setup_xdp_prog
//     int xsk_setup_xdp_prog(int ifindex, int *xsks_map_fd)
func Xsk_setup_xdp_prog(ifindex int32, xsks_map_fd *int32) int32 {
	return C.xsk_setup_xdp_prog()
}

// Xsk_socket__update_xskmap is a wrapper of xsk_socket__update_xskmap
//     int xsk_socket__update_xskmap(struct xsk_socket *xsk, int xsks_map_fd)
func Xsk_socket__update_xskmap(xsk *Struct_xsk_socket, xsks_map_fd int32) int32 {
	return C.xsk_socket__update_xskmap()
}

// Xsk_umem__create is a wrapper of xsk_umem__create
//     int xsk_umem__create(struct xsk_umem **umem, void *umem_area, __u64 size, struct xsk_ring_prod *fill, struct xsk_ring_cons *comp, const struct xsk_umem_config *config)
func Xsk_umem__create(umem *Struct_xsk_umem, umem_area unsafe.Pointer, size uint64, fill *Struct_xsk_ring_prod, comp *Struct_xsk_ring_cons, config *Struct_xsk_umem_config) int32 {
	return C.xsk_umem__create()
}

// Xsk_umem__create_v0_0_2 is a wrapper of xsk_umem__create_v0_0_2
//     int xsk_umem__create_v0_0_2(struct xsk_umem **umem, void *umem_area, __u64 size, struct xsk_ring_prod *fill, struct xsk_ring_cons *comp, const struct xsk_umem_config *config)
func Xsk_umem__create_v0_0_2(umem *Struct_xsk_umem, umem_area unsafe.Pointer, size uint64, fill *Struct_xsk_ring_prod, comp *Struct_xsk_ring_cons, config *Struct_xsk_umem_config) int32 {
	return C.xsk_umem__create_v0_0_2()
}

// Xsk_umem__create_v0_0_4 is a wrapper of xsk_umem__create_v0_0_4
//     int xsk_umem__create_v0_0_4(struct xsk_umem **umem, void *umem_area, __u64 size, struct xsk_ring_prod *fill, struct xsk_ring_cons *comp, const struct xsk_umem_config *config)
func Xsk_umem__create_v0_0_4(umem *Struct_xsk_umem, umem_area unsafe.Pointer, size uint64, fill *Struct_xsk_ring_prod, comp *Struct_xsk_ring_cons, config *Struct_xsk_umem_config) int32 {
	return C.xsk_umem__create_v0_0_4()
}

// Xsk_socket__create is a wrapper of xsk_socket__create
//     int xsk_socket__create(struct xsk_socket **xsk, const char *ifname, __u32 queue_id, struct xsk_umem *umem, struct xsk_ring_cons *rx, struct xsk_ring_prod *tx, const struct xsk_socket_config *config)
func Xsk_socket__create(xsk *Struct_xsk_socket, ifname string, queue_id uint32, umem *Struct_xsk_umem, rx *Struct_xsk_ring_cons, tx *Struct_xsk_ring_prod, config *Struct_xsk_socket_config) int32 {
	return C.xsk_socket__create()
}

// Xsk_socket__create_shared is a wrapper of xsk_socket__create_shared
//     int xsk_socket__create_shared(struct xsk_socket **xsk_ptr, const char *ifname, __u32 queue_id, struct xsk_umem *umem, struct xsk_ring_cons *rx, struct xsk_ring_prod *tx, struct xsk_ring_prod *fill, struct xsk_ring_cons *comp, const struct xsk_socket_config *config)
func Xsk_socket__create_shared(xsk_ptr *Struct_xsk_socket, ifname string, queue_id uint32, umem *Struct_xsk_umem, rx *Struct_xsk_ring_cons, tx *Struct_xsk_ring_prod, fill *Struct_xsk_ring_prod, comp *Struct_xsk_ring_cons, config *Struct_xsk_socket_config) int32 {
	return C.xsk_socket__create_shared()
}

// Xsk_umem__delete is a wrapper of xsk_umem__delete
//     int xsk_umem__delete(struct xsk_umem *umem)
func Xsk_umem__delete(umem *Struct_xsk_umem) int32 {
	return C.xsk_umem__delete()
}

// Xsk_socket__delete is a wrapper of xsk_socket__delete
//     void xsk_socket__delete(struct xsk_socket *xsk)
func Xsk_socket__delete(xsk *Struct_xsk_socket)  {
	 C.xsk_socket__delete()
}
