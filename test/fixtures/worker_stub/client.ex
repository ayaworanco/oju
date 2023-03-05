defmodule Oluwoye.Fixtures.WorkerStub.Client do
  use GenServer

  @default_port 7070

  def start_link(options) do
    GenServer.start_link(__MODULE__, :no_state, options)
  end

  def init(:no_state) do
    {:ok, socket} = :gen_tcp.connect(@default_port)
    {:ok, socket}
  end
end
