package Cache

const ontype = "memory"

func config() int {
	switch ontype {
	case "memory":
		return 1

	case "redis":
		return 2

	}
	return 0
}
