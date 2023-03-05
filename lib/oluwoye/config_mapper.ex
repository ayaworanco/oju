defmodule Oluwoye.ConfigMapper do
  def get_config_from_file(path \\ "./applications.yaml") do
    case YamlElixir.read_from_file(path) do
      {:ok, yaml} ->
        yaml

      {:error, reason} ->
        %Oluwoye.Exceptions.ConfigMapperError{msg: reason}
    end
  end
end
