package libbpf

// #include <bpf/btf.h>
import "C"


	LIBBPF_API void btf__free(struct btf *btf);

	LIBBPF_API struct btf *btf__new(const void *data, __u32 size);

	LIBBPF_API struct btf *btf__new_split(const void *data, __u32 size, struct btf *base_btf);

	LIBBPF_API struct btf *btf__new_empty(void);

	LIBBPF_API struct btf *btf__new_empty_split(struct btf *base_btf);

	LIBBPF_API struct btf *btf__parse(const char *path, struct btf_ext **btf_ext);

	LIBBPF_API struct btf *btf__parse_split(const char *path, struct btf *base_btf);

	LIBBPF_API struct btf *btf__parse_elf(const char *path, struct btf_ext **btf_ext);

	LIBBPF_API struct btf *btf__parse_elf_split(const char *path, struct btf *base_btf);

	LIBBPF_API struct btf *btf__parse_raw(const char *path);

	LIBBPF_API struct btf *btf__parse_raw_split(const char *path, struct btf *base_btf);

	LIBBPF_API int btf__finalize_data(struct bpf_object *obj, struct btf *btf);

	LIBBPF_API int btf__load(struct btf *btf);

	LIBBPF_API __s32 btf__find_by_name(const struct btf *btf,
				   const char *type_name);

	LIBBPF_API __s32 btf__find_by_name_kind(const struct btf *btf,
					const char *type_name, __u32 kind);

	LIBBPF_API __u32 btf__get_nr_types(const struct btf *btf);

	LIBBPF_API const struct btf *btf__base_btf(const struct btf *btf);

	LIBBPF_API const struct btf_type *btf__type_by_id(const struct btf *btf,
						  __u32 id);

	LIBBPF_API size_t btf__pointer_size(const struct btf *btf);

	LIBBPF_API int btf__set_pointer_size(struct btf *btf, size_t ptr_sz);

	LIBBPF_API enum btf_endianness btf__endianness(const struct btf *btf);

	LIBBPF_API int btf__set_endianness(struct btf *btf, enum btf_endianness endian);

	LIBBPF_API __s64 btf__resolve_size(const struct btf *btf, __u32 type_id);

	LIBBPF_API int btf__resolve_type(const struct btf *btf, __u32 type_id);

	LIBBPF_API int btf__align_of(const struct btf *btf, __u32 id);

	LIBBPF_API int btf__fd(const struct btf *btf);

	LIBBPF_API void btf__set_fd(struct btf *btf, int fd);

	LIBBPF_API const void *btf__get_raw_data(const struct btf *btf, __u32 *size);

	LIBBPF_API const char *btf__name_by_offset(const struct btf *btf, __u32 offset);

	LIBBPF_API const char *btf__str_by_offset(const struct btf *btf, __u32 offset);

	LIBBPF_API int btf__get_from_id(__u32 id, struct btf **btf);

	LIBBPF_API int btf__get_map_kv_tids(const struct btf *btf, const char *map_name,
				    __u32 expected_key_size,
				    __u32 expected_value_size,
				    __u32 *key_type_id, __u32 *value_type_id);

	LIBBPF_API struct btf_ext *btf_ext__new(__u8 *data, __u32 size);

	LIBBPF_API void btf_ext__free(struct btf_ext *btf_ext);

	LIBBPF_API const void *btf_ext__get_raw_data(const struct btf_ext *btf_ext,
					     __u32 *size);

	LIBBPF_API LIBBPF_DEPRECATED("btf_ext__reloc_func_info was never meant as a public API and has wrong assumptions embedded in it;

	LIBBPF_API LIBBPF_DEPRECATED("btf_ext__reloc_line_info was never meant as a public API and has wrong assumptions embedded in it;

	LIBBPF_API __u32 btf_ext__func_info_rec_size(const struct btf_ext *btf_ext);

	LIBBPF_API __u32 btf_ext__line_info_rec_size(const struct btf_ext *btf_ext);

	LIBBPF_API struct btf *libbpf_find_kernel_btf(void);

	LIBBPF_API int btf__find_str(struct btf *btf, const char *s);

	LIBBPF_API int btf__add_str(struct btf *btf, const char *s);

	LIBBPF_API int btf__add_type(struct btf *btf, const struct btf *src_btf,
			     const struct btf_type *src_type);

	LIBBPF_API int btf__add_int(struct btf *btf, const char *name, size_t byte_sz, int encoding);

	LIBBPF_API int btf__add_float(struct btf *btf, const char *name, size_t byte_sz);

	LIBBPF_API int btf__add_ptr(struct btf *btf, int ref_type_id);

	LIBBPF_API int btf__add_array(struct btf *btf,
			      int index_type_id, int elem_type_id, __u32 nr_elems);

	LIBBPF_API int btf__add_struct(struct btf *btf, const char *name, __u32 sz);

	LIBBPF_API int btf__add_union(struct btf *btf, const char *name, __u32 sz);

	LIBBPF_API int btf__add_field(struct btf *btf, const char *name, int field_type_id,
			      __u32 bit_offset, __u32 bit_size);

	LIBBPF_API int btf__add_enum(struct btf *btf, const char *name, __u32 bytes_sz);

	LIBBPF_API int btf__add_enum_value(struct btf *btf, const char *name, __s64 value);

	LIBBPF_API int btf__add_fwd(struct btf *btf, const char *name, enum btf_fwd_kind fwd_kind);

	LIBBPF_API int btf__add_typedef(struct btf *btf, const char *name, int ref_type_id);

	LIBBPF_API int btf__add_volatile(struct btf *btf, int ref_type_id);

	LIBBPF_API int btf__add_const(struct btf *btf, int ref_type_id);

	LIBBPF_API int btf__add_restrict(struct btf *btf, int ref_type_id);

	LIBBPF_API int btf__add_func(struct btf *btf, const char *name,
			     enum btf_func_linkage linkage, int proto_type_id);

	LIBBPF_API int btf__add_func_proto(struct btf *btf, int ret_type_id);

	LIBBPF_API int btf__add_func_param(struct btf *btf, const char *name, int type_id);

	LIBBPF_API int btf__add_var(struct btf *btf, const char *name, int linkage, int type_id);

	LIBBPF_API int btf__add_datasec(struct btf *btf, const char *name, __u32 byte_sz);

	LIBBPF_API int btf__add_datasec_var_info(struct btf *btf, int var_type_id,
					 __u32 offset, __u32 byte_sz);

	LIBBPF_API int btf__dedup(struct btf *btf, struct btf_ext *btf_ext,
			  const struct btf_dedup_opts *opts);

	LIBBPF_API struct btf_dump *btf_dump__new(const struct btf *btf,
					  const struct btf_ext *btf_ext,
					  const struct btf_dump_opts *opts,
					  btf_dump_printf_fn_t printf_fn);

	LIBBPF_API void btf_dump__free(struct btf_dump *d);

	LIBBPF_API int btf_dump__dump_type(struct btf_dump *d, __u32 id);

	LIBBPF_API int
btf_dump__emit_type_decl(struct btf_dump *d, __u32 id,
			 const struct btf_dump_emit_type_decl_opts *opts);

