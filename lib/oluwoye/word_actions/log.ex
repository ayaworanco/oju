defmodule Oluwoye.WordActions.Log do
  require Logger

  # TODO: here we will process this log
  def handle(%Awo.Words.Log{level: level, key: key, log: log} = _word) do
    Logger.info("#{level}: LOG from #{key} -> #{log}")
  end
end
