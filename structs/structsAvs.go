package structs

import "time"

type Metadata struct {
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

type ResultRow struct {
	ETHxTVL            float64 `json:"ETHx_TVL"`
	OETHTVL            float64 `json:"OETH_TVL"`
	AnkrETHTVL         float64 `json:"ankrETH_TVL"`
	AVSContractAddress string  `json:"avs_contract_address"`
	AVSName            string  `json:"avs_name"`
	BeaconChainETHTVL  float64 `json:"beacon_chain_ETH_TVL"`
	CbETHTVL           float64 `json:"cbETH_TVL"`
	LsETHTVL           float64 `json:"lsETH_TVL"`
	METHTVL            float64 `json:"mETH_TVL"`
	NumOperators       int     `json:"num_operators"`
	NumStakers         int     `json:"num_stakers"`
	OsETHTVL           float64 `json:"osETH_TVL"`
	REHTVL             float64 `json:"rETH_TVL"`
	SfrxETHTVL         float64 `json:"sfrxETH_TVL"`
	StETHTVL           float64 `json:"stETH_TVL"`
	SwETHTVL           float64 `json:"swETH_TVL"`
	TotalTVL           float64 `json:"total_TVL"`
	WBETHTVL           float64 `json:"wBETH_TVL"`
}

type ExecutionResult struct {
	Rows []ResultRow `json:"rows"`
}

type Response struct {
	ExecutionID         string          `json:"execution_id"`
	QueryID             int             `json:"query_id"`
	IsExecutionFinished bool            `json:"is_execution_finished"`
	State               string          `json:"state"`
	SubmittedAt         time.Time       `json:"submitted_at"`
	ExpiresAt           time.Time       `json:"expires_at"`
	ExecutionStartedAt  time.Time       `json:"execution_started_at"`
	ExecutionEndedAt    time.Time       `json:"execution_ended_at"`
	Result              ExecutionResult `json:"result"`
	Metadata            Metadata        `json:"metadata"`
	NextURI             string          `json:"next_uri"`
	NextOffset          int             `json:"next_offset"`
}
