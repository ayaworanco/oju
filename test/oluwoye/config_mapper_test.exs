defmodule Oluwoye.Test.ConfigMapperTest do
  use ExUnit.Case

  test "should load applications.yaml" do
    config = Oluwoye.ConfigMapper.get_config_from_file("test/fixtures/applications.yaml")

    assert %{"applications" => [%{"name" => "worker_php"}, %{"name" => "ruby_on_trails"}]} =
             config
  end
end
