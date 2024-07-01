package structs

type ResponseOp struct {
	ExecutionID         string     `json:"execution_id"`
	QueryID             int        `json:"query_id"`
	IsExecutionFinished bool       `json:"is_execution_finished"`
	State               string     `json:"state"`
	SubmittedAt         string     `json:"submitted_at"`
	ExpiresAt           string     `json:"expires_at"`
	ExecutionStartedAt  string     `json:"execution_started_at"`
	ExecutionEndedAt    string     `json:"execution_ended_at"`
	Result              ResultOp   `json:"result"`
	Metadata            MetadataOp `json:"metadata"`
	NextURI             string     `json:"next_uri"`
	NextOffset          int        `json:"next_offset"`
}

type ResultOp struct {
	Rows []RowOp `json:"rows"`
}

type RowOp struct {
	ETHxTVL                 float64 `json:"ETHx_TVL"`
	OETHTVL                 float64 `json:"OETH_TVL"`
	AnkrETHTVL              float64 `json:"ankrETH_TVL"`
	BeaconChainETHTVL       float64 `json:"beacon_chain_ETH_TVL"`
	CbETHTVL                float64 `json:"cbETH_TVL"`
	NumStakers              int     `json:"num_stakers"`
	LsETHTVL                float64 `json:"lsETH_TVL"`
	METHTVL                 float64 `json:"mETH_TVL"`
	OperatorContractAddress string  `json:"operator_contract_address"`
	OperatorName            string  `json:"operator_name"`
	OsETHTVL                float64 `json:"osETH_TVL"`
	RETHTVL                 float64 `json:"rETH_TVL"`
	SfrxETHTVL              float64 `json:"sfrxETH_TVL"`
	StETHTVL                float64 `json:"stETH_TVL"`
	SwETHTVL                float64 `json:"swETH_TVL"`
	TotalTVL                float64 `json:"total_TVL"`
	WbETHTVL                float64 `json:"wBETH_TVL"`
}

type MetadataOp struct {
	ColumnNames         []string `json:"column_names"`
	ColumnTypes         []string `json:"column_types"`
	RowCount            int      `json:"row_count"`
	ResultSetBytes      int      `json:"result_set_bytes"`
	TotalRowCount       int      `json:"total_row_count"`
	TotalResultSetBytes int      `json:"total_result_set_bytes"`
	DatapointCount      int      `json:"datapoint_count"`
	PendingTimeMillis   int      `json:"pending_time_millis"`
	ExecutionTimeMillis int      `json:"execution_time_millis"`
}
