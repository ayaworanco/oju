defmodule Oluwoye.Fixtures.WorkerStub.Client do
  use GenServer

  @default_port 7070

  def start_link(options) do
    GenServer.start_link(__MODULE__, :no_state, options)
  end

  def init(:no_state) do
    {:ok, map} = YamlElixir.read_from_file("test/fixtures/worker_stub/applications.yaml")

    key =
      Map.get(map, "applications")
      |> List.first()
      |> Map.get("key")

    {:ok, socket} = :gen_tcp.connect({127, 0, 0, 1}, @default_port, [:binary, active: true])

    :gen_tcp.send(socket, "AUTH:[key=\"#{key}\"]")
    {:ok, socket}
  end
end
