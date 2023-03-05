defmodule Oluwoye.ConfigMapper do
  def from_file(path \\ "./applications.yaml")
  def from_file(""), do: %Oluwoye.Exceptions.ConfigMapperError{msg: "No file path passed"}
  def from_file(nil), do: %Oluwoye.Exceptions.ConfigMapperError{msg: "No file path passed"}

  def from_file(path), do: validate_schema(path)

  defp validate_schema(path) do
    case YamlElixir.read_from_file(path) do
      {:ok, yaml} ->
        yaml

      {:error, reason} ->
        %Oluwoye.Exceptions.ConfigMapperError{msg: reason}
    end
    |> case do
      %{"applications" => list_of_apps} when list_of_apps != [] ->
        has_valid_apps?(list_of_apps)

      %{} ->
        %Oluwoye.Exceptions.ConfigMapperError{msg: "YAML is empty"}
    end
  end

  defp has_valid_apps?(nil),
    do: %Oluwoye.Exceptions.ConfigMapperError{msg: "No applications found"}

  defp has_valid_apps?(applications) do
    Enum.map(applications, fn
      app ->
        case app do
          %{"name" => _name, "key" => _key} -> true
          _ -> false
        end
    end)
    |> Enum.member?(false)
    |> case do
      true -> %Oluwoye.Exceptions.ConfigMapperError{msg: "Applications not valid"}
      false -> applications
    end
  end
end
