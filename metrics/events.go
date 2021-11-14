package metrics

import (
	"fmt"
)

type RecordAcceptEventAggregated struct {
	IsNAT      bool
	RecordType string
	Count      int
}

func (r RecordAcceptEventAggregated) ToEvent() Event {
	return Event{
		EventType: "thread_record_accepted",
		EventData: map[string]interface{}{
			"record_type": r.RecordType,
			"is_nat":      r.IsNAT,
			"count":       r.Count,
		},
	}
}

func (r RecordAcceptEventAggregated) Key() string {
	return fmt.Sprintf("RecordAcceptEventAggregated%s%v", r.RecordType, r.IsNAT)
}

func (r RecordAcceptEventAggregated) Aggregate(other EventAggregatable) EventAggregatable {
	ev, ok := other.(RecordAcceptEventAggregated)
	// going here we already check the keys, so let's not do this another time
	if !ok {
		return r
	}
	r.Count += ev.Count
	return r
}

type ChangesetEvent struct {
	Diff int64
}

func (c ChangesetEvent) ToEvent() Event {
	return Event{
		EventType: "changeset_applied",
		EventData: map[string]interface{}{
			// we send diff, and not timestamps of records, because we cannot filter
			// them in Amplitude (unless we will have access to SQL there)
			"diff_current_time_vs_first": c.Diff,
		},
	}
}

type ReindexType int

const (
	ReindexTypeThreads ReindexType = iota
	ReindexTypeFiles
	ReindexTypeBundledRelations
	ReindexTypeBundledTypes
	ReindexTypeBundledObjects
	ReindexTypeBundledTemplates
	ReindexTypeOutdatedHeads
)

type ReindexEvent struct {
	ReindexType    ReindexType
	Total          int
	Success        int
	SpentMs        int
	IndexesRemoved bool
}

func (c ReindexEvent) ToEvent() Event {
	return Event{
		EventType: "store_reindex",
		EventData: map[string]interface{}{
			"spent_ms":   c.SpentMs,
			"total":      c.Total,
			"failed":     c.Total - c.Success,
			"type":       c.ReindexType,
			"ix_removed": c.IndexesRemoved,
		},
	}
}

type RecordCreateEvent struct {
	PrepareMs       int64
	NewRecordMs     int64
	LocalEventBusMs int64
	PushMs          int64
}

func (c RecordCreateEvent) ToEvent() Event {
	return Event{
		EventType: "record_create",
		EventData: map[string]interface{}{
			"prepare_ms":    c.PrepareMs,
			"new_record_ms": c.NewRecordMs,
			"local_ms":      c.LocalEventBusMs,
			"push_ms":       c.PushMs,
			"total_ms":      c.PrepareMs + c.NewRecordMs + c.LocalEventBusMs + c.PushMs,
		},
	}
}

type BlockSplit struct {
	AlgorithmMs int64
	ApplyMs     int64
}

func (c BlockSplit) ToEvent() Event {
	return Event{
		EventType: "block_merge",
		EventData: map[string]interface{}{
			"algorithm_ms": c.AlgorithmMs,
			"apply_ms":     c.ApplyMs,
			"total_ms":     c.AlgorithmMs + c.ApplyMs,
		},
	}
}

type BlockMerge struct {
	AlgorithmMs int64
	ApplyMs     int64
}

func (c BlockMerge) ToEvent() Event {
	return Event{
		EventType: "block_split",
		EventData: map[string]interface{}{
			"algorithm_ms": c.AlgorithmMs,
			"apply_ms":     c.ApplyMs,
			"total_ms":     c.AlgorithmMs + c.ApplyMs,
		},
	}
}

type CreateObjectEvent struct {
	SetDetailsMs            int64
	GetWorkspaceBlockWaitMs int64
	WorkspaceCreateMs       int64
	SmartblockCreateMs      int64
	SmartblockType          int
}

func (c CreateObjectEvent) ToEvent() Event {
	return Event{
		EventType: "create_object",
		EventData: map[string]interface{}{
			"set_details_ms":              c.SetDetailsMs,
			"get_workspace_block_wait_ms": c.GetWorkspaceBlockWaitMs,
			"workspace_create_ms":         c.WorkspaceCreateMs,
			"smartblock_create_ms":        c.SmartblockCreateMs,
			"total_ms":                    c.SetDetailsMs + c.GetWorkspaceBlockWaitMs + c.WorkspaceCreateMs + c.SmartblockCreateMs,
			"smartblock_type":             c.SmartblockType,
		},
	}
}

type OpenBlockEvent struct {
	GetBlockMs     int64
	DataviewMs     int64
	ApplyMs        int64
	ShowMs         int64
	FileWatcherMs  int64
	SmartblockType int
}

func (c OpenBlockEvent) ToEvent() Event {
	return Event{
		EventType: "open_block",
		EventData: map[string]interface{}{
			"get_block_ms":       c.GetBlockMs,
			"dataview_notify_ms": c.DataviewMs,
			"apply_ms":           c.ApplyMs,
			"show_ms":            c.ShowMs,
			"file_watchers_ms":   c.FileWatcherMs,
			"total_ms":           c.GetBlockMs + c.DataviewMs + c.ApplyMs + c.ShowMs + c.FileWatcherMs,
			"smartblock_type":    c.SmartblockType,
		},
	}
}

type ProcessThreadsEvent struct {
	WaitTimeMs int64
}

func (c ProcessThreadsEvent) ToEvent() Event {
	return Event{
		EventType: "process_threads",
		EventData: map[string]interface{}{
			"wait_time_ms": c.WaitTimeMs,
		},
	}
}

type AccountRecoverEvent struct {
	SpentMs              int
	TotalThreads         int
	SimultaneousRequests int
}

func (c AccountRecoverEvent) ToEvent() Event {
	return Event{
		EventType: "account_recover",
		EventData: map[string]interface{}{
			"spent_ms":              c.SpentMs,
			"total_threads":         c.TotalThreads,
			"simultaneous_requests": c.SimultaneousRequests,
		},
	}
}
