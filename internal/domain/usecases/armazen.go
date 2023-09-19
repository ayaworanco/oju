package usecases

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"
	"io"
	"oju/internal/domain/entities"
	"os"
	"path/filepath"
	"sort"
	"strings"
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

func decompress(buffer []byte) ([]byte, error) {
	reader, error_on_new_reader := gzip.NewReader(bytes.NewReader(buffer))
	if error_on_new_reader != nil {
		return []byte{}, error_on_new_reader
	}

	data, error_on_read := io.ReadAll(reader)
	if error_on_read != nil {
		return []byte{}, error_on_read
	}
	reader.Close()

	return data, nil
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

func LoadSystem() (entities.System, error) {
	data, error_on_load_data := load_file()
	if error_on_load_data != nil {
		return entities.System{}, error_on_load_data
	}

	decompressed, error_on_decompress := decompress(data)

	if error_on_decompress != nil {
		return entities.System{}, error_on_decompress
	}

	entity := entities.System{}
	decoder := gob.NewDecoder(bytes.NewReader(decompressed))
	error_on_decode := decoder.Decode(&entity)

	if error_on_decode != nil {
		return entities.System{}, error_on_decode
	}

	return entity, nil
}

func load_file() ([]byte, error) {
	path, path_error := os.Executable()
	if path_error != nil {
		return []byte{}, path_error
	}
	executable_path := filepath.Dir(path) + "/data"

	env_data_dir := os.Getenv("OUTPUT_DATA_DIR")
	if env_data_dir != "" {
		executable_path = env_data_dir + "/data"
	}

	file_name, file_name_error := get_last_data_file(executable_path)

	if file_name_error != nil {
		return []byte{}, file_name_error
	}

	data, err_on_read := os.ReadFile(executable_path + "/" + file_name)

	if err_on_read != nil {
		return []byte{}, err_on_read
	}

	return data, nil
}

func get_last_data_file(path string) (string, error) {
	dir, dir_error := os.ReadDir(path)

	if dir_error != nil {
		return "", dir_error
	}

	var timers []time.Time

	for _, entry := range dir {
		entry_timer := strings.Replace(entry.Name(), ".dat", "", 1)
		entry_timer = strings.Replace(entry_timer, "-03:00", "", 1)
		timing, error_timning := time.Parse("2006-01-02T15:04:05", entry_timer)

		if error_timning != nil {
			return "", error_timning
		}

		timers = append(timers, timing)
	}

	sort.SliceStable(timers, func(i, j int) bool {
		return timers[i].After(timers[j])
	})

	return strings.Replace(timers[0].Format(time.RFC3339), "Z", "-03:00.dat", 1), nil
}
