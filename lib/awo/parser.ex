defmodule Awo.Parser do
  @auth "AUTH"
  @error "ERROR"
  @log "LOG"
  @ok "OK"

  alias Awo.Words.Log
  alias Awo.Words.Ok
  alias Awo.Words.Error
  alias Awo.Words.Auth

  def parse(string) do
    # begin the parse of an word 
    # ex: AUTH:[key=123456]
    [word, data] = String.split(string, ":")

    case String.upcase(word) do
      @auth -> generate_word(@auth, data)
      @error -> generate_word(@error, data)
      @ok -> generate_word(@ok, data)
      @log -> generate_word(@log, data)
    end
  end

  defp generate_word("AUTH", data) do
    %{"key" => key} = serialize_data(data)
    %Auth{key: key}
  end

  defp generate_word("ERROR", data) do
    %{"msg" => msg} = serialize_data(data)
    %Error{msg: msg}
  end

  defp generate_word("LOG", data) do
    %{"level" => level, "key" => key, "log" => log} = serialize_data(data)
    %Log{level: level, key: key, log: log}
  end

  defp generate_word("OK", data) do
    %{"msg" => msg} = serialize_data(data)
    %Ok{msg: msg}
  end

  defp serialize_data(data) do
    list_of_data =
      data
      |> String.replace("[", "")
      |> String.replace("]", "")
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
