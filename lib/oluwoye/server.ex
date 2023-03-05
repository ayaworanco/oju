defmodule Oluwoye.Server do
  use GenServer
  require Logger

  def start_link(options) do
    GenServer.start_link(__MODULE__, options, name: __MODULE__)
  end

  def init([port: port, applications_file: file] = _opts) do
    %{"applications" => apps} = Oluwoye.ConfigMapper.from_file(file)
    {:ok, socket} = :gen_tcp.listen(port, [:binary, active: true, packet: :line, reuseaddr: true])
    send(self(), :accept)

    Logger.info("Oluwoye::SERVER started at #{port}")
    {:ok, %{socket: socket, apps: apps}}
  end

  def handle_info(:accept, %{socket: socket} = state) do
    {:ok, _} = :gen_tcp.accept(socket)
    Logger.info("Client connected")

    {:noreply, state}
  end

  def handle_info({:tcp, _tcp_socket, data}, %{socket: _socket} = state) do
    case Awo.Parser.parse(data) do
      %Awo.ParserError{msg: msg} -> Logger.error(msg)
      packet -> Logger.debug("#{inspect(packet)}")
    end

    {:noreply, state}
  end

  def handle_info({:tcp_closed, _tcp_socket}, %{socket: _socket} = state) do
    Logger.info("Client closing")

    {:noreply, state}
  end
end
