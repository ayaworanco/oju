defmodule Oluwoye.Test.ConfigMapperTest do
  use ExUnit.Case

  describe "When success config mapping" do
    test "should load applications.yaml" do
      config = Oluwoye.ConfigMapper.from_file("test/fixtures/applications.yaml")

      assert [%{"name" => "worker_php"}, %{"name" => "ruby_on_trails"}] = config
    end
  end

  describe "When erroring Config Mapper" do
    test "should return error on nil" do
      config = Oluwoye.ConfigMapper.from_file(nil)

      assert %Oluwoye.Exceptions.ConfigMapperError{msg: msg} = config
      assert "No file path passed" = msg
    end

    test "should return error on empty string" do
      config = Oluwoye.ConfigMapper.from_file("")

      assert %Oluwoye.Exceptions.ConfigMapperError{msg: msg} = config
      assert "No file path passed" = msg
    end

    test "should return error if is a wrong schema" do
      config = Oluwoye.ConfigMapper.from_file("test/fixtures/error_applications.yaml")

      assert %Oluwoye.Exceptions.ConfigMapperError{msg: msg} = config
      assert "Applications not valid" = msg
    end

    test "should have not valid apps" do
      config = Oluwoye.ConfigMapper.from_file("test/fixtures/no_applications.yaml")
      assert %Oluwoye.Exceptions.ConfigMapperError{msg: msg} = config
      assert "YAML is empty" = msg
    end

    test "should have not valid apps by nil" do
      config = Oluwoye.ConfigMapper.from_file("test/fixtures/nil_applications.yaml")
      assert %Oluwoye.Exceptions.ConfigMapperError{msg: msg} = config
      assert "No applications found" = msg
    end
  end
end
