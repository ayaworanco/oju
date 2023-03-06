defmodule Awo.Parser do
  @log "log"

  alias Awo.Words.Log

  @doc """
  This will parse a TCP packet incoming from an application to a word struct
  """

  def parse(string) do
    with [word, key, data] <- String.split(string, ":") do
      case String.downcase(word) do
        @log -> generate_word(@log, data, key)
      end
    else
      _ -> %Awo.ParserError{}
    end
  end

  defp generate_word("log", data, key) do
    %{"level" => level, "log" => log} = serialize_data(data)
    %Log{level: level, key: key, log: log}
  end

  defp serialize_data(data) do
    list_of_data =
      data
      |> String.replace("{", "")
      |> String.replace("}", "")
      |> String.split(",")

    Enum.map(
      list_of_data,
      fn
        "" ->
          {"msg", nil}

        str ->
          [key, value] = String.split(str, "=")
          replaced_value = String.replace(value, "\"", "")
          {key, replaced_value}
      end
    )
    |> Enum.into(%{})
  end
end
