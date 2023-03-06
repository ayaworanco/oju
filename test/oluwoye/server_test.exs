defmodule Oluwoye.Test.ServerTest do
  use ExUnit.Case

  setup do
    {:ok, pid} =
      Oluwoye.Server.start_link(
        port: 7070,
        applications_file: "test/fixtures/worker_stub/applications.yaml"
      )

    %{server: pid}
  end

  describe "When connects and log" do
    test "should save log in a folder", %{
      server: server
    } do
      {:ok, _client_pid} = Oluwoye.Fixtures.WorkerStub.Client.start_link([])
    end
  end
end
