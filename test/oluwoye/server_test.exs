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

  describe "When do AUTH by one client" do
    test "should return to client an OK response with message that client was authenticated", %{
      server: server
    } do
      {:ok, _client_pid} = Oluwoye.Fixtures.WorkerStub.Client.start_link([])
      [%{"name" => worker_client_name}] = Oluwoye.Server.get_authenticated_apps()
      assert worker_client_name != ""
      refute is_nil(worker_client_name)
    end
  end
end
