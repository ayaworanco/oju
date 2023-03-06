defmodule Oluwoye.Server do
  use GenServer
  require Logger

  def start_link(options) do
    GenServer.start_link(__MODULE__, options, name: __MODULE__)
  end

  def init([port: port, applications_file: file] = _opts) do
    case Oluwoye.ConfigMapper.from_file(file) do
      %Oluwoye.Exceptions.ConfigMapperError{msg: reason} ->
        {:error, reason}

      apps ->
        start_server(port, apps)
    end
  end

  # API functions

  # Server functions
  def handle_info(:accept, %{socket: socket} = state) do
    {:ok, _} = :gen_tcp.accept(socket)
    Logger.info("Client connected")

    {:noreply, state}
  end

  @doc """
  This handle info will get the packet and parse to an WORD struct
  then this gonna be passed to a action function to be executed
  """
  def handle_info({:tcp, _tcp_socket, data}, %{apps: apps} = state) do
    case Awo.Parser.parse(data) do
      %Awo.ParserError{msg: msg} ->
        Logger.error(msg)

      packet ->
        Logger.debug("#{inspect(packet)}")
        action(packet, apps)
    end

    {:noreply, state}
  end

  def handle_info({:tcp_closed, _tcp_socket}, %{socket: _socket} = state) do
    Logger.info("Client closing")

    {:noreply, state}
  end

  # Helper private functions
  defp start_server(port, apps) do
    {:ok, socket} = :gen_tcp.listen(port, [:binary, active: true, packet: :line, reuseaddr: true])

    send(self(), :accept)
    Logger.info("Oluwoye::SERVER started at #{port}")
    {:ok, %{socket: socket, apps: apps}}
  end

  defp action(packet, apps) do
    case packet do
      %Awo.Words.Log{key: key} ->
        check_from_key_and_handle_log(key, apps, packet)
    end
  end

  defp check_from_key_and_handle_log(key, apps, packet) do
    case Enum.find(apps, fn app -> app["key"] == key end) do
      %{"name" => _name, "key" => _key} ->
        Oluwoye.WordActions.Log.handle(packet)

      _ ->
        %Oluwoye.Exceptions.AuthorizationError{msg: "KEY #{key} is not from this packet"}
    end
  end
end
