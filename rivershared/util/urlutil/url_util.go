package urlutil

import "net/url"

// DatabaseSQLCompatibleURL returns databaseURL, but with any known pgx-specific
// URL parameters removed so that it'll be parseable by `database/sql`.
func DatabaseSQLCompatibleURL(databaseURL string) string {
	parsedURL, err := url.Parse(databaseURL)
	if err != nil {
		panic(err)
	}

	query := parsedURL.Query()
	for _, key := range []string{
		"pool_health_check_period",
		"pool_max_conn_idle_time",
		"pool_max_conn_lifetime",
		"pool_max_conn_lifetime_jitter",
		"pool_max_conns",
		"pool_min_conns",
		"pool_min_idle_conns",
	} {
		query.Del(key)
	}

	parsedURL.RawQuery = query.Encode()

	return parsedURL.String()
}
