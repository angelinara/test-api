package parser

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadDescription(t *testing.T) {
	content := "#!/bin/bash\n# Login with email and password\ncurl -s -X POST \"http://localhost\"\n"

	tmp, err := os.CreateTemp("", "*.sh")
	require.NoError(t, err)
	defer os.Remove(tmp.Name())

	_, err = tmp.WriteString(content)
	require.NoError(t, err)
	tmp.Close()

	description, err := ReadDescription(tmp.Name())
	require.NoError(t, err)
	assert.Equal(t, "Login with email and password", description)
}
