package supabase

import (
	"fmt"
	"github.com/mqqff/absensi-app/domain/errx"
	"github.com/mqqff/absensi-app/internal/infra/env"
	"github.com/mqqff/absensi-app/pkg/log"
	storage "github.com/nedpals/supabase-go"
	"io"
)

type SupabaseInterface interface {
	UploadFile(bucket, path, mimeType string, body io.Reader) (string, error)
	DeleteFile(bucket string, paths ...string) error
}

type SupabaseStruct struct {
	client *storage.Client
}

var Supabase = getSupabase()

func getSupabase() SupabaseInterface {
	url := fmt.Sprintf("%s", env.AppEnv.SupabaseURL)
	client := storage.CreateClient(url, env.AppEnv.SupabaseSecret)

	return &SupabaseStruct{client}
}

func (s *SupabaseStruct) UploadFile(bucket, path, mimeType string, body io.Reader) (string, error) {
	err := safeWrapper(func() error {
		s.client.Storage.From(bucket).Upload(path, body, &storage.FileUploadOptions{ContentType: mimeType})
		return nil
	})

	if err != nil {
		log.Error(log.LogInfo{
			"error":  err.Error(),
			"bucket": bucket,
			"path":   path,
		}, "[Supabase][UploadFile] failed to upload file")
		return "", err
	}

	publicURL := s.client.Storage.From(bucket).GetPublicUrl(path).SignedUrl

	return publicURL, nil
}

func (s *SupabaseStruct) DeleteFile(bucket string, paths ...string) error {
	err := safeWrapper(func() error {
		s.client.Storage.From(bucket).Remove(paths)
		return nil
	})

	if err != nil {
		log.Error(log.LogInfo{
			"error":  err.Error(),
			"bucket": bucket,
			"path":   paths,
		}, "[Supabase][DeleteFile] failed to delete file")
		return err
	}

	return nil
}

func safeWrapper(f func() error) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errx.ErrInternalServer
		}
	}()
	return f()
}
