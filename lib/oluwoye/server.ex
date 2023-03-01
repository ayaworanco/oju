defmodule Oluwoye.Server do
  use GenServer

  @spec start_link(options :: tuple()) :: {:ok, pid()} | {:error, any()}
  def start_link(options) do
    GenServer.start_link(__MODULE__, options, name: __MODULE__)
  end

  def init({port, options}) do
    {:ok, socket} = :ssl.listen(port, options)
    {:ok, socket}
  end

  def handle_info({:ssl, socket, data}, socket) do
    IO.puts("Received data: #{data}")

    {:noreply, socket}
  end
end
