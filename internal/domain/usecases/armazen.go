package usecases

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"
	"os"
	"path/filepath"
	"time"
)

func encode_to_bytes[T any](entity T) ([]byte, error) {
	buffer := bytes.Buffer{}
	encoding := gob.NewEncoder(&buffer)
	encoding_error := encoding.Encode(entity)

	if encoding_error != nil {
		return []byte{}, encoding_error
	}

	return buffer.Bytes(), nil
}

func compress(buffer []byte) []byte {
	compress_buffer := bytes.Buffer{}
	compressed := gzip.NewWriter(&compress_buffer)

	compressed.Write(buffer)
	compressed.Close()

	return compress_buffer.Bytes()
}

func write_to_file(buffer []byte) error {
	folder_path, folder_path_error := generate_data_folder()
	if folder_path_error != nil {
		return folder_path_error
	}

	output_file, output_file_error := output_system_file(folder_path)
	if output_file_error != nil {
		return output_file_error
	}

	created_file, create_error := os.Create(output_file)

	if create_error != nil {
		return create_error
	}

	created_file.Write(buffer)
	created_file.Close()

	return nil
}

func generate_data_folder() (string, error) {
	path, path_error := os.Executable()
	if path_error != nil {
		return "", path_error
	}

	executable_path := filepath.Dir(path) + "/data"

	env_data_dir := os.Getenv("OUTPUT_DATA_DIR")
	if env_data_dir != "" {
		executable_path = env_data_dir + "/data"
	}

	if _, folder_not_exist_error := os.Stat(executable_path); os.IsNotExist(folder_not_exist_error) {
		make_dir_error := os.Mkdir(executable_path, os.ModePerm)
		if make_dir_error != nil {
			return "", make_dir_error
		}
	}

	return executable_path, nil
}

func output_system_file(folder_path string) (string, error) {
	return folder_path + "/" + time.Now().Format(time.RFC3339) + ".dat", nil
}

func save[T any](entity T) error {

	system_to_bytes, error_encoding := encode_to_bytes[T](entity)
	if error_encoding != nil {
		return error_encoding
	}

	system_compressed := compress(system_to_bytes)
	error_on_writing := write_to_file(system_compressed)
	if error_on_writing != nil {
		return error_encoding
	}
	return nil
}
