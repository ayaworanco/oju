package armazen

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"
	"oju/internal/system"
	"os"
	"path/filepath"
	"time"
)

func EncodeSystemToBytes(sys system.System) ([]byte, error) {
	buffer := bytes.Buffer{}
	gob.Register(system.System{})
	encoding := gob.NewEncoder(&buffer)
	encoding_error := encoding.Encode(sys)

	if encoding_error != nil {
		return []byte{}, encoding_error
	}

	return buffer.Bytes(), nil
}

func Compress(buffer []byte) []byte {
	zip_buffer := bytes.Buffer{}
	zipped := gzip.NewWriter(&zip_buffer)

	zipped.Write(buffer)
	zipped.Close()
	return zip_buffer.Bytes()
}

func WriteToFile(buffer []byte) error {
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
