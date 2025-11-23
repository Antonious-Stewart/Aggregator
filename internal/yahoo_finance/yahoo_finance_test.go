package yahoo_finance_test

import (
	"github.com/Antonious-Stewart/Aggregator/internal/db"
	"github.com/Antonious-Stewart/Aggregator/internal/yahoo_finance"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewYahooFinance(t *testing.T) {
	runner := &db.Database{}
	yh := yahoo_finance.New(runner)

	assert.NotNil(t, yh)
	assert.Equal(t, runner, yh.SqlRunner)
}
