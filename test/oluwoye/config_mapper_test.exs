defmodule Oluwoye.Test.ConfigMapperTest do
  use ExUnit.Case

  test "should load applications.yaml" do
    config = Oluwoye.ConfigMapper.map("test/fixtures/applications.yaml")
    assert [%{name: "worker_php"}] == config
  end
end
