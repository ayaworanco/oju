defmodule Oluwoye.WordActions.Auth do
  def handle(%Awo.Words.Auth{key: key} = _word), do: Oluwoye.Server.authorize_by_key(key)
end
