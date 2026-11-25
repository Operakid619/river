package urlutil

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDatabaseSQLCompatibleURL(t *testing.T) {
	t.Parallel()

	input := "postgres://user:pass@localhost/db?sslmode=require&application_name=worker&pool_max_conns=20&pool_min_conns=2&pool_min_idle_conns=1&pool_max_conn_lifetime=1h&pool_max_conn_lifetime_jitter=1m&pool_max_conn_idle_time=30m&pool_health_check_period=1m"
	parsed, err := url.Parse(DatabaseSQLCompatibleURL(input))
	require.NoError(t, err)
	require.Equal(t, url.Values{"application_name": {"worker"}, "sslmode": {"require"}}, parsed.Query())
	require.Equal(t, "localhost", parsed.Host)
	require.Equal(t, "user:pass", parsed.User.String())
	require.Equal(t, "/db", parsed.Path)
}
