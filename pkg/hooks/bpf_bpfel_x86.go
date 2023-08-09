// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64
// +build 386 amd64

package hooks

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

// loadBpf returns the embedded CollectionSpec for bpf.
func loadBpf() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_BpfBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load bpf: %w", err)
	}

	return spec, err
}

// loadBpfObjects loads bpf and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*bpfObjects
//	*bpfPrograms
//	*bpfMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadBpfObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadBpf()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// bpfSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfSpecs struct {
	bpfProgramSpecs
	bpfMapSpecs
}

// bpfSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfProgramSpecs struct {
	K_connect4                       *ebpf.ProgramSpec `ebpf:"k_connect4"`
	K_connect6                       *ebpf.ProgramSpec `ebpf:"k_connect6"`
	K_getpeername4                   *ebpf.ProgramSpec `ebpf:"k_getpeername4"`
	K_getpeername6                   *ebpf.ProgramSpec `ebpf:"k_getpeername6"`
	SyscallProbeEntryAccept          *ebpf.ProgramSpec `ebpf:"syscall__probe_entry_accept"`
	SyscallProbeEntryAccept4         *ebpf.ProgramSpec `ebpf:"syscall__probe_entry_accept4"`
	SyscallProbeEntryBind            *ebpf.ProgramSpec `ebpf:"syscall__probe_entry_bind"`
	SyscallProbeEntryClose           *ebpf.ProgramSpec `ebpf:"syscall__probe_entry_close"`
	SyscallProbeEntryRead            *ebpf.ProgramSpec `ebpf:"syscall__probe_entry_read"`
	SyscallProbeEntryRecvfrom        *ebpf.ProgramSpec `ebpf:"syscall__probe_entry_recvfrom"`
	SyscallProbeEntrySendto          *ebpf.ProgramSpec `ebpf:"syscall__probe_entry_sendto"`
	SyscallProbeEntryTcpV4Connect    *ebpf.ProgramSpec `ebpf:"syscall__probe_entry_tcp_v4_connect"`
	SyscallProbeEntryTcpV4PreConnect *ebpf.ProgramSpec `ebpf:"syscall__probe_entry_tcp_v4_pre_connect"`
	SyscallProbeEntryTcpV6Connect    *ebpf.ProgramSpec `ebpf:"syscall__probe_entry_tcp_v6_connect"`
	SyscallProbeEntryTcpV6PreConnect *ebpf.ProgramSpec `ebpf:"syscall__probe_entry_tcp_v6_pre_connect"`
	SyscallProbeEntryUdpPreConnect   *ebpf.ProgramSpec `ebpf:"syscall__probe_entry_udp_pre_connect"`
	SyscallProbeEntryWrite           *ebpf.ProgramSpec `ebpf:"syscall__probe_entry_write"`
	SyscallProbeEntryWritev          *ebpf.ProgramSpec `ebpf:"syscall__probe_entry_writev"`
	SyscallProbeRetAccept            *ebpf.ProgramSpec `ebpf:"syscall__probe_ret_accept"`
	SyscallProbeRetAccept4           *ebpf.ProgramSpec `ebpf:"syscall__probe_ret_accept4"`
	SyscallProbeRetClose             *ebpf.ProgramSpec `ebpf:"syscall__probe_ret_close"`
	SyscallProbeRetRead              *ebpf.ProgramSpec `ebpf:"syscall__probe_ret_read"`
	SyscallProbeRetRecvfrom          *ebpf.ProgramSpec `ebpf:"syscall__probe_ret_recvfrom"`
	SyscallProbeRetSendto            *ebpf.ProgramSpec `ebpf:"syscall__probe_ret_sendto"`
	SyscallProbeRetTcpV4Connect      *ebpf.ProgramSpec `ebpf:"syscall__probe_ret_tcp_v4_connect"`
	SyscallProbeRetTcpV6Connect      *ebpf.ProgramSpec `ebpf:"syscall__probe_ret_tcp_v6_connect"`
	SyscallProbeRetWrite             *ebpf.ProgramSpec `ebpf:"syscall__probe_ret_write"`
	SyscallProbeRetWritev            *ebpf.ProgramSpec `ebpf:"syscall__probe_ret_writev"`
}

// bpfMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfMapSpecs struct {
	ActiveAcceptArgsMap       *ebpf.MapSpec `ebpf:"active_accept_args_map"`
	ActiveCloseArgsMap        *ebpf.MapSpec `ebpf:"active_close_args_map"`
	ActiveReadArgsMap         *ebpf.MapSpec `ebpf:"active_read_args_map"`
	ActiveWriteArgsMap        *ebpf.MapSpec `ebpf:"active_write_args_map"`
	AppPidMap                 *ebpf.MapSpec `ebpf:"app_pid_map"`
	ConnInfoMap               *ebpf.MapSpec `ebpf:"conn_info_map"`
	CurrentSockMap            *ebpf.MapSpec `ebpf:"current_sock_map"`
	DestInfoMap               *ebpf.MapSpec `ebpf:"dest_info_map"`
	DockerCmdMap              *ebpf.MapSpec `ebpf:"docker_cmd_map"`
	GlobalNsPidInfoMap        *ebpf.MapSpec `ebpf:"global_nsPid_info_map"`
	InodeMap                  *ebpf.MapSpec `ebpf:"inode_map"`
	KeployKernelPidMap        *ebpf.MapSpec `ebpf:"keploy_kernel_pid_map"`
	KeployModeMap             *ebpf.MapSpec `ebpf:"keploy_mode_map"`
	KeployNamespacePidMap     *ebpf.MapSpec `ebpf:"keploy_namespace_pid_map"`
	KeployNsPidInfoMap        *ebpf.MapSpec `ebpf:"keploy_nsPid_info_map"`
	ProxyInfoMap              *ebpf.MapSpec `ebpf:"proxy_info_map"`
	RedirectProxyMap          *ebpf.MapSpec `ebpf:"redirect_proxy_map"`
	SocketCloseEvents         *ebpf.MapSpec `ebpf:"socket_close_events"`
	SocketDataEventBufferHeap *ebpf.MapSpec `ebpf:"socket_data_event_buffer_heap"`
	SocketDataEvents          *ebpf.MapSpec `ebpf:"socket_data_events"`
	SocketOpenEvents          *ebpf.MapSpec `ebpf:"socket_open_events"`
	TaskStructMap             *ebpf.MapSpec `ebpf:"task_struct_map"`
}

// bpfObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfObjects struct {
	bpfPrograms
	bpfMaps
}

func (o *bpfObjects) Close() error {
	return _BpfClose(
		&o.bpfPrograms,
		&o.bpfMaps,
	)
}

// bpfMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfMaps struct {
	ActiveAcceptArgsMap       *ebpf.Map `ebpf:"active_accept_args_map"`
	ActiveCloseArgsMap        *ebpf.Map `ebpf:"active_close_args_map"`
	ActiveReadArgsMap         *ebpf.Map `ebpf:"active_read_args_map"`
	ActiveWriteArgsMap        *ebpf.Map `ebpf:"active_write_args_map"`
	AppPidMap                 *ebpf.Map `ebpf:"app_pid_map"`
	ConnInfoMap               *ebpf.Map `ebpf:"conn_info_map"`
	CurrentSockMap            *ebpf.Map `ebpf:"current_sock_map"`
	DestInfoMap               *ebpf.Map `ebpf:"dest_info_map"`
	DockerCmdMap              *ebpf.Map `ebpf:"docker_cmd_map"`
	GlobalNsPidInfoMap        *ebpf.Map `ebpf:"global_nsPid_info_map"`
	InodeMap                  *ebpf.Map `ebpf:"inode_map"`
	KeployKernelPidMap        *ebpf.Map `ebpf:"keploy_kernel_pid_map"`
	KeployModeMap             *ebpf.Map `ebpf:"keploy_mode_map"`
	KeployNamespacePidMap     *ebpf.Map `ebpf:"keploy_namespace_pid_map"`
	KeployNsPidInfoMap        *ebpf.Map `ebpf:"keploy_nsPid_info_map"`
	ProxyInfoMap              *ebpf.Map `ebpf:"proxy_info_map"`
	RedirectProxyMap          *ebpf.Map `ebpf:"redirect_proxy_map"`
	SocketCloseEvents         *ebpf.Map `ebpf:"socket_close_events"`
	SocketDataEventBufferHeap *ebpf.Map `ebpf:"socket_data_event_buffer_heap"`
	SocketDataEvents          *ebpf.Map `ebpf:"socket_data_events"`
	SocketOpenEvents          *ebpf.Map `ebpf:"socket_open_events"`
	TaskStructMap             *ebpf.Map `ebpf:"task_struct_map"`
}

func (m *bpfMaps) Close() error {
	return _BpfClose(
		m.ActiveAcceptArgsMap,
		m.ActiveCloseArgsMap,
		m.ActiveReadArgsMap,
		m.ActiveWriteArgsMap,
		m.AppPidMap,
		m.ConnInfoMap,
		m.CurrentSockMap,
		m.DestInfoMap,
		m.DockerCmdMap,
		m.GlobalNsPidInfoMap,
		m.InodeMap,
		m.KeployKernelPidMap,
		m.KeployModeMap,
		m.KeployNamespacePidMap,
		m.KeployNsPidInfoMap,
		m.ProxyInfoMap,
		m.RedirectProxyMap,
		m.SocketCloseEvents,
		m.SocketDataEventBufferHeap,
		m.SocketDataEvents,
		m.SocketOpenEvents,
		m.TaskStructMap,
	)
}

// bpfPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfPrograms struct {
	K_connect4                       *ebpf.Program `ebpf:"k_connect4"`
	K_connect6                       *ebpf.Program `ebpf:"k_connect6"`
	K_getpeername4                   *ebpf.Program `ebpf:"k_getpeername4"`
	K_getpeername6                   *ebpf.Program `ebpf:"k_getpeername6"`
	SyscallProbeEntryAccept          *ebpf.Program `ebpf:"syscall__probe_entry_accept"`
	SyscallProbeEntryAccept4         *ebpf.Program `ebpf:"syscall__probe_entry_accept4"`
	SyscallProbeEntryBind            *ebpf.Program `ebpf:"syscall__probe_entry_bind"`
	SyscallProbeEntryClose           *ebpf.Program `ebpf:"syscall__probe_entry_close"`
	SyscallProbeEntryRead            *ebpf.Program `ebpf:"syscall__probe_entry_read"`
	SyscallProbeEntryRecvfrom        *ebpf.Program `ebpf:"syscall__probe_entry_recvfrom"`
	SyscallProbeEntrySendto          *ebpf.Program `ebpf:"syscall__probe_entry_sendto"`
	SyscallProbeEntryTcpV4Connect    *ebpf.Program `ebpf:"syscall__probe_entry_tcp_v4_connect"`
	SyscallProbeEntryTcpV4PreConnect *ebpf.Program `ebpf:"syscall__probe_entry_tcp_v4_pre_connect"`
	SyscallProbeEntryTcpV6Connect    *ebpf.Program `ebpf:"syscall__probe_entry_tcp_v6_connect"`
	SyscallProbeEntryTcpV6PreConnect *ebpf.Program `ebpf:"syscall__probe_entry_tcp_v6_pre_connect"`
	SyscallProbeEntryUdpPreConnect   *ebpf.Program `ebpf:"syscall__probe_entry_udp_pre_connect"`
	SyscallProbeEntryWrite           *ebpf.Program `ebpf:"syscall__probe_entry_write"`
	SyscallProbeEntryWritev          *ebpf.Program `ebpf:"syscall__probe_entry_writev"`
	SyscallProbeRetAccept            *ebpf.Program `ebpf:"syscall__probe_ret_accept"`
	SyscallProbeRetAccept4           *ebpf.Program `ebpf:"syscall__probe_ret_accept4"`
	SyscallProbeRetClose             *ebpf.Program `ebpf:"syscall__probe_ret_close"`
	SyscallProbeRetRead              *ebpf.Program `ebpf:"syscall__probe_ret_read"`
	SyscallProbeRetRecvfrom          *ebpf.Program `ebpf:"syscall__probe_ret_recvfrom"`
	SyscallProbeRetSendto            *ebpf.Program `ebpf:"syscall__probe_ret_sendto"`
	SyscallProbeRetTcpV4Connect      *ebpf.Program `ebpf:"syscall__probe_ret_tcp_v4_connect"`
	SyscallProbeRetTcpV6Connect      *ebpf.Program `ebpf:"syscall__probe_ret_tcp_v6_connect"`
	SyscallProbeRetWrite             *ebpf.Program `ebpf:"syscall__probe_ret_write"`
	SyscallProbeRetWritev            *ebpf.Program `ebpf:"syscall__probe_ret_writev"`
}

func (p *bpfPrograms) Close() error {
	return _BpfClose(
		p.K_connect4,
		p.K_connect6,
		p.K_getpeername4,
		p.K_getpeername6,
		p.SyscallProbeEntryAccept,
		p.SyscallProbeEntryAccept4,
		p.SyscallProbeEntryBind,
		p.SyscallProbeEntryClose,
		p.SyscallProbeEntryRead,
		p.SyscallProbeEntryRecvfrom,
		p.SyscallProbeEntrySendto,
		p.SyscallProbeEntryTcpV4Connect,
		p.SyscallProbeEntryTcpV4PreConnect,
		p.SyscallProbeEntryTcpV6Connect,
		p.SyscallProbeEntryTcpV6PreConnect,
		p.SyscallProbeEntryUdpPreConnect,
		p.SyscallProbeEntryWrite,
		p.SyscallProbeEntryWritev,
		p.SyscallProbeRetAccept,
		p.SyscallProbeRetAccept4,
		p.SyscallProbeRetClose,
		p.SyscallProbeRetRead,
		p.SyscallProbeRetRecvfrom,
		p.SyscallProbeRetSendto,
		p.SyscallProbeRetTcpV4Connect,
		p.SyscallProbeRetTcpV6Connect,
		p.SyscallProbeRetWrite,
		p.SyscallProbeRetWritev,
	)
}

func _BpfClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed bpf_bpfel_x86.o
var _BpfBytes []byte
