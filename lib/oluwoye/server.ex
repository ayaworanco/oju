defmodule Oluwoye.Server do
  use GenServer

  def start_link(options) do
    GenServer.start_link(__MODULE__, options, name: __MODULE__)
  end

  def init(_opts) do
    {:ok, []}
  end

  def handle_info({:ssl, socket, data}, socket) do
    IO.puts("Received data: #{data}")

    {:noreply, socket}
  end
end
