package ytconfig

import "time"

type BundleControllerServer struct {
	CommonServer
	Cluster                      string        `yson:"cluster"`
	BundleScanPeriod             time.Duration `yson:"bundle_scan_period"`
	BundleScanTransactionTimeout time.Duration `yson:"bundle_scan_transaction_timeout"`

	HulkRequestTimeout              time.Duration `yson:"hulk_request_timeout"`
	CellRemovalTimeout              time.Duration `yson:"cell_removal_timeout"`
	NodeAssignmentTimeout           time.Duration `yson:"spare_node_assignment_timeout"`
	MuteTabletCellsCheckGracePeriod time.Duration `yson:"mute_tablet_cells_check_grace_period"`

	// NYPath::TYPath RootPath;
	RootPath string `yson:"root_path"`

	HasInstanceAllocatorService  bool   `yson:"has_instance_allocator_service"`
	HulkAllocationsPath          string `yson:"hulk_allocations_path"`
	HulkAllocationsHistoryPath   string `yson:"hulk_allocations_history_path"`
	HulkDeallocationsPath        string `yson:"hulk_deallocations_history_path"`
	HulkDeallocationsHistoryPath string `yson:"hulk_deallocations_history_path"`

	DecommissionReleasedNodes bool `yson:"decommission_released_nodes"`
	EnableSpareNodeAssignment bool `yson:"enable_spare_node_assignment"`

	NodeCountPerCell         int   `yson:"node_count_per_cell"`
	ChunkCountPerCell        int   `yson:"chunk_count_per_cell"`
	JournalDiskSpacePerCell  int64 `yson:"journal_disk_space_per_cell"`
	SnapshotDiskSpacePerCell int64 `yson:"snapshot_disk_space_per_cell"`
	MinNodeCount             int   `yson:"min_node_count"`
	MinChunkCount            int   `yson:"min_chunk_count"`

	ReallocateInstanceBudget int `yson:"reallocate_instance_budget"`

	RemoveInstanceCypressNodeAfter time.Duration `yson:"remove_instance_cypress_node_after"`
	OfflineInstanceGracePeriod     time.Duration `yson:"offline_instance_grace_period"`

	EnableNetworkLimits bool `yson:"enable_network_limits"`

	SkipJailedBundles bool `yson:"skip_jailed_bundles"`

	EnableChaosBundleManagement bool         `yson:"enable_chaos_bundle_management"`
	ChaosConfig                 *ChaosConfig `yson:"chaos_config"`
}

type ChaosConfig struct {
	TabletCellClusters []string `yson:"tablet_cell_clusters"`
	ChaosCellClusters  []string `yson:"chaos_cell_clusters"`
	ClockClusterTag    CellTag  `yson:"clock_cluster_tag"`

	AlphaChaosCluster string `yson:"alpha_chaos_cluster"`
	BetaChaosCluster  string `yson:"beta_chaos_cluster"`
}

type CellTag uint16
