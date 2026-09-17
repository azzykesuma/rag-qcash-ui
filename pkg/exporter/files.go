package exporter

import (
	"bytes"
	"os"
	"path/filepath"
	"runtime"
)

// WriteIfChanged preserves the modification time of identical exports. The
// temporary file lives beside its destination so replacement stays on one volume.
func WriteIfChanged(path string, data []byte) error {
	if err := recoverFile(path); err != nil {
		return err
	}
	old, err := os.ReadFile(path)
	if err == nil && bytes.Equal(old, data) {
		return nil
	}
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	tempDir := filepath.Join(filepath.Dir(path), ".vault-tmp")
	if err := os.MkdirAll(tempDir, 0700); err != nil {
		return err
	}
	f, err := os.CreateTemp(tempDir, "export-*")
	if err != nil {
		return err
	}
	defer os.Remove(tempDir)
	defer os.Remove(f.Name())
	if _, err := f.Write(data); err != nil {
		f.Close()
		return err
	}
	if err := f.Sync(); err != nil {
		f.Close()
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return replaceFile(f.Name(), path)
}

func replaceFile(tempPath, destination string) error {
	if runtime.GOOS != "windows" {
		return os.Rename(tempPath, destination)
	}
	if err := recoverFile(destination); err != nil {
		return err
	}
	backup := backupPath(destination)
	_ = os.Remove(backup)
	if err := os.Rename(destination, backup); err != nil && !os.IsNotExist(err) {
		return err
	}
	if err := os.Rename(tempPath, destination); err != nil {
		_ = os.Rename(backup, destination)
		return err
	}
	if err := os.Remove(backup); err != nil && !os.IsNotExist(err) {
		return err
	}
	return nil
}

func backupPath(destination string) string {
	return filepath.Join(filepath.Dir(destination), ".vault-tmp", filepath.Base(destination)+".backup")
}

func recoverFile(destination string) error {
	backup := backupPath(destination)
	if _, err := os.Stat(destination); err == nil {
		if removeErr := os.Remove(backup); removeErr != nil && !os.IsNotExist(removeErr) {
			return removeErr
		}
		return nil
	} else if !os.IsNotExist(err) {
		return err
	}
	if _, err := os.Stat(backup); err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	return os.Rename(backup, destination)
}
