defmodule Oluwoye.Test.ServerTest do
  use ExUnit.Case

  describe "When do AUTH by one client" do
    setup do
      {:ok, pid} = Oluwoye.Server.start_link(port: 7070)
      %{server: pid}
    end

    test "should return to client an OK response with message that client was authenticated", %{
      server: server
    } do
      {:ok, client_pid} = Oluwoye.Fixtures.WorkerStub.Client.start_link([])
    end
  end
end
