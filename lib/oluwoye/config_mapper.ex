defmodule Oluwoye.ConfigMapper do
  def get_config_from_file(path \\ "./applications.yaml") do
    case File.read(path) do
      {:ok, binary_file} ->
        nil

      {:error, :enoent} ->
        %Oluwoye.Exceptions.ConfigMapperError{msg: "No file exists"}
    end
  end
end
