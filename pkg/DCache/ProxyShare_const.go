//Package DCache comment
// This file war generated by tars2go 1.1
// Generated from ProxyShare.tars
package DCache

//const as define in tars file
const (
	ET_SUCC                  int32  = 0
	ET_SYS_ERR               int32  = -1
	ET_MODULE_NAME_INVALID   int32  = -2
	ET_KEY_INVALID           int32  = -3
	ET_KEY_AREA_ERR          int32  = -4
	ET_KEY_TYPE_ERR          int32  = -5
	ET_NO_DATA               int32  = -6
	ET_COMMAND_INVALID       int32  = -7
	ET_FORBID_DEL            int32  = -8
	ET_SERVER_TYPE_ERR       int32  = -9
	ET_DATA_EXIST            int32  = -10
	ET_FORBID_OPT            int32  = -11
	ET_INPUT_PARAM_ERROR     int32  = -12
	ET_ONLY_KEY              int32  = -13
	ET_DIRTY_DATA            int32  = -14
	ET_DATA_VER_MISMATCH     int32  = -15
	ET_PARTIAL_FAIL          int32  = -16
	ET_GZIP_UNCOMPRESS_ERR   int32  = -17
	ET_CACHE_TYPE_ERR        int32  = -18
	ET_SYNC_SET_SAME         int32  = -19
	ET_SET_SYNC_NOENABLE     int32  = -20
	ET_MEM_FULL              int32  = -21
	ET_DATA_TOO_MUCH         int32  = -22
	ET_READ_ONLY             int32  = -23
	ET_DB_ERR                int32  = -24
	ET_CACHE_ERR             int32  = -25
	ET_PARAM_REDUNDANT       int32  = -26
	ET_PARAM_MISSING         int32  = -27
	ET_PARAM_TOO_LONG        int32  = -28
	ET_PARAM_DIGITAL_ERR     int32  = -29
	ET_PARAM_NOT_EXIST       int32  = -30
	ET_PARAM_OP_ERR          int32  = -31
	ET_PARAM_UKEY_MISSING    int32  = -32
	ET_PARAM_UPDATE_UKEY_ERR int32  = -33
	ET_PARAM_LIMIT_VALUE_ERR int32  = -34
	ET_ERASE_DIRTY_ERR       int32  = -35
	ET_DB_TO_CACHE_ERR       int32  = -36
	DVER                     string = "@DataVer"
	EXPIRETIME               string = "@ExpireTime"
	SCOREVALUE               string = "@ScoreValue"
	VALUE_SUCC               int32  = 0
	VALUE_NO_DATA            int32  = 1
	SET_SUCC                 int32  = 0
	SET_ERROR                int32  = -1
	SET_DATA_VER_MISMATCH    int32  = -2
	DEL_SUCC                 int32  = 0
	DEL_ERROR                int32  = -1
	DEL_DATA_VER_MISMATCH    int32  = -2
	WRITE_SUCC               int32  = 0
	WRITE_ERROR              int32  = -1
	WRITEL_DATA_VER_MISMATCH int32  = -2
)