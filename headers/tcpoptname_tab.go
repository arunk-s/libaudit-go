package headers

// Location: include/uapi/linux/tcp.h
var TcpOptLookup = map[int]string{
	1:  "TCP_NODELAY",
	2:  "TCP_MAXSEG",
	3:  "TCP_CORK",
	4:  "TCP_KEEPIDLE",
	5:  "TCP_KEEPINTVL",
	6:  "TCP_KEEPCNT",
	7:  "TCP_SYNCNT",
	8:  "TCP_LINGER2",
	9:  "TCP_DEFER_ACCEPT",
	10: "TCP_WINDOW_CLAMP",
	11: "TCP_INFO",
	12: "TCP_QUICKACK",
	13: "TCP_CONGESTION",
	14: "TCP_MD5SIG",
	15: "TCP_COOKIE_TRANSACTIONS",
	16: "TCP_THIN_LINEAR_TIMEOUTS",
	17: "TCP_THIN_DUPACK",
	18: "TCP_USER_TIMEOUT",
	19: "TCP_REPAIR",
	20: "TCP_REPAIR_QUEUE",
	21: "TCP_QUEUE_SEQ",
	22: "TCP_REPAIR_OPTIONS",
	23: "TCP_FASTOPEN",
	24: "TCP_TIMESTAMP",
	25: "TCP_NOTSENT_LOWAT",
}
