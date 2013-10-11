define ['golem'], (Golem) ->

  class Chat
    constructor: ->
      @WsConn = new Golem.Connection("ws://" + CONFIG.baseUrl + "/ws", CONFIG.wsDebug)
      @WsConn.on "open", @open
      @WsConn.on "message", @message

    open: =>
      @WsConn.emit("subscribe", {channel: CONFIG.projChanName})

    message: (data) =>
      console.log data

  Controllers =
    chatSubscribe: ->
      window.chat = new Chat()

  return Controllers
