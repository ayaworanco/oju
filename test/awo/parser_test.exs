defmodule Awo.Test.ParserTest do
  use ExUnit.Case

  describe "AUTH" do
    test "should return a auth word" do
      payload = "AUTH:[key=123456]"
      parsed_value = Awo.Parser.parse(payload)

      assert %Awo.Words.Auth{key: data_key} = parsed_value
      refute is_nil(data_key)
      assert Map.get(parsed_value, :key) == "123456"
    end
  end

  describe "ERROR" do
    test "should return a error word" do
      payload = "ERROR:[]"
      parsed_value = Awo.Parser.parse(payload)

      assert %Awo.Words.Error{msg: nil} = parsed_value
    end

    test "should return a error word with a message" do
      payload = "ERROR:[msg=\"some tests\"]"
      parsed_value = Awo.Parser.parse(payload)

      assert %Awo.Words.Error{msg: "some tests"} = parsed_value
    end
  end

  describe "LOG" do
    test "should return a log word" do
      payload = "LOG:[level=DEBUG,key=123456,log=\"testing\"]"
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

  describe "OK" do
    test "should return an ok word" do
      payload = "OK:[]"
      parsed_value = Awo.Parser.parse(payload)

      assert %Awo.Words.Ok{msg: nil} = parsed_value
    end

    test "should return a ok word with a message" do
      payload = "OK:[msg=\"some tests\"]"
      parsed_value = Awo.Parser.parse(payload)

      assert %Awo.Words.Ok{msg: "some tests"} = parsed_value
    end
  end
end
