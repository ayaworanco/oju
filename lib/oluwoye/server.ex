defmodule Oluwoye.Server do
  use GenServer
  require Logger

  def start_link(options) do
    GenServer.start_link(__MODULE__, options, name: __MODULE__)
  end

  def init([port: port] = _opts) do
    {:ok, socket} = :gen_tcp.listen(port, [:binary, active: true, packet: :line, reuseaddr: true])
    send(self(), :accept)

    Logger.info("Oluwoye::SERVER started at #{port}")
    {:ok, socket}
  end

  def handle_info(:accept, socket) do
    {:ok, _} = :gen_tcp.accept(socket)
    Logger.info("Client connected")

    {:noreply, socket}
  end

  def handle_info({:tcp, _tcp_socket, data}, socket) do
    case Awo.Parser.parse(data) do
      %Awo.ParserError{msg: msg} -> Logger.error(msg)
      packet -> Logger.debug("#{inspect(packet)}")
    end

    {:noreply, socket}
  end

  def handle_info({:tcp_closed, _tcp_socket}, socket) do
    Logger.info("Client closing")

    {:noreply, socket}
  end
end
