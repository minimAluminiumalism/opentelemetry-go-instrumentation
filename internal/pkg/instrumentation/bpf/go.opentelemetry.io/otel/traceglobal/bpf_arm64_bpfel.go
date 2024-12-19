// Code generated by bpf2go; DO NOT EDIT.
//go:build arm64

package global

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type bpfOtelSpanT struct {
	StartTime uint64
	EndTime   uint64
	Sc        bpfSpanContext
	Psc       bpfSpanContext
	SpanName  bpfSpanNameT
	Status    struct {
		Code        uint32
		Description struct{ Buf [64]int8 }
	}
	Attributes struct {
		Attrs [16]struct {
			ValLength uint16
			Vtype     uint8
			Reserved  uint8
			Key       [32]int8
			Value     [128]int8
		}
		ValidAttrs uint8
	}
	TracerId bpfTracerIdT
	_        [3]byte
}

type bpfSliceArrayBuff struct{ Buff [1024]uint8 }

type bpfSpanContext struct {
	TraceID    [16]uint8
	SpanID     [8]uint8
	TraceFlags uint8
	Padding    [7]uint8
}

type bpfSpanNameT struct{ Buf [64]int8 }

type bpfTracerIdT struct {
	Name      [128]int8
	Version   [32]int8
	SchemaUrl [128]int8
}

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
	bpfVariableSpecs
}

// bpfProgramSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfProgramSpecs struct {
	UprobeEnd           *ebpf.ProgramSpec `ebpf:"uprobe_End"`
	UprobeSetAttributes *ebpf.ProgramSpec `ebpf:"uprobe_SetAttributes"`
	UprobeSetName       *ebpf.ProgramSpec `ebpf:"uprobe_SetName"`
	UprobeSetStatus     *ebpf.ProgramSpec `ebpf:"uprobe_SetStatus"`
	UprobeStart         *ebpf.ProgramSpec `ebpf:"uprobe_Start"`
	UprobeStartReturns  *ebpf.ProgramSpec `ebpf:"uprobe_Start_Returns"`
}

// bpfMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfMapSpecs struct {
	ActiveSpansBySpanPtr      *ebpf.MapSpec `ebpf:"active_spans_by_span_ptr"`
	AllocMap                  *ebpf.MapSpec `ebpf:"alloc_map"`
	Events                    *ebpf.MapSpec `ebpf:"events"`
	GoContextToSc             *ebpf.MapSpec `ebpf:"go_context_to_sc"`
	GolangMapbucketStorageMap *ebpf.MapSpec `ebpf:"golang_mapbucket_storage_map"`
	OtelSpanStorageMap        *ebpf.MapSpec `ebpf:"otel_span_storage_map"`
	ProbeActiveSamplerMap     *ebpf.MapSpec `ebpf:"probe_active_sampler_map"`
	SamplersConfigMap         *ebpf.MapSpec `ebpf:"samplers_config_map"`
	SliceArrayBuffMap         *ebpf.MapSpec `ebpf:"slice_array_buff_map"`
	SpanNameByContext         *ebpf.MapSpec `ebpf:"span_name_by_context"`
	TracerIdByContext         *ebpf.MapSpec `ebpf:"tracer_id_by_context"`
	TracerIdStorageMap        *ebpf.MapSpec `ebpf:"tracer_id_storage_map"`
	TracerPtrToIdMap          *ebpf.MapSpec `ebpf:"tracer_ptr_to_id_map"`
	TrackedSpansBySc          *ebpf.MapSpec `ebpf:"tracked_spans_by_sc"`
}

// bpfVariableSpecs contains global variables before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfVariableSpecs struct {
	AttrTypeBool                    *ebpf.VariableSpec `ebpf:"attr_type_bool"`
	AttrTypeBoolslice               *ebpf.VariableSpec `ebpf:"attr_type_boolslice"`
	AttrTypeFloat64                 *ebpf.VariableSpec `ebpf:"attr_type_float64"`
	AttrTypeFloat64slice            *ebpf.VariableSpec `ebpf:"attr_type_float64slice"`
	AttrTypeInt64                   *ebpf.VariableSpec `ebpf:"attr_type_int64"`
	AttrTypeInt64slice              *ebpf.VariableSpec `ebpf:"attr_type_int64slice"`
	AttrTypeInvalid                 *ebpf.VariableSpec `ebpf:"attr_type_invalid"`
	AttrTypeString                  *ebpf.VariableSpec `ebpf:"attr_type_string"`
	AttrTypeStringslice             *ebpf.VariableSpec `ebpf:"attr_type_stringslice"`
	BucketsPtrPos                   *ebpf.VariableSpec `ebpf:"buckets_ptr_pos"`
	EndAddr                         *ebpf.VariableSpec `ebpf:"end_addr"`
	Hex                             *ebpf.VariableSpec `ebpf:"hex"`
	IsRegistersAbi                  *ebpf.VariableSpec `ebpf:"is_registers_abi"`
	StartAddr                       *ebpf.VariableSpec `ebpf:"start_addr"`
	TotalCpus                       *ebpf.VariableSpec `ebpf:"total_cpus"`
	TracerDelegatePos               *ebpf.VariableSpec `ebpf:"tracer_delegate_pos"`
	TracerIdContainsSchemaURL       *ebpf.VariableSpec `ebpf:"tracer_id_contains_schemaURL"`
	TracerIdContainsScopeAttributes *ebpf.VariableSpec `ebpf:"tracer_id_contains_scope_attributes"`
	TracerNamePos                   *ebpf.VariableSpec `ebpf:"tracer_name_pos"`
	TracerProviderPos               *ebpf.VariableSpec `ebpf:"tracer_provider_pos"`
	TracerProviderTracersPos        *ebpf.VariableSpec `ebpf:"tracer_provider_tracers_pos"`
}

// bpfObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfObjects struct {
	bpfPrograms
	bpfMaps
	bpfVariables
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
	ActiveSpansBySpanPtr      *ebpf.Map `ebpf:"active_spans_by_span_ptr"`
	AllocMap                  *ebpf.Map `ebpf:"alloc_map"`
	Events                    *ebpf.Map `ebpf:"events"`
	GoContextToSc             *ebpf.Map `ebpf:"go_context_to_sc"`
	GolangMapbucketStorageMap *ebpf.Map `ebpf:"golang_mapbucket_storage_map"`
	OtelSpanStorageMap        *ebpf.Map `ebpf:"otel_span_storage_map"`
	ProbeActiveSamplerMap     *ebpf.Map `ebpf:"probe_active_sampler_map"`
	SamplersConfigMap         *ebpf.Map `ebpf:"samplers_config_map"`
	SliceArrayBuffMap         *ebpf.Map `ebpf:"slice_array_buff_map"`
	SpanNameByContext         *ebpf.Map `ebpf:"span_name_by_context"`
	TracerIdByContext         *ebpf.Map `ebpf:"tracer_id_by_context"`
	TracerIdStorageMap        *ebpf.Map `ebpf:"tracer_id_storage_map"`
	TracerPtrToIdMap          *ebpf.Map `ebpf:"tracer_ptr_to_id_map"`
	TrackedSpansBySc          *ebpf.Map `ebpf:"tracked_spans_by_sc"`
}

func (m *bpfMaps) Close() error {
	return _BpfClose(
		m.ActiveSpansBySpanPtr,
		m.AllocMap,
		m.Events,
		m.GoContextToSc,
		m.GolangMapbucketStorageMap,
		m.OtelSpanStorageMap,
		m.ProbeActiveSamplerMap,
		m.SamplersConfigMap,
		m.SliceArrayBuffMap,
		m.SpanNameByContext,
		m.TracerIdByContext,
		m.TracerIdStorageMap,
		m.TracerPtrToIdMap,
		m.TrackedSpansBySc,
	)
}

// bpfVariables contains all global variables after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfVariables struct {
	AttrTypeBool                    *ebpf.Variable `ebpf:"attr_type_bool"`
	AttrTypeBoolslice               *ebpf.Variable `ebpf:"attr_type_boolslice"`
	AttrTypeFloat64                 *ebpf.Variable `ebpf:"attr_type_float64"`
	AttrTypeFloat64slice            *ebpf.Variable `ebpf:"attr_type_float64slice"`
	AttrTypeInt64                   *ebpf.Variable `ebpf:"attr_type_int64"`
	AttrTypeInt64slice              *ebpf.Variable `ebpf:"attr_type_int64slice"`
	AttrTypeInvalid                 *ebpf.Variable `ebpf:"attr_type_invalid"`
	AttrTypeString                  *ebpf.Variable `ebpf:"attr_type_string"`
	AttrTypeStringslice             *ebpf.Variable `ebpf:"attr_type_stringslice"`
	BucketsPtrPos                   *ebpf.Variable `ebpf:"buckets_ptr_pos"`
	EndAddr                         *ebpf.Variable `ebpf:"end_addr"`
	Hex                             *ebpf.Variable `ebpf:"hex"`
	IsRegistersAbi                  *ebpf.Variable `ebpf:"is_registers_abi"`
	StartAddr                       *ebpf.Variable `ebpf:"start_addr"`
	TotalCpus                       *ebpf.Variable `ebpf:"total_cpus"`
	TracerDelegatePos               *ebpf.Variable `ebpf:"tracer_delegate_pos"`
	TracerIdContainsSchemaURL       *ebpf.Variable `ebpf:"tracer_id_contains_schemaURL"`
	TracerIdContainsScopeAttributes *ebpf.Variable `ebpf:"tracer_id_contains_scope_attributes"`
	TracerNamePos                   *ebpf.Variable `ebpf:"tracer_name_pos"`
	TracerProviderPos               *ebpf.Variable `ebpf:"tracer_provider_pos"`
	TracerProviderTracersPos        *ebpf.Variable `ebpf:"tracer_provider_tracers_pos"`
}

// bpfPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfPrograms struct {
	UprobeEnd           *ebpf.Program `ebpf:"uprobe_End"`
	UprobeSetAttributes *ebpf.Program `ebpf:"uprobe_SetAttributes"`
	UprobeSetName       *ebpf.Program `ebpf:"uprobe_SetName"`
	UprobeSetStatus     *ebpf.Program `ebpf:"uprobe_SetStatus"`
	UprobeStart         *ebpf.Program `ebpf:"uprobe_Start"`
	UprobeStartReturns  *ebpf.Program `ebpf:"uprobe_Start_Returns"`
}

func (p *bpfPrograms) Close() error {
	return _BpfClose(
		p.UprobeEnd,
		p.UprobeSetAttributes,
		p.UprobeSetName,
		p.UprobeSetStatus,
		p.UprobeStart,
		p.UprobeStartReturns,
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
//go:embed bpf_arm64_bpfel.o
var _BpfBytes []byte
