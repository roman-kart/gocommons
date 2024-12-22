package filesys_test

import (
	"github.com/roman-kart/gocommons/filesys"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIsDirectory(t *testing.T) {
	require.NoError(t, filesys.IsDirectory("test_data/empty_dir"))
	require.ErrorIs(t, filesys.IsDirectory("test_data/empty_file"), filesys.ErrPathIsNotDirectory)
	require.ErrorIs(t, filesys.IsDirectory("test_data/none"), filesys.ErrPathNotExists)
}
