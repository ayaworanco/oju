defmodule Awo.Test.ParserTest do
  use ExUnit.Case

  describe "erroring parser" do
    test "should return an error when a non correct message is incoming" do
      payload = "a"
      parsed_value = Awo.Parser.parse(payload)

      assert %Awo.ParserError{msg: msg} = parsed_value
      refute is_nil(msg)
      assert msg == "Error parsing packet"
    end
  end

  describe "LOG" do
    test "should return a log word" do
      payload = "log:123456:{level=DEBUG,log=\"testing\"}"
      parsed_value = Awo.Parser.parse(payload)

      assert %Awo.Words.Log{level: level, key: key, log: log} = parsed_value
      refute is_nil(level)
      refute is_nil(key)
      refute is_nil(log)

      assert %Awo.Words.Log{
               level: "DEBUG",
               key: "123456",
               log: "testing"
             } == parsed_value
    end
  end
end
