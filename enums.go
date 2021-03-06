package libbpf

// enums required by libbpf APIs

type Enum_bpf_attach_type int

const (
	BPF_CGROUP_INET_INGRESS Enum_bpf_attach_type = iota
	BPF_CGROUP_INET_EGRESS
	BPF_CGROUP_INET_SOCK_CREATE
	BPF_CGROUP_SOCK_OPS
	BPF_SK_SKB_STREAM_PARSER
	BPF_SK_SKB_STREAM_VERDICT
	BPF_CGROUP_DEVICE
	BPF_SK_MSG_VERDICT
	BPF_CGROUP_INET4_BIND
	BPF_CGROUP_INET6_BIND
	BPF_CGROUP_INET4_CONNECT
	BPF_CGROUP_INET6_CONNECT
	BPF_CGROUP_INET4_POST_BIND
	BPF_CGROUP_INET6_POST_BIND
	BPF_CGROUP_UDP4_SENDMSG
	BPF_CGROUP_UDP6_SENDMSG
	BPF_LIRC_MODE2
	BPF_FLOW_DISSECTOR
	BPF_CGROUP_SYSCTL
	BPF_CGROUP_UDP4_RECVMSG
	BPF_CGROUP_UDP6_RECVMSG
	BPF_CGROUP_GETSOCKOPT
	BPF_CGROUP_SETSOCKOPT
	BPF_TRACE_RAW_TP
	BPF_TRACE_FENTRY
	BPF_TRACE_FEXIT
	BPF_MODIFY_RETURN
	BPF_LSM_MAC
	BPF_TRACE_ITER
	BPF_CGROUP_INET4_GETPEERNAME
	BPF_CGROUP_INET6_GETPEERNAME
	BPF_CGROUP_INET4_GETSOCKNAME
	BPF_CGROUP_INET6_GETSOCKNAME
	BPF_XDP_DEVMAP
	BPF_CGROUP_INET_SOCK_RELEASE
	BPF_XDP_CPUMAP
	BPF_SK_LOOKUP
	BPF_XDP
	__MAX_BPF_ATTACH_TYPE
)

type Enum_bpf_func_id int

const (
	BPF_FUNC_unspec Enum_bpf_func_id = iota
	BPF_FUNC_map_lookup_elem
	BPF_FUNC_map_update_elem
	BPF_FUNC_map_delete_elem
	BPF_FUNC_probe_read
	BPF_FUNC_ktime_get_ns
	BPF_FUNC_trace_printk
	BPF_FUNC_get_prandom_u32
	BPF_FUNC_get_smp_processor_id
	BPF_FUNC_skb_store_bytes
	BPF_FUNC_l3_csum_replace
	BPF_FUNC_l4_csum_replace
	BPF_FUNC_tail_call
	BPF_FUNC_clone_redirect
	BPF_FUNC_get_current_pid_tgid
	BPF_FUNC_get_current_uid_gid
	BPF_FUNC_get_current_comm
	BPF_FUNC_get_cgroup_classid
	BPF_FUNC_skb_vlan_push
	BPF_FUNC_skb_vlan_pop
	BPF_FUNC_skb_get_tunnel_key
	BPF_FUNC_skb_set_tunnel_key
	BPF_FUNC_perf_event_read
	BPF_FUNC_redirect
	BPF_FUNC_get_route_realm
	BPF_FUNC_perf_event_output
	BPF_FUNC_skb_load_bytes
	BPF_FUNC_get_stackid
	BPF_FUNC_csum_diff
	BPF_FUNC_skb_get_tunnel_opt
	BPF_FUNC_skb_set_tunnel_opt
	BPF_FUNC_skb_change_proto
	BPF_FUNC_skb_change_type
	BPF_FUNC_skb_under_cgroup
	BPF_FUNC_get_hash_recalc
	BPF_FUNC_get_current_task
	BPF_FUNC_probe_write_user
	BPF_FUNC_current_task_under_cgroup
	BPF_FUNC_skb_change_tail
	BPF_FUNC_skb_pull_data
	BPF_FUNC_csum_update
	BPF_FUNC_set_hash_invalid
	BPF_FUNC_get_numa_node_id
	BPF_FUNC_skb_change_head
	BPF_FUNC_xdp_adjust_head
	BPF_FUNC_probe_read_str
	BPF_FUNC_get_socket_cookie
	BPF_FUNC_get_socket_uid
	BPF_FUNC_set_hash
	BPF_FUNC_setsockopt
	BPF_FUNC_skb_adjust_room
	BPF_FUNC_redirect_map
	BPF_FUNC_sk_redirect_map
	BPF_FUNC_sock_map_update
	BPF_FUNC_xdp_adjust_meta
	BPF_FUNC_perf_event_read_value
	BPF_FUNC_perf_prog_read_value
	BPF_FUNC_getsockopt
	BPF_FUNC_override_return
	BPF_FUNC_sock_ops_cb_flags_set
	BPF_FUNC_msg_redirect_map
	BPF_FUNC_msg_apply_bytes
	BPF_FUNC_msg_cork_bytes
	BPF_FUNC_msg_pull_data
	BPF_FUNC_bind
	BPF_FUNC_xdp_adjust_tail
	BPF_FUNC_skb_get_xfrm_state
	BPF_FUNC_get_stack
	BPF_FUNC_skb_load_bytes_relative
	BPF_FUNC_fib_lookup
	BPF_FUNC_sock_hash_update
	BPF_FUNC_msg_redirect_hash
	BPF_FUNC_sk_redirect_hash
	BPF_FUNC_lwt_push_encap
	BPF_FUNC_lwt_seg6_store_bytes
	BPF_FUNC_lwt_seg6_adjust_srh
	BPF_FUNC_lwt_seg6_action
	BPF_FUNC_rc_repeat
	BPF_FUNC_rc_keydown
	BPF_FUNC_skb_cgroup_id
	BPF_FUNC_get_current_cgroup_id
	BPF_FUNC_get_local_storage
	BPF_FUNC_sk_select_reuseport
	BPF_FUNC_skb_ancestor_cgroup_id
	BPF_FUNC_sk_lookup_tcp
	BPF_FUNC_sk_lookup_udp
	BPF_FUNC_sk_release
	BPF_FUNC_map_push_elem
	BPF_FUNC_map_pop_elem
	BPF_FUNC_map_peek_elem
	BPF_FUNC_msg_push_data
	BPF_FUNC_msg_pop_data
	BPF_FUNC_rc_pointer_rel
	BPF_FUNC_spin_lock
	BPF_FUNC_spin_unlock
	BPF_FUNC_sk_fullsock
	BPF_FUNC_tcp_sock
	BPF_FUNC_skb_ecn_set_ce
	BPF_FUNC_get_listener_sock
	BPF_FUNC_skc_lookup_tcp
	BPF_FUNC_tcp_check_syncookie
	BPF_FUNC_sysctl_get_name
	BPF_FUNC_sysctl_get_current_value
	BPF_FUNC_sysctl_get_new_value
	BPF_FUNC_sysctl_set_new_value
	BPF_FUNC_strtol
	BPF_FUNC_strtoul
	BPF_FUNC_sk_storage_get
	BPF_FUNC_sk_storage_delete
	BPF_FUNC_send_signal
	BPF_FUNC_tcp_gen_syncookie
	BPF_FUNC_skb_output
	BPF_FUNC_probe_read_user
	BPF_FUNC_probe_read_kernel
	BPF_FUNC_probe_read_user_str
	BPF_FUNC_probe_read_kernel_str
	BPF_FUNC_tcp_send_ack
	BPF_FUNC_send_signal_thread
	BPF_FUNC_jiffies64
	BPF_FUNC_read_branch_records
	BPF_FUNC_get_ns_current_pid_tgid
	BPF_FUNC_xdp_output
	BPF_FUNC_get_netns_cookie
	BPF_FUNC_get_current_ancestor_cgroup_id
	BPF_FUNC_sk_assign
	BPF_FUNC_ktime_get_boot_ns
	BPF_FUNC_seq_printf
	BPF_FUNC_seq_write
	BPF_FUNC_sk_cgroup_id
	BPF_FUNC_sk_ancestor_cgroup_id
	BPF_FUNC_ringbuf_output
	BPF_FUNC_ringbuf_reserve
	BPF_FUNC_ringbuf_submit
	BPF_FUNC_ringbuf_discard
	BPF_FUNC_ringbuf_query
	BPF_FUNC_csum_level
	BPF_FUNC_skc_to_tcp6_sock
	BPF_FUNC_skc_to_tcp_sock
	BPF_FUNC_skc_to_tcp_timewait_sock
	BPF_FUNC_skc_to_tcp_request_sock
	BPF_FUNC_skc_to_udp6_sock
	BPF_FUNC_get_task_stack
	BPF_FUNC_load_hdr_opt
	BPF_FUNC_store_hdr_opt
	BPF_FUNC_reserve_hdr_opt
	BPF_FUNC_inode_storage_get
	BPF_FUNC_inode_storage_delete
	BPF_FUNC_d_path
	BPF_FUNC_copy_from_user
	BPF_FUNC_snprintf_btf
	BPF_FUNC_seq_printf_btf
	BPF_FUNC_skb_cgroup_classid
	BPF_FUNC_redirect_neigh
	BPF_FUNC_per_cpu_ptr
	BPF_FUNC_this_cpu_ptr
	BPF_FUNC_redirect_peer
	BPF_FUNC_task_storage_get
	BPF_FUNC_task_storage_delete
	BPF_FUNC_get_current_task_btf
	BPF_FUNC_bprm_opts_set
	BPF_FUNC_ktime_get_coarse_ns
	BPF_FUNC_ima_inode_hash
	BPF_FUNC_sock_from_file
	__BPF_FUNC_MAX_ID
)

type Enum_bpf_map_type int

const (
	BPF_MAP_TYPE_UNSPEC Enum_bpf_map_type = iota
	BPF_MAP_TYPE_HASH
	BPF_MAP_TYPE_ARRAY
	BPF_MAP_TYPE_PROG_ARRAY
	BPF_MAP_TYPE_PERF_EVENT_ARRAY
	BPF_MAP_TYPE_PERCPU_HASH
	BPF_MAP_TYPE_PERCPU_ARRAY
	BPF_MAP_TYPE_STACK_TRACE
	BPF_MAP_TYPE_CGROUP_ARRAY
	BPF_MAP_TYPE_LRU_HASH
	BPF_MAP_TYPE_LRU_PERCPU_HASH
	BPF_MAP_TYPE_LPM_TRIE
	BPF_MAP_TYPE_ARRAY_OF_MAPS
	BPF_MAP_TYPE_HASH_OF_MAPS
	BPF_MAP_TYPE_DEVMAP
	BPF_MAP_TYPE_SOCKMAP
	BPF_MAP_TYPE_CPUMAP
	BPF_MAP_TYPE_XSKMAP
	BPF_MAP_TYPE_SOCKHASH
	BPF_MAP_TYPE_CGROUP_STORAGE
	BPF_MAP_TYPE_REUSEPORT_SOCKARRAY
	BPF_MAP_TYPE_PERCPU_CGROUP_STORAGE
	BPF_MAP_TYPE_QUEUE
	BPF_MAP_TYPE_STACK
	BPF_MAP_TYPE_SK_STORAGE
	BPF_MAP_TYPE_DEVMAP_HASH
	BPF_MAP_TYPE_STRUCT_OPS
	BPF_MAP_TYPE_RINGBUF
	BPF_MAP_TYPE_INODE_STORAGE
	BPF_MAP_TYPE_TASK_STORAGE
)

type Enum_bpf_perf_event_ret int

const (
	LIBBPF_PERF_EVENT_DONE  Enum_bpf_perf_event_ret = 0
	LIBBPF_PERF_EVENT_ERROR Enum_bpf_perf_event_ret = -1
	LIBBPF_PERF_EVENT_CONT  Enum_bpf_perf_event_ret = -2
)

type Enum_bpf_prog_type int

const (
	BPF_PROG_TYPE_UNSPEC Enum_bpf_prog_type = iota
	BPF_PROG_TYPE_SOCKET_FILTER
	BPF_PROG_TYPE_KPROBE
	BPF_PROG_TYPE_SCHED_CLS
	BPF_PROG_TYPE_SCHED_ACT
	BPF_PROG_TYPE_TRACEPOINT
	BPF_PROG_TYPE_XDP
	BPF_PROG_TYPE_PERF_EVENT
	BPF_PROG_TYPE_CGROUP_SKB
	BPF_PROG_TYPE_CGROUP_SOCK
	BPF_PROG_TYPE_LWT_IN
	BPF_PROG_TYPE_LWT_OUT
	BPF_PROG_TYPE_LWT_XMIT
	BPF_PROG_TYPE_SOCK_OPS
	BPF_PROG_TYPE_SK_SKB
	BPF_PROG_TYPE_CGROUP_DEVICE
	BPF_PROG_TYPE_SK_MSG
	BPF_PROG_TYPE_RAW_TRACEPOINT
	BPF_PROG_TYPE_CGROUP_SOCK_ADDR
	BPF_PROG_TYPE_LWT_SEG6LOCAL
	BPF_PROG_TYPE_LIRC_MODE2
	BPF_PROG_TYPE_SK_REUSEPORT
	BPF_PROG_TYPE_FLOW_DISSECTOR
	BPF_PROG_TYPE_CGROUP_SYSCTL
	BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE
	BPF_PROG_TYPE_CGROUP_SOCKOPT
	BPF_PROG_TYPE_TRACING
	BPF_PROG_TYPE_STRUCT_OPS
	BPF_PROG_TYPE_EXT
	BPF_PROG_TYPE_LSM
	BPF_PROG_TYPE_SK_LOOKUP
)

type Enum_bpf_stats_type int

const (
	BPF_STATS_RUN_TIME Enum_bpf_stats_type = 0
)

type Enum_btf_endianness int

const (
	BTF_LITTLE_ENDIAN Enum_btf_endianness = 0
	BTF_BIG_ENDIAN    Enum_btf_endianness = 1
)

type Enum_btf_func_linkage int

const (
	BTF_FUNC_STATIC Enum_btf_func_linkage = 0
	BTF_FUNC_GLOBAL Enum_btf_func_linkage = 1
	BTF_FUNC_EXTERN Enum_btf_func_linkage = 2
)

type Enum_btf_fwd_kind int

const (
	BTF_FWD_STRUCT Enum_btf_fwd_kind = 0
	BTF_FWD_UNION  Enum_btf_fwd_kind = 1
	BTF_FWD_ENUM   Enum_btf_fwd_kind = 2
)
